package database

import (
	"errors"
	"go-app/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnDB(dsn string, dbtype config.DBTYPE) (*gorm.DB, error) {
	var (
		dialector gorm.Dialector
		db        *gorm.DB
		err       error
	)

	dialector = GenDialector(dsn, dbtype)

	db, err = gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.Pool.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Pool.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(config.Pool.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(config.Pool.ConnMaxLifetime)
	return db, nil
}

func GenDialector(dsn string, dbtype config.DBTYPE) (dialector gorm.Dialector) {
	switch dbtype {
	case config.DBTYPE_MYSQL:
		dialector = mysql.Open(dsn)
	case config.DBTYPE_MSSQL:
		dialector = sqlserver.Open(dsn)
	case config.DBTYPE_POSTGRES:
		dialector = postgres.Open(dsn)
	case config.DBTYPE_SQLITE:
		dialector = sqlite.Open(dsn)
	default:
		panic(errors.New("请配置正确的数据库类型"))
	}
	return
}
