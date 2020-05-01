package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	//var connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUser, dbPassword, dbName,)
	return sql.Open("mysql", "bestprice:a123456@tcp(127.0.0.1:3306)/mysql")
}
