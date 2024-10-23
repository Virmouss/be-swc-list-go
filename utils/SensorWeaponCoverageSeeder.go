package utils

import (
	"be-swc-list/app/model"
	"log"
	"time"

	"gorm.io/gorm"
)

// type SensorWeaponCoverageSeeder struct {
// 	db *gorm.DB
// }

// func NewSensorWeaponCoverageSeeder() *SensorWeaponCoverageSeeder {

// 	db, err := db.ConnectDB()
// 	if err != nil {
// 		log.Fatalf(db, "Failed to connect to the database: %v", err)
// 	}

// 	return &SensorWeaponCoverageSeeder{db: db}
// }

func Seed(db *gorm.DB) {
	var count int64

	db.AutoMigrate(&model.SensorWeaponCoverage{})
	db.Model(&model.SensorWeaponCoverage{}).Count(&count)

	if count == 0 {
		CreateSWCModel(db, "Gun STBD", "Min. Range", "Weapon", "NM", 0, "non-air")
		CreateSWCModel(db, "Gun STBD", "Max. Range", "Weapon", "NM", 1.6198, "non-air")
		CreateSWCModel(db, "Gun STBD", "Start Bearing", "Weapon", "Deg", 10.0, "non-air")
		CreateSWCModel(db, "Gun STBD", "End Bearing", "Weapon", "Deg", 170.0, "non-air")

		CreateSWCModel(db, "Gun PORT", "Min. Range", "Weapon", "NM", 0, "non-air")
		CreateSWCModel(db, "Gun PORT", "Max. Range", "Weapon", "NM", 1.6198, "non-air")
		CreateSWCModel(db, "Gun PORT", "Start Bearing", "Weapon", "Deg", 190.0, "non-air")
		CreateSWCModel(db, "Gun PORT", "End Bearing", "Weapon", "Deg", 350.0, "non-air")

		CreateSWCModel(db, "Sea Eagle Optic", "Min. Range", "Sensor", "NM", 0, "non-air")
		CreateSWCModel(db, "Sea Eagle Optic", "Max. Range", "Sensor", "NM", 6.6069, "non-air")
		CreateSWCModel(db, "Sea Eagle Optic", "Start Bearing", "Sensor", "Deg", 10.0, "non-air")
		CreateSWCModel(db, "Sea Eagle Optic", "End Bearing", "Sensor", "Deg", 350.0, "non-air")

		CreateSWCModel(db, "Sea Eagle Laser", "Min. Range", "Sensor", "NM", 0, "non-air")
		CreateSWCModel(db, "Sea Eagle Laser", "Max. Range", "Sensor", "NM", 6.6069, "non-air")
		CreateSWCModel(db, "Sea Eagle Laser", "Start Bearing", "Sensor", "Deg", 10.0, "non-air")
		CreateSWCModel(db, "Sea Eagle Laser", "End Bearing", "Sensor", "Deg", 350.0, "non-air")
	} else {
		log.Println("Users table already has data, skipping seeding")
	}

}

func CreateSWCModel(db *gorm.DB, group string, item string, types string, unit string, value float64, environment string) {
	data := []model.SensorWeaponCoverage{
		{CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), DefaultValue: value, GroupSWC: group, IsDeleted: false, Item: item, Type: types, Unit: unit, Value: value, Environment: environment},
	}

	if err := db.Create(&data).Error; err != nil {
		log.Fatalf("Could not seed SWC Data: %v", err)
	}
}
