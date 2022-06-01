-- noinspection SqlNoDataSourceInspectionForFile

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS questions (
    id INT NOT NULL PRIMARY KEY,
    question text,
    answer text,
    authors text,
    comments text
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS questions;
-- +goose StatementEnd
