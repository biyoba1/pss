package repository

import (
	"fmt"
)

func (m *MessagePostgres) UpdateMessageProcessed(id int, processed bool) error {
	query := fmt.Sprintf("UPDATE %s SET processed=$1 WHERE id=$2", messagesTable)
	_, err := m.db.Exec(query, id, processed)
	return err
}
