package constants

const (
	ProjectName           = "go-spotify-cli"
	ServerUrl             = "http://localhost:4949"
	Port                  = ":4949"
	SpotifyPlayerEndpoint = "https://api.spotify.com/v1/me"
	SpotifySearchEndpoint = "https://api.spotify.com/v1/search"
)

const (
	UserModifyPlaybackStateRoute         = "/user-modify-playback-state-auth"
	UserModifyPlaybackStateRouteCallback = "/user-modify-playback-state-auth-callback"
	UserModifyPlaybackStateScope         = "user-modify-playback-state"
)

const (
	UserReadPlaybackStateRoute         = "/user-read-playback-state-auth"
	UserReadPlaybackStateRouteCallback = "/user-read-playback-state-auth-callback"
	UserReadPlaybackState              = "user-read-playback-state"
)

type TokenType string

const (
	ModifyToken TokenType = "ModifyToken"
	ReadToken   TokenType = "ReadToken"
)
