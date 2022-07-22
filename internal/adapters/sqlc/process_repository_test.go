package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/entity/enum/process_type"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

//TODO use testcontainers
func TestProcessRepository_AddProcess(t *testing.T) {
	connStr := "user=postgres password=password dbname=trazavino sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	require.NoError(t, err)

	ctx := context.Background()

	pr, _ := entity.NewProcess(uuid.New().String(), time.Now().Round(time.Second), process_type.Reception.String())
	repo := NewProcessRepository(db)
	err = repo.AddProcess(ctx, pr)
	require.NoError(t, err)
}
