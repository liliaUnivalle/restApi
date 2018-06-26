
package main_test

import (
    "os"
    "testing"
    "log"

    "net/http"
    "net/http/httptest"
    "encoding/json"
    "bytes"
    "strconv"
    "fmt"

    "."
)

var a main.Aplicacion

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
    if _, err := a.BD.Exec(queryCreacionTabla); err != nil {
        log.Fatal(err)
    }
}

func limpiarTabla() {
    a.BD.Exec("DELETE FROM recetas")
    a.BD.Exec("ALTER TABLE recetas AUTO_INCREMENT = 1")
}


func TestMain(m *testing.M) {
    a = main.Aplicacion{}

    a.Inicializar("go", "password", "Go")

    asegurarExistenciaTabla()

    code := m.Run()

    limpiarTabla()

    os.Exit(code)
}



func ejecutarRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    return rr
}

func verificarCodigoResponsive(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("se esperaba codigo de respuesta %d. obtuvo %d\n", expected, actual)
    }
}


//Para comprobar si una tabla está vacía
func TestTablaVacia(t *testing.T) {
    limpiarTabla()

    req, _ := http.NewRequest("GET", "/recetas", nil)//solicitud get
    response := ejecutarRequest(req)

    verificarCodigoResponsive(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "[]" {
        t.Errorf("Esperaba un array vacío y se obtuvo: %s", body)
    }
}


//Prueba de receta no existente

func TestObtenerRecetaNoExistente(t *testing.T) {
    limpiarTabla()
    req, _ := http.NewRequest("GET", "/receta/45", nil)
    response := ejecutarRequest(req)
    verificarCodigoResponsive(t, http.StatusNotFound, response.Code)
    var m map[string]string//______________???
    json.Unmarshal(response.Body.Bytes(), &m)
    if m["error"] != "Receta no encontrada" {
        t.Errorf("Se esperaba error para decir que la receta no existe pero se obtuvo: '%s'", m["error"])
    }
}

//Prueba de la creación de una receta

func TestCrearReceta(t *testing.T) {
    limpiarTabla()

    //for i:=0; i<10; i++ {
        
    payload := []byte(`{"nombre":"test receta","ingredientes":"test receta", "preparacion":"test receta", "tiempo":"30min"}`)

    req, _ := http.NewRequest("POST", "/receta", bytes.NewBuffer(payload))
    response := ejecutarRequest(req)

    verificarCodigoResponsive(t, http.StatusCreated, response.Code)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    if m["nombre"] != "test receta" {
        t.Errorf("Esperaba una receta con el nombre 'test receta' y obtuvo: '%v'", m["nombre"])
    }

    if m["ingredientes"] != "test receta" {
        t.Errorf("Esperaba una receta con ingredientes'test receta' y obtuvo: '%v'", m["ingredientes"])
    }

    if m["preparacion"] != "test receta" {
      
        t.Errorf("Esperaba una receta con preparacion 'test receta' y obtuvo: '%v'", m["preparacion"])
    }
    if m["tiempo"] != "30min" {
        t.Errorf("Esperaba una receta con tiempo '30min' y obtuvo: '%v'", m["tiempo"])
    }

    // // el id se compara con 1.0 porque Unmarshal de JSON convierte los números en flotantes, cuando el objetivo es un map[string] {}
    if m["id"] != 1.0 {
        t.Errorf("Esperaba una receta con '1'y obtuvo:'%v'", m["id"])
        }
    //}
}

//test para la obtención de un producto

func TestObtenerReceta(t *testing.T) {
    limpiarTabla()
    agregarRecetas(1)

    req, _ := http.NewRequest("GET", "/receta/1", nil)
    response := ejecutarRequest(req)

    verificarCodigoResponsive(t, http.StatusOK, response.Code)
}

func agregarRecetas(count int) {
    if count < 1 {
        count = 1
    }
    
    for i := 0; i < count; i++ {

        statement := fmt.Sprintf("INSERT INTO recetas(nombre, ingredientes, preparacion, tiempo) VALUES('%s', '%s','%s','%s')", ("User " + strconv.Itoa(i+1)), ("Ingredientes " + strconv.Itoa(i+1)), ("Preparacion " + strconv.Itoa(i+1)),strconv.Itoa(((i+1) * 10)))
        a.BD.Exec(statement)
    }
}

//Prueba para actualizar un producto
func TestActualizarProducto(t *testing.T) {
    limpiarTabla()
    agregarRecetas(1)

    req, _ := http.NewRequest("GET", "/receta/1", nil)
    response := ejecutarRequest(req)
    var originalProduct map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &originalProduct)

    payload := []byte(`{"nombre":"test receta - actualizado","ingredientes":"ingredientes actualizado", "preparacion":"preparacion actualizado", "tiempo":"20min"}`)

    req, _ = http.NewRequest("PUT", "/receta/1", bytes.NewBuffer(payload))
    response = ejecutarRequest(req)

    verificarCodigoResponsive(t, http.StatusOK, response.Code)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    if m["id"] != originalProduct["id"] {
        t.Errorf("se espera que se mantengan igual (%v). se obtuvo %v", originalProduct["id"], m["id"])
    }

    if m["nombre"] == originalProduct["nombre"] {
        t.Errorf("se esperaba que el nombre cambiara '%v' por '%v'. se obtuvo '%v'", originalProduct["nombre"], m["name"], m["name"])
    }

    
    if m["ingredientes"] == originalProduct["ingredientes"] {
        t.Errorf("se esperaba que los ingredientes cambiaran '%v' por '%v'. se obtuvo '%v'", originalProduct["ingredientes"], m["name"], m["name"])
    }

    if m["preparacion"] == originalProduct["preparacion"] {
        t.Errorf("se esperaba que la preparacion cambiaran '%v' por '%v'. se obtuvo '%v'", originalProduct["preparacion"], m["name"], m["name"])
    }

    if m["tiempo"] == originalProduct["tiempo"] {
        t.Errorf("se esperaba que el tiempo cambiara '%v' por '%v'. se obtuvo '%v'", originalProduct["tiempo"], m["name"], m["name"])
    }
}


// Prueba para eliminar una receta
func TestBorrarReceta(t *testing.T) {
    limpiarTabla()
    agregarRecetas(1)

    req, _ := http.NewRequest("GET", "/receta/1", nil)
    response := ejecutarRequest(req)
    verificarCodigoResponsive(t, http.StatusOK, response.Code)

    req, _ = http.NewRequest("DELETE", "/receta/1", nil)
    response = ejecutarRequest(req)

    verificarCodigoResponsive(t, http.StatusOK, response.Code)

    req, _ = http.NewRequest("GET", "/receta/1", nil)
    response = ejecutarRequest(req)
    verificarCodigoResponsive(t, http.StatusNotFound, response.Code)
}

