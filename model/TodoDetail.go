package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	TodoDetail struct {
		ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		UserId   string        `json:"UserId" bson:"UserId"`
		TaskName string        `json:"TaskName" bson:"TaskName"`
		TaskDesc string        `json:"TaskDesc" bson:"TaskDesc"`
		Notes    string        `json:"Notes" bson:"Notes"`
		TaskDate time.Time     `json:"TaskDate" bson:"TaskDate"`
		Status   int           `json:"Status" bson:"Status"`
	}
)
