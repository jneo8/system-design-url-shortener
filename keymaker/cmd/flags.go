package main

import (
	"github.com/spf13/cobra"
)

func addMongoFlags(cmd *cobra.Command) {
	cmd.Flags().String("mongo_db", "keymaker", "mongodb database name")
	cmd.Flags().String("mongo_dsn", "mongodb://mongo:mongo123@localhost:27017", "mongodb dsn string")
	cmd.Flags().String("mongo_key_collection", "shortURL", "mongodb short url collection name")
}
