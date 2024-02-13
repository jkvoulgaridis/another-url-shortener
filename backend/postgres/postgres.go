package postgres


import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UrlShortModel struct {
	gorm.Model
	OriginalUrl string
	ShortUrl string
}

func GetDbConnection() (*gorm.DB, error){
	creds := "host=postgres user=admin password=admin dbname=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(creds), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&UrlShortModel{})
}


func CreateSamplesInDB(db *gorm.DB){
	db.Create(&UrlShortModel{OriginalUrl: "http://original-1.com", ShortUrl: "short1"})
	db.Create(&UrlShortModel{OriginalUrl: "http://original-2.com", ShortUrl: "short2"})
}

