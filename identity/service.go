package identity

import (
	"github.com/withzeus/mugi-identity/core"
)

type Service struct {
	ds *Datastore
	h  core.Helper
}

func NewService(ds *Datastore, h core.Helper) *Service {
	return &Service{ds: ds, h: h}
}

func (s *Service) Create(md Model) (*Response, error) {
	i, err := s.ds.Create(md)

	if err != nil {
		return nil, err
	}
	return i.ToResponse(), nil
}
