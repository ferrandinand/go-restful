package common

import (
	"os"

	"gopkg.in/mgo.v2"
)

type mongo struct {
	Tasks *mgo.Collection
}

var DB *mongo

func connectDB() {
	database := os.Getenv("DB_NAME")
	hostname := os.Getenv("DB_HOST")

	session, err := mgo.Dial(hostname)
	if err != nil {
		panic(err)
	}

	DB = &mongo{session.DB(database).C("tasks")}

}
