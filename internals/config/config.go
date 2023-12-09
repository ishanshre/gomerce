// This package consists of global configurations

package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig consist of global configs.
// This package should not import from other internals package
// to avoid import cycle loop
type AppConfig struct {
	InProduction  bool
	UseRedis      bool
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
	ContactEmail  string
	Addr          string
	Port          int
	DbString      string
	Dsn           string
	RedisHost     string
}
