CREATE TABLE videos (
    id SERIAL PRIMARY KEY,               -- Auto-incrementing unique ID for each video
    user_id INT NOT NULL,                -- Foreign key referencing the user who uploaded the video
    title VARCHAR(255) NOT NULL,         -- Title of the video
    description TEXT,                    -- Description of the video
    file_path TEXT NOT NULL,             -- Path to the video file
    thumbnail_path TEXT,                 -- Path to the thumbnail image
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Video upload time
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Last update time
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE -- Foreign key constraint
);

-- Add a trigger to update the `updated_at` field on any change
CREATE OR REPLACE FUNCTION update_video_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_video_updated_at
BEFORE UPDATE ON videos
FOR EACH ROW
EXECUTE FUNCTION update_video_updated_at_column();