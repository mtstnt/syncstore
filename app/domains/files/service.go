package files

type fileRepository interface {
}

type Service struct {
	repo fileRepository
}

func NewService(r fileRepository) Service {
	return Service{
		repo: r,
	}
}
