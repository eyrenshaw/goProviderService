package docdb

import (
	"fmt"
	"testing"
	//	"gopkg.in/mgo.v2/bson"
)

func TestDb1(t *testing.T) {

	items, err := GetPersonByName("Bubba", 0, 0)
	if err != "" {
		panic(err)
	}

	fmt.Printf("%d\n", len(items))
	for index := 0; index < len(items); index++ {
		fmt.Println("index: %d Value: %d\n", index, items[index])
	}
}
