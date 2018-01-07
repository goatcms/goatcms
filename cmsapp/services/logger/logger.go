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

// LoggerFactory create new logger instance
func LoggerFactory(dp dependency.Provider) (interface{}, error) {
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

func (logger *Logger) DevLog(format string, data ...interface{}) {
	if logger.devLvl {
		logger.log.Printf(format, data...)
	}
}

func (logger *Logger) TestLog(format string, data ...interface{}) {
	if logger.testLvl {
		logger.log.Printf(format, data...)
	}
}

func (logger *Logger) ProdLog(format string, data ...interface{}) {
	if logger.prodLvl {
		logger.log.Printf(format, data...)
	}
}

func (logger *Logger) ErrorLog(format string, data ...interface{}) {
	logger.log.Fatalf(format, data...)
}
