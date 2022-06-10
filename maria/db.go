package maria

import (
	"fmt"

	"github.com/wnjoon/go-orm-db/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type statuses struct {
	ID   uint   `sql:"AUTO_INCREMENT"`
	TxId string `sql:"type:varchar(100)"`
	Code string `sql:"type:varchar(100)"`
}

func Connect(id string, pwd string, host string, port string, dbName string) *gorm.DB {

	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", id, pwd, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	util.HandleErr(err, util.ErrOpenDB)

	sqlDB, err := db.DB()
	util.HandleErr(err, util.ErrSqlDB)
	setConnectionPool(sqlDB)

	return db
}

func Insert(txId string, code string, db *gorm.DB) {
	status := statuses{
		TxId: txId,
		Code: code,
	}

	result := db.Create(&status)
	fmt.Println("status.ID : ", status.ID, "\n",
		"status.TxHash : ", status.TxId, "\n",
		"status.StatusCode : ", status.Code, "\n",
		"result.RowsAffected : ", result.RowsAffected, "\n",
		"result.Error : ", result.Error)
}

func Query(txId string, db *gorm.DB) {
	var status statuses
	db.Where("tx_id = ?", txId).First(&status)
	fmt.Println(status.Code)

}
