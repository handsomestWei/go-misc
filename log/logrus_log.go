package log

import "github.com/sirupsen/logrus"

// 初始化后直接使用logrus.xxx
func NewLogWithLogrus(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)
}
