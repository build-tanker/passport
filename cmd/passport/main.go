package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/build-tanker/passport/pkg/common/appcontext"
	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/postgres"
	"github.com/build-tanker/passport/pkg/common/server"
	"github.com/build-tanker/passport/pkg/translate"
)

func main() {
	config := config.NewConfig([]string{".", "..", "../.."})
	ctx := appcontext.NewAppContext(config)
	db := postgres.NewPostgres(config.Database().ConnectionURL(), config.Database().MaxPoolSize())
	server := server.NewServer(ctx, db)

	log.Println(translate.T("passport:app:start"))

	app := cli.NewApp()
	app.Name = "passport"
	app.Version = "0.0.1"
	app.Usage = translate.T("passport:app:usage")

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: translate.T("passport:cli:start"),
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:  "migrate",
			Usage: translate.T("passport:cli:migrate"),
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations(ctx)
			},
		},
		{
			Name:  "rollback",
			Usage: translate.T("passport:cli:rollback"),
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration(ctx)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
