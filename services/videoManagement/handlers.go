package videoManagement

import (
	"Video-Streaming-API/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type VideoHandler struct {
	storage *VideoStorage
}

func NewVideoHandler(storage *VideoStorage) *VideoHandler {
	return &VideoHandler{storage: storage}
}

func (h *VideoHandler) GetVideos(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(config.EmailKey).(string)
	userID, err := h.storage.GetUserIDByEmail(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	videos, thumbnail_paths, err := h.storage.GetVideos(userID)
	if err != nil {
		http.Error(w, "Failed to fetch videos", http.StatusInternalServerError)
		return
	}

	type VideoResponse struct {
		Video     Video  `json:"video"`
		Thumbnail string `json:"thumbnail"`
	}

	var response []VideoResponse
	for i, video := range videos {
		response = append(response, VideoResponse{
			Video:     video,
			Thumbnail: "/uploads/" + thumbnail_paths[i],
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *VideoHandler) UpdateVideo(w http.ResponseWriter, r *http.Request) {
	var video VideoPayload

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid video ID", http.StatusBadRequest)
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
	thumnail, thumnailHandler, err := r.FormFile("thumbnail")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer thumnail.Close()

	oldThumbnailPath, err := h.storage.GetVideoThumbnail(id)
	if err != nil {
		http.Error(w, "Failed to fetch video thumbnail", http.StatusInternalServerError)
		return
	}

	err = os.Remove(oldThumbnailPath)
	if err != nil {
		http.Error(w, "Failed to delete old thumbnail", http.StatusInternalServerError)
		return
	}

	temp := strings.Split(oldThumbnailPath, "/")

	video.Title = r.FormValue("title")
	video.Description = r.FormValue("description")
	video.Thumbnail_path = temp[0] + "/" + temp[1] + "/" + temp[2] + "/" + thumnailHandler.Filename

	thumnailDST, err := os.Create(video.Thumbnail_path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer thumnailDST.Close()

	if _, err := io.Copy(thumnailDST, thumnail); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	video.ID = id
	if err := h.storage.UpdateVideo(&video); err != nil {
		http.Error(w, "Failed to update video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Video updated successfully"})
}

func (h *VideoHandler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(config.EmailKey).(string)
	userID, err := h.storage.GetUserIDByEmail(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid video ID", http.StatusBadRequest)
		return
	}

	if err := h.storage.DeleteVideo(id); err != nil {
		http.Error(w, "Failed to delete video", http.StatusInternalServerError)
		return
	}

	path := "uploads/" + strconv.Itoa(userID) + "/" + vars["id"]
	fmt.Println(path)
	err = os.RemoveAll(path)
	if err != nil {
		http.Error(w, "Failed to delete video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"message": "Video deleted successfully"})
}
