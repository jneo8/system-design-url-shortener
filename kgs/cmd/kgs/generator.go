package main

import (
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	generatorCMD.Flags().Int("n", 100000, "Number of unique key to be generated.")
	generatorCMD.Flags().Int("length", 6, "Length of unique key.")
}

var base62Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type generatorCMDOpts struct {
	dig.In
	N      int `name:"n"`
	Length int `name:"length"`
}

var generatorCMD = &cobra.Command{
	Use:   "gen",
	Short: "Generator unique key for database.",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{}
		return mermaid.Run(cmd, generatorRunable, initializers...)
	},
}

func generatorRunable(logger *log.Logger, opts generatorCMDOpts) error {
	logger.Info(opts)
	for n := 0; n <= opts.N; n++ {
		key := randStringRunes(opts.Length)
		logger.Info(key)
	}
	return nil
}

func randStringRunes(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = base62Letters[rand.Intn(len(base62Letters))]
	}
	return string(b)
}
