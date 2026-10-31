package player

import (
	"errors"
	"testing"

	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
)

func spotifyNoActiveDeviceError() error {
	apiError := cmdTypes.SpotifyAPIError{}
	apiError.Detail.Error.Message = noActiveDeviceMessage
	return apiError
}

func TestRunWithDeviceFallbackRetriesAfterActivation(t *testing.T) {
	var commandCalls int
	var activationCalls int
	err := runWithDeviceFallback(func() error {
		commandCalls++
		if commandCalls == 1 {
			return spotifyNoActiveDeviceError()
		}
		return nil
	}, func() {
		activationCalls++
	})

	if err != nil {
		t.Fatalf("runWithDeviceFallback() error = %v, want nil", err)
	}
	if commandCalls != 2 {
		t.Errorf("command called %d times, want 2", commandCalls)
	}
	if activationCalls != 1 {
		t.Errorf("device activation called %d times, want 1", activationCalls)
	}
}

func TestRunWithDeviceFallbackDoesNotRetryOtherErrors(t *testing.T) {
	wantErr := errors.New("request failed")
	var commandCalls int
	var activationCalls int
	err := runWithDeviceFallback(func() error {
		commandCalls++
		return wantErr
	}, func() {
		activationCalls++
	})

	if !errors.Is(err, wantErr) {
		t.Fatalf("runWithDeviceFallback() error = %v, want %v", err, wantErr)
	}
	if commandCalls != 1 {
		t.Errorf("command called %d times, want 1", commandCalls)
	}
	if activationCalls != 0 {
		t.Errorf("device activation called %d times, want 0", activationCalls)
	}
}

func TestRunWithDeviceFallbackReturnsRetryError(t *testing.T) {
	var commandCalls int
	var activationCalls int
	wantErr := spotifyNoActiveDeviceError()
	err := runWithDeviceFallback(func() error {
		commandCalls++
		return spotifyNoActiveDeviceError()
	}, func() {
		activationCalls++
	})

	if err == nil || err.Error() != wantErr.Error() {
		t.Fatalf("runWithDeviceFallback() error = %v, want %v", err, wantErr)
	}
	if commandCalls != 2 {
		t.Errorf("command called %d times, want 2", commandCalls)
	}
	if activationCalls != 1 {
		t.Errorf("device activation called %d times, want 1", activationCalls)
	}
}
