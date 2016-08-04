package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
  Bookful struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    WukongDocId uint64 `json:"wukongdocid" bson::"wukongdocid"` 
    Author []string `json:"author" bson:"author"`
    Image string `json:"image" bson:"image"`
    Title string `json:"title" bson:"title"`
  }
)
