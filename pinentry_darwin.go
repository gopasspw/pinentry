//go:build darwin
// +build darwin

package pinentry

import (
	"os/exec"

	"github.com/gopasspw/pinentry/gpgconf"
)

// GetBinary always returns pinentry-mac
func GetBinary() string {
	// check, whether the returned path acutally exists
	if _, err := exec.LookPath("pinentry-mac"); err == nil {
		return "pinentry-mac"
	}
	if p, err := gpgconf.Path("pinentry"); err == nil && p != "" {
		return p
	}
	return "pinentry-mac"
}
