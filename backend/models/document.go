package models

import "gopkg.in/mgo.v2/bson"

import "time"

type Document struct {
	ID               bson.ObjectId  `bson:"_id" json:"id"`
	CreatedBy        Employee       `bson:"employee" json:"employee"`
	CreatedAt        time.Time      `bson:"created_at" json:"created_at"`
	Associations     []PersonDetail `bson:"associations" json:"associations"`
	FinalDestination string         `bson:"final_destination" json:"final_destination"`
	Priority         string         `bson:"priority" json:"priority"`
	NextEmployees    []Employee     `bson:"next_employees" json:"next_employees"`
	Remarks          string         `bson:"remarks" json:"remarks"`
	Completed        bool           `bson:"completed" json:"completed"`
	Error            bool           `bson:"error" json:"error"`
}

type PersonDetail struct {
	Person           Employee  `bson:"person" json:"person"`
	SignaturePending bool      `bson:"signature_pending" json:"signature_pending"`
	SignedAt         time.Time `bson:"signed_at" json:"signed_at"`
	Remarks          string    `bson:"remarks" json:"remarks"`
	Priority         int       `bson:"priority" json:"priority"`
}
