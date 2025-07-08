package task

import (
	"example.com/example/pkg/logging"
	"github.com/robfig/cron/v3"
	"time"
)

// Start 启动定时任务
func Start() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 30s", func() {
		logging.Info(time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()
}
