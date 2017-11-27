package models

import "gopkg.in/mgo.v2/bson"
import "../common"

type Task struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"desc"`
}

var Tasks = new(tasks)

type tasks struct{}

func (tasks) FindAll() ([]*Task, error) {
	var ts []*Task
	return ts, common.DB.Tasks.Find(nil).All(&ts)
}

func (tasks) FindOne(id string) (*Task, error) {
	var t *Task
	return t, common.DB.Tasks.Find(bson.M{"_id": bson.IsObjectIdHex(id)}).One(&t)
}

func (tasks) Create(name string, desc string) (*Task, error) {
	t := &Task{
		ID:          bson.NewObjectId(),
		Name:        name,
		Description: desc,
	}
	if err := common.DB.Tasks.Insert(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (tasks) Update(id, name, desc string) error {

	if err := common.DB.Tasks.UpdateId(bson.IsObjectIdHex(id),
		bson.M{"$set": bson.M{
			"name": name,
			"desc": desc,
		}}); err != nil {
		return err
	}
	return nil
}

func (tasks) DeleteById(id string) error {
	if err := common.DB.Tasks.RemoveId(bson.IsObjectIdHex(id)); err != nil {
		return err
	}
	return nil
}
