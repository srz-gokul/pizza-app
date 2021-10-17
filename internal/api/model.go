package api

import (
	"pizza-app/internal/data"
	"pizza-app/internal/sms"
)

// App is the top level struct
type App struct {
	Repo data.Repo
	Msg  sms.Messenger
}
