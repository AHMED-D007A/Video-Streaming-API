-- Drop the trigger first
DROP TRIGGER IF EXISTS trigger_update_video_updated_at ON videos;

-- Drop the function associated with the trigger
DROP FUNCTION IF EXISTS update_video_updated_at_column;

-- Drop the videos table
DROP TABLE IF EXISTS videos;