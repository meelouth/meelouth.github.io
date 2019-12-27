-- +migrate Up

CREATE TABLE games(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    video VARCHAR(50) NOT NULL,
    icon VARCHAR(50),
    background VARCHAR(50),
    description VARCHAR(50) NOT NULL
);

CREATE TABLE senders(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    skype_id VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL,
    email VARCHAR(30) NOT NULL
);

CREATE TABLE sent_games(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    id_game bigint,
    id_sender bigint,
    FOREIGN KEY (id_game) REFERENCES games(id),
    FOREIGN KEY (id_sender) REFERENCES senders(id)
);

-- +migrate Down

DROP TABLE sent_games;
DROP TABLE senders;
DROP TABLE games;
