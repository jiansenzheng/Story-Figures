package webapp

import (
	"iflytek.com/mongotxt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func GetBook(name string) mongotxt.Book {
	session, err := mgo.Dial("127.0.0.1:5565")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	bookC := session.DB("gomboss").C("book")
	result := mongotxt.Book{}
	err = bookC.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}
