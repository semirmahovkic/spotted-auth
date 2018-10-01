DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id varchar(64) PRIMARY KEY,
    username varchar(50) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255),
    created_at timestamp DEFAULT NOW(),
    UNIQUE(username),
    UNIQUE(email)
);

CREATE OR REPLACE FUNCTION create_user(id varchar(64), username varchar(50), email varchar(255), password varchar(255))
RETURNS VOID
LANGUAGE SQL
AS $$
    INSERT INTO users VALUES (id, username, email, password);
$$;