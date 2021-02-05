package log

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 参考https://github.com/uber-go/zap目录改为go.uber.org/zap
func NewLogWithZap(level string) *zap.Logger {
	var js string
	js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "console",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, level)

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	Logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return Logger

}
