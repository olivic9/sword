package databases

import (
	"context"
	"database/sql"
	"fmt"
	"sword-project/pkg/configs"
	"sword-project/pkg/logging"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func GetMysqlDatabase(ctx context.Context) *sql.DB {
	if db == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			configs.MysqlCfg.UserName,
			configs.MysqlCfg.Password,
			configs.MysqlCfg.Url,
			configs.MysqlCfg.Port,
			configs.MysqlCfg.Database))

		if err != nil {
			logging.Logger.Fatal(ctx, err, logging.Metadata{})
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

	}
	return db
}
