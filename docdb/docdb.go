package docdb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

var (
	mgoSession   *mgo.Session
	databaseName = "test"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

// this is a higher order function.. needs to be somewhere else as i understand it better.
func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}

func SearchPerson(q interface{}, skip int, limit int) (searchResults []Person, searchErr string) {
	searchErr = ""
	searchResults = []Person{}
	query := func(c *mgo.Collection) error {
		fn := c.Find(q).Skip(skip).Limit(limit).All(&searchResults)
		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(&searchResults)
		}
		return fn
	}
	search := func() error {
		return withCollection("people", query)
	}
	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return
}

func GetPersonByName(lastName string, skip int, limit int) (searchResults []Person, searchErr string) {
	searchResults, searchErr = SearchPerson(bson.M{"name": lastName}, skip, limit)
	return
}

func GetPersonByFullName(lastName string, firstName string, skip int, limit int) (searchResults []Person, searchErr string) {
	searchResults, searchErr = SearchPerson(bson.M{
		"lastName":  bson.RegEx{"^" + lastName, "i"},
		"firstName": bson.RegEx{"^" + firstName, "i"}}, skip, limit)
	return
}

func GetPersonByExactFullName(lastName string, firstName string, skip int, limit int) (searchResults []Person, searchErr string) {
	searchResults, searchErr = SearchPerson(bson.M{"lastName": lastName, "firstName": firstName}, skip, limit)
	return
}

// func main() {

// 	logtofile.WriteMessage("Staring...")
// 	session, err := mgo.Dial("localhost")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("test").C("people")
// 	err = c.Insert(&Person{"Bubba", "+55 53 8116 9639"},
// 		&Person{"Cla", "+55 53 8402 8510"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	result := Person{}
// 	err = c.Find(bson.M{"name": "Bubba"}).One(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Phone:", result.Phone)
// 	logtofile.WriteMessage("Ending...")

// }
