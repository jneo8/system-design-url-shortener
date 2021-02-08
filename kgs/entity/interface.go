package entity

// APIService ...
type APIService interface {
	Run() error
}

// KeyRepository ...
type KeyRepository interface {
	Init() error
	KeyBatchInsert(key []string) (int, error) // Batch insert key.
	KeyBatchUpsert(key []string) (int, error) // Batch insert key.
	Close() error
	GetKey(expire int64) (string, error)
	// GetKeys() ([]string, error)
}
