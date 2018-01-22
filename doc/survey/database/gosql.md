# Go database/sql package

## Basic usage

- `dbHandle, err := sql.Open("drivername", "dsn")`
  - it does NOT establish real connection
  - [ ] we should wrap this up so user won't know which exact driver we are using, and dsn is constructed from config struct
- `stmt, err := db.Prepare("INSERT INTO user (user, pwd) VALUES (?, ?)")`
  - [x] does the prepare communicate with database server? Yes 
    - https://github.com/go-sql-driver/mysql/blob/master/connection.go#L161:34
    - https://github.com/gocraft/dbr says for `db.Query` mysql driver would create a prepared statement and throw it away
  - [ ] different dialect for place holder `$` instead of `?` in postgres
    - [ ] I remember there is one package handling it
- `res, err := stmt.Exec("jack", "123")`
  - res only has `LastInsertId` and `RowsAffected` , need to use `Query` if rows are needed
- `rows, err := db.Qyuery("SELECT * FROM user)`
  - `for rows.Next()` `rows.Scan(&user, &pwd)` scan and update value, this is where ORM etc. jumps in
  - xo https://github.com/xo/xo/blob/master/examples/django/sqlite3/authgroup.xo.go
  - https://github.com/gocraft/dbr
  
````go
// database/sql.go

// QueryerContext is an optional interface that may be implemented by a Conn.
//
// If a Conn does not implement QueryerContext, the sql package's DB.Query will
// first prepare a query, execute the statement, and then close the
// statement.
//
// QueryerContext may return ErrSkip.
//
// QueryerContext must honor the context timeout and return when the context is canceled.
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}
````

````go
// database/driver/driver.go
type StmtExecContext interface {
        // ExecContext executes a query that doesn't return rows, such
        // as an INSERT or UPDATE.
        //
        // ExecContext must honor the context timeout and return when it is canceled.
        ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}
type StmtQueryContext interface {
        // QueryContext executes a query that may return rows, such as a
        // SELECT.
        //
        // QueryContext must honor the context timeout and return when it is canceled.
        QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}
````

## Prepared statement

- it seems there are three round trips, prepare, query, remove prepared
  - [x] I need to call `stmt.Close` right? It will send command to server to close statement
    - https://github.com/go-sql-driver/mysql/blob/master/statement.go#L25
- http://go-database-sql.org/prepared.html
  - [ ] https://www.vividcortex.com/blog/2014/11/19/analyzing-prepared-statement-performance-with-vividcortex/
- [ ] place holder dialect

## Error

- http://go-database-sql.org/errors.html
  - iterate result, `sql.ErrNoRows`
  - driver specific module
- http://go-database-sql.org/surprises.html
  - multiple statement support

when running new query when reading query result, new connection will be used

````go
rows, err := db.Query("select * from tbl1") // Uses connection 1
for rows.Next() {
	err = rows.Scan(&myvariable)
	// The following line will NOT use connection 1, which is already in-use
	db.Query("select * from tbl2 where id = ?", myvariable)
}
````

transactions are bound to just one connection, so you can no longer run new query without finish reading existing query

````go
tx, err := db.Begin()
rows, err := tx.Query("select * from tbl1") // Uses tx's connection
for rows.Next() {
	err = rows.Scan(&myvariable)
	// ERROR! tx's connection is already busy!
	tx.Query("select * from tbl2 where id = ?", myvariable)
}
````

## Ref

- https://github.com/VividCortex/go-database-sql-tutorial
  - Unknown columns http://go-database-sql.org/varcols.html useful when building shell and admin UI
- https://www.vividcortex.com/blog/2015/09/22/common-pitfalls-go/
- https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.0.html