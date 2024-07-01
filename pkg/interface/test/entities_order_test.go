package entities_test

import (
	"fmt"
	"testing"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/assert"
)

func TestOrder_NextStatus(t *testing.T) {
	tests := []struct {
		name           string
		initialStatus  string
		expectedStatus string
		wantErr        bool
		expectedErrMsg string
	}{
		{
			name:           "move from New to Paid",
			initialStatus:  "New",
			expectedStatus: "Paid",
			wantErr:        false,
		},
		{
			name:           "already in Done status",
			initialStatus:  "Done",
			expectedStatus: "Done",
			wantErr:        true,
			expectedErrMsg: "order is already done",
		},
		{
			name:           "not correct status",
			initialStatus:  "test",
			expectedStatus: "test",
			wantErr:        true,
			expectedErrMsg: "invalid order status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &entities.Order{Status: tt.initialStatus}
			err := order.NextStatus()
			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(tt.expectedErrMsg)
				fmt.Println("err", err.Error())
				assert.Equal(t, tt.expectedErrMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStatus, order.Status)
			}
		})
	}
}

func TestOrder_NextPaidToDone(t *testing.T) {
	tests := []struct {
		name           string
		initialStatus  string
		expectedStatus string
		wantErr        bool
		expectedErrMsg string
	}{
		{
			name:           "move from Paid to Done",
			initialStatus:  "Paid",
			expectedStatus: "Done",
			wantErr:        false,
		},
		{
			name:           "invalid transition from New to Done",
			initialStatus:  "New",
			expectedStatus: "New",
			wantErr:        true,
			expectedErrMsg: "invalid order status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &entities.Order{Status: tt.initialStatus}
			err := order.NextPaidToDone()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStatus, order.Status)
			}
		})
	}
}
