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
	"math"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	generatorCMD.Flags().Int("batch_size", 20000, "Batch insert size.")
	generatorCMD.Flags().Int("length", 6, "Length of unique key.")
	generatorCMD.Flags().Int("worker_n", 4, "Number of generator worker.")
	addMongoFlags(generatorCMD)
}

var baseLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

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
	Length    int `name:"length"`
	BatchSize int `name:"batch_size"`
	WorkerN   int `name:"worker_n"`
}

func generatorRunable(logger *log.Logger, opts generatorCMDOpts, repo entity.KeyRepository) error {
	// Close repo conn.
	defer repo.Close()
	if err := repo.Init(); err != nil {
		return err
	}

	// Get worker working range.
	workerCoverRanges, err := getWorkerCoverRanges(opts.WorkerN)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	pb, pbWg := progressbar.NewWithWaitGroup()
	itemCh := make(chan string, opts.WorkerN)

	wg.Add(1)
	go func() {
		defer wg.Done()
		keys := []string{}
		for k := range itemCh {
			keys = append(keys, k)
			if len(keys) >= opts.BatchSize {
				_, err := repo.KeyBatchInsert(keys)
				if err != nil {
					logger.Error(err)
				}
				keys = keys[:0]
			}
		}
		if len(keys) > 0 {
			_, err := repo.KeyBatchInsert(keys)
			if err != nil {
				logger.Error(err)
			}
			keys = keys[:0]
		}
	}()

	// Generator
	baseN := int(math.Pow(float64(len(baseLetters)), float64(opts.Length-1)))
	for workerID, coverRange := range workerCoverRanges {
		pbWg.Add(1)
		wg.Add(1)
		go func(workerID int, coverRange []int) {
			defer pbWg.Done()
			defer wg.Done()
			bar := progressbar.AddBar(pb, (coverRange[1]+1)*baseN-coverRange[0]*baseN, fmt.Sprintf("Worker_%d", workerID))
			for i := coverRange[0] * baseN; i < (coverRange[1]+1)*baseN; i++ {
				start := time.Now()
				n := i
				s := ""
				for n > 0 {
					s = s + string(baseLetters[n%len(baseLetters)])
					n = n / len(baseLetters)
				}
				itemCh <- s
				bar.Increment()
				bar.DecoratorEwmaUpdate(time.Since(start))
			}
		}(workerID, coverRange)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		pbWg.Wait()
		close(itemCh)
	}()
	wg.Wait()
	return nil
}

func getWorkerCoverRanges(workerN int) ([][]int, error) {
	if workerN > len(baseLetters) {
		return nil, fmt.Errorf("Worker number couldn't > %d", len(baseLetters))
	}

	workerCoverRanges := [][]int{}
	quotient := len(baseLetters) / workerN
	for i := 0; i < workerN; i++ {
		workerCoverRanges = append(
			workerCoverRanges,
			[]int{i * quotient, (i+1)*quotient - 1},
		)
	}
	if len(baseLetters)%workerN != 0 {
		workerCoverRanges = append(
			workerCoverRanges,
			[]int{len(baseLetters) - len(baseLetters)%workerN, len(baseLetters) - 1},
		)
	}
	return workerCoverRanges, nil
}

func randStringRunes(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = baseLetters[rand.Intn(len(baseLetters))]
	}
	return string(b)
}
