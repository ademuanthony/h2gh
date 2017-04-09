package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Init_20170409_055503 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Init_20170409_055503{}
	m.Created = "20170409_055503"
	migration.Register("Init_20170409_055503", m)
}

// Run the migrations
func (m *Init_20170409_055503) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Init_20170409_055503) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
