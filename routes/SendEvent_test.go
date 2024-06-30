package routes

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestSendEvent(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Testcase 1",
			args: args{
				c: &fiber.Ctx{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendEvent(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SendEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
