package main

import (
    "os"
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("clave: "+os.Getenv("MYSQL_ENV_MYSQL_ROOT_PASSWORD"))
    fmt.Println("ip: "+os.Getenv("MYSQL_PORT_3306_TCP_ADDR"))
    fmt.Println("port: "+os.Getenv("MYSQL_PORT_3306_TCP_PORT"))
    dsn := "root:"+os.Getenv("MYSQL_ENV_MYSQL_ROOT_PASSWORD") +
            "@tcp("+os.Getenv("MYSQL_PORT_3306_TCP_ADDR") +
            ":"+os.Getenv("MYSQL_PORT_3306_TCP_PORT") + ")/demo_go"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Create DB
    // db, err = db.Exec("CREATE DATABASE demogo; USE demogo; CREATE TABLE SquareNum (number int(11) PRIMARY KEY, squareNumber varchar(20));")

    // Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

    // Prepare statement for reading data
    stmtOut, err := db.Prepare("SELECT squareNumber FROM squareNum WHERE number = ?")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()

    // Insert square numbers for 0-24 in the database
    for i := 0; i < 25; i++ {
        _, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }

    var squareNum int // we "scan" the result in here

    // Query the square-number of 13
    err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The square number of 13 is: %d", squareNum)

    // Query another number.. 1 maybe?
    err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The square number of 1 is: %d", squareNum)
}
