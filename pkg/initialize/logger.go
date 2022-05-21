package initialize

import "go.uber.org/zap"
// NewDevelopment builds a development Logger that writes DebugLevel and above
// logs to standard error in a human-friendly format.
func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
