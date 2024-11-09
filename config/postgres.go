package config

import (
	"database/sql"
	"fmt"

	"github.com/oiime/logrusbun"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type PostgresConfig struct {
	Host     string `required:"true"`
	Port     int    `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Database string `required:"true"`
	Debug    bool   `default:"false"`
}

func (c *PostgresConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.Database)
}

func (c *PostgresConfig) NewClient(log logrus.FieldLogger) (bun.IDB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.DSN())))
	client := bun.NewDB(sqldb, pgdialect.New())

	if c.Debug {
		client.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	if log != nil {
		client.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{
			Logger: log,
		}))
	}

	rows, err := client.Query("SELECT 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return client, nil
}
