package main

import (
	"fmt"

	"github.com/suryakencana007/mimir/log"
)

func main() {
	fmt.Println("Hello world")

	// init log to assign memory addres
	log.ZapInit()

	// error loging using mimir/log
	log.Error("Query Find By ID", log.Field("Query", "tes"))

	// warning loging using mimir/log
	log.Warn("Query Find By ID", log.Field("Query", "tes"))
}
