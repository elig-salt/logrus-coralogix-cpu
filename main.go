package main

import (
	"fmt"
	"os"
	"time"

	"net/http"
	_ "net/http/pprof"

	coralogix "github.com/coralogix/go-coralogix-sdk"
	logrus "github.com/sirupsen/logrus"
)

func attachCoralogixHook(log *logrus.Logger, privateKey, applicationKey, subsystemKey string) {
	CoralogixHook := coralogix.NewCoralogixHook(
		privateKey, applicationKey, subsystemKey,
	)

	// defer CoralogixHook.Close() // No need to close in our case
	log.AddHook(CoralogixHook)
}

func runLogs(log *logrus.Logger) {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		log.WithFields(logrus.Fields{
			"hello":       "world",
			"verisonDate": time.Now(),
		}).Info("Hello World Message")
	}
}

func newLogger() *logrus.Logger {
	var log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	return log
}

func runProfiler() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func main() {
	runProfiler()

	log := newLogger()
	privateKey, applicationKey, subsystemKey :=
		os.Getenv("CORALOGIX_PRIVATE_KEY"),
		os.Getenv("CORALOGIX_APPLICATION_NAME"),
		os.Getenv("CORALOGIX_SUBSYSTEM_NAME")

	if privateKey == "" || applicationKey == "" || subsystemKey == "" {
		panic("Please specificy env vars: CORALOGIX_PRIVATE_KEY, CORALOGIX_APPLICATION_NAME, CORALOGIX_SUBSYSTEM_NAME")
	}

	// COMMENT THE NEXT LINE OUT - and see cpu difference
	attachCoralogixHook(log, privateKey, applicationKey, subsystemKey)

	go runLogs(log)
	<-time.After(2 * time.Minute)
}
