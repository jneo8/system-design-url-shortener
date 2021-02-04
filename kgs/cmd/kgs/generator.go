package main

import (
	"fmt"
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/system-design-url-shortener/kgs/entity"
	"github.com/system-design-url-shortener/kgs/pkg/progressbar"
	"github.com/system-design-url-shortener/kgs/pkg/repository/mongo"
	"go.uber.org/dig"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	generatorCMD.Flags().Int("n", 100000, "Number of unique key to be generated.")
	generatorCMD.Flags().Int("batch_size", 10000, "Batch insert size.")
	generatorCMD.Flags().Int("length", 6, "Length of unique key.")
	generatorCMD.Flags().Int("worker_n", 5, "Number of generator worker.")
	addMongoFlags(generatorCMD)
}

var base62Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var generatorCMD = &cobra.Command{
	Use:   "gen",
	Short: "Generator unique key for database.",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{
			mongo.New,
		}
		return mermaid.Run(cmd, generatorRunable, initializers...)
	},
}

type generatorCMDOpts struct {
	dig.In
	N         int `name:"n"`
	Length    int `name:"length"`
	BatchSize int `name:"batch_size"`
	WorkerN   int `name:"worker_n"`
}

func generatorRunable(logger *log.Logger, opts generatorCMDOpts, repo entity.KeyRepository) error {
	defer repo.Close()
	if err := repo.Init(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	pb, pbWg := progressbar.NewWithWaitGroup()
	itemCh := make(chan string, opts.WorkerN)
	total := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		keys := []string{}
		for k := range itemCh {
			keys = append(keys, k)
			if len(keys) >= opts.BatchSize || total+len(keys) == opts.N {
				numInsert, err := repo.KeyBatchInsert(keys)
				if err != nil {
					logger.Error(err)
				}
				logger.Info(total)
				total += numInsert
				keys = keys[:0]
			}
		}
	}()

	for w := 0; w <= opts.WorkerN; w++ {
		pbWg.Add(1)
		wg.Add(1)
		go func(workerID int) {
			defer pbWg.Done()
			defer wg.Done()
			bar := progressbar.AddBar(pb, opts.N/opts.WorkerN, fmt.Sprintf("Worker_%d", workerID))
			for total <= opts.N {
				start := time.Now()
				itemCh <- randStringRunes(opts.Length)
				bar.Increment()
				bar.DecoratorEwmaUpdate(time.Since(start))
			}
		}(w)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		pbWg.Wait()
		close(itemCh)
	}()
	wg.Wait()

	logger.Infof("Total insert %d", total)
	return nil
}

func randStringRunes(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = base62Letters[rand.Intn(len(base62Letters))]
	}
	return string(b)
}
