-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS correct_answers (
    user_id int REFERENCES users (id) not null,
    answered_questions int[]
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS correct_answers;
-- +goose StatementEnd
