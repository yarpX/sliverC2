package db

import (
	"github.com/bishopfox/sliver/server/db/models"
	"gorm.io/gorm"
)

var (
	// ErrRecordNotFound - Record not found error
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// ImplantBuilds - Return all implant builds
func ImplantBuilds() ([]*models.ImplantBuild, error) {
	dbSession := Session()
	builds := []*models.ImplantBuild{}
	result := dbSession.Where(&models.ImplantBuild{}).Find(&builds)
	return builds, result.Error
}

// ImplantBuildByName - Fetch implant build by name
func ImplantBuildByName(name string) (*models.ImplantBuild, error) {
	dbSession := Session()
	build := models.ImplantBuild{}
	result := dbSession.Where(&models.ImplantBuild{Name: name}).First(&build)
	return &build, result.Error
}

// ImplantBuildNames - Fetch a list of all build names
func ImplantBuildNames() ([]string, error) {
	dbSession := Session()
	builds := []*models.ImplantBuild{}
	result := dbSession.Where(&models.ImplantBuild{}).Find(&builds)
	if result.Error != nil {
		return []string{}, result.Error
	}
	names := []string{}
	for _, build := range builds {
		names = append(names, build.Name)
	}
	return names, result.Error
}

// Profiles - Fetch a map of name<->profiles current in the database
func Profiles() ([]*models.ImplantProfile, error) {
	profiles := []*models.ImplantProfile{}
	dbSession := Session()
	result := dbSession.Where(&models.ImplantProfile{}).Find(&profiles)
	return profiles, result.Error
}

// ImplantProfileByName - Fetch implant build by name
func ImplantProfileByName(name string) (*models.ImplantProfile, error) {
	dbSession := Session()
	profile := models.ImplantProfile{}
	result := dbSession.Where(&models.ImplantProfile{Name: name}).First(&profile)
	return &profile, result.Error
}

// ImplantProfileNames - Fetch a list of all build names
func ImplantProfileNames() ([]string, error) {
	dbSession := Session()
	profiles := []*models.ImplantProfile{}
	result := dbSession.Where(&models.ImplantProfile{}).Find(&profiles)
	if result.Error != nil {
		return []string{}, result.Error
	}
	names := []string{}
	for _, build := range profiles {
		names = append(names, build.Name)
	}
	return names, result.Error
}

// ProfileByName - Fetch a single profile from the database
func ProfileByName(name string) (*models.ImplantProfile, error) {
	dbProfile := &models.ImplantProfile{}
	dbSession := Session()
	result := dbSession.Where(&models.ImplantProfile{Name: name}).Find(&dbProfile)
	return dbProfile, result.Error
}

// ListCanaries - List of all embedded canaries
func ListCanaries() ([]*models.DNSCanary, error) {
	canaries := []*models.DNSCanary{}
	dbSession := Session()
	result := dbSession.Where(&models.DNSCanary{}).Find(&canaries)
	return canaries, result.Error
}

// CanaryByDomain - Check if a canary exists
func CanaryByDomain(domain string) (*models.DNSCanary, error) {
	dbSession := Session()
	canary := models.DNSCanary{}
	result := dbSession.Where(&models.DNSCanary{Domain: domain}).First(&canary)
	return &canary, result.Error
}

// WebsiteByName - Get website by name
func WebsiteByName(name string) (*models.Website, error) {
	website := models.Website{}
	dbSession := Session()
	result := dbSession.Where(&models.Website{Name: name}).First(&website)
	if result.Error != nil {
		return nil, result.Error
	}
	return &website, nil
}
