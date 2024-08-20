package setting

import (
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// log settings
	LogMode       string
	RouterLogMode string
	XormLogMode   string
	LogLevel      string
	LogRootPath   string
)

var defaultLogFile = "app_logrus.log"
var routerLogFile = "router.log"
var xormLogFile = "xorm.log"
var RouterLogger *logrus.Logger
var XormLogger *logrus.Logger

func InitLogConfig() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		logrus.Fatalf("Fail to get section 'log': %v", err)
	}
	LogMode = sec.Key("MODE").MustString("file")
	RouterLogMode = sec.Key("ROUTER").MustString("file")
	XormLogMode = sec.Key("XORM").MustString("file")
	LogLevel = sec.Key("LEVEL").MustString("info")
	LogRootPath = sec.Key("ROOT_PATH").String()
}

// GetDefaultLogFile returns the default log file.
func GetDefaultLogFile() string {
	return LogRootPath + "/" + defaultLogFile
}

// GetRouterLogFile returns the router log file.
func GetRouterLogFile() string {
	return LogRootPath + "/" + routerLogFile
}

// GetXormLogFile returns the xorm log file.
func GetXormLogFile() string {
	return LogRootPath + "/" + xormLogFile
}

// NewLogService initializes the log service.
func NewLogService() {
	newLogService()
	newRouterLogService()
	newXormLogService()
}

// newLogService initializes the log service.
func newLogService() {
	SetLogEntry("default", LogMode)
}

// newRouterLogService initializes the router log service.
func newRouterLogService() {
	SetLogEntry("router", RouterLogMode)
}

// newXormLogService initializes the xorm log service.
func newXormLogService() {
	SetLogEntry("xorm", XormLogMode)
}

// SetLogEntry sets the log entry.
func SetLogEntry(logType, mode string) {
	// Create a new logger
	logger := logrus.New()

	var fileName string
	switch logType {
	case "router":
		RouterLogger = logger
		fileName = GetRouterLogFile()
	case "xorm":
		XormLogger = logger
		fileName = GetXormLogFile()
	default:
		fileName = GetDefaultLogFile()
	}

	// Report caller
	logger.SetReportCaller(true)
	// Set log level
	level, _ := logrus.ParseLevel(LogLevel)
	logger.SetLevel(level)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return path.Base(frame.Function) + "()", frame.File + ":" + strconv.Itoa(frame.Line)
		},
	})

	// Set log output
	var output io.Writer
	if mode == "file" {
		output = &lumberjack.Logger{
			Filename:  fileName,
			MaxSize:   256,
			MaxAge:    180,
			Compress:  true,
			LocalTime: true,
		}
	} else {
		output = os.Stdout
	}
	logger.SetOutput(output)
}
