-- +goose Up
-- +goose StatementBegin
CREATE DATABASE if NOT EXISTS careerise;
USE careerise;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE if EXISTS careerise;
-- +goose StatementEnd
