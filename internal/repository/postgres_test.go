package repository_test

import (
	"context"
	"database/sql"
	"os"
	"test_for_goforge/internal/repository"
	"testing"

	_ "github.com/lib/pq"
)

func setupTestDB(t *testing.T) *sql.DB {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestAddNumberAndGetAllSorted(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.New(db)
	ctx := context.Background()

	// Очистка таблицы перед тестом
	_, err := db.ExecContext(ctx, "TRUNCATE TABLE numbers")
	if err != nil {
		t.Fatal(err)
	}

	testNumbers := []int{5, 3, 8, 1}
	for _, n := range testNumbers {
		if err := repo.AddNumber(ctx, n); err != nil {
			t.Fatalf("AddNumber error: %v", err)
		}
	}

	nums, err := repo.GetAllSorted(ctx)
	if err != nil {
		t.Fatalf("GetAllSorted error: %v", err)
	}

	expected := []int{1, 3, 5, 8}
	if len(nums) != len(expected) {
		t.Fatalf("expected %d numbers, got %d", len(expected), len(nums))
	}
	for i, v := range expected {
		if nums[i] != v {
			t.Errorf("expected nums[%d] = %d, got %d", i, v, nums[i])
		}
	}
}
