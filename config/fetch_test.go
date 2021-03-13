package config

import "testing"

func TestConfig_Fetch(t *testing.T) {
	tests := []struct {
		name    string
		conf    *Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test 1",
			conf:    New(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.conf.Fetch(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Fetch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
