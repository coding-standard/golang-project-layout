package log

import (
	"fmt"
	"time"

	"github.com/coding-standard/golang-project-layout/pkg/settings"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", settings.AppSetting.RuntimeRootPath, settings.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		settings.AppSetting.LogSaveName,
		time.Now().Format(settings.AppSetting.TimeFormat),
		settings.AppSetting.LogFileExt,
	)
}
