package command

import (
	"context"
	"testing"

	"github.com/alejandroik/trazavino-api/internal/domain/repository"
)

func TestRegisterReception(t *testing.T) {
	type fields struct {
		processRepository   repository.ProcessRepository
		receptionRepository repository.ReceptionRepository
	}
	type args struct {
		ctx context.Context
		cmd RegisterReception
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := registerReceptionHandler{
				processRepository:   tt.fields.processRepository,
				receptionRepository: tt.fields.receptionRepository,
			}
			if err := h.Handle(tt.args.ctx, tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
