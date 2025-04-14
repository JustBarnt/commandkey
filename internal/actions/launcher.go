package actions

import (
	"fmt"
	"os/exec"
	"runtime"
)

func LaunchEditor() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("notepad.exe")
	case "darwin":
		cmd = exec.Command("open", "-a", "TextEdit")
	default:
		cmd = exec.Command("xdg-open", ".")
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to launch editor:", err)
	}
}
