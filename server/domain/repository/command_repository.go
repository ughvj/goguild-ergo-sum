package repository

type CoomandRepository interface {
	Get(key string) (string)
}