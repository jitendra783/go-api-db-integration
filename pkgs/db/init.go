package db

import (
	"context"
	"fmt"
	"go-api/pkgs/config"
	"log"
	"time"

	elog "go-api/pkgs/logger"

	oracle "github.com/godoes/gorm-oracle"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbLogger struct{}

// utils function to connect oracle db
func ConnectOracleDb(schemaConfig string) *gorm.DB {
	schemaConfig = "db.oracle"
	c := config.GetConfig()
	dsn := oracle.BuildUrl(c.GetString(schemaConfig+".host"),
		c.GetInt(schemaConfig+".port"),
		c.GetString(schemaConfig+".db"),
		c.GetString(schemaConfig+".user"),
		c.GetString(schemaConfig+".password"),
		nil)
	log.Println("------------dsn:", dsn)
	db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, /// Removed defuault transaction used by GORM to faster the query and execution
		Logger:                 customLogger(),
	})
	if err != nil {
		elog.Log().Warn("failed to connect oracle connection")
		//elog.Log().Warn("database connection error ", zap.Error(err))
		return nil
	}

	// Set the current schema for the session
	schema := fmt.Sprintf("ALTER SESSION SET CURRENT_SCHEMA = %s", c.GetString(schemaConfig+".schema"))
	err = db.Exec(schema).Error
	if err != nil {
		fmt.Println("Error setting schema:", err)
		elog.Log().Warn("database connection error ")
		return nil
	}
	elog.Log().Info("Oracle Database Connected")
	return db
}
func customLogger() logger.Interface {
	return dbLogger{}
}

func (d dbLogger) Error(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("error", data), zap.Any("description", others))
}

func (d dbLogger) Info(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Warn(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	query, others := fc()
	if err != nil {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others), zap.Error(err))
	} else {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others))
	}
}

func (d dbLogger) LogMode(l logger.LogLevel) logger.Interface {
	return d
}
