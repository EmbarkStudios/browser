package browser

import "os/exec"

func openBrowser(url string) error {
	// Unfortunately, due to how WSL works, we can't differentiate between WSL and Linux via build tags
	// So we try xdg-open first, and if that fails, use wslview
	if err := runCmd("xdg-open", url); err != nil {
		if wslErr := runCmd("wslview", url); wslErr != nil {
			return err
		}
	}

	return nil
}

func setFlags(cmd *exec.Cmd) {}
