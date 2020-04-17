package main

import (
	"errors"
	"os"
	"time"

	log "github.com/Codehardt/go-logging"
)

/*
Output of this program is:

{"_HOSTNAME":"my-hostname","_IP":"123.123.123.123","_LVL":"Info","_TIME":"Fri Apr 17 11:40:52 2020","_message":"Starting"}
{"_HOSTNAME":"my-hostname","_IP":"123.123.123.123","_LVL":"Debug","_TIME":"Fri Apr 17 11:40:52 2020","_message":"calculating sum of a and b","a":2,"b":3}
{"_HOSTNAME":"my-hostname","_IP":"123.123.123.123","_LVL":"Notice","_TIME":"Fri Apr 17 11:40:52 2020","_message":"sum of a and b","a":2,"b":3,"sum":5}
{"_HOSTNAME":"my-hostname","_IP":"123.123.123.123","_LVL":"Error","_TIME":"Fri Apr 17 11:40:52 2020","_message":"could not do something with error","error":"no such file or directory"}
{"_HOSTNAME":"my-hostname","_IP":"123.123.123.123","_LVL":"Info","_TIME":"Fri Apr 17 11:40:52 2020","_message":"Exiting"}
*/

func init() {
	opts := []log.Option{
		log.OptionEnableDebug(true),
		log.OptionWithTimeFormat(time.ANSIC),
		log.OptionEnableLocalTime(true),
		log.OptionWithLevelKey("_LVL"),
		log.OptionWithTimeKey("_TIME"),
		log.OptionWithMessageKey("_message"),
		log.OptionWithStaticKV("_HOSTNAME", "my-hostname"),
		log.OptionWithStaticKV("_IP", "123.123.123.123"),
		log.OptionWithWriter(os.Stdout),
	}
	logger := log.New(log.FormatterJSON, opts...)
	log.SetLogger(logger)
}

func main() {
	log.Info("Starting")
	defer log.Info("Exiting")
	a := 2
	b := 3
	log.Debug("calculating sum of a and b", "a", a, "b", b)
	sum := a + b
	log.Notice("sum of a and b", "a", a, "b", b, "sum", sum)
	if err := doSomethingWithError(); err != nil {
		log.Error("could not do something with error", "error", err)
	} else {
		log.Notice("successfully done something without error")
	}
}

func doSomethingWithError() error {
	return errors.New("no such file or directory")
}
