package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	"github.com/tally-it/tallyd/config"
	"github.com/tally-it/tallyd/handler"
	"github.com/tally-it/tallyd/log"
	"github.com/tally-it/tallyd/repository/ldap"
	"github.com/tally-it/tallyd/repository/sql"
	"github.com/tally-it/tallyd/router"

	"github.com/sirupsen/logrus"
)

var (
	pkgLogger  = log.New("main")
	isDebug    = flag.Bool("debug", false, "enables debug mode")
	configPath = flag.String("config", os.Getenv("GOPATH")+"/src/github.com/tally-it/tallyd/misc/config/config.toml", "path to config")
)

func main() {
	logger := pkgLogger.ForFunc(context.Background(), "main")
	logger.Info("startup")

	flag.Parse()
	if *isDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	conf, err := config.ReadFile(*configPath)
	if err != nil {
		logger.Fatal("bailing")
	}

	db, err := sql.New(conf.Mysql)
	if err != nil {
		logger.Fatal("bailing")
	}

	l, err := ldap.New(conf.LDAP)
	if err != nil {
		logger.Fatal("bailing")
	}

	err = bootstrap(db, conf.Bootstrap)
	if err != nil {
		logger.Fatal("bailing")
	}

	authorizer := &handler.JWTAuthorizer{
		Secret: conf.JWT.Secret,
	}

	r := router.NewRouter(handler.New(db, l, authorizer))

	logger.Info("running...")

	logger.WithError(http.ListenAndServe(":8080", r)).Error("bailing")
}

func bootstrap(db *sql.Mysql, conf *config.Bootstrap) error {
	logger := pkgLogger.ForFunc(context.Background(), "bootstrap")
	logger.Debug("enter bootstrap")

	users, err := db.GetUserCount(context.Background())
	if err != nil {
		return err
	}

	// if there are users we don't need to bootstrap
	if users != 0 {
		logger.Info("no bootstrapping needed")
		return nil
	}

	logger.Info("user database empty, adding bootstrap user")

	_, err = db.AddLocalUser(context.Background(), conf.User, "", conf.Password, true)
	return err
}
