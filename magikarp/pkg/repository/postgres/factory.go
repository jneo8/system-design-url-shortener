package postgres

import (
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/magikarp/entity"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config ...
type Config struct {
	dig.In
	DSN                  string `name:"postgres_dsn"`
	PreferSimpleProtocol bool   `name:"postgres_prefer_simple_protocol"`
}

// New is a simple factory func return URLBackend.
func New(logger *log.Logger, config Config) (entity.URLBackend, error) {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  config.DSN,                  // data source name, refer https://github.com/jackc/pgx
				PreferSimpleProtocol: config.PreferSimpleProtocol, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	return &repo{DB: db, Logger: logger, Config: config}, nil
}
