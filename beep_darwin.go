// +build darwin,!linux,!windows,!js

package beeep

import (
	"os"
	"os/exec"
)

var (
	// Default frequency, Hz, middle A
	DefaultFreq = 0.0
	// Default duration in milliseconds
	DefaultDuration = 0
)

// Beep beeps the pc speaker (https://en.wikipedia.org/wiki/PC_speaker).
// On macOS this just sends Bell character (https://en.wikipedia.org/wiki/Bell_character)
func Beep(freq float64, duration int) error {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		// Output the only beep we can
		_, err = os.Stdout.Write([]byte{7})
		return err
	}

	cmd := exec.Command(osa, "-e", `tell application "System Events" to beep`)
	return cmd.Start()
}
