package worker

import (
	"backend-noted/domain"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func SetupCron(repo domain.NoteRepository) *cron.Cron {
	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		log.Println("[WORKER] cek Cleaning DB")
		today := time.Now().Format("2006-01-02")
		
		rows, err := repo.DeleteExpired(today)
		if err != nil {
			log.Println("[WORKER] Error:", err)
			return
		}
		log.Printf("[WORKER] Berhasil menghapus %d data kadaluarsa.\n", rows)
	})
	c.Start()
	return c
}