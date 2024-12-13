CREATE DATABASE orderdb;

\c orderdb

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- psql -U postgres -d orderdb -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'
