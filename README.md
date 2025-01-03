# Video-Streaming-API

The Video Streaming API is a web API that enables users to upload, store, and stream videos on demand. This project uses Go for the backend and supports adaptive bitrate streaming using HLS (HTTP Live Streaming).

## Features

- **Video Uploading**: Users can upload videos which are then transcoded to multiple resolutions.
- **Adaptive Bitrate Streaming**: Videos are chunked into HLS segments and served with a master playlist for adaptive streaming.
- **User Authentication**: Secure user authentication using JWT.
- **Video Management**: Manage video metadata and file paths in the database.

## Project Structure

- **bin**: Compiled binary files.
- **cmd**: Main application entry point.
- **config**: Configuration settings for the application.
- **database**: Database connection.
- **migrations**: Database schema migrations.
- **server**: HTTP server setup, middleware and routes.
- **services**: Business logic for authentication, video management, and video streaming.
- **utils**: Utility functions for handling JWT tokens and file uploads.

### Technologies

- Go 1.23.4 or later
- PostgreSQL
- FFmpeg

### Running the Application

1. Build the project:
    ```sh
    make build
    ```

2. Run the project:
    ```sh
    make run
    ```

### API Endpoints

- **Authenticaion**: `POST /api/v1/signup`, `POST /api/v1/login`
- **Video Management**: `GET /api/v1/uvideos`, `PUT /api/v1/uvideos/{id:[0-9]+}`, `DELETE /api/v1/uvideos/{id:[0-9]+}`
- **Upload Video**: `POST /api/v1/upload`
- **Stream Video**: `GET /api/v1/stream/{videoID}/master.m3u8`
- **Stream Video Segment**: `GET /api/v1/stream/{videoID}/{segment}`

### Example

To test the video streaming, open streaming.html in your browser and replace the videoID with an actual video ID from your database.

### Environment Variables

The following environment variables need to be set in the .env file:

- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `JWT_SECRET`: Secret key for JWT

### Dependencies

This project uses the following dependencies:

- [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
- [github.com/gorilla/mux](https://github.com/gorilla/mux)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)
- [github.com/lib/pq](https://github.com/lib/pq)
- [golang.org/x/crypto](https://golang.org/x/crypto)

### Main Components

- **Main Entry Point**: [main.go](http://_vscodecontentref_/17)
- **Configuration**: [config.go](http://_vscodecontentref_/18)
- **Database Connection**: [connection.go](http://_vscodecontentref_/19)
- **Server Setup**: [server.go](http://_vscodecontentref_/20)
- **Routes**: [routes.go](http://_vscodecontentref_/21)
- **Middleware**: [middleware.go](http://_vscodecontentref_/22)

### Services

- **Authentication**: [authentication](http://_vscodecontentref_/23)
- **Video Management**: [videoManagement](http://_vscodecontentref_/24)
- **Video Streaming**: [videoStreaming](http://_vscodecontentref_/25)
- **Video Uploading**: [videoUploading](http://_vscodecontentref_/28)

### Utility Functions

- **JWT Token Handling**: [JWT_token.go](http://_vscodecontentref_/30)
