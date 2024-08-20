package setting

import (
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	AppName string
	AppEnv  string
)

// IsLocal returns true if the app is running in local environment
func IsLocal() bool {
	return strings.EqualFold(AppEnv, "local")
}

func InitAppConfig() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		logrus.Fatalf("Fail to get section 'app': %v", err)
	}
	AppName = sec.Key("APP_NAME").String()
	AppEnv = sec.Key("APP_ENV").MustString("")
}
