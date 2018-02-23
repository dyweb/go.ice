# Postgres

Table

- don't use quote around table name
  - https://stackoverflow.com/questions/6331504/omitting-the-double-quote-to-do-query-on-postgresql
  
Query

- place holder `$1`, `$2` ...

Shell

- `\c icehub` to use a database
- `\dt` show tables

Drivers

- https://github.com/jackc/pgx offers a native interface similar to database/sql that offers better performance and more features
  - supports JSON and JSONB