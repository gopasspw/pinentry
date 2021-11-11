// +build !darwin,!windows

package pinentry

import (
	"github.com/gopasspw/pinentry/gpgconf"
	"os/exec"
)

// GetBinary returns the binary name
func GetBinary() string {
	if p, err := gpgconf.Path("pinentry"); err == nil && p != "" {
		// check, whether the returned path acutally exists
		if _, err := exec.LookPath(p); err == nil {
			return p
		}
	}
	return "pinentry"
}
