
package main

import (
    "os"
    "testing"
    "log"

    "."
)

var a main.app

//Query para la creación de la tabla recetas en la base de datos con los datos de la receta

const queryCreacionTabla = `CREATE TABLE IF NOT EXISTS recetas
(
id SERIAL,
nombre TEXT NOT NULL,
ingredientes TEXT NOT NULL,
preparacion TEXT NOT NULL,
tiempo TEXT NOT NULL,
CONSTRAINT recetas_pkey PRIMARY KEY (id)
)`

//Método que asefura la existencia de la tabla recetas en la base de datos

func asegurarExistenciaTabla() {
    if _, err := a.DB.Exec(queryCreacionTabla); err != nil {
        log.Fatal(err)
    }
}

func limpiarTabla() {
    a.DB.Exec("DELETE FROM recetas")
    a.DB.Exec("ALTER SEQUENCE recetas_id_seq RESTART WITH 1")
}

func TestMain(m *testing.M) {
    a = main.Aplicacion{}
    a.Inicializar(
        os.Getenv("TEST_DB_USERNAME"),
        os.Getenv("TEST_DB_PASSWORD"),
        os.Getenv("TEST_DB_NAME"))

    asegurarExistenciaTabla()

    code := m.Run()

    limpiarTabla()

    os.Exit(code)
}