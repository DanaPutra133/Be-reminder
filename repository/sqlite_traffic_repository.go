package repository

import (
	"backend-noted/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sqliteTrafficRepository struct {
	db *gorm.DB
}

func NewSqliteTrafficRepository(db *gorm.DB) domain.TrafficRepository {
	return &sqliteTrafficRepository{db}
}

func (r *sqliteTrafficRepository) UpsertTraffic(stat *domain.TrafficStat, incGet, incPost, incPut, incDel int) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "timestamp"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"get":        gorm.Expr("get + ?", incGet),
			"post":       gorm.Expr("post + ?", incPost),
			"put":        gorm.Expr("put + ?", incPut),
			"delete_req": gorm.Expr("delete_req + ?", incDel),
		}),
	}).Create(stat).Error
}

func (r *sqliteTrafficRepository) GetStats(limit int) ([]domain.TrafficStat, error) {
	var stats []domain.TrafficStat
	err := r.db.Order("timestamp desc").Limit(limit).Find(&stats).Error
	return stats, err
}