DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id varchar(64) PRIMARY KEY,
    username varchar(50) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255),
    created_at integer DEFAULT EXTRACT(EPOCH FROM NOW()),
    UNIQUE(username),
    UNIQUE(email)
);

CREATE OR REPLACE FUNCTION create_user(id varchar(64), username varchar(50), email varchar(255), password varchar(255))
RETURNS VOID
LANGUAGE SQL
AS $$
    INSERT INTO users VALUES (id, username, email, password);
$$;

CREATE OR REPLACE FUNCTION get_user_by_email(_email varchar(255))
RETURNS TABLE (id varchar(64), username varchar(50), email varchar(255), password varchar(255), created_at integer)
LANGUAGE SQL
AS $$
    SELECT * FROM users WHERE email = _email LIMIT 1;
$$;