package routes

import (
	"testing"

	"github.com/Sandy143toce/webhook-multiplexer/models"
)

func TestCreateEndpointDb(t *testing.T) {
	type args struct {
		endpoint           models.Endpoint
		CustomerEndpointId string
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
				endpoint: models.Endpoint{
					WebhookID: "WH-123",
					URL:       "https://test.com",
				},
				CustomerEndpointId: "CE-123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateEndpointDb(tt.args.endpoint, tt.args.CustomerEndpointId); (err != nil) != tt.wantErr {
				t.Errorf("CreateEndpointDb() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
