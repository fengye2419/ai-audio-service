package models

import (
	"context"
	"fmt"

	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/setting"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	x *xorm.EngineGroup

	// HasEngine specifies if we have a xorm.Engine
	HasEngine bool
)

func NewEngine(ctx context.Context) (err error) {
	if err = SetEngine(); err != nil {
		return err
	}

	x.SetDefaultContext(ctx)

	if err = x.Ping(); err != nil {
		return err
	}

	return nil
}

func SetEngine() (err error) {
	x, err = GetEngineGroup()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	x.SetMapper(names.GonicMapper{})
	x.SetLogger(NewSQLLogger())
	x.ShowSQL(setting.Database.LogSQL)
	x.SetMaxOpenConns(setting.Database.MaxOpenConns)
	x.SetMaxIdleConns(setting.Database.MaxIdleConns)
	x.SetConnMaxLifetime(setting.Database.ConnMaxLifetime)
	return nil
}

func GetEngineGroup() (*xorm.EngineGroup, error) {
	connStrings, err := setting.DBConnStrings()
	if err != nil {
		return nil, err
	}

	var engine *xorm.EngineGroup

	if engine, err = xorm.NewEngineGroup(setting.Database.Type, connStrings, xorm.LeastConnPolicy()); err != nil {
		return nil, err
	}

	engine.Dialect().SetParams(map[string]string{"rowFormat": "DYNAMIC"})
	engine.SetSchema(setting.Database.Schema)
	return engine, nil
}
