package identity

type IIdentityService interface {
	Create(i Identity)
}

type IdentityService struct {
	db DB
}

func NewIdentityService(db DB) *IdentityService {
	return &IdentityService{db: db}
}

func (s *IdentityService) Create(id Identity) (*IdentityResponse, error) {
	i, err := s.db.Create(id)

	if err != nil {
		return nil, err
	}
	return i.ToResponse(), nil
}
