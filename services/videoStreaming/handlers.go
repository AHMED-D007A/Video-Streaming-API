package videoStreaming

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type StreamingHandler struct {
	storage *StreamingStorage
}

func NewStreamingHandler(storage *StreamingStorage) *StreamingHandler {
	return &StreamingHandler{storage: storage}
}

func (h *StreamingHandler) ServePlaylist(w http.ResponseWriter, r *http.Request) {
	videoIDStr := strings.TrimPrefix(r.URL.Path, "/api/v1/stream/")
	videoIDStr = strings.TrimSuffix(videoIDStr, "/master.m3u8")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		http.Error(w, "Invalid video ID", http.StatusBadRequest)
		return
	}

	filePath, err := h.storage.GetVideoFilePath(videoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	playlistPath := filepath.Join(filepath.Dir(filePath), "master.m3u8")
	fmt.Println(playlistPath)
	http.ServeFile(w, r, playlistPath)
}

func (h *StreamingHandler) ServeSegment(w http.ResponseWriter, r *http.Request) {
	// Extract the video ID and segment path from the URL
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 5 {
		http.Error(w, "Invalid segment path", http.StatusBadRequest)
		return
	}

	videoIDStr := urlParts[4]
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		http.Error(w, "Invalid video ID", http.StatusBadRequest)
		return
	}

	filePath, err := h.storage.GetVideoFilePath(videoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	segmentPath := filepath.Join(filepath.Dir(filePath), strings.Join(urlParts[5:], "/"))
	fmt.Println(segmentPath)
	http.ServeFile(w, r, segmentPath)
}
