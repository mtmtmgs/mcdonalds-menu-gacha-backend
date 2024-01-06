package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/uptrace/bun/extra/bundebug"
)

const (
	Dev  = "development"
	Prod = "production"
)

type Env struct {
	Dsn         string
	QueryHook   *bundebug.QueryHook
	LoggerLevel log.Lvl
}

func NewAppEnv() Env {
	godotenv.Load()
	return getEnv()
}

func NewCmdEnv(path string) Env {
	godotenv.Load(path)
	return getEnv()
}

func getEnv() Env {
	env := os.Getenv("ENV")

	var e Env
	switch env {
	case Dev:
		e.Dsn = getDsn(env)
		e.QueryHook = getQueryHook(env)
		e.LoggerLevel = getLoggerLevel(env)
	case Prod:
		e.Dsn = getDsn(env)
		e.QueryHook = getQueryHook(env)
		e.LoggerLevel = getLoggerLevel(env)
	default:
		e.Dsn = getDsn(Dev)
		e.QueryHook = getQueryHook(Dev)
		e.LoggerLevel = getLoggerLevel(Dev)
	}
	return e
}

/*
DBアクセスURL
PostgreSQL
*/
func getDsn(env string) string {
	switch env {
	case Dev:
		return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)
	case Prod:
		// Cloud SQL（パブリックIP）
		return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)
	default:
		return ""
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
