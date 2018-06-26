#Api rest en Go

clonar este proyecto 

ejecutar la instrucción:

-- $export GOBIN=(direccion)restApi/

Donde direccion = ubicación de la carpeta

Obtener dependencias mux y mysql

-- restApi$ go get github.com/gorilla/mux
-- restApi$ go get "github.com/go-sql-driver/mysql"


Para ejecutar las pruebas, deben pasarse los datos de la bd en el llamado a la funcion inicializar en main_test.co

Para obtener ejecutable

-- restApi$ go get 

Ejecutar pruebas

-- restApi$ go test -v

Para correr el proyecto sin las pruebas
 
-- restApi$ ./restApi
