package setting

import (
	"github.com/sirupsen/logrus"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/util"
	"gopkg.in/ini.v1"
)

var (
	Cfg        *ini.File
	ConfigFile string
)

// SetConfigFile sets the config file
func SetConfigFile(file string) {
	ConfigFile = file
}

// Setup initializes the configuration
func Setup() {
	Cfg = ini.Empty()
	isFile, err := util.IsFile(ConfigFile)
	if err != nil {
		logrus.Fatalf("Fail to check file '%s': %v", ConfigFile, err)
	}
	if isFile {
		if err := Cfg.Append(ConfigFile); err != nil {
			logrus.Fatalf("Fail to load config file: %v", err)
		}
	} else {
		logrus.Fatalf("'%s' is not found", ConfigFile)
	}

	InitAppConfig()
	InitOverseerConfig()
	InitLogConfig()
	InitDBConfig()
}

// NewService initializes the services
func NewService() {
	NewLogService()
}
