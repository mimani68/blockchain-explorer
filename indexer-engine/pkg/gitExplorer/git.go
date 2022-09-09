package gitexplorer

import (
	"os/exec"
	"strings"

	"app.io/pkg/logHandler"
)

func ShowGitHash() string {
	cmd := exec.Command("git", "log", "--oneline")
	stdout, err := cmd.Output()
	if err != nil {
		logHandler.Log(logHandler.ERROR, err.Error())
	}
	var result = string(stdout)
	lines := strings.Split(result, "\n")
	cols := strings.Split(lines[0], " ")
	if cols[0] == "" {
		return ""
	} else {
		return cols[0]
	}
}
