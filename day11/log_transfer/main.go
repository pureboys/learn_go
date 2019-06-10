package main

import "github.com/astaxie/beego/logs"

func main() {

	filename := "/home/oliver/go/src/demo/day11/log_transfer/conf/log_agent.ini"
	err := initConfig("ini", filename)
	if err != nil {
		panic(err)
		return
	}

	err = initLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		panic(err)
		return
	}

	err = initKafka(logConfig.KafkaAddr, logConfig.KafkaTopic)
	if err != nil {
		logs.Error("init kafka failed, err: %v", err)
		return
	}

	logs.Debug("init kafka success")

	/*

		err = initES()
		if err != nil {
			logs.Error("init es failed, err: %v", err)
			return
		}

		err = run()
		if err != nil {
			logs.Error("run failed, err: %v", err)
			return
		}

		logs.Warn("warning, log_transfer is exited")
	*/
}
