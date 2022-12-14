package entity

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&RecordTimeOut{},
		&Case{},
		&Car{})

	db = database

	//Car
	car_1 := Car{
		Car_ID:  "C0001",
		Name:    "ambulance1",
		Company: "ยานยนต์จำกัด มหาชน",
		TypeCar: "C01",
	}
	db.Model(&Car{}).Create(&car_1)

	car_2 := Car{
		Car_ID:  "C0002",
		Name:    "ambulance2",
		Company: "ฮอนด้าจำกัด มหาชน",
		TypeCar: "C02",
	}
	db.Model(&Car{}).Create(&car_2)

	//Case
	case_1 := Case{
		Case_ID:       "C00001",
		Case_Name:     "รถยนต์ชนรถมอไซค์",
		Type_Case:     "ด่วน",
		Location:      "มทส ต4",
		Whistleblower: "Bee",
	}
	db.Model(&Case{}).Create(&case_1)

	case_2 := Case{
		Case_ID:       "C00002",
		Case_Name:     "รถล้มลงคลอง",
		Type_Case:     "ด่วน",
		Location:      "ข้าง ๆ สระสามแสน",
		Whistleblower: "Bam",
	}
	db.Model(&Case{}).Create(&case_2)

	//RecordTimeOut
	recordtimeout_1 := RecordTimeOut{

		RecordTimeOutID: "R0001",
		ODO_Meter:       122,
		Case:            case_1,
		Car:             car_1,
		TimeOUT:         time.Now(),
	}
	db.Model(&RecordTimeOut{}).Create(&recordtimeout_1)

	recordtimeout_2 := RecordTimeOut{

		RecordTimeOutID: "R0001",
		ODO_Meter:       122,
		Case:            case_1,
		Car:             car_1,
		TimeOUT:         time.Now(),
	}
	db.Model(&RecordTimeOut{}).Create(&recordtimeout_2)

}
