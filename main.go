package main

import (
	"providerservice/logtofile"
	"providerservice/messagebus"
)

func init() {
	if !logtofile.FileExists("testlogfile.txt") {
		logtofile.CreateFile()
	}
}

func main() {

	logtofile.WriteMessage("Provider Micro service started.")
	messagebus.ReceiveRpc("rpc_queue")

}
