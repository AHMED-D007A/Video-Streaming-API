package videoUploading

import (
	"Video-Streaming-API/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type UploadHandler struct {
	storage *UploadStorage
}

func NewUploadHandler(storage *UploadStorage) *UploadHandler {
	return &UploadHandler{
		storage: storage,
	}
}

func (h *UploadHandler) UploadVideo(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(config.EmailKey).(string)
	userID, err := h.storage.GetUserIDByEmail(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Max upload size is 1 GB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<30)

	// Parse the multipart form
	err = r.ParseMultipartForm(1 << 29) // 512 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("video")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get the file from the form
	thumnail, thumnailHandler, err := r.FormFile("thumbnail")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer thumnail.Close()

	// Store metadata in the database to get the video ID
	video := &Video{
		UserID:      userID,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		FilePath:    "", // Will be updated later
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	videoID, err := h.storage.StoreVideo(video)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a directory named with the user ID and video ID
	videoDir := filepath.Join("./uploads", fmt.Sprintf("%d", userID), fmt.Sprintf("%d", videoID))
	if err := os.MkdirAll(videoDir, os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the file to the new directory
	finalFilePath := filepath.Join(videoDir, handler.Filename)
	dst, err := os.Create(finalFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the thumnail to the new directory
	finalThumnailPath := filepath.Join(videoDir, thumnailHandler.Filename)
	thumnailDST, err := os.Create(finalThumnailPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer thumnailDST.Close()

	if _, err := io.Copy(thumnailDST, thumnail); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the video file path in the database
	video.FilePath = finalFilePath
	if err := h.storage.UpdateVideoFilePath(videoID, finalFilePath, finalThumnailPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Video uploaded successfully"})
}
