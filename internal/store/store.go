package store

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type URL struct {
	ShortCode string `gorm:"primaryKey"`
	LongUrl   string
}

type Store struct {
	db *gorm.DB
}

func New() *Store {
	db, err := gorm.Open(sqlite.Open("data/urls.db"), &gorm.Config{})
	if err != nil {
		return nil
	}
	db.AutoMigrate(&URL{})
	return &Store{db: db}
}

func (data *Store) Write(code, longUrl string) {
	urlStruct := URL{code, longUrl}
	data.db.Create(&urlStruct)
}

func (data *Store) Read(code string) string {
	urlStruct := URL{}
	result := data.db.First(&urlStruct, "short_code = ?", code)

	if result.Error != nil {
		return ""
	}

	return urlStruct.LongUrl
}
