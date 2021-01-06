package main

import (
	// nested "github.com/antonfisher/nested-logrus-formatter"

	log "github.com/shenjing023/llog"
	// log "github.com/sirupsen/logrus"
)

func main() {
	// if err := log.SetFileLogger(
	// 	"test.log"+"-%Y%m%d%H%M",
	// 	log.WithCaller(true),
	// 	log.WithMaxAge(30*time.Second),
	// 	log.WithRotationTime(time.Second*10),
	// 	log.WithJSON(true),
	// 	log.WithPrettyPrint(true),
	// ); err != nil {
	// 	fmt.Errorf("init log error %v", err)
	// 	return
	// }
	log.SetConsoleLogger(
		log.WithCaller(true),
		log.WithLevel(log.TraceLevel),
		// log.WithJSON(true),
		// log.WithMaxAge(30*time.Second),
	)
	// log.SetReportCaller(true)

	// 打印日志
	// log.Debug("调试信息")
	log.WithField("component", "rest").Warn("warn message")
	log.Info("提示信息")
	// log.Warn("警告信息")
	log.Error("错误信息")
	// log.Printf("hello world")
	log.WithField("component", "rest").Warn("warn message")
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// var wg sync.WaitGroup
	// wg.Add(100)
	// f := func() {
	// 	for i := 0; i < 100; i++ {
	// 		defer wg.Done()
	// 		log.WithFields(log.Fields{

	// 			"animal": "walrus",

	// 			"number": i,
	// 		}).Info("A walrus appears")

	// 		log.Error("xxxxxxxxxxxxxxx")
	// 		log.Error("xxxxxxxxxxxxxxx2")
	// 		log.Error("xxxxxxxxxxxxxxx3")

	// 		time.Sleep(time.Second)
	// 	}

	// }
	// f()
	// wg.Wait()

}
