package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Goods_20190731_115144 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Goods_20190731_115144{}
	m.Created = "20190731_115144"

	migration.Register("Goods_20190731_115144", m)
}

// Run the migrations
func (m *Goods_20190731_115144) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE goods(`id` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL,`image` varchar(128) NOT NULL)")
}

// Reverse the migrations
func (m *Goods_20190731_115144) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `goods`")
}
