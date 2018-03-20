package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/postgres"
	"github.com/build-tanker/passport/pkg/common/server"
)

func main() {
	conf := config.New([]string{".", "..", "../.."})
	db := postgres.New(conf.ConnectionURL(), conf.MaxPoolSize())
	server := server.New(conf, db)

	log.Println("Starting passport")

	app := cli.NewApp()
	app.Name = "passport"
	app.Version = "0.0.1"
	app.Usage = "this service saves files and makes them available for distribution"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations(conf)
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration(conf)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
