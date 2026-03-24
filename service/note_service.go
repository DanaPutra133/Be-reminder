package service

import (
	"backend-noted/domain"
	"errors"
	"time"
)

type noteService struct {
	repo domain.NoteRepository
}

func NewNoteService(repo domain.NoteRepository) domain.NoteService {
	return &noteService{repo}
}

func (s *noteService) CreateNote(note *domain.Note) error {
	inputDate, err := time.Parse("2006-01-02", note.Tanggal)
	if err != nil {
		return errors.New("format tanggal salah, gunakan YYYY-MM-DD")
	}

	today := time.Now().Truncate(24 * time.Hour)
	if inputDate.Before(today) {
		return errors.New("tanggal sudah lewat, tidak bisa menambahkan noted")
	}

	return s.repo.Create(note)
}

func (s *noteService) GetNotes(jid string) ([]domain.Note, error) {
	return s.repo.GetByJidGrub(jid)
}

func (s *noteService) UpdateNote(jid string, data map[string]interface{}) (int64, error) {
	return s.repo.Update(jid, data)
}

func (s *noteService) DeleteNote(jid string) (int64, error) {
	return s.repo.Delete(jid)
}