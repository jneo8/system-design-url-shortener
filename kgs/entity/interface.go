package entity

// KeyRepository ...
type KeyRepository interface {
	KeyBatchInsert(key []string) (int, error) // Batch insert key.
}
