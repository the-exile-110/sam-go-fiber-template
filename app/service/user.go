package service

import (
	"app/database/schema"
)

func (s *Service) ListUsers() (*[]schema.Users, error) {
	res := make([]schema.Users, 0)

	if err := s.mysqlDB.Table("users").Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
