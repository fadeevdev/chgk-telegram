-- +goose Up
-- +goose StatementBegin
CREATE TABLE [IF NOT EXISTS] users (
    id INT NOT NULL PRIMARY KEY,
    username VARCHAR ( 50 ),
    firstname VARCHAR ( 255 ),
    lastname VARCHAR ( 255 )
    );
CREATE TABLE users_top (
    user_id FOREIGN KEY (id)
        REFERENCES users (id),
    answered INT,
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE [IF EXISTS] users;
DROP TABLE [IF EXISTS] users_top
-- +goose StatementEnd
