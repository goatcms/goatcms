package logger

import (
	"log"
	"os"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/dependency"
)

// Logger show logger messages
type Logger struct {
	devLvl  bool
	testLvl bool
	prodLvl bool
	log     *log.Logger
}

// Factory create new logger instance
func Factory(dp dependency.Provider) (interface{}, error) {
	var deps struct {
		Loglvl string `argument:"?loglvl"`
	}
	if err := dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	logger := &Logger{
		devLvl:  deps.Loglvl == "dev",
		testLvl: deps.Loglvl == "dev" || deps.Loglvl == "test",
		prodLvl: true,
		log:     log.New(os.Stdout, " ", log.Ltime|log.Ldate),
	}
	logger.TestLog("Log level: %s", deps.Loglvl)
	return services.Logger(logger), nil
}

// DevLog print dev level logs
func (logger *Logger) DevLog(format string, data ...interface{}) {
	if logger.devLvl {
		logger.log.Printf(format, data...)
	}
}

// TestLog print test level logs
func (logger *Logger) TestLog(format string, data ...interface{}) {
	if logger.testLvl {
		logger.log.Printf(format, data...)
	}
}

// ProdLog print prod level logs
func (logger *Logger) ProdLog(format string, data ...interface{}) {
	if logger.prodLvl {
		logger.log.Printf(format, data...)
	}
}

// ErrorLog print error level logs
func (logger *Logger) ErrorLog(format string, data ...interface{}) {
	logger.log.Printf("ERROR: "+format, data...)
}

// IsProdLVL return true if prod level messages v set
func (logger *Logger) IsProdLVL() bool {
	return logger.prodLvl
}

// IsDevLVL return true if dev level messages is set
func (logger *Logger) IsDevLVL() bool {
	return logger.devLvl
}

// IsTestLVL return true if test level messages is set
func (logger *Logger) IsTestLVL() bool {
	return logger.devLvl
}
