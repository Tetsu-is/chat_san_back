package service

import (
	"chat_san/model"
	"context"
	"database/sql"
)

type MessageService struct {
	db *sql.DB
}

func (s *MessageService) CreateMessage(ctx context.Context, text string) (*model.Message, error) {
	const (
		insert  = `INSERT INTO messages(text) VALUES(?)`
		confirm = `SELECT id, text, created_at, updated_at FROM messages WHERE id = ?`
	)

	if _, err := s.db.PrepareContext(ctx, insert); err != nil {
		return nil, err
	}

	res, err := s.db.ExecContext(ctx, insert, text)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, confirm, id)

	var msg model.Message
	if err := row.Scan(&msg.ID, &msg.Text, &msg.CreatedAt, &msg.UpdatedAt); err != nil {
		return nil, err
	}

	return &msg, nil
}

func (s *MessageService) ReadMessage(ctx context.Context, offset, limit int64) ([]*model.Message, error) {
	const (
		read = `SELECT id, text, created_at, updated_at FROM messages WHERE id > ? ORDER BY id LIMIT ?`
	)

	s.db.PrepareContext(ctx, read)

	messages := []*model.Message{}

	rows, err := s.db.QueryContext(ctx, read, offset, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var msg model.Message
		rows.Scan(&msg.ID, &msg.Text, &msg.CreatedAt, &msg.UpdatedAt)
		messages = append(messages, &msg)
	}

	return messages, nil
}
