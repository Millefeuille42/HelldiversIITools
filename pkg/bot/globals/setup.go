package globals

import (
	"os"
)

func SetGlobals() {
	DatabaseDriver = os.Getenv("HDII__BOT__DB_DRIVER")
	if DatabaseDriver == "" {
		DatabaseDriver = "sqlite3"
	}

	DatabaseDSN = os.Getenv("HDII__BOT__DB_DSN")
	if DatabaseDSN == "" {
		DatabaseDSN = "file:./db.sqlite3?_foreign_keys=ON"
	}

	BotToken = os.Getenv("HDII__BOT__TOKEN")

	ApiUrl = os.Getenv("HDII__BOT__API")
}
