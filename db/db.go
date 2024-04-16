package db

import (
	"finance/models"
	"finance/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase(host string) *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(host), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	DB = dbConn
	if err != nil {
		panic(err)
	}

	// Migrate(dbConn)

	utils.Logger.Println("::successfully connected to DB::")
	return dbConn

}

func Migrate(db *gorm.DB) {
	utils.Logger.Println("::starting migration::")
	err := db.AutoMigrate(&models.User{}, &models.Location{}, &models.Org{},
		&models.Register{}, &models.Transaction{}, &models.TransactionCategory{}, &models.Payee{}, &models.Tag{}, &models.TxTag{}, &models.TxLines{})

	if err != nil {
		utils.Logger.Printf("Error %s", err.Error())
	}
	utils.Logger.Println("::migration completed successfully::")
}
