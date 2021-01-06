package main

import (
	log "github.com/shenjing023/llog"
)

func main() {
	log.Debug("调试信息")
	log.Info("提示信息")
	log.Warn("警告信息")
	log.Error("错误信息")
	log.WithField("key", "value").Warn("warn message")
	log.WithFields(log.Fields{
		"a": "b",
		"c": 1,
	}).Info("xxxxxxxxxxxxxxx")
}
