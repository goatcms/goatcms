package logger

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
)

// Logger show logger messages
type Logger struct {
	devLVL  bool
	testLVL bool
	prodLVL bool
	log     *log.Logger
}

// Factory create new logger instance
func Factory(dp dependency.Provider) (interface{}, error) {
	var deps struct {
		LogLVL string `argument:"?loglvl"`
	}
	if err := dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	logger := &Logger{
		devLVL:  deps.LogLVL == "dev",
		testLVL: deps.LogLVL == "dev" || deps.LogLVL == "test",
		prodLVL: true,
		log:     log.New(os.Stdout, " ", log.Ltime|log.Ldate),
	}
	logger.TestLog("Log level: %s", deps.LogLVL)
	return services.Logger(logger), nil
}

// DevLog print dev level logs
func (logger *Logger) DevLog(format string, data ...interface{}) {
	if logger.devLVL {
		logger.printf(format, data...)
	}
}

// TestLog print test level logs
func (logger *Logger) TestLog(format string, data ...interface{}) {
	if logger.testLVL {
		logger.printf(format, data...)
	}
}

// ProdLog print prod level logs
func (logger *Logger) ProdLog(format string, data ...interface{}) {
	if logger.prodLVL {
		logger.printf(format, data...)
	}
}

// ErrorLog print error level logs
func (logger *Logger) ErrorLog(format string, data ...interface{}) {
	logger.printf("ERROR: "+format, data...)
}

// IsProdLVL return true if prod level messages v set
func (logger *Logger) IsProdLVL() bool {
	return logger.prodLVL
}

// IsDevLVL return true if dev level messages is set
func (logger *Logger) IsDevLVL() bool {
	return logger.devLVL
}

// IsTestLVL return true if test level messages is set
func (logger *Logger) IsTestLVL() bool {
	return logger.devLVL
}

// print value to logs
func (logger *Logger) printf(format string, data ...interface{}) {
	var (
		err      error
		jsonData []string
	)
	jsonData = make([]string, len(data))
	if logger.devLVL {
		for i, val := range data {
			var (
				typeName string
				msg      string
				json     string
			)
			typeName = reflect.TypeOf(val).Name()
			if e, ok := val.(error); ok {
				msg = e.Error()
			}
			if json, err = varutil.ObjectToJSON(val); err != nil {
				panic(err)
			}
			jsonData[i] = fmt.Sprintf("\n%v (%s) %s -> %s\n", i, typeName, msg, json)
		}
		format += strings.Join(jsonData, "")
	}
	logger.log.Printf(format, data...)
}
