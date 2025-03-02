# xk6-sql-driver-oracle

Database driver extension for [xk6-sql](https://github.com/grafana/xk6-sql) k6 extension to support Oracle database. 
Uses [go-ora](https://github.com/sijms/go-ora).

## Example

```JavaScript file=examples/example.js
import sql from "k6/x/sql";
import driver from "k6/x/sql/driver/oracle";

const db = sql.open(driver, "oracle://oracle:oracle@localhost:1521/FREEPDB1");

export function setup() {
    db.exec(`
        CREATE TABLE roster
        (
            id          NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            given_name  VARCHAR2(255) NOT NULL,
            family_name VARCHAR2(255) NOT NULL
        )
    `);
}

export function teardown() {
    db.exec("DROP TABLE roster PURGE");
    db.close();
}

export default function () {
    let inserted = 0;
    const inserts = [
        "INSERT INTO roster (given_name, family_name) VALUES ('Peter', 'Pan')",
        "INSERT INTO roster (given_name, family_name) VALUES ('Wendy', 'Darling')",
        "INSERT INTO roster (given_name, family_name) VALUES ('Tinker', 'Bell')",
        "INSERT INTO roster (given_name, family_name) VALUES ('James', 'Hook')",
    ];
    for (const insertion of inserts) {
        const result = db.exec(insertion);
        inserted += result.rowsAffected();
    }
    console.log(`${inserted} rows inserted`);

    let rows = db.query("SELECT * FROM roster WHERE given_name = :1", "Peter");
    for (const row of rows) {
        console.log(`${row.FAMILY_NAME}, ${row.GIVEN_NAME}`);
    }
}
```

## Usage
Check the [xk6-sql documentation](https://github.com/grafana/xk6-sql) on how to use this database driver.