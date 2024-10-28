package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectDB() {
	var err error
	db, err = sql.Open("mysql", "root:@#$senha$#@@tcp(127.0.0.1:3306)/bdcrud")
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao verificar a conexão com o banco de dados:", err)
	}
	fmt.Println("Conexão com o banco de dados estabelecida.")
}
