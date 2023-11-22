package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Post_20231122_124250 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Post_20231122_124250{}
	m.Created = "20231122_124250"

	migration.Register("Post_20231122_124250", m)
}

// Run the migrations
func (m *Post_20231122_124250) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE post(id serial primary key,title TEXT NOT NULL,body TEXT NOT NULL)")
}

// Reverse the migrations
func (m *Post_20231122_124250) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE post")
}
