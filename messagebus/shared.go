package messagebus

import (
	"fmt"
	"log"
	"providerservice/logtofile"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		logtofile.WriteMessage("ERROR: " + msg)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
