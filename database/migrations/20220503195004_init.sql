-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL PRIMARY KEY,
    username VARCHAR ( 50 ),
    firstname VARCHAR ( 255 ),
    lastname VARCHAR ( 255 )
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
