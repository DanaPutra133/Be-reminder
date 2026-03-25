package service

import (
	"backend-noted/domain"
	"time"
)

type trafficService struct {
	repo domain.TrafficRepository
}

func NewTrafficService(repo domain.TrafficRepository) domain.TrafficService {
	return &trafficService{repo}
}

func (s *trafficService) GetServerStats() ([]domain.TrafficStat, error) {
	stats, err := s.repo.GetStats(60)
	if err != nil {
		return nil, err
	}

	statMap := make(map[int64]domain.TrafficStat)
	for _, st := range stats {
		statMap[st.Timestamp] = st
	}

	var continuousStats []domain.TrafficStat
	
	now := time.Now()
	roundedNow := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())

	for i := 59; i >= 0; i-- {
		ts := roundedNow.Add(-time.Duration(i) * time.Minute).UnixMilli()
		
		if val, ok := statMap[ts]; ok {
			continuousStats = append(continuousStats, val)
		} else {
			continuousStats = append(continuousStats, domain.TrafficStat{
				Timestamp: ts,
				GET:       0,
				POST:      0,
				PUT:       0,
				DELETEReq: 0,
			})
		}
	}

	return continuousStats, nil
}