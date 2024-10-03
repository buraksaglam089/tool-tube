# Spotify to YouTube Playlist Converter

This application allows you to convert Spotify playlists to YouTube playlists. It's developed using Go (chi, goauth, gorilla sessions) on the backend and React (Vite, Redux, Radix UI) on the frontend.

## Requirements

- Go (1.16 or higher)
- Node.js (14 or higher)
- npm or yarn
- Google Developer Console account
- Spotify Developer account

## Installation

### Backend (Go)

1. Clone the project:

   ```
   git clone https://github.com/buraksaglam089/tool-tube.git
   cd tool-tube
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Create a `.env` file and add the necessary environment variables:

   ```
   SPOTIFY_CLIENT_ID=your_spotify_client_id
   SPOTIFY_CLIENT_SECRET=your_spotify_client_secret
   GOOGLE_CLIENT_ID=your_google_client_id
   GOOGLE_CLIENT_SECRET=your_google_client_secret
   SESSION_KEY=your_session_key
   ```

4. Run the application:
   ```
   make run
   ```

### Frontend (React-Vite)

1. Navigate to the frontend directory:

   ```
   cd client
   ```

2. Install dependencies:

   ```
   npm install
   # or
   yarn
   ```

3. Start the application in development mode:
   ```
   npm run dev
   # or
   yarn dev
   ```

## Google Developer Console Setup

1. Go to the [Google Developer Console](https://console.developers.google.com/) and create a new project.

2. Navigate to the "Credentials" tab and select "Create Credentials" > "OAuth client ID".

3. Choose "Web application" as the application type.

4. Add `http://localhost:5173` to "Authorized JavaScript origins" (Vite's default port).

5. Add `http://localhost:8080/auth/google/callback` to "Authorized redirect URIs" (your Go application's callback URL).

6. Add the generated Client ID and Client Secret to your `.env` file.

7. In the "OAuth consent screen" tab, configure your app and add the necessary YouTube API permissions.

## Spotify Developer Setup

1. Go to the [Spotify Developer Dashboard](https://developer.spotify.com/dashboard/) and create a new application.

2. In your application's settings, add `http://localhost:8080/auth/spotify/callback` to the "Redirect URIs" section.

3. Add the Client ID and Client Secret to your `.env` file.

## Usage

1. Open `http://localhost:5173` in your browser.

2. Log in with your Spotify and Google accounts.

3. Select the Spotify playlist you want to convert.

4. Click the convert button and wait for your YouTube playlist to be created.

## Notes

- Both the backend and frontend servers need to be running simultaneously for the application to work properly.
- Be mindful of API rate limits. Converting large playlists may take some time.
- If you encounter any issues, please open an issue on the project repository.
