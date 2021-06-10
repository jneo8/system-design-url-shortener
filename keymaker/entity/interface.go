package entity

// APIService ...
type APIService interface {
	Run() error
}

// KeyRepository ...
type KeyRepository interface {
	KeyBatchInsert(key []string) (int, error) // Batch insert key.
	KeyBatchUpsert(key []string) (int, error) // Batch insert key.
	Close() error
	GetKey() (string, error)
	Ping() error
	CreateIndexes() error
	// GetKeys(expire int64) ([]string, error)
}
