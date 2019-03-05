CREATE TABLE users
(
  id       BIGSERIAL PRIMARY KEY NOT NULL,
  username VARCHAR(2048)         NOT NULL,
  password VARCHAR(2048)         NOT NULL
);
CREATE UNIQUE INDEX idx_unique_users_username ON users (username);


CREATE TABLE sessions
(
  id     BIGSERIAL PRIMARY KEY NOT NULL,
  cookie VARCHAR(2048)         NOT NULL,
  expiry BIGINT                NOT NULL
);
CREATE UNIQUE INDEX idx_unique_sessions_cookie ON sessions (cookie);