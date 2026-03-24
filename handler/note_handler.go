package handler

import (
	"backend-noted/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	Service domain.NoteService
}

func NewNoteHandler(r *gin.Engine, service domain.NoteService) {
	handler := &NoteHandler{Service: service}

	r.POST("/noted", handler.Create)
	r.GET("/noted", handler.Get)
	r.PATCH("/noted", handler.Update)
	r.DELETE("/noted", handler.Delete)
}

func (h *NoteHandler) Create(c *gin.Context) {
	var note domain.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	if err := h.Service.CreateNote(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "sukses", "message": "Noted berhasil disimpan", "data": note})
}

func (h *NoteHandler) Get(c *gin.Context) {
	jidGrub := c.Query("jidgrub")
	if jidGrub == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter ?jidgrub= wajib diisi"})
		return
	}

	notes, err := h.Service.GetNotes(jidGrub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "sukses", "data": notes})
}

func (h *NoteHandler) Update(c *gin.Context) {
	jidGrub := c.Query("jidgrub")
	if jidGrub == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter ?jidgrub= wajib diisi"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	rows, err := h.Service.UpdateNote(jidGrub, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data dengan jidgrub tersebut tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "sukses", "message": "Data berhasil diupdate"})
}

func (h *NoteHandler) Delete(c *gin.Context) {
	jidGrub := c.Query("jidgrub")
	if jidGrub == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter ?jidgrub= wajib diisi"})
		return
	}

	rows, err := h.Service.DeleteNote(jidGrub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "sukses", "message": "Berhasil menghapus noted untuk grub tersebut"})
}