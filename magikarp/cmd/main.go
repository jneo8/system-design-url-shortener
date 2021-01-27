package main

import (
	"fmt"
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var cmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{}
		return mermaid.Run(cmd, runable, initializers...)
	},
}

func runable(logger *log.Logger) error {
	logger.Info(123)
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
