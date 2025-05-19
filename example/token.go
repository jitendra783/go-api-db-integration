package main

import (
	"sync"

	"gorm.io/gorm"
)

var db1 *gorm.DB
var wg sync.WaitGroup

type DataRow struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Info      string `gorm:"column:info" json:"info"`
	Processed bool   `gorm:"column:processed" json:"processed"`
	Status    string `gorm:"column:status" json:"status"`
}
