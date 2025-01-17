package db

import (
	"fmt"
	"path/filepath"

	"github.com/ArtalkJS/Artalk/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func OpenSQLite(filename string, gormConfig *gorm.Config) (*gorm.DB, error) {
	if filename == "" {
		return nil, fmt.Errorf("please set `db.file` option in config file to specify a sqlite database path")
	}
	if err := utils.EnsureDir(filepath.Dir(filename)); err != nil {
		return nil, err
	}
	return gorm.Open(sqlite.Open(filename), gormConfig)
}

func OpenMySql(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), gormConfig)
}

func OpenPostgreSQL(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), gormConfig)
}

func OpenSqlServer(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(dsn), gormConfig)
}
