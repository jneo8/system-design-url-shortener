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

type repo struct {
	Config Config
	DB     *gorm.DB
	Logger *log.Logger
}

func (r *repo) InitSchema() error {
	if err := r.DB.AutoMigrate(&URL{}); err != nil {
		return err
	}
	return nil
}

func (r *repo) NewURL(url entity.URL) (entity.URL, error) {
	newURL := URL{
		UserID:      url.UserID,
		OriginalURL: url.OriginalURL,
		ShortURL:    url.ShortURL,
		ExpireTime:  url.ExpireTime,
	}
	r.Logger.Infof("%#v", newURL)
	if result := r.DB.Create(&newURL); result.Error != nil {
		return entity.URL{}, result.Error
	}
	r.Logger.Infof("%#v", newURL)

	return entity.URL{
		UserID:      newURL.UserID,
		URLID:       newURL.Model.ID,
		OriginalURL: newURL.OriginalURL,
		ShortURL:    newURL.ShortURL,
		ExpireTime:  newURL.ExpireTime,
		CreateTime:  newURL.Model.CreatedAt.UnixNano(),
	}, nil
}

func (r *repo) DeleteURL(urlID int64) error {
	return nil
}

func (r *repo) GetURLByShortURL(shortURL string) (entity.URL, error) {
	return entity.URL{}, nil
}
