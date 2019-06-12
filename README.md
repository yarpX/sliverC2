Sliver
======

![Sliver](/.github/images/sliver.jpeg)

⚠️ __Warning:__ Sliver is currently in __alpha__, you've been warned :) and please consider [contributing](/CONTRIBUTING.md)

Sliver is a general purpose cross-platform implant framework that supports C2 over Mutual-TLS, HTTP(S), and DNS. Implants are dynamically compiled with unique X.509 certificates signed by a per-instance certificate authority generated when you first run the binary.

### Features

* Dynamic code generation
* Compile-time obfuscation
* Local and remote process injection
* Anti-anti-anti-forensics
* Secure C2 over mTLS, HTTP(S), and DNS
* Windows process migration
* Windows user token manipulation
* Multiplayer-mode
* Procedurally generated C2 over HTTP
* Let's Encrypt integration
* In-memory .NET assembly execution
* [DNS Canary](https://github.com/BishopFox/sliver/wiki/DNS-C2#dns-canaries) Blue Team Detection

### Getting Started

Download the latest [release](https://github.com/BishopFox/sliver/releases) and see the Sliver [wiki](https://github.com/BishopFox/sliver/wiki/Getting-Started) for a quick tutorial on basic setup and usage. To get the very latest and greatest compile from source.

### Compile From Source

Do a `git clone` of the Sliver repo into your local `$GOPATH/github.com/bishopfox/sliver` and then run the `build.py` script (requires Docker), or for details see the [wiki](https://github.com/BishopFox/sliver/wiki/Compile-From-Source).

### Source Code

The source code repo contains the following directories:

 * `assets/` - Static assets that are embedded into the server binary, generated by `go-assets.sh`
 * `client/` - Client code, the majority of this code is also used by the server
 * `protobuf/` - Protobuf code
 * `server/` - Server-side code
 * `sliver/` - Implant code, rendered by the server at runtime
 * `util/` - Utility functions that may be shared by the server and client

### License - GPLv3

Sliver is licesed under [GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html), some subcomponets have seperate licenses. See their respective subdirectories in this project for details.
