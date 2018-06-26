package main

import (
    "database/sql"
    "fmt"
)

//Definición del modelo de Receta
type Receta struct{
	Id int           `json:"id"`
	Nombre string         `json:"nombre"`
	Ingredientes string   `json:"ingredientes"`
	Preparacion string  `json:"preparacion"`
	Tiempo string            `json:"tiempo"`
}

//Campos anónimos correspondientes a la estructura receta 
func (r *Receta) obtenerReceta(bd *sql.DB) error {
  statement := fmt.Sprintf("SELECT nombre, ingredientes,preparacion,tiempo FROM recetas WHERE id=%d", r.Id)
  return bd.QueryRow(statement).Scan(&r.Nombre, &r.Ingredientes, &r.Preparacion, &r.Tiempo)
}

func (r *Receta) actualizarReceta(bd *sql.DB) error {
  statement := fmt.Sprintf("UPDATE recetas SET  nombre='%s', ingredientes='%s',preparacion='%s',tiempo='%s' WHERE id=%d", r.Nombre, r.Ingredientes, r.Preparacion, r.Tiempo, r.Id)
  _, err := bd.Exec(statement)
  return err
}

func (r *Receta) eliminarReceta(bd *sql.DB) error {
  statement := fmt.Sprintf("DELETE FROM recetas WHERE id=%d", r.Id)
  _, err := bd.Exec(statement)
  return err
}

func (r *Receta) crearReceta(bd *sql.DB) error {
  statement := fmt.Sprintf("INSERT INTO recetas(nombre,ingredientes,preparacion,tiempo) VALUES('%s','%s','%s','%s')", r.Nombre, r.Ingredientes, r.Preparacion, r.Tiempo)
    _, err := bd.Exec(statement)
    if err != nil {
        return err
    }
    err = bd.QueryRow("SELECT LAST_INSERT_ID()").Scan(&r.Id)
    if err != nil {
        return err
    }
    return nil
}

//Método que obtiene la lista de todas la recetas existentes

func obtenerRecetas(bd *sql.DB, start, cont int) ([]Receta, error) {
  statement := fmt.Sprintf("SELECT id, nombre, Ingredientes, preparacion, tiempo FROM recetas LIMIT %d OFFSET %d", cont, start)
    rows, err := bd.Query(statement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    recetas := []Receta{}
    for rows.Next() {
        var r Receta
        if err := rows.Scan(&r.Id, &r.Nombre, &r.Ingredientes, &r.Preparacion, &r.Tiempo); err != nil {
            return nil, err
        }
        recetas = append(recetas, r)
    }
    return recetas, nil
}