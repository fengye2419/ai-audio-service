package models

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/setting"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/util"
	xormlog "xorm.io/xorm/log"
)

type SQLLogger struct {
	logger *logrus.Logger
}

func (s *SQLLogger) BeforeSQL(ctx xormlog.LogContext) {
}

func (s *SQLLogger) AfterSQL(ctx xormlog.LogContext) {
	var funcVal, fileVal string
	pc := make([]uintptr, 20)
	_ = runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc)
	for f, again := frames.Next(); again; f, again = frames.Next() {
		if strings.Contains(f.Function, "/models") {
			funcVal, fileVal = f.Function, fmt.Sprintf("%s:%d", f.File, f.Line)
			break
		}
	}
	s.logger.WithFields(logrus.Fields{
		"SQL":         ctx.SQL,
		"Args":        ctx.Args,
		"ExecuteTime": ctx.ExecuteTime,
		"Err":         ctx.Err,
		"func":        funcVal,
		"file":        fileVal,
		"request_id":  util.GetRequestID(ctx.Ctx),
	}).Info("xorm-sql")
}

func (s *SQLLogger) Debugf(format string, v ...interface{}) {
	s.logger.Debugf(format, v...)
}

func (s *SQLLogger) Errorf(format string, v ...interface{}) {
	s.logger.Errorf(format, v...)
}

func (s *SQLLogger) Infof(format string, v ...interface{}) {
	s.logger.Infof(format, v...)
}

func (s *SQLLogger) Warnf(format string, v ...interface{}) {
	s.logger.Warnf(format, v...)
}

func (s *SQLLogger) Level() xormlog.LogLevel {
	switch s.logger.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return xormlog.LOG_DEBUG
	case logrus.ErrorLevel:
		return xormlog.LOG_ERR
	case logrus.InfoLevel:
		return xormlog.LOG_INFO
	case logrus.WarnLevel:
		return xormlog.LOG_WARNING
	default:
		return xormlog.LOG_OFF
	}
}

func (s *SQLLogger) SetLevel(l xormlog.LogLevel) {
}

func (s *SQLLogger) ShowSQL(show ...bool) {
}

func (s *SQLLogger) IsShowSQL() bool {
	return true
}

func NewSQLLogger() xormlog.ContextLogger {
	return &SQLLogger{
		logger: setting.XormLogger,
	}
}
