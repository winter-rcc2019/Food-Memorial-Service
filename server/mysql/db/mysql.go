package db

import (
	"database/sql"
	"fmt"
	"os"
    "strings"
    
	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error){
	connectionString := getConnectionString()
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
        return nil, err
    }
    return db, nil
}

func getParamString(param string, defaultValue string) string {
    env := os.Getenv(param)
    if env != "" {
        return env
    }
    return defaultValue
}

func getConnectionString() string {
    host := getParamString("MYSQL_DB_HOST", "winter-rcc_mysql_host")
    port := getParamString("MYSQL_PORT", "3306")
    user := getParamString("MYSQL_USER", "root")
    pass := getParamString("MYSQL_PASSWORD", "password")
    dbname := getParamString("MYSQL_DB", "food_memory")
    protocol := getParamString("MYSQL_PROTOCOL", "tcp")
    dbargs := getParamString("MYSQL_DBARGS", " ")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
	} else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
        user, pass, protocol, host, port, dbname, dbargs)
}