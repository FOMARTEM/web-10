package api

type Usecase interface {
	FetchCount() (int, error)
	IncrementCount(int) error
}
