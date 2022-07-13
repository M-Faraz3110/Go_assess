package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// func GetDBCtx(dsn string) *sqlx.DB {
// 	if dsn == "" {
// 		err := fmt.Errorf("missing database dsn variable from enviroment")
// 		panic(err)
// 	} else {
// 		if db, err := sqlx.Connect("postgres", dsn); err != nil {
// 			panic(err)
// 		} else {
// 			db.MustExec(schema)
// 			return db
// 		}
// 	}
// }

func GetDBCtx() *sqlx.DB {
	if db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=Salmon123"); err != nil {
		panic(err)
	} else {
		db.MustExec(schema)
		return db
	}

}
