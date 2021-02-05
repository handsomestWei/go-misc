package time

import (
	"github.com/robfig/cron"
	"time"
)

// 创建定时任务。spec例*/1 * * * * ?
func NewTask(spec string, f func()) {
	go func() {
		task := cron.New()
		task.AddFunc(spec, f)
		task.Start()
		select {}
	}()
}

// 创建延迟启动的定时器。避免多个定时器同时启动引发性能问题
func NewTickerDelay(f func(), d, delay time.Duration, stopCh <-chan struct{}) {
	if delay > 0 {
		time.Sleep(delay)
	}

	NewTicker(f, d, stopCh)
}

// 创建定时器
func NewTicker(f func(), d time.Duration, stopCh <-chan struct{}) {
	go func() {
		ticker := time.NewTicker(d)
		for {
			select {
			case <-stopCh:
				ticker.Stop()
				return
			case <-ticker.C:
				f()
			}
		}
	}()
}
