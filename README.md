# 🎵 Go Spotify CLI
<img src="assets/emoji.png" width="300" height="300">

Go Spotify CLI is a command-line interface tool built with GoLang (v1.21) that interfaces with the Spotify Web API, allowing users to control playback, manage devices, and more.

## ⬇️ Installation

### 🍏 `Mac`
Brew package manager is required to install Go Spotify CLI. More information on steps to download brew can be found here: `https://brew.sh/`

Once brew is installed, commands steps are required.

```bash
brew tap Envoy49/homebrew-go-spotify-cli
```
```
brew install go-spotify-cli
```

### 🐧 `Linux`

🧪 To be tested with Homebrew installation

### 🪟 `Windows`

🚧 To be developed using Chocolatey package or similar


## 📌 Prerequisites

- **Spotify Account** is required to obtain `ClientId` and `ClientSecret`.

## 🔧 Configuration

Before using the Go Spotify CLI:

1. **Obtain Credentials**:
    - Visit the [Spotify Developer Dashboard](https://developer.spotify.com/dashboard/applications).
    - Log in and create a new app.
    - Once the app is created, you can retrieve the `ClientId` and `ClientSecret` from the dashboard of the created application.


2. **Input Credentials**:
    - Execute any command using the Go Spotify CLI.
    - On first execution, there will be a prompt asking you to enter the `ClientId` and `ClientSecret`.
    - After entering these details, they will be saved in the `.go-spotify-cli` folder in the root directory for future use.

## 🔑 Authentication

🚀 On the first run, Go Spotify CLI will initiate an authentication process through the Spotify interface. 
A browser window will open, requesting access grant. Once access is granted, Spotify will issue a 1-hour auth token along with a refresh token. 
The refresh token will be used to obtain a new token after the original token has expired. 
This ensures that browser authentication is no longer required after initial access has been granted.

**Note**: 💾 Tokens are stored in the `.go-spotify-cli` folder of root directory.

## 📟 Commands

🎶 **`play`**: Starts playback on the current device.

⏸️ **`pause`**: Pauses playback on the current device.

⏩ **`next`**: Skips to the next track.

⏪ **`previous`**: Returns to the previous track.

🔊 **`volume`**: Adjusts volume (0-100). Usage: `volume [0-100]`.

📱 **`device`**: Activates a specific device. Usage: `device [DEVICE_ID]`.

🔍 **`search`**: Search any `Tracks` and `Episodes`. Searching `Artists`, `Albums`, `Playlists`, `Shows` are not available yet. Any selected song from search result will be added to the current queue.

🔄 **`flush-all-tokens`**: This command will delete all token saved in `.go-spotify-cli` folder. Further commands will require a new browser authentication.

## 🌍 Endpoints

The CLI communicates with the following Spotify API Endpoints:

1. `/v1/me/player/play`
2. `/v1/me/player/pause`
3. `/v1/me/player/next`
4. `/v1/me/player/previous`
5. `/v1/me/player/volume`
6. `/v1/me/player/devices`
7. `/v1/search`

We are planning to add more endpoints with the help of the community.

## 🤝 Contributing

Your contributions light up our world! 🌟 Feel free to submit pull requests or raise issues.
There are still a lot of endpoints which can be implemented and a lot of room for improvement.

## 💻 Local Development

1. Install Go version 1.21 or above
2. Pull repo
3. run `go mod tidy`
4. Follow `Configuration` steps mentioned above
5. Now commands can be executed from `cmd/gsc` folder e.g. `go run cmd/gsc/main.go play`

## 📝 TODO List

1. Add liked songs command
2. Add search options for `Artists`, `Albums`, `Playlists`, `Shows`
3. Add tests(use race flag to detect race conditions)
4. Add more commands reflecting Spotify Api
5. Test on Windows and Linux. Tested and developed only on Mac.

## 📜 License

This project is under the MIT License. Dive into the `LICENSE` file for more.
