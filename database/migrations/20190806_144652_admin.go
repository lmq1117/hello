package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Admin_20190806_144652 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Admin_20190806_144652{}
	m.Created = "20190806_144652"

	migration.Register("Admin_20190806_144652", m)
}

// Run the migrations
func (m *Admin_20190806_144652) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	//m.SQL("CREATE TABLE IF NOT EXISTS admin (" +
	//	"`id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT," +
	//	"`name` varchar(255) NOT NULL DEFAULT '' ," +
	//	"`pwd` varchar(128) NOT NULL DEFAULT '' " +
	//	")")
}

// Reverse the migrations
func (m *Admin_20190806_144652) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `admin`")
}
