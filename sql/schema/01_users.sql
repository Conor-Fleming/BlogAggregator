-- +goose up
CREATE TABLE users (
    id int,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name text,
    PRIMARY KEY(id)
)

-- +goose down
DROP TABLE users;