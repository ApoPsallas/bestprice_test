package sql

import (
	"bestprice_test/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(c *config.MySQLConfig) (*sql.DB, error) {
	var connectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB)
	return sql.Open("mysql", connectionString)
}
