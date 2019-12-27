package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kurenkoff/bstu_backend/model"
)

const (
	SelectGame = `
	select id, name, video, icon, background, description 
	from games 
	where id = ? limit 1
`

	SelectGames = `
	select id, name, video, icon, background, description 
	from games 
	`

	InsertGame = `
	insert into games(name, video, icon, background, description)
	values(:name, :video, :icon, :background, :description)
	`
)

// GetGame возвращает объект Game с id == id
func (s *SQL) GetGame(id string) (*model.Game, error) {
	var games []model.Game

	err := s.Conn.Select(&games, SelectGame, id)
	if err != nil {
		return nil, err
	}

	if len(games) > 0 {
		return &games[0], nil
	}

	return nil, nil
}

// GetGames возвращает все игры, сохраненные в бд
func (s *SQL) GetGames() (*[]model.Game, error) {
	var games []model.Game

	err := s.Conn.Select(&games, SelectGames)
	if err != nil {
		return nil, err
	}

	return &games, nil
}

// InsertGame дописывает в конец таблицы games объект game
func (s *SQL) InsertGame(game model.Game) (int64, error) {
	query, args, err := sqlx.Named(InsertGame, game)
	if err != nil {
		return 0, err
	}

	res, err := s.Conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
