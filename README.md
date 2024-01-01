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

# ⬇️ Installation

### 🍏 `Mac`
Brew package manager is required to install Go Spotify CLI. More information on steps to download brew can be found here: `https://brew.sh/`

Once brew is installed, below steps are required for installation.

###  `Homebrew Installation Guide`

```bash
brew tap Envoy49/homebrew-go-spotify-cli
```
```
brew install go-spotify-cli
```

###  `Update`

```bash
brew update
```

```
brew upgrade go-spotify-cli
```



### 🐧 `Linux`

🧪 To be tested with Homebrew installation

### 🪟 `Windows`

Sure, here's a formatted and ready-to-use README for Chocolatey installation steps for Go Spotify CLI:



###  `Chocolatey Installation Guide`

### Installing Chocolatey

### Open PowerShell as Administrator
- Right-click on the Start button.
- Click on “Windows PowerShell (Admin)” or “Command Prompt (Admin)” if PowerShell is not available.

### Run the Installation Command
Paste the following command into the PowerShell window and press Enter:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

### Close and Reopen PowerShell
- After the installation is complete, close the PowerShell window.
- Reopen a new PowerShell window as an administrator to start using Chocolatey.

### Verify Installation
To verify that Chocolatey is installed, run:

```powershell
choco -?
```

### Using Chocolatey to Manage Go Spotify CLI

### Installation
To install Go Spotify CLI, run the following command from the command line or from PowerShell:

```powershell
choco install go-spotify-cli --version=<version here>
```

### Upgrade
To upgrade Go Spotify CLI, run the following command from the command line or from PowerShell:

```powershell
choco upgrade go-spotify-cli --version=<version here>
```

### Uninstallation
To uninstall Go Spotify CLI, run the following command from the command line or from PowerShell:

```powershell
choco uninstall go-spotify-cli --version=<version here>
```


#### Remember to replace `<version here>` with the specific version number of Go Spotify CLI you want to install, upgrade, or uninstall. This README provides clear and concise instructions for users to manage Go Spotify CLI via Chocolatey.

---

# 📌 Prerequisites

- **Spotify Account** is required to obtain `ClientId` and `ClientSecret`.

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

    **Note**: If secrets entered are wrong although validation is in place, `flush-secrets` can be used to delete saved secrets.

---

# 🔑 Authentication

🚀 On the first run, Go Spotify CLI will initiate an authentication process through the Spotify interface. 
A browser window will open, requesting access grant. Once access is granted, Spotify will issue a 1-hour auth token along with a refresh token. 
The refresh token will be used to obtain a new token after the original token has expired. 
This ensures that browser authentication is no longer required after initial access has been granted.

**Note**: 📝 Tokens are stored in the `.go-spotify-cli` folder of root directory.

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

**Note**: 📝`To make executing commands more convenient, aliases can be utilized.`

---

# 🌍 Endpoints

The CLI communicates with the following Spotify API Endpoints:

1. `/v1/me/player/play`
2. `/v1/me/player/pause`
3. `/v1/me/player/next`
4. `/v1/me/player/previous`
5. `/v1/me/player/volume`
6. `/v1/me/player/devices`
7. `/v1/me/player/devices`
8. `/v1/me/tracks`

**Note**: 📝 More endpoints and functionality will be added once this project gains 25 stars.

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
5. Now commands can be executed from `cmd/gsc` folder e.g. `go run cmd/gsc/main.go play`

---

# 📝 TODO List

1. Add liked songs command
2. Add search options for `Artists`, `Albums`, `Playlists`, `Shows`
3. Add tests(use race flag to detect race conditions)
4. Add more commands reflecting Spotify Api
5. Test on Windows and Linux. Tested and developed only on Mac.

---

# 📜 License

This project is under the MIT License. Dive into the `LICENSE` file for more.
