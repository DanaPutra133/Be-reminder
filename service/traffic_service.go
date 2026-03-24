package service

import "backend-noted/domain"

type trafficService struct {
	repo domain.TrafficRepository
}

func NewTrafficService(repo domain.TrafficRepository) domain.TrafficService {
	return &trafficService{repo}
}

func (s *trafficService) GetServerStats() ([]domain.TrafficStat, error) {
	return s.repo.GetStats(100)
}