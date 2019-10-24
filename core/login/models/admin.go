package models

import (
	"common/gorm"
	"log"
)

type Admin struct {
	Id   uint   `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"column:name"`
}

func SaveInfo() {
	if err := common.Gorm.Create(&Admin{Name: "xixi"}).Error; err != nil {
		log.Fatal(err)
	}

}
