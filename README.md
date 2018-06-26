#Api rest en Go

clonar este proyecto 

ejecutar la instrucción:
<div class="highlight highlight-source-shell">
 <pre> 
 $export GOBIN=(direccion)restApi/
 </pre>
</div>
Donde direccion = ubicación de la carpeta

Obtener dependencias mux y mysql
<div class="highlight highlight-source-shell">
 <pre>
 restApi$ go get github.com/gorilla/mux
 restApi$ go get "github.com/go-sql-driver/mysql"
 </pre>
</div>

Para ejecutar las pruebas, deben pasarse los datos de la bd en el llamado a la funcion inicializar en main_test.go

Para obtener ejecutable
<div class="highlight highlight-source-shell">
 <pre>
 restApi$ go get 
 </pre>
</div>

Ejecutar pruebas
<div class="highlight highlight-source-shell">
 <pre>
-- restApi$ go test -v
</pre>
</div>

Para correr el proyecto sin las pruebas
 <div class="highlight highlight-source-shell">
 <pre>
-- restApi$ ./restApi
</pre>
</div>
