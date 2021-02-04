package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCMD.AddCommand(generatorCMD)
	rootCMD.PersistentFlags().String("log_level", "info", "Logger Level")
}

var rootCMD = &cobra.Command{
	Use:   "run",
	Short: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(123)
		return nil
	},
}

func main() {
	if err := rootCMD.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
