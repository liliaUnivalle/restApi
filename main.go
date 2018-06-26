
package main

import "os"

func main() {
    a := Aplicacion{}
    a.Inicializar(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))

    a.Correr(":8080")
}