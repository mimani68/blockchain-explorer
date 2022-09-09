package db

import (
	"os"

	"app.io/config"
	"app.io/internal/data/domain"
	"app.io/pkg/logHandler"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type Database struct {
	Db         *gorm.DB
	Connection string
}

func NewDatabase(cfg config.Config) Database {
	// https://gorm.io/docs/connecting_to_the_database.html
	logHandler.Log(logHandler.INFO, `Connecting to "`+cfg.Database.Type+`"`, "DB")
	db, err := gorm.Open(clickhouse.Open(config.GetDbDSN(cfg)), &gorm.Config{})
	if err != nil {
		logHandler.Log(logHandler.ERROR, `Connecting error `+err.Error(), "DB")
		os.Exit(0)
	}
	return Database{
		Db:         db,
		Connection: config.GetDbDSN(cfg),
	}
}

func (d *Database) AutoMigrate() {
	err := d.Db.AutoMigrate(
		&domain.Block{},
		&domain.Transaction{},
	)
	if err != nil {
		logHandler.Log(logHandler.ERROR, err.Error())
	} else {
		logHandler.Log(logHandler.INFO, "Migration model was successful", "DB")
	}
}

func (d *Database) PreProcess() {
	SCHEMA_NAME := "app"
	sql := `CREATE SCHEMA IF NOT EXISTS ` + SCHEMA_NAME
	d.Db.Exec(sql)
	logHandler.Log(logHandler.INFO, "Preprocess was successful", "DB")
}

func (d *Database) CheckAccessOfUser() bool {
	tx := d.Db.Exec(`
	SELECT * FROM information_schema.table_privileges where grantor = 'postgres' and table_catalog = 'indexer' and table_schema = 'app' and privilege_type = any(array['DELETE', 'TRUNCATE', 'TRIGGER', 'REFERENCES']) ;
	`)
	if tx.Error != nil {
		logHandler.Log(logHandler.ERROR, tx.Error.Error())
	}
	if tx.RowsAffected > 0 {
		logHandler.Log(logHandler.DEBUG, "Database User has unlimited permission.", "DB")
	}
	return true
}
