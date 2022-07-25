package sqlc

import (
	"context"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestAddBottling(t *testing.T) {
	connStr := "user=postgres password=password dbname=trazavino sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	require.NoError(t, err)

	ctx := context.Background()

	winery, err := entity.NewWinery(uuid.NewString(), "Budeger")
	require.NoError(t, err)
	wnRepo := NewWineryRepository(db)
	err = wnRepo.AddWinery(ctx, winery)
	require.NoError(t, err)

	tr, err := entity.NewTruck(uuid.NewString(), "EEOO-990", winery.ID())
	require.NoError(t, err)
	truckRepo := NewTruckRepository(db)
	err = truckRepo.AddTruck(ctx, tr)
	require.NoError(t, err)

	vy, err := entity.NewVineyard(uuid.NewString(), "Los Andes", winery.ID())
	require.NoError(t, err)
	vyRepo := NewVineyardRepository(db)
	err = vyRepo.AddVineyard(ctx, vy)
	require.NoError(t, err)

	gt, err := entity.NewGrapeType(uuid.NewString(), "Rosa", winery.ID())
	require.NoError(t, err)
	gtRepo := NewGrapeTypeRepository(db)
	err = gtRepo.AddGrapeType(ctx, gt)
	require.NoError(t, err)

	rec, err := entity.NewReception(uuid.NewString(), time.Now(), winery.ID(), tr.ID(), vy.ID(), gt.ID(), 2500, 5)
	require.NoError(t, err)
	recRepo := NewReceptionRepository(db)
	err = recRepo.AddReception(ctx, rec)
	require.NoError(t, err)

	wh, err := entity.NewWarehouse(uuid.NewString(), "AAA#4322", true, winery.ID())
	require.NoError(t, err)
	whRepo := NewWarehouseRepository(db)
	err = whRepo.AddWarehouse(ctx, wh)
	require.NoError(t, err)

	mc, err := entity.NewMaceration(uuid.NewString(), time.Now(), winery.ID(), rec.UUID(), wh.ID())
	require.NoError(t, err)
	mcRepo := NewMacerationRepository(db)
	err = mcRepo.AddMaceration(ctx, mc)
	require.NoError(t, err)

	tk, err := entity.NewTank(uuid.NewString(), "B3224FF", true, winery.ID())
	require.NoError(t, err)
	tkRepo := NewTankRepository(db)
	err = tkRepo.AddTank(ctx, tk)
	require.NoError(t, err)

	fr, err := entity.NewFermentation(uuid.NewString(), time.Now(), winery.ID(), wh.ID(), tk.ID())
	require.NoError(t, err)
	frRepo := NewFermentationRepository(db)
	err = frRepo.AddFermentation(ctx, fr)
	require.NoError(t, err)

	ck, err := entity.NewCask(uuid.NewString(), "C#4432", "Roble", true, winery.ID())
	require.NoError(t, err)
	ckRepo := NewCaskRepository(db)
	err = ckRepo.AddCask(ctx, ck)
	require.NoError(t, err)

	age, err := entity.NewAgeing(uuid.NewString(), time.Now(), winery.ID(), tk.ID(), ck.ID())
	require.NoError(t, err)
	ageRepo := NewAgeingRepository(db)
	err = ageRepo.AddAgeing(ctx, age)
	require.NoError(t, err)

	wine, err := entity.NewWine(uuid.NewString(), "4000 Black", winery.ID())
	require.NoError(t, err)
	wineRepo := NewWineRepository(db)
	err = wineRepo.AddWine(ctx, wine)
	require.NoError(t, err)

	bot, err := entity.NewBottling(uuid.NewString(), time.Now(), winery.ID(), ck.ID(), wine.ID(), 100)
	require.NoError(t, err)
	botRepo := NewBottlingRepository(db)
	err = botRepo.AddBottling(ctx, bot)
	require.NoError(t, err)
}
