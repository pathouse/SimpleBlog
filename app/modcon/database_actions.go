package modcon

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"reflect"
)

func AutoMigrate(db *gorm.DB) {
	for _, model := range []interface{}{User{}, Post{}} {
		if err := db.AutoMigrate(model).Error; err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Auto migrating", reflect.TypeOf(model).Name(), "...")
		}
	}
}

func CreateRecord(db *gorm.DB, i interface{}) error {
	if err := db.Save(i).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRecord(db *gorm.DB, old, new interface{}) error {
	if err := db.Model(old).Updates(new).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRecord(db *gorm.DB, i interface{}) error {
	if err := db.Delete(i).Error; err != nil {
		return err
	}
	return nil
}

// boolean nf is true if we want to report RecordNotFound errors
func FindByMap(db *gorm.DB, m map[string]interface{}, i interface{}, nf bool) error {
	err := db.Where(m).Find(i).Error
	if err != nil && (nf || err != gorm.RecordNotFound) {
		return err
	}
	return nil
}

// boolean nf is true if we want to report RecordNotFound errors
func FindAll(db *gorm.DB, i interface{}, nf bool) error {
	err := db.Find(i).Error
	if err != nil && (nf || err != gorm.RecordNotFound) {
		return err
	}
	return nil
}
