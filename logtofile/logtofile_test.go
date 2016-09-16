package logtofile

import (
	"log"
	"os"
	"testing"
)

func TestLogging(t *testing.T) {
	if !FileExists("logfile") {
		CreateFile()
	}
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
}
