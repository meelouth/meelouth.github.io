package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kurenkoff/bstu_backend/model"
)

const (
	SelectSender = `
	select id, skype_id, name, email 
	from senders
	where id = ? limit 1
	`

	SelectSenders = `
	select id, skype_id, name, email 
	from senders
	`

	InsertSender = `
	insert into senders(skype_id, name, email)
				 values(:skype_id, :name, :email)
`
)

// GetSender возвращает объект Sender с id == id
func (s *SQL) GetSender(id string) (*model.Sender, error) {
	var sender model.Sender

	err := s.Conn.Select(&sender, SelectSender, id)
	if err != nil {
		return nil, err
	}

	return &sender, nil
}

// GetSenders возвращает всех Sender сохраненных в базе данных
func (s *SQL) GetSenders() (*[]model.Sender, error) {
	var senders []model.Sender

	err := s.Conn.Select(&senders, SelectSenders)
	if err != nil {
		return nil, err
	}

	return &senders, nil
}

func (s *SQL) InsertSender(sender model.Sender) (int64, error) {
	query, args, err := sqlx.Named(InsertSender, sender)
	if err != nil {
		return 0, err
	}

	res, err := s.Conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
