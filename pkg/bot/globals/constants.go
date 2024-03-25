package globals

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	DatabaseDriver = ""
	DatabaseDSN    = ""

	BotToken = ""

	ApiUrl = ""

	NumberPrinter = message.NewPrinter(language.English)
)
