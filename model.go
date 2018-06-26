package main

import (
    "database/sql"
    "errors"
)

//Definición del modelo de Receta
type Receta struct{
	Id string             `json:"id"`
	Nombre string         `json:"nombre"`
	Ingredientes string   `json:"ingredientes"`
	Instrucciones string  `json:"instrucciones"`
	Tiempo int            `json:"tiempo"`
}

//Campos anónimos correspondientes a la estructura receta 
func (p *Receta) getReceta(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (p *Receta) updateReceta(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (p *Receta) deleteReceta(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (p *Receta) createReceta(db *sql.DB) error {
  return errors.New("Not implemented")
}

//Método que obtiene la lista de todas la recetas existentes

func getRecetas(db *sql.DB, start, count int) ([]Receta, error) {
  return nil, errors.New("Not implemented")
}