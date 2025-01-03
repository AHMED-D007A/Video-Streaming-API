package videoStreaming

type Video struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	FilePath      string `json:"file_path"`
	ThumbnailPath string `json:"thumbnail_path"`
}
