package domain

type Note struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	JidGrub string `gorm:"type:varchar(100);index" json:"jidGrub"`
	Jam     string `gorm:"type:varchar(10)" json:"jam"`
	Tanggal string `gorm:"type:date" json:"tanggal"`
	Noted   string `gorm:"type:text" json:"noted"`
	Img     string `gorm:"type:varchar(255)" json:"img"`
}

type NoteRepository interface {
	Create(note *Note) error
	GetByJidGrub(jid string) ([]Note, error)
	Update(id uint, jid string, data map[string]interface{}) (int64, error)
	Delete(id uint, jid string) (int64, error)
	DeleteExpired(date string) (int64, error)
}

type NoteService interface {
	CreateNote(note *Note) error
	GetNotes(jid string) ([]Note, error)
	UpdateNote(id uint, jid string, data map[string]interface{}) (int64, error)
	DeleteNote(id uint, jid string) (int64, error)
}