package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kurenkoff/bstu_backend/model"

	migrate "github.com/rubenv/sql-migrate"
)

const InsertSentGame = `
	insert into sent_games(id_game, id_sender)
					values(?, ?)
`

type SQL struct {
	Conn *sqlx.DB
}

// New возвращает новый экземпляр типа SQL, применяет миграции
func New(DSN string) *SQL {
	DB, err := sqlx.Connect("mysql", DSN)
	for err != nil {
		DB, err = sqlx.Connect("mysql", DSN)
	}

	sql := &SQL{Conn: DB}
	sql.migrate()
	return sql
}

// migrate применяет миграции описанные в data/migrations/
func (s *SQL) migrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "data/migrations",
	}

	n, err := migrate.Exec(s.Conn.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

// InsertSendGame дописывает в конец таблицы sent_games объект game
func (s *SQL) InsertSendGame(game model.SentGame) (int64, error) {
	res, err := s.Conn.Exec(InsertSentGame, game.GameID, game.SenderID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
