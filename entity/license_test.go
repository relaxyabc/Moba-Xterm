package entity

import (
	"fmt"
	"testing"
)

func TestGetLicenseKey(t *testing.T) {
	license := License{
		UserName:     "wanna",
		LicenseType:  1,
		MajorVersion: "11",
		MinorVersion: "0",
		count:        1,
		unknown:      0,
		OpenGames:    false,
		OpenPlugin:   false,
	}
	fmt.Println(license.GetLicenseKey())
}
