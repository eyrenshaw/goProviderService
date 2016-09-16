package messagebus

import "testing"

func TestSendingToQueue(t *testing.T) {

	SendMessage("hello", "Bubba")
}

func TestReceivingFromQueue(t *testing.T) {

	//	Receive("hello")
}
