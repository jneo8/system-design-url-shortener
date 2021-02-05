package entity

// KeyRepository ...
type KeyRepository interface {
	Init() error
	KeyBatchInsert(key []string) (int, error) // Batch insert key.
	KeyBatchUpsert(key []string) (int, error) // Batch insert key.
	Close() error
}
