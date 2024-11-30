package user_segment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"segment-manager/internal/store/segment"
)

type Executer interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type UserSegment struct {
	Executer Executer
}

func New(executer Executer) *UserSegment {
	return &UserSegment{
		Executer: executer,
	}
}

func (s *UserSegment) UpsertUserSegments(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error {
	tx, err := s.Executer.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("для userID %d ошибка при начале транзакции: %w", userID, err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Получение ID сегментов пользователя
	query, args, err := buildSelectQuery(userID)
	if err != nil {
		return fmt.Errorf("ошибка при построении SQL-запроса: %w", err)
	}
	rows, err := s.Executer.QueryContext(ctx, query, args...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("ошибка при получении текущих ID сегментов: %w", err)
	}

	currentIDs, err := scanRowsToIDs(rows)
	if err != nil {
		return err
	}

	// Получение ID сегментов для добавления
	addSegmentIDs, err := s.getSegmentIDs(ctx, slugsToAdd)
	if err != nil {
		return fmt.Errorf("ошибка при получении ID сегментов для добавления: %w", err)
	}

	// Получение ID сегментов для удаления
	deleteSegmentIDs, err := s.getSegmentIDs(ctx, slugsToDelete)
	if err != nil {
		return fmt.Errorf("ошибка при получении ID сегментов для удаления: %w", err)
	}

	query, args, err = buildUpsertQuery(userID, getUpdatedIDs(currentIDs, addSegmentIDs, deleteSegmentIDs))
	_, err = s.Executer.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении сегментов пользователя: %w", err)
	}

	return nil
}

func (s *UserSegment) SelectUserSegments(ctx context.Context, userID int64) ([]UserSegmentDB, error) {
	query, args, err := buildSelectUserSegmentsQuery(userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при построении SQL-запроса: %w", err)
	}

	rows, _ := s.Executer.QueryContext(ctx, query, args...)

	return scanRowsToSegments(rows)
}

// Вспомогательная функция для получения списка ID сегментов по слагам
func (s *UserSegment) getSegmentIDs(ctx context.Context, slugs []string) ([]int64, error) {
	if len(slugs) == 0 {
		return nil, nil // Нет слагов для обработки
	}

	query, args, err := segment.BuildSelectQuery(slugs)
	if err != nil {
		return nil, fmt.Errorf("ошибка при построении SQL-запроса: %w", err)
	}

	rows, err := s.Executer.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	return scanRowsToIDs(rows)
}

// scanRowsToIDs сканирует *sql.Rows и возвращает срез ID
func scanRowsToIDs(rows *sql.Rows) ([]int64, error) {
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по строкам: %w", err)
	}

	return ids, nil
}

// scanRowsToIDs сканирует *sql.Rows и возвращает срез Segments
func scanRowsToSegments(rows *sql.Rows) ([]UserSegmentDB, error) {
	defer rows.Close()

	userSegments := make([]UserSegmentDB, 0)
	for rows.Next() {
		userSegment := UserSegmentDB{}
		if err := rows.Scan(&userSegment.ID, &userSegment.Slug); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		userSegments = append(userSegments, userSegment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по строкам: %w", err)
	}

	return userSegments, nil
}

func getUpdatedIDs(currentIDs, addIDs, deleteIDs []int64) []int64 {
	updatedIDs := make(map[int64]struct{})

	for _, v := range currentIDs {
		updatedIDs[v] = struct{}{}
	}

	for _, v := range addIDs {
		updatedIDs[v] = struct{}{}
	}

	for _, v := range deleteIDs {
		delete(updatedIDs, v)
	}

	result := make([]int64, 0, len(updatedIDs))
	for k := range updatedIDs {
		result = append(result, k)
	}

	return result
}
