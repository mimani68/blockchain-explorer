package job

import (
	"app.io/pkg/logHandler"
)

func SystemCheck() {
	logHandler.Log(logHandler.INFO, "Check system", "JOB")
}
