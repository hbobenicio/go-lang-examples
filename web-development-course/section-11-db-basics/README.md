# section-11-db-basics

This project tries to implement good db use cases for postgresql and golang,
according to the lessons of the section.

## Lessons Learned

- `*sql.DB` Represents a connection pool to the database
- There are some ways to organize the access of the DB Pool
- The PostgreSQL doesn't support the LastInsertedID feature, because not it only works on tables with sequencies, but not always a sequence is required. You can get it using the SQL `RETURNING` Clause though.
- `func (*sql.DB) Prepare(query string)` sends the template query to the server for it to prepare for eventual parameters. This is **bound to the connection**. This can become a bottleneck in concurrent scenarious if connections could become busy or disconnected between the preparement and the parameter dispatch (which would cause a repreparation on a new connection)
- DB Prepared Statements in transactions are bound to transactions
- The syntax for placeholder parameters in prepared statements is database-specific
- Avoid DB Prepared Statements

## Links and Resources

- [Organizing DB Access](https://www.alexedwards.net/blog/organising-database-access)
- [More about Prepared Statements](http://go-database-sql.org/prepared.html)
