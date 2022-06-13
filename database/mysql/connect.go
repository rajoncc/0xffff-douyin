package mysql

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
    var err error
    db, err = gorm.Open(mysql.Open(DBCONFIG), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
        SkipDefaultTransaction: true,
    })
    if err != nil {
        panic(err)
    }
    //defer db.Close()
}

func Conn() *gorm.DB {
    return db
}
