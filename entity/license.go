package entity

import (
	"fmt"
	"strings"
)

const LicenseTemplate = "%v#%v|%v%v#%v#%v3%v6%v#%v#%v#%v#"

//"%d#%s|%d%d#%d#%d3%d6%d#%d#%d#%d#";

type License struct {
	UserName     string
	LicenseType  int
	MajorVersion string
	MinorVersion string
	count        int
	unknown      int
	OpenGames    bool
	OpenPlugin   bool
}

func NewLicense(userName string, licenseType int, version string, openGame, openPlugin bool) License {
	split := strings.Split(version, ".")
	return License{
		UserName:     userName,
		LicenseType:  licenseType,
		MajorVersion: split[0],
		MinorVersion: split[1],
		count:        0,
		unknown:      0,
		OpenGames:    openGame,
		OpenPlugin:   openPlugin,
	}

}

func (license *License) GetLicenseKey() string {
	license.count = 1
	game := 0
	plugin := 0
	if license.OpenGames {
		game = 1
	}
	if license.OpenPlugin {
		plugin = 1
	}
	return fmt.Sprintf(LicenseTemplate, license.LicenseType,
		license.UserName, license.MajorVersion, license.MinorVersion, license.count,
		license.MajorVersion, license.MinorVersion, license.MinorVersion,
		license.unknown, game, plugin,
	)

}
