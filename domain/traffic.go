package domain

// Model Database untuk Traffic
type TrafficStat struct {
	Timestamp int64 `gorm:"primaryKey;column:timestamp" json:"timestamp"`
	GET       int   `gorm:"column:get" json:"get"`
	POST      int   `gorm:"column:post" json:"post"`
	PUT       int   `gorm:"column:put" json:"put"`
	DELETEReq int   `gorm:"column:delete_req" json:"deleteReq"`
}

type TrafficRepository interface {
	UpsertTraffic(stat *TrafficStat, incGet, incPost, incPut, incDel int) error
	GetStats(limit int) ([]TrafficStat, error)
}

type TrafficService interface { 
	GetServerStats() ([]TrafficStat, error)
}