package uploading

import (
	"Video-Streaming-API/utils"
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
	userEmail := r.Context().Value(utils.EmailKey).(string)
	userID, err := h.storage.GetUserIDByEmail(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Headers:", r.Header)
	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
	fmt.Println("Request Method:", r.Method)

	// Max upload size is 1 GB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<30)

	// Parse the multipart form
	err = r.ParseMultipartForm(1 << 29) // 512 MB
	if err != nil {
		fmt.Println(err)
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

	// Create a directory for storing the uploaded files, named with the user ID
	uploadDir := filepath.Join("./uploads", fmt.Sprintf("%d", userID))
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a file on the server
	filePath := filepath.Join(uploadDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the server
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store metadata in the database
	video := &Video{
		UserID:      userID,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		FilePath:    filePath,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := h.storage.StoreVideo(video); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Video uploaded successfully"))
}
