# Go Spotify CLI

Go Spotify CLI is a command-line interface tool built with GoLang (v1.20) that interfaces with the Spotify Web API, allowing users to control playback, manage devices, and more.

## Prerequisites

- **GoLang v1.20**: Ensure you have this version installed on your machine.
- **Spotify Developer Account**: Required to obtain `ClientId` and `ClientSecret`.

## Configuration

Before using Go Spotify CLI, users need to obtain `ClientId` and `ClientSecret`:

1. Visit the [Spotify Developer Dashboard](https://developer.spotify.com/dashboard/applications).
2. Log in and create a new app.
3. Once the app is created, you'll be provided with a `ClientId` and `ClientSecret`.

### Storing Credentials

Both `ClientId` and `ClientSecret` are stored in the `.go-spotify-cli` folder in the root directory. The CLI will use these credentials for authentication and other necessary operations.

## Authentication

When you run the Go Spotify CLI for the first time, it will prompt an authentication process. A browser window will open, requesting access to your Spotify account. After granting access, you'll receive a token.

**Note**: This token has a 1-hour expiry. Once it's expired, you'll need to re-authenticate. The token is securely stored in the `.go-spotify-cli` directory.

## Commands

### `play`

Starts playback on the current device.

### `pause`

Pauses playback on the current device.

### `next`

Skips to the next track in the playback queue.

### `previous`

Returns to the previous track in the playback queue.

### `volume`

Adjust the volume of the current device. Usage: `volume [0-100]`.

### `device`

Activates a specific device. If the device is off or in a non-active state, this command will "wake" it and set it as the current playback device. Usage: `device [DEVICE_ID]`.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or raise issues if you find any.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
