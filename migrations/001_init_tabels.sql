-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS users
(
    user_id   SERIAL PRIMARY KEY,
    username  varchar(80) UNIQUE  NOT NULL,
    name      varchar(256)        NOT NULL,
    email     varchar(256) UNIQUE NOT NULL,
    password  varchar(75)         NOT NULL,
    biography varchar(256)        NOT NULL,
    gender    smallint            NOT NULL
);
---- create above / drop below ----
DROP TABLE IF EXISTS users;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
