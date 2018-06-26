package main

import (
    "database/sql"//bd
    "net/http"
    "encoding/json"

    "fmt"
    "log"
    "strconv"
    "os"
    "bufio"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
 )

type Aplicacion struct {
    Router *mux.Router
    BD     *sql.DB
}

func (a *Aplicacion) Inicializar(usuario, password, dbnombre string) { //Inicializa los datos
	datosConexion :=
        fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", usuario, password, dbnombre)

    var err error
    a.BD, err = sql.Open("mysql", datosConexion)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
    a.creacionUrls()
}

//Métodos controladores 

func (a *Aplicacion) Correr(addr string) { //corre la aplicación
    log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *Aplicacion) ObtenerDatos() (nombre, nombreUsuario, password string){
    reader:= bufio.NewReader(os.Stdin)
    fmt.Println("Ingresa tu el nombre de la base de datos: ")
    nombre, err := reader.ReadString('\n')
    if err != nil{
     fmt.Println(err)
    }

    fmt.Println("Ingresa el nombre de usuario: ")
    nombreUsuario, err2 := reader.ReadString('\n')
    if err2 != nil{
     fmt.Println(err)
    }

    fmt.Println("Ingresa tu password: ")
    password, err3 := reader.ReadString('\n')
    if err3 != nil{
     fmt.Println(err)
    }

    return nombre, nombreUsuario, password
}

func responderConJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func responderConError(w http.ResponseWriter, code int, message string) {
    responderConJSON(w, code, map[string]string{"error": message})
}


func (a *Aplicacion) obtenerReceta(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        responderConError(w, http.StatusBadRequest, " ID de receta inválido")
        return
    }

    re := Receta{Id: id}
    if err := re.obtenerReceta(a.BD); err != nil {
        switch err {
        case sql.ErrNoRows:
            responderConError(w, http.StatusNotFound, "Receta no encontrada")
        default:
            responderConError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    responderConJSON(w, http.StatusOK, re)
}

func (a *Aplicacion) obtenerRecetas(w http.ResponseWriter, r *http.Request) {
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
        count = 10
    }
    if start < 0 {
        start = 0
    }

    recetas, err := obtenerRecetas(a.BD, start, count)
    if err != nil {
        responderConError(w, http.StatusInternalServerError, err.Error())
        return
    }

    responderConJSON(w, http.StatusOK, recetas)
}

//supone que el cuerpo de la solicitud en un json
func (a *Aplicacion) crearReceta(w http.ResponseWriter, r *http.Request) {
    var re Receta
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&re); err != nil {
        responderConError(w, http.StatusBadRequest, "invalida petición")
        return
    }
    defer r.Body.Close()

    if err := re.crearReceta(a.BD); err != nil {
        responderConError(w, http.StatusInternalServerError, err.Error())
        return
    }

    responderConJSON(w, http.StatusCreated, re)
}

func (a *Aplicacion) actualizarReceta(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        responderConError(w, http.StatusBadRequest, "ID de receta inválido")
        return
    }

    var re Receta
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&re); err != nil {
        responderConError(w, http.StatusBadRequest, "inválida solicitud")
        return
    }
    defer r.Body.Close()
    re.Id = id

    if err := re.actualizarReceta(a.BD); err != nil {
        responderConError(w, http.StatusInternalServerError, err.Error())
        return
    }

    responderConJSON(w, http.StatusOK, re)
}

func (a *Aplicacion) borrarReceta(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        responderConError(w, http.StatusBadRequest, "ID de Receta inválido")
        return
    }

    re := Receta{Id: id}
    if err := re.eliminarReceta(a.BD); err != nil {
        responderConError(w, http.StatusInternalServerError, err.Error())
        return
    }

    responderConJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *Aplicacion) creacionUrls() {
    a.Router.HandleFunc("/recetas", a.obtenerRecetas).Methods("GET")
    a.Router.HandleFunc("/receta", a.crearReceta).Methods("POST")
    a.Router.HandleFunc("/receta/{id:[0-9]+}", a.obtenerReceta).Methods("GET")
    a.Router.HandleFunc("/receta/{id:[0-9]+}", a.actualizarReceta).Methods("PUT")
    a.Router.HandleFunc("/receta/{id:[0-9]+}", a.borrarReceta).Methods("DELETE")
}




