package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/alejandroik/trazavino-api/internal/domain/entity/enum/process_type"
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

	pr, _ := entity.NewProcess(0, time.Now(), time.Time{}, "", "", process_type.Reception.String(), 0)
	ctx := context.Background()

	repo := NewProcessRepository(db)
	process, err := repo.AddProcess(ctx, pr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(process)

	tr, _ := entity.NewTruck(0, "EEOO-990")
	truckRepo := NewTruckRepository(db)
	truck, err := truckRepo.AddTruck(ctx, tr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(truck)

	vy, _ := entity.NewVineyard(0, "Los Andes")
	vyRepo := NewVineyardRepository(db)
	vineyard, err := vyRepo.AddVineyard(ctx, vy)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vineyard)

	gt, _ := entity.NewGrapeType(0, "Rosa")
	gtRepo := NewGrapeTypeRepository(db)
	grapeType, err := gtRepo.AddGrapeType(ctx, gt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(grapeType)

	rec, _ := entity.NewReception(process.ID(), truck.ID(), vineyard.ID(), grapeType.ID(), 5, 5)
	recRepo := NewReceptionRepository(db)
	reception, err := recRepo.AddReception(ctx, rec)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reception)
}
