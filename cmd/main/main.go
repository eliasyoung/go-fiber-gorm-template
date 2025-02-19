package main

import (
	"github.com/eliasyoung/fiber-flavor/config"
	"github.com/eliasyoung/fiber-flavor/internal/db"
	"go.uber.org/zap"
)

func main() {

	log := zap.Must(zap.NewProduction()).Sugar()
	defer log.Sync()

	cfg := config.InitConfig()

	conn, err := db.ConnectDB(cfg.DBConfig.DBHost, cfg.DBConfig.DBUser, cfg.DBConfig.DBPassword, cfg.DBConfig.DBName, cfg.DBConfig.DBPort)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.RunMigrations(conn)
	if err != nil {
		log.Fatalln(err)
	}

	store := db.InitStore(conn)

	app := &application{
		config: cfg,
		store:  store,
		logger: log,
	}

	fiberApp := app.mount()

	log.Fatal(app.run(fiberApp))
}
