package api

import (
	"context"
	"fmt"
	"go-api/pkgs/config"
	e "go-api/pkgs/error"
	"go-api/pkgs/logger"
	serv "go-api/pkgs/service"
	"log"
	"net/http"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var srv *http.Server
var ctx context.Context
var databases []*gorm.DB

func Start() error {
	ctx = context.Background()
	e.ErrorInit()

	var pgSql *gorm.DB = nil
	databases = make([]*gorm.DB, 0)
	//pgSql = db.ConnectOracleDb("")
	databases = append(databases, pgSql)
	serviceObj := serv.NewServiceGroupObject(pgSql)

	startRouter(serviceObj)
	return nil
}

func startRouter(obj serv.ServiceGroupLayer) {
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetConfig().GetInt("server.port")),
		Handler: RouterInit(obj, logger.Log()), //getRouter set the api specs for version-1 routes
	}
	// run api router
	log.Println("Listening and serving at port", srv.Addr)
	logger.Log().Info("starting router")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log().Fatal("Error starting server", zap.Error(err))
	}
	//logger.Log().Info("server working")
}
