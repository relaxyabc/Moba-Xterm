package layout

import (
	"errors"
	"fyne.io/fyne/v2"
	"mobaxterm/entity"
	"mobaxterm/register"
	"regexp"
	"strings"
)

var compile *regexp.Regexp

func init() {
	compile, _ = regexp.Compile("^\\d+\\.\\d+$")
}
func NewBlankValidation(reason string) fyne.StringValidator {

	return func(text string) error {
		if text == "" || strings.Trim(text, " ") == "" {
			return errors.New(reason)
		}
		return nil // Nothing to validate with, same as having no validator.
	}
}

func validationAndRegister(window fyne.Window) {
	uname := userName.Text
	if uname == "" || strings.Trim(uname, " ") == "" {
		window.Canvas().Focus(userName)
		return
	}

	version := targetVersion.Text
	if !compile.MatchString(version) {
		window.Canvas().Focus(targetVersion)
		return
	}
	license := entity.NewLicense(uname, 1, version, game.Checked, plugin.Checked)
	register.Register(license.GetLicenseKey(), filePath.Text)
}
