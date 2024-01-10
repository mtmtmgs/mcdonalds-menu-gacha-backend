package env

import (
	"database/sql"
	"fmt"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

const (
	Dev  = "development"
	Prod = "production"
)

type Env struct {
	DBConnect   *sql.DB
	QueryHook   *bundebug.QueryHook
	LoggerLevel log.Lvl
}

func NewAppEnv() Env {
	_ = godotenv.Load()
	return getEnv()
}

func NewCmdEnv(path string) Env {
	_ = godotenv.Load(path)
	return getEnv()
}

func getEnv() Env {
	env := os.Getenv("ENV")

	var e Env
	switch env {
	case Dev:
		e.DBConnect = getDBConnect(env)
		e.QueryHook = getQueryHook(env)
		e.LoggerLevel = getLoggerLevel(env)
	case Prod:
		e.DBConnect = getDBConnect(env)
		e.QueryHook = getQueryHook(env)
		e.LoggerLevel = getLoggerLevel(env)
	default:
		e.DBConnect = getDBConnect(Dev)
		e.QueryHook = getQueryHook(Dev)
		e.LoggerLevel = getLoggerLevel(Dev)
	}
	return e
}

/*
DBコネクト（PostgreSQL）
開発: Dockerコンテナ
本番: Cloud SQL
*/
func getDBConnect(env string) *sql.DB {
	switch env {
	case Dev:
		sqldb := sql.OpenDB(pgdriver.NewConnector((pgdriver.WithDSN(
			fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
				os.Getenv("POSTGRES_USER"),
				os.Getenv("POSTGRES_PASSWORD"),
				os.Getenv("POSTGRES_HOST"),
				os.Getenv("POSTGRES_PORT"),
				os.Getenv("POSTGRES_DB")),
		))))
		return sqldb
	case Prod:
		keyJson := []byte(os.Getenv("CLOUD_SQL_USER_KEY_JSON"))
		cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithCredentialsJSON(keyJson))
		if err != nil {
			log.Fatal(err)
		}
		sqldb, err := sql.Open("cloudsql-postgres",
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("CLOUD_SQL_INSTANCE_CONNECTION_NAME"),
				os.Getenv("CLOUD_SQL_USER"),
				os.Getenv("CLOUD_SQL_PASSWORD"),
				os.Getenv("CLOUD_SQL_DB"),
			))
		if err != nil {
			_ = cleanup()
			log.Fatal(err)
		}
		_ = cleanup()
		return sqldb
	default:
		return nil
	}
}

/*
DBクエリ実行前後の処理
*/
func getQueryHook(env string) *bundebug.QueryHook {
	switch env {
	case Dev:
		return bundebug.NewQueryHook(
			bundebug.WithEnabled(true),
			bundebug.WithVerbose(true),
			// bundebug.FromEnv("BUNDEBUG"),
		)
	case Prod:
		return bundebug.NewQueryHook(
			bundebug.WithEnabled(true),
			// bundebug.FromEnv("BUNDEBUG"),
		)
	default:
		return nil
	}
}

/*
アプリのログレベル
*/
func getLoggerLevel(env string) log.Lvl {
	switch env {
	case Dev:
		return log.DEBUG
	case Prod:
		return log.INFO
	default:
		return 0
	}
}
