package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//TODO use testcontainers
func TestProcessRepository_AddProcess(t *testing.T) {
	connStr := "user=postgres password=password dbname=trazavino sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	pr, _ := entity.NewProcess(time.Now(), time.Time{}, "", "", entity.TypeReception.String(), 0)
	ctx := context.Background()

	repo := NewProcessRepository(db)
	process, err := repo.AddProcess(ctx, pr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(process)

	rec, _ := entity.NewReception(process, nil, 5, 5)
	recRepo := NewReceptionRepository(db)
	reception, err := recRepo.AddReception(ctx, rec)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reception)
}
