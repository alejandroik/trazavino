package sqlc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	gosql "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func TestProcessMysqlRepository_AddProcess(t *testing.T) {
	cfg := &gosql.Config{
		Addr:      fmt.Sprintf("%v:%v", "localhost", "3306"),
		DBName:    "trazavino",
		User:      "root",
		Passwd:    "",
		ParseTime: true,
	}
	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		t.Error(err)
	}

	pr, _ := entity.NewProcess(time.Now(), time.Time{}, "", "", entity.TypeReception.String(), 0)
	ctx := context.Background()

	repo := NewProcessMysqlRepository(db)
	process, err := repo.AddProcess(ctx, pr)
	if err != nil {
		t.Error(err)
	}
	t.Log(process)
}
