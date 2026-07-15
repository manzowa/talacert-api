package config

import (
	"fmt"
	"talacert-api/internal/logger"
	"time"

	gormlogger "gorm.io/gorm/logger" // aliase to avoid conflict with our logger

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(
	cfg *Config,
) (*gorm.DB, error) {

	if err := CreateDatabaseIfNotExists(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})
	if err != nil {
		logger.ErrorLogger.Error("failed to connect database", "error", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 🔥 Pool settings (important en prod)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func CreateDatabaseIfNotExists(
	host,
	port,
	user,
	password,
	dbName string,
) error {

	// Connexion au serveur MySQL sans sélectionner de base
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci",
		dbName,
	)

	return db.Exec(query).Error
}
