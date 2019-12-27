package model

type Game struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Video       string `json:"video" db:"video"`
	Icon        string `json:"icon" db:"icon"`
	Background  string `json:"background" db:"background"`
	Description string `json:"description" db:"description"`
}

type Sender struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	SkypeID string `json:"skype_id" db:"skype_id"`
	Email   string `json:"email" db:"email"`
}

type SentGame struct {
	ID       int   `json:"id" db:"id"`
	GameID   int64 `json:"id_game" db:"id_game"`
	SenderID int64 `json:"id_sender" db:"id_sender"`
}

type AddRequest struct {
	Sender Sender `json:"sender"`
	Game   Game   `json:"game"`
}
