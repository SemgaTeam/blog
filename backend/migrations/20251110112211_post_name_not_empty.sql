-- +goose Up
-- +goose StatementBegin
ALTER TABLE posts
ADD CONSTRAINT name_not_empty CHECK (trim(name) <> '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE posts
DROP CHECK name_not_empty;
-- +goose StatementEnd
