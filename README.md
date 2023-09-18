# 🎵 Go Spotify CLI

Go Spotify CLI is a command-line interface tool built with GoLang (v1.20) that interfaces with the Spotify Web API, allowing users to control playback, manage devices, and more.

## 📌 Prerequisites

- **GoLang v1.20 or above**: Ensure you have this version installed on your machine.
- **Spotify Account**: Required to obtain `ClientId` and `ClientSecret`.

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

🚀 On the first run, Go Spotify CLI will prompt an authentication process. A browser window will open, requesting access to your Spotify account. After granting access, you'll receive a token.

**Note**: 🕐 This token has a 1-hour expiry. Re-authenticate once expired. Token is securely stored in the `.go-spotify-cli` directory.

## 🎛 Commands

🎶 **`play`**: Starts playback on the current device.

⏸ **`pause`**: Pauses playback on the current device.

⏩ **`next`**: Skips to the next track.

⏪ **`previous`**: Returns to the previous track.

🔊 **`volume`**: Adjusts volume (0-100). Usage: `volume [0-100]`.

📱 **`device`**: Activates a specific device. Usage: `device [DEVICE_ID]`.

## 🌍 Endpoints

The CLI communicates with the following Spotify API Endpoints:

1. `/v1/me/player/play`
2. `/v1/me/player/pause`
3. `/v1/me/player/next`
4. `/v1/me/player/previous`
5. `/v1/me/player/volume`
6. `/v1/me/player/devices`

We are hoping to add more endpoints with the help of the community.

## 🤝 Contributing

Your contributions light up our world! 🌟 Feel free to submit pull requests or raise issues.

## 📜 License

This project is under the MIT License. Dive into the `LICENSE` file for more.
