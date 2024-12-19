package usecase

type Provider interface {
	FetchCount() (int, error)
	CheckCountExist() (bool, error)
	UpdateCount(int) error
}
