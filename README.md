<div align="center">
  <h1>🎵 Go Spotify CLI</h1>
</div>

<div align="center">
    <img src="assets/emoji.png" width="300" height="300" alt="Go Spotify CLI Logo">
</div>

<div align="center">
    <a href="https://www.buymeacoffee.com/envoy49">
        <img src="https://img.buymeacoffee.com/button-api/?text=BuyMeACoffee&emoji=&slug=envoy49&button_colour=51e2f5&font_colour=000000&font_family=Cookie&outline_colour=000000&coffee_colour=FFDD00" alt="Buy Me A Coffee">
    </a>
</div>

Go Spotify CLI is a command-line interface tool built with GoLang (v1.21) that interfaces with the Spotify Web API, allowing users to control playback, manage devices, and more.

---

# 📌 Prerequisites

- **Spotify Account** is required to obtain `ClientId` and `ClientSecret`.

---

# ⬇️ Installation

### 🍏 `Mac`, 🐧 `Linux`, 🪟 `Windows`

Note: 📝 Previous installation methods using Homebrew and Winget have been phased out in favor of the more direct `go install` method.

### To install `go-spotify-cli`, follow these steps:

- Install Go version 1.21 or above https://go.dev/
- To install latest version run the following command
```bash
go install github.com/envoy49/go-spotify-cli@latest
```
- To install a specific version of go-spotify-cli, first obtain the release number from the [Releases](https://github.com/Envoy49/go-spotify-cli/releases) page. Then, use the following command in the terminal, replacing <version> with the release number you obtained:

```bash
go install github.com/envoy49/go-spotify-cli@<version>
```

After the installation is complete, open a new terminal and run `go-spotify-cli --version`.

> **Note** 📝: `go-spotify-cli --version` command is hardcoded at this point until the issue with dynamic assignment is resolved.
---

# 🔧 Configuration

Before using the Go Spotify CLI:

1. **Obtain Credentials**:
    - Visit the [Spotify Developer Dashboard](https://developer.spotify.com/dashboard/applications).
    - Log in and create a new app.
    - Once the app is created, you can retrieve the `ClientId` and `ClientSecret` from the dashboard of the created application.


2. **Input Credentials**:
    - Execute any command using the Go Spotify CLI.
    - On first execution, there will be a prompt asking you to enter the `ClientId` and `ClientSecret`.
    - After entering these details, they will be saved in the `.go-spotify-cli` folder in the root directory for future use.

>**Note**: 📝 If secrets entered are wrong although validation is in place, `flush-secrets` can be used to delete saved secrets.

---

# 🔑 Authentication

🚀 On the first run, Go Spotify CLI will initiate an authentication process through the Spotify interface. 
A browser window will open, requesting access grant. Once access is granted, Spotify will issue a 1-hour auth token along with a refresh token. 
The refresh token will be used to obtain a new token after the original token has expired. 
This ensures that browser authentication is no longer required after initial access has been granted.

>**Note**: 📝 Tokens are stored in the `.go-spotify-cli` folder of root directory.

---
# 📟 Commands usage

Type `go-spotify-cli` + `<command>`

▶️ **`play`**: Starts playback on the current device.

⏸️ **`pause`**: Pauses playback on the current device.

⏩ **`next`**: Skips to the next track.

⏪ **`previous`**: Returns to the previous track.

🔊 **`volume`**: Adjusts volume (0-100). Usage: `example: volume -v=80`.

📱 **`device`**: Activates a specific device from provided options. E.g. laptop, tablet, phone etc.

💾 **`saved`**: Prints a list of saved tracks and allows to play selected track.

🔍 **`search`**: Search any `Tracks` and `Episodes`. Searching `Artists`, `Albums`, `Playlists`, `Shows` are not available yet. Any selected song from search result will be added to the current queue.

🔄 **`flush-tokens`**: This command will delete all token saved in `.go-spotify-cli` folder. Further commands will require a new browser authentication.

🔄 **`flush-secrets`**: This command will delete all secrets saved in `.go-spotify-cli` folder.

>**Note**: 📝`To make executing commands more convenient, aliases can be utilized.`

---

# 🌍 Endpoints

The CLI communicates with the following Spotify API Endpoints:

1. `/v1/me/player/play`
2. `/v1/me/player/pause`
3. `/v1/me/player/next`
4. `/v1/me/player/previous`
5. `/v1/me/player/volume`
6. `/v1/me/player/devices`
7. `/v1/me/tracks`

~~**Note**: 📝 More endpoints and functionality will be added once this project gains 25 stars.~~

#### Repo gained more stars than expected, so adding new functionality will be prioritised. 

---

# 🤝 Contributing

Your contributions light up our world! 🌟 Feel free to submit pull requests or raise issues.
There are still a lot of endpoints which can be implemented and a lot of room for improvement.

---

# 💻 Local Development

1. Install Go version 1.21 or above https://go.dev/
2. Clone repo
3. Run `go mod tidy`
4. Follow `Configuration` steps mentioned above
5. Now commands can be executed from root directory e.g. `go run main.go play`

---

# 📝 TODO List

1. Add search options for `Artists`, `Albums`, `Playlists`, `Shows`.
2. Add tests(use race flag to detect race conditions).
3. Add more commands reflecting Spotify Api.
4. Clean up global variables.
5. Get rid of constants, common folders and refactor code.

---

# 📜 License

This project is under the MIT License. Dive into the `LICENSE` file for more.
