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

func TestFermentationRepository_AddFermentation(t *testing.T) {
	connStr := "user=postgres password=password dbname=trazavino sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	require.NoError(t, err)

	ctx := context.Background()

	tr, err := entity.NewTruck(uuid.NewString(), "EEOO-990")
	require.NoError(t, err)
	truckRepo := NewTruckRepository(db)
	err = truckRepo.AddTruck(ctx, tr)
	require.NoError(t, err)

	vy, err := entity.NewVineyard(uuid.NewString(), "Los Andes")
	require.NoError(t, err)
	vyRepo := NewVineyardRepository(db)
	err = vyRepo.AddVineyard(ctx, vy)
	require.NoError(t, err)

	gt, err := entity.NewGrapeType(uuid.NewString(), "Rosa")
	require.NoError(t, err)
	gtRepo := NewGrapeTypeRepository(db)
	err = gtRepo.AddGrapeType(ctx, gt)
	require.NoError(t, err)

	rec, err := entity.NewReception(uuid.NewString(), time.Now(), tr.ID(), vy.ID(), gt.ID(), 2500, 5)
	require.NoError(t, err)
	recRepo := NewReceptionRepository(db)
	err = recRepo.AddReception(ctx, rec)
	require.NoError(t, err)

	wh, err := entity.NewWarehouse(uuid.NewString(), "AAA#4322", true)
	require.NoError(t, err)
	whRepo := NewWarehouseRepository(db)
	err = whRepo.AddWarehouse(ctx, wh)
	require.NoError(t, err)

	mc, err := entity.NewMaceration(uuid.NewString(), time.Now(), rec.UUID(), wh.ID())
	require.NoError(t, err)
	mcRepo := NewMacerationRepository(db)
	err = mcRepo.AddMaceration(ctx, mc)
	require.NoError(t, err)

	tk, err := entity.NewTank(uuid.NewString(), "B3224FF", true)
	require.NoError(t, err)
	tkRepo := NewTankRepository(db)
	err = tkRepo.AddTank(ctx, tk)
	require.NoError(t, err)

	fr, err := entity.NewFermentation(uuid.NewString(), time.Now(), wh.ID(), tk.ID())
	require.NoError(t, err)
	frRepo := NewFermentationRepository(db)
	err = frRepo.AddFermentation(ctx, fr)
	require.NoError(t, err)
}