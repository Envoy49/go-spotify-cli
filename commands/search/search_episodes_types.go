package search

import "go-spotify-cli/common"

type Episodes struct {
	Body
	Items []EpisodeItem `json:"items"`
}

type EpisodeItem struct {
	AudioPreviewURL      string              `json:"audio_preview_url"`
	Description          string              `json:"description"`
	HTMLDescription      string              `json:"html_description"`
	DurationMS           int                 `json:"duration_ms"`
	Explicit             bool                `json:"explicit"`
	ExternalURLs         map[string]string   `json:"external_urls"`
	Href                 string              `json:"href"`
	ID                   string              `json:"id"`
	Images               []common.Image      `json:"images"`
	IsExternallyHosted   bool                `json:"is_externally_hosted"`
	IsPlayable           bool                `json:"is_playable"`
	Language             string              `json:"language"`
	Languages            []string            `json:"languages"`
	Name                 string              `json:"name"`
	ReleaseDate          string              `json:"release_date"`
	ReleaseDatePrecision string              `json:"release_date_precision"`
	ResumePoint          EpisodeResumePoint  `json:"resume_point"`
	Type                 string              `json:"type"`
	URI                  string              `json:"uri"`
	Restrictions         EpisodeRestrictions `json:"restrictions"`
}

type EpisodeResumePoint struct {
	FullyPlayed      bool `json:"fully_played"`
	ResumePositionMS int  `json:"resume_position_ms"`
}

type EpisodeRestrictions struct {
	Reason string `json:"reason"`
}
