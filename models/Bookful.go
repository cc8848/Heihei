package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
  Bookful struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Author []string `json:"author" bson:"author"`
    Image string `json:"image" bson:"image"`
    Title string `json:"title" bson:"title"`
  }
)
