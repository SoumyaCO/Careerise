-- +goose Up
CREATE USER IF NOT EXISTS 'bgengine'@'%' IDENTIFIED BY 'password';
GRANT ALL ON careerise.* TO 'bgengine'@'%';

-- +goose Down
REVOKE ALL ON careerise.* FROM 'bgengine'@'%';
DROP USER 'bgengine'@'%';
