package main

import (
	"jneo8/system-design-url-shortener/keymaker/api"
	"jneo8/system-design-url-shortener/keymaker/entity"
	"jneo8/system-design-url-shortener/keymaker/pkg/repository/mongo"

	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	addMongoFlags(serverCMD)
}

var serverCMD = &cobra.Command{
	Use:   "server",
	Short: "Run keymaker server",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{
			api.New,
			api.Register,
			mongo.New,
		}
		return mermaid.Run(cmd, serverRunable, initializers...)
	},
}

func serverRunable(logger *log.Logger, apiService entity.APIService, repo entity.KeyRepository) error {
	defer repo.Close()
	if err := apiService.Run(); err != nil {
		return err
	}
	return nil
}
