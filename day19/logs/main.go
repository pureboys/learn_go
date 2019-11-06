package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	//var log = logrus.New()
	//// You could set this to any `io.Writer` such as a file
	//file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	// log.Out = file
	//} else {
	// log.Info("Failed to log to file, using default stderr")
	//}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.Error("hello")

}
