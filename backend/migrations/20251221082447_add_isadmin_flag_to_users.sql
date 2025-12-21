-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD is_admin BOOL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP is_admin;
-- +goose StatementEnd
