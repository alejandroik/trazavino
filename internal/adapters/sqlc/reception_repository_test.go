package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestReceptionRepository_AddReception(t *testing.T) {
	connStr := "user=postgres password=password dbname=trazavino sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	require.NoError(t, err)

	ctx := context.Background()

	wn, err := entity.NewWinery(uuid.NewString(), "Budeger")
	require.NoError(t, err)
	wnRepo := NewWineryRepository(db)
	err = wnRepo.AddWinery(ctx, wn)
	require.NoError(t, err)

	tr, err := entity.NewTruck(uuid.NewString(), "EEOO-990", wn.ID())
	require.NoError(t, err)
	truckRepo := NewTruckRepository(db)
	err = truckRepo.AddTruck(ctx, tr)
	require.NoError(t, err)

	vy, err := entity.NewVineyard(uuid.NewString(), "Los Andes", wn.ID())
	require.NoError(t, err)
	vyRepo := NewVineyardRepository(db)
	err = vyRepo.AddVineyard(ctx, vy)
	require.NoError(t, err)

	gt, err := entity.NewGrapeType(uuid.NewString(), "Rosa", wn.ID())
	require.NoError(t, err)
	gtRepo := NewGrapeTypeRepository(db)
	err = gtRepo.AddGrapeType(ctx, gt)
	require.NoError(t, err)

	rec, err := entity.NewReception(uuid.NewString(), time.Now(), wn.ID(), tr.ID(), vy.ID(), gt.ID(), 2500, 5)
	require.NoError(t, err)
	recRepo := NewReceptionRepository(db)
	err = recRepo.AddReception(ctx, rec)
	require.NoError(t, err)
}
