package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"test_for_goforge/internal/handler"
	"testing"

	"context"
	"errors"
)

// Мок репозитория для тестирования handler
type mockRepo struct {
	addNumberCalled bool
	numbers         []int
	addErr          error
	getErr          error
}

func (m *mockRepo) AddNumber(ctx context.Context, num int) error {
	m.addNumberCalled = true
	return m.addErr
}

func (m *mockRepo) GetAllSorted(ctx context.Context) ([]int, error) {
	return m.numbers, m.getErr
}

func TestAddNumberHandler_Success(t *testing.T) {
	repo := &mockRepo{numbers: []int{1, 2, 3}}
	h := handler.New(repo)

	reqBody := map[string]int{"number": 2}
	bodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/number", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddNumber(w, req)

	if !repo.addNumberCalled {
		t.Errorf("expected AddNumber to be called")
	}
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	var nums []int
	if err := json.NewDecoder(res.Body).Decode(&nums); err != nil {
		t.Fatalf("decode response error: %v", err)
	}
	if len(nums) != 3 || nums[0] != 1 || nums[2] != 3 {
		t.Errorf("unexpected numbers response: %v", nums)
	}
}

func TestAddNumberHandler_BadRequest(t *testing.T) {
	repo := &mockRepo{}
	h := handler.New(repo)

	req := httptest.NewRequest(http.MethodPost, "/number", bytes.NewReader([]byte(`invalid json`)))
	w := httptest.NewRecorder()

	h.AddNumber(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}
}

func TestAddNumberHandler_RepositoryError(t *testing.T) {
	repo := &mockRepo{addErr: errors.New("db error")}
	h := handler.New(repo)

	reqBody := map[string]int{"number": 1}
	bodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/number", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddNumber(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", res.StatusCode)
	}
}
