package setting

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Database = struct {
		Type             string
		Host             string
		Name             string
		User             string
		Passwd           string
		Schema           string
		SSLMode          string
		LogSQL           bool
		Charset          string
		Timeout          int // seconds
		DBConnectRetries int
		DBConnectBackoff time.Duration
		MaxIdleConns     int
		MaxOpenConns     int
		ConnMaxLifetime  time.Duration
	}{
		Type:    "mysql",
		Timeout: 500,
	}
)

func InitDBConfig() {
	defaultCharset := "utf8mb4"
	sec := Cfg.Section("database")
	Database.Type = sec.Key("TYPE").String()
	Database.Host = sec.Key("HOST").String()
	Database.Name = sec.Key("NAME").String()
	Database.User = sec.Key("USER").String()
	Database.Passwd = sec.Key("PASSWD").String()
	Database.Schema = sec.Key("SCHEMA").String()
	Database.SSLMode = sec.Key("SSL_MODE").MustString("disable")
	Database.Charset = sec.Key("CHARSET").In(defaultCharset, []string{"utf8", "utf8mb4"})
	Database.MaxIdleConns = sec.Key("MAX_IDLE_CONNS").MustInt(2)
	Database.ConnMaxLifetime = sec.Key("CONN_MAX_LIFETIME").MustDuration(3 * time.Second)
	Database.MaxOpenConns = sec.Key("MAX_OPEN_CONNS").MustInt(0)
	Database.LogSQL = sec.Key("LOG_SQL").MustBool(true)
	Database.DBConnectRetries = sec.Key("DB_RETRIES").MustInt(10)
	Database.DBConnectBackoff = sec.Key("DB_RETRY_BACKOFF").MustDuration(3 * time.Second)
}

// DBConnStrings returns database connection strings
func DBConnStrings() ([]string, error) {
	connStr := make([]string, 0)
	var Param = "?"
	if strings.Contains(Database.Name, Param) {
		Param = "&"
	}
	switch Database.Type {
	case "mysql":
		hosts := strings.Split(Database.Host, ",")
		connType := "tcp"
		tls := Database.SSLMode
		if tls == "disable" { // allow (Postgres-inspired) default value to work in MySQL
			tls = "false"
		}
		for _, host := range hosts {
			if len(host) > 0 && host[0] == '/' { // looks like a unix socket
				connType = "unix"
			}
			connStr = append(connStr, fmt.Sprintf("%s:%s@%s(%s)/%s%scharset=%s&parseTime=true&tls=%s",
				Database.User, Database.Passwd, connType, host, Database.Name, Param, Database.Charset, tls))
		}
	default:
		return []string{""}, fmt.Errorf("unknown database type: %s", Database.Type)
	}

	return connStr, nil
}
