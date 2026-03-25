package domain

type TrafficStat struct {
	Timestamp int64 `gorm:"primaryKey;column:timestamp" json:"timestamp"`
	GET       int   `gorm:"column:get" json:"GET"`
	POST      int   `gorm:"column:post" json:"POST"`
	PUT       int   `gorm:"column:put" json:"PUT"`
	DELETEReq int   `gorm:"column:delete_req" json:"DELETE"`
}

type TrafficRepository interface {
	UpsertTraffic(stat *TrafficStat, incGet, incPost, incPut, incDel int) error
	GetStats(limit int) ([]TrafficStat, error)
}
type TrafficService interface {
	GetServerStats() ([]TrafficStat, error)
}