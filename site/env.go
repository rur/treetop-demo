package site

import (
	"log"
)

// Env is used to pass site-wide configuration and resources to handlers
type Env struct {
	// settings
	HTTPS bool

	// loggers
	ErrorLog *log.Logger
	WarnLog  *log.Logger
	InfoLog  *log.Logger
}
