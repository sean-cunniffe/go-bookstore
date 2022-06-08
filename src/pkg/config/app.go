package config

import (
	// "os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func Connect() {
	url := getEnvironmentVarOrExit("DB_URL")
	log.Println("Connecting to DB url ", url)
	for {
		d, err := gorm.Open("mysql", url)
		if err != nil {
			log.Println(err)
		} else {
			db = d
			break
		}
		time.Sleep(time.Second)
	}

}

func getEnvironmentVarOrExit(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment Variable %v is required", key)
	}
	return value
}

func GetDb() *gorm.DB {
	return db
}
