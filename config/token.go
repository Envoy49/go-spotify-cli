package config

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

const (
	UserLibraryReadRoute         = "/user-library-read-auth"
	UserLibraryReadRouteCallback = "/user-library-read-auth-callback"
	UserLibraryRead              = "user-library-read"
)

type TokenType string

const (
	ModifyToken TokenType = "ModifyToken"
	ReadToken   TokenType = "ReadToken"
	LibraryRead TokenType = "LibraryRead"
)

type CombinedTokenStructure struct {
	ModifyToken      UserModifyTokenStructure      `yaml:"ModifyToken"`
	ReadToken        UserReadTokenStructure        `yaml:"ReadToken"`
	LibraryReadToken UserLibraryReadTokenStructure `yaml:"LibraryReadToken"`
}

type UserModifyTokenStructure struct {
	UserModifyToken          string `yaml:"UserModifyToken"`
	UserModifyRefreshToken   string `yaml:"UserModifyRefreshToken"`
	UserModifyTokenExpiresIn int64  `yaml:"UserModifyTokenExpiresIn"`
}

type UserReadTokenStructure struct {
	UserReadToken          string `yaml:"UserReadToken"`
	UserReadRefreshToken   string `yaml:"UserReadRefreshToken"`
	UserReadTokenExpiresIn int64  `yaml:"UserReadTokenExpiresIn"`
}

type UserLibraryReadTokenStructure struct {
	UserLibraryReadToken          string `yaml:"UserLibraryReadToken"`
	UserLibraryReadRefreshToken   string `yaml:"UserLibraryReadRefreshToken"`
	UserLibraryReadTokenExpiresIn int64  `yaml:"UserLibraryReadTokenExpiresIn"`
}
