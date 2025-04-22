package user

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	s := Service{repo}
	return &s
}

func (s *Service) EmailExists(u *User) (bool, error) {
	user, err := s.Repo.FindByEmail(u)

	if err != nil {
		return false, err
	}

	if user != nil {
		return true, nil
	}

	return false, nil
}
