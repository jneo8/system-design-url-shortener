package main

import (
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/system-design-url-shortener/keymaker/api"
	"github.com/system-design-url-shortener/keymaker/entity"
	"github.com/system-design-url-shortener/keymaker/pkg/repository/mongo"
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
	if err := repo.Init(); err != nil {
		return err
	}

	if err := apiService.Run(); err != nil {
		return err
	}
	return nil
}
