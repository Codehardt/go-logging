package main

import (
	"errors"

	log "github.com/Codehardt/go-logging"
)

/*
Output of this program is:

2020-04-17T08:59:51Z [INF] Starting
2020-04-17T08:59:51Z [NOT] sum of a and b A: 2 B: 3 SUM: 5
2020-04-17T08:59:51Z [ERR] could not do something with error ERROR: no such file or directory
2020-04-17T08:59:51Z [INF] Exiting
*/

func main() {
	log.Info("Starting")
	defer log.Info("Exiting")
	a := 2
	b := 3
	sum := 5
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
