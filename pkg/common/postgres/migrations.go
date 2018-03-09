package postgres

import (
	_ "github.com/lib/pq" // postgres driver
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file" // get db migration from path

	"database/sql"

	"github.com/build-tanker/passport/pkg/common/appcontext"
	"github.com/build-tanker/passport/pkg/translate"
)

const migrationsPath = "file://./pkg/postgres/migrations"

// RunDatabaseMigrations - run the next migration, needs to be run multiple times if there are multiple
func RunDatabaseMigrations(ctx *appcontext.AppContext) error {
	db, err := sql.Open("postgres", ctx.GetConfig().Database().ConnectionURL())

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		ctx.GetLogger().Infoln(translate.T("postgres:migration:up:fail"))
		return nil
	}

	if err != nil {
		return err
	}

	ctx.GetLogger().Infoln(translate.T("postgres:migration:up:success"))
	return nil
}

// RollbackDatabaseMigration - rollback the database migration
func RollbackDatabaseMigration(ctx *appcontext.AppContext) error {
	m, err := migrate.New(migrationsPath, ctx.GetConfig().Database().ConnectionURL())
	if err != nil {
		return err
	}

	if err := m.Steps(-1); err != nil {
		ctx.GetLogger().Infoln(translate.T("postgres:migration:down:fail"))
		return nil
	}

	ctx.GetLogger().Infoln(translate.T("postgres:migration:down:success"))
	return nil
}
