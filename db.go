package outbarriers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func (ctx *Context) StartDB() error {

	var err error
	ctx.DB, err = gorm.Open("mysql", DBSTRING)
	if err != nil {
		log.Printf("ERROR opening DB\n")
		return err
	}
	log.Printf("\tDB opened!")
	ctx.DB.SingularTable(true)
	//ctx.DB.LogMode(true)
	//ctx.DB.AutoMigrate(ipubs.GoTest{})
	return nil
}
