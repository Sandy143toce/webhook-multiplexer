package routes

import (
	"testing"

	"github.com/Sandy143toce/webhook-multiplexer/models"
)

func TestCreateWebhookDb(t *testing.T) {
	type args struct {
		webhook   models.Webhook
		WebhookId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Testcase 1",
			args: args{
				webhook: models.Webhook{
					Name: "Test Webhook",
					URL:  "https://test.com",
				},
				WebhookId: "WH-123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateWebhookDb(tt.args.webhook, tt.args.WebhookId); (err != nil) != tt.wantErr {
				t.Errorf("CreateWebhookDb() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
