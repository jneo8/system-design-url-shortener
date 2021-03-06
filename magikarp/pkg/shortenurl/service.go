package shortenurl

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/magikarp/entity"
	"io"
	"strconv"
	"time"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

type service struct {
	Backend      entity.Backend
	CacheBackend entity.CacheBackend
	Config       Config
	Logger       *log.Logger
}

func (s *service) NewUser(user entity.User) (entity.User, error) {
	return s.Backend.NewUser(user)
}

func (s *service) ShortenURL(originalURL string, expireTime int64, userID *uuid.UUID) (entity.URL, error) {
	if err := validateURL(originalURL); err != nil {
		return entity.URL{}, err
	}
	encodeURL, err := s.encodeURL(originalURL, userID)
	if err != nil {
		return entity.URL{}, err
	}
	url := entity.URL{
		UserID:      userID,
		OriginalURL: originalURL,
		ShortURL:    encodeURL,
		ExpireTime:  expireTime,
	}
	return s.Backend.NewURL(url)
}

func (s *service) DeleteURL(urlID int64) error {
	return nil
}

func (s *service) GetByShortURL(shortURL string) (string, error) {
	// Get url from cache.
	if originalURL, ok := s.CacheBackend.Get(shortURL); ok {
		s.Logger.Debugf("Get url from cache: %s -> %s", shortURL, originalURL)
		return originalURL, nil
	}

	// If url not in cache, get from cache.
	url, err := s.Backend.GetURLByShortURL(shortURL)
	// Set cache.
	if err == nil {
		if err := s.CacheBackend.Set(url); err != nil {
			s.Logger.Warning(err)
		}
	}
	return url.OriginalURL, err
}

// Same user get the same url back.
// If user not login, return the encoded url with uniquenessKey.
// uniquenessKey now generate by random. It could be a increasing sequence number ( save in zookeeper)
func (s *service) encodeURL(originalURL string, userID *uuid.UUID) (string, error) {
	h := md5.New()

	if userID != nil {
		if _, err := io.WriteString(h, userID.String()); err != nil {
			return "", err
		}
	} else {
		// Not login.
		if _, err := io.WriteString(h, gofakeit.UUID()); err != nil {
			return "", err
		}
		if _, err := io.WriteString(h, strconv.FormatInt(time.Now().UnixNano(), 10)); err != nil {
			return "", err
		}
	}
	if _, err := io.WriteString(h, originalURL); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil))[:s.Config.URLLength], nil
}

func (s *service) UserLogin(user entity.User) (entity.User, bool) {
	return s.Backend.UserLogin(user)
}
