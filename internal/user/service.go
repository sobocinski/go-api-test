package user

type Service struct {
	repository UserRepository
}

// NewService - returns a new comment service
func NewService(ur UserRepository) *Service {
	return &Service{
		repository: ur,
	}
}

func (s *Service) GetUser(id UserId) (User, error) {
	return s.repository.GetById(id)
}

func (s *Service) RegisterUser(u User) (*User, error) {
	var err error
	u.Password, err = hashAndSaltPwd(u.Password)
	if err != nil {
		return nil, err
	}

	return s.repository.Create(u)
}

func (s *Service) GetAll() ([]User, error) {
	return s.repository.GetAll()
}
