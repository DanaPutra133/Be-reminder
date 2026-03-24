package repository

import (
	"backend-noted/domain"

	"gorm.io/gorm"
)

type sqliteNoteRepository struct {
	db *gorm.DB
}

func NewSqliteNoteRepository(db *gorm.DB) domain.NoteRepository {
	return &sqliteNoteRepository{db}
}

func (r *sqliteNoteRepository) Create(note *domain.Note) error {
	return r.db.Create(note).Error
}

func (r *sqliteNoteRepository) GetByJidGrub(jid string) ([]domain.Note, error) {
	var notes []domain.Note
	err := r.db.Where("jid_grub = ?", jid).Find(&notes).Error
	return notes, err
}

func (r *sqliteNoteRepository) Update(jid string, data map[string]interface{}) (int64, error) {
	res := r.db.Model(&domain.Note{}).Where("jid_grub = ?", jid).Updates(data)
	return res.RowsAffected, res.Error
}

func (r *sqliteNoteRepository) Delete(jid string) (int64, error) {
	res := r.db.Where("jid_grub = ?", jid).Delete(&domain.Note{})
	return res.RowsAffected, res.Error
}

func (r *sqliteNoteRepository) DeleteExpired(date string) (int64, error) {
	res := r.db.Where("tanggal < ?", date).Delete(&domain.Note{})
	return res.RowsAffected, res.Error
}