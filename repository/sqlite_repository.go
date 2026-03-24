package repository

import (
	"backend-noted/domain"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type sqliteNoteRepository struct {
	db *gorm.DB
	rdb *redis.Client
}

func NewSqliteNoteRepository(db *gorm.DB, rdb *redis.Client) domain.NoteRepository {
	return &sqliteNoteRepository{db, rdb}
}

func (r *sqliteNoteRepository) Create(note *domain.Note) error {
	err := r.db.Create(note).Error
	if err == nil {
		r.rdb.Del(context.Background(), "notes:"+note.JidGrub) 
	}
	return err
}

func (r *sqliteNoteRepository) GetByJidGrub(jid string) ([]domain.Note, error) {
	ctx := context.Background()
	cacheKey := "notes:" + jid
	var notes []domain.Note
	cachedData, err := r.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(cachedData), &notes)
		return notes, nil
	}
	err = r.db.Where("jid_grub = ?", jid).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	notesJSON, _ := json.Marshal(notes)
	r.rdb.Set(ctx, cacheKey, notesJSON, 1*time.Hour)

	return notes, nil
}

func (r *sqliteNoteRepository) Update(id uint, jid string, data map[string]interface{}) (int64, error) {
	res := r.db.Model(&domain.Note{}).Where("id = ? AND jid_grub = ?", id, jid).Updates(data)
	if res.Error == nil && res.RowsAffected > 0 {
		r.rdb.Del(context.Background(), "notes:"+jid) 
	}
	return res.RowsAffected, res.Error
}

func (r *sqliteNoteRepository) Delete(id uint, jid string) (int64, error) {
	res := r.db.Where("id = ? AND jid_grub = ?", id, jid).Delete(&domain.Note{})
	if res.Error == nil && res.RowsAffected > 0 {
		r.rdb.Del(context.Background(), "notes:"+jid) 
	}
	return res.RowsAffected, res.Error
}

func (r *sqliteNoteRepository) DeleteExpired(date string) (int64, error) {
	var affectedNotes []domain.Note
	r.db.Where("tanggal < ?", date).Find(&affectedNotes)
	res := r.db.Where("tanggal < ?", date).Delete(&domain.Note{})
	if res.Error == nil && res.RowsAffected > 0 {
		ctx := context.Background()
		for _, note := range affectedNotes {
			r.rdb.Del(ctx, "notes:"+note.JidGrub)
		}
	}
	
	return res.RowsAffected, res.Error
}