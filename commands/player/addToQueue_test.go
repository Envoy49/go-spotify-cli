package player

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func noActiveDeviceError() error {
	apiErr := cmdTypes.SpotifyAPIError{}
	apiErr.Detail.Error.Message = noActiveDeviceMessage
	return apiErr
}

func TestRecoverQueuePlaybackPlaysSelectedTrackOnSelectedDevice(t *testing.T) {
	type capturedRequest struct {
		method        string
		path          string
		query         string
		authorization string
		body          string
	}

	var requests []capturedRequest
	previousClient := http.DefaultClient
	http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		requests = append(requests, capturedRequest{
			method:        request.Method,
			path:          request.URL.Path,
			query:         request.URL.RawQuery,
			authorization: request.Header.Get("Authorization"),
			body:          string(body),
		})
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("")),
			Request:    request,
		}, nil
	})}
	t.Cleanup(func() { http.DefaultClient = previousClient })

	const (
		accessToken = "test-access-token"
		deviceID    = "selected-device"
		trackURI    = "spotify:track:selected"
	)
	var events []string
	startedSelectedTrack, err := recoverQueuePlayback(
		noActiveDeviceError(),
		func() (string, error) {
			events = append(events, "select device")
			return deviceID, nil
		},
		func(selectedDeviceID string, startPlayback bool) error {
			events = append(events, "transfer playback")
			if selectedDeviceID != deviceID {
				t.Fatalf("transfer used device %q, want %q", selectedDeviceID, deviceID)
			}
			if startPlayback {
				t.Fatal("device transfer resumed the previous track")
			}
			return transferPlayback(accessToken, selectedDeviceID, startPlayback)
		},
		func(selectedDeviceID string) error {
			events = append(events, "play selected track")
			return playTrack(accessToken, trackURI, selectedDeviceID)
		},
	)
	if err != nil {
		t.Fatalf("recoverQueuePlayback() error = %v", err)
	}
	if !startedSelectedTrack {
		t.Fatal("recoverQueuePlayback() did not report that it started the selected track")
	}
	if want := []string{"select device", "transfer playback", "play selected track"}; !reflect.DeepEqual(events, want) {
		t.Fatalf("recovery events = %v, want %v", events, want)
	}
	if len(requests) != 2 {
		t.Fatalf("Spotify requests = %d, want 2", len(requests))
	}

	if request := requests[0]; request.method != http.MethodPut || request.path != "/v1/me/player" {
		t.Fatalf("transfer request = %s %s, want PUT /v1/me/player", request.method, request.path)
	}
	if requests[0].authorization != "Bearer "+accessToken {
		t.Fatalf("transfer authorization = %q", requests[0].authorization)
	}
	var transferBody struct {
		DeviceIDs []string `json:"device_ids"`
		Play      bool     `json:"play"`
	}
	if err := json.Unmarshal([]byte(requests[0].body), &transferBody); err != nil {
		t.Fatalf("decode transfer body: %v", err)
	}
	if !reflect.DeepEqual(transferBody.DeviceIDs, []string{deviceID}) || transferBody.Play {
		t.Fatalf("transfer body = %+v, want selected device with play=false", transferBody)
	}

	if request := requests[1]; request.method != http.MethodPut || request.path != "/v1/me/player/play" {
		t.Fatalf("play request = %s %s, want PUT /v1/me/player/play", request.method, request.path)
	}
	if got := requests[1].authorization; got != "Bearer "+accessToken {
		t.Fatalf("play authorization = %q", got)
	}
	if got := requests[1].query; got != "device_id="+deviceID {
		t.Fatalf("play query = %q, want device_id=%s", got, deviceID)
	}
	var playBody struct {
		URIs []string `json:"uris"`
	}
	if err := json.Unmarshal([]byte(requests[1].body), &playBody); err != nil {
		t.Fatalf("decode play body: %v", err)
	}
	if !reflect.DeepEqual(playBody.URIs, []string{trackURI}) {
		t.Fatalf("played URIs = %v, want [%s]", playBody.URIs, trackURI)
	}
}

func TestRecoverQueuePlaybackDoesNotRecoverOtherErrors(t *testing.T) {
	originalErr := errors.New("Spotify is unavailable")
	called := false
	startedSelectedTrack, err := recoverQueuePlayback(
		originalErr,
		func() (string, error) { called = true; return "", nil },
		func(string, bool) error { called = true; return nil },
		func(string) error { called = true; return nil },
	)
	if startedSelectedTrack {
		t.Fatal("recoverQueuePlayback() reported playback for an unrelated error")
	}
	if err != originalErr {
		t.Fatalf("recoverQueuePlayback() error = %v, want original error %v", err, originalErr)
	}
	if called {
		t.Fatal("recovery actions ran for an unrelated error")
	}
}

func TestRecoverQueuePlaybackStopsWhenDeviceSelectionFails(t *testing.T) {
	selectionErr := errors.New("device selection canceled")
	transferCalled := false
	playCalled := false
	startedSelectedTrack, err := recoverQueuePlayback(
		noActiveDeviceError(),
		func() (string, error) { return "", selectionErr },
		func(string, bool) error { transferCalled = true; return nil },
		func(string) error { playCalled = true; return nil },
	)
	if startedSelectedTrack {
		t.Fatal("recoverQueuePlayback() reported playback after selection failed")
	}
	if err != selectionErr {
		t.Fatalf("recoverQueuePlayback() error = %v, want selection error %v", err, selectionErr)
	}
	if transferCalled || playCalled {
		t.Fatalf("transfer called = %t, play called = %t after selection failed", transferCalled, playCalled)
	}
}
