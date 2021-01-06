# llog
log dependent on logrus and file-rotatelogs

## Example
- console output
```go
import (
	log "github.com/shenjing023/llog"
)

func main(
    log.SetConsoleLogger(
		log.WithCaller(true),
		log.WithLevel(log.TraceLevel),
	)
    log.Debug("调试信息")
	log.Info("提示信息")
	log.Warn("警告信息")
	log.Error("错误信息")
	log.WithField("key", "value").Warn("warn message")
	log.WithFields(log.Fields{
		"a": "b",
		"c": 1,
	}).Info("xxxxxxxxxxxxxxx")
)
```
![20210106233252.png](https://i.loli.net/2021/01/06/PYOWf6QhXLqmAz2.png)

```go
import (
	log "github.com/shenjing023/llog"
)

func main(
    log.SetConsoleLogger(
		log.WithLevel(log.TraceLevel),
        log.WithJSON(true),
	)
    log.Debug("调试信息")
	log.Info("提示信息")
	log.Warn("警告信息")
	log.Error("错误信息")
	log.WithField("key", "value").Warn("warn message")
	log.WithFields(log.Fields{
		"a": "b",
		"c": 1,
	}).Info("xxxxxxxxxxxxxxx")
)
```
```
{"level":"debug","msg":"调试信息","time":"2021/01/06 23:36:14.730"}
{"level":"info","msg":"提示信息","time":"2021/01/06 23:36:14.731"}
{"level":"warning","msg":"警告信息","time":"2021/01/06 23:36:14.731"}
{"level":"error","msg":"错误信息","time":"2021/01/06 23:36:14.731"}
{"key":"value","level":"warning","msg":"warn message","time":"2021/01/06 23:36:14.731"}
{"a":"b","c":1,"level":"info","msg":"xxxxxxxxxxxxxxx","time":"2021/01/06 23:36:14.731"}
```
- file output
```go
package main

import (
	log "github.com/shenjing023/llog"
)

func main() {
	if err := log.SetFileLogger(
		"test.log"+"-%Y%m%d%H%M",
		log.WithCaller(true),
		log.WithMaxAge(30*time.Second),
		log.WithRotationTime(time.Second*10),
		log.WithJSON(true),
		log.WithPrettyPrint(true),
	); err != nil {
		fmt.Errorf("init log error %v", err)
		return
	}

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
```


## Thanks
- nested-logrus-formatter
- beego log
