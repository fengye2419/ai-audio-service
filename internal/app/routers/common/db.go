package common

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/app/models"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/setting"
)

func InitDBEngine(ctx context.Context) (err error) {
	logrus.Infof("Beginning ORM engine initialization.")
	for i := 0; i < setting.Database.DBConnectRetries; i++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("aborted due to shutdown:\nin retry ORM engine initialization")
		default:
		}
		logrus.Infof("ORM engine initialization attempt #%d/%d...", i+1, setting.Database.DBConnectRetries)
		if err = models.NewEngine(ctx); err == nil {
			break
		} else if i == setting.Database.DBConnectRetries-1 {
			return err
		}
		logrus.Errorf("ORM engine initialization attempt #%d/%d failed. Error: %v", i+1, setting.Database.DBConnectRetries, err)
		logrus.Infof("Backing off for %d seconds", int64(setting.Database.DBConnectBackoff/time.Second))
		time.Sleep(setting.Database.DBConnectBackoff)
	}
	models.HasEngine = true
	return nil
}
