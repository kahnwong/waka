package wakatime

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  string
		wantErr bool
	}{
		{
			name:    "valid API key",
			apiKey:  "test-api-key",
			wantErr: false,
		},
		{
			name:    "empty API key",
			apiKey:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Error("NewClient() returned nil client")
			}
			if !tt.wantErr {
				if client.baseURL != "https://wakatime.com" {
					t.Errorf("NewClient() baseURL = %v, want https://wakatime.com", client.baseURL)
				}
				if client.client == nil {
					t.Error("NewClient() http client is nil")
				}
				if client.authorizationHeader == "" {
					t.Error("NewClient() authorization header is empty")
				}
			}
		})
	}
}
