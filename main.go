
package main

	

func main() {
    a := Aplicacion{}
    
    nombre, nombreUsuario, password := a.ObtenerDatos()
    a.Inicializar(nombreUsuario, password, nombre)

    a.Correr(":8080")
}