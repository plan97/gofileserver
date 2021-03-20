package main

import (
	"context"
	"testing"
	"time"
)

func Test_runfs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())

					go func() {
						time.Sleep(5 * time.Second)
						cancel()
					}()

					return ctx
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runfs(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("runfs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
