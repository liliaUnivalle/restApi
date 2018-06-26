package main

import (
    "database/sql"//bd

    "fmt"
    "log"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
 )

type Aplicacion struct {
    Router *mux.Router
    BD     *sql.DB
}

func (a *Aplicacion) Inicializar(usuario, password, dbnombre string) { //Inicializa los datos
	datosConexion :=
        fmt.Sprintf("user=%s password=%s dbname=%s", usuario, password, dbnombre)

    var err error
    a.BD, err = sql.Open("mysql", datosConexion)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
}

func (a *Aplicacion) Correr(addr string) { //corre la aplicación
}