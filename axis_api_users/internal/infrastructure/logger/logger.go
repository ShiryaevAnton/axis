package logger

import (
	"github.com/ShiryaevAnton/axis/axis_api_utils/apilogger"
)

var (
	log apilogger.Logger
)

//NewLogger ...
func NewLogger(logLevel string, logOutput string) error {
	zapLogger, err := apilogger.InitLogger(logLevel, logOutput)
	if err != nil {
		return err
	}
	log = apilogger.NewLogger(zapLogger)
	return nil
}

//GetLogger ...
func GetLogger() apilogger.Logger {
	return log
}
