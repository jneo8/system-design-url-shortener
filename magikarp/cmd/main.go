package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/system-design-url-shortener/magikarp/api"
	"github.com/system-design-url-shortener/magikarp/entity"
	"github.com/system-design-url-shortener/magikarp/pkg/repository/cache"
	"github.com/system-design-url-shortener/magikarp/pkg/repository/postgres"
	"github.com/system-design-url-shortener/magikarp/pkg/shortenurl"
	"os"
)

func init() {
	cmd.Flags().String("log_level", "DEBUG", "Logger Level")
	cmd.Flags().String("api_dev_key", "abcdefg", "API dev key")
	// Postgres
	cmd.Flags().String("postgres_dsn", "postgres://postgres:postgres_pwd@localhost:5432/magikarp", "data source name, refer https://github.com/jackc/pgx")
	cmd.Flags().Bool("postgres_prefer_simple_protocol", true, "disables implicit prepared statement usage. By default pgx automatically uses the extended protocol")
	// shortenURL
	cmd.Flags().Int("shortenurl_url_length", 6, "Encoded URL length.")
	// API
	cmd.Flags().Int("session_size", 6, "maximum number of idle connections")
	cmd.Flags().String("session_password", "", "redis-password")
	cmd.Flags().String("session_network", "tcp", "tcp or udp")
	cmd.Flags().String("session_address", "localhost:6379", "host:port")
	cmd.Flags().StringSlice("session_key_pairs", []string{"secret"}, "The first key in a pair is used for authentication and the second for encryption. The encryption key can be set to nil or omitted in the last pair, but the authentication key is required in all pairs.")
}

var cmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	Long:  "run " + appName,
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{
			shortenurl.New,
			postgres.New,
			cache.New,
			api.NewRedisSessionStore,
			api.RegisterAPI,
		}
		return mermaid.Run(cmd, runable, initializers...)
	},
}

func runable(cfg *viper.Viper, logger *log.Logger, shortenURLService entity.ShortenURLService, engine *gin.Engine) error {
	if err := engine.Run(); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func main() {
	fmt.Println(logo)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Date: %s\n", date)
	fmt.Printf("Commit: %s\n", commit)
	if err := cmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
