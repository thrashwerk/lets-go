### Running MySQL in a container instead of the host

##### Connecting to the database
```bash
# from the host
mysql -u root -p -h 127.0.0.1

mysql -u web -p -h 127.0.0.1 -D snippetbox

# from the container
podman exec -it snippetbox-mysql mysql -u root -p

podman exec -it snippetbox-mysql mysql -u web -p
```

##### Setting up the container
```bash
podman volume create snippetbox-mysql-data
podman run --name snippetbox-mysql --network=host --volume snippetbox-mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=pasas123 -d mysql
```

##### Setting up the database, table, sample data and user
```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE snippetbox;

CREATE TABLE snippets (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Basho',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n- Natsume Soseki',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\n shows my father''s face.\n\n- Marukami Kijo',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

-- localhost is only for connecting through the socket
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pasas456';

-- this is for connecting to the container from the host
CREATE USER 'web'@'127.0.0.1';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'127.0.0.1';
ALTER USER 'web'@'127.0.0.1' IDENTIFIED BY 'pasas456';
```

##### Setup for stateful sessions
```sql
USE snippetbox;

CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```

##### Setup for user auth
```sql
USE snippetbox;

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```
