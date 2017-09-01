package ArchivoModel

import (
	"html/template"

	"gopkg.in/mgo.v2/bson"
)

//#########################< ESTRUCTURAS >##############################

//EKeyArchivo Estructura de campo de Archivo
type EKeyArchivo struct {
	Key      string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ECerArchivo Estructura de campo de Archivo
type ECerArchivo struct {
	Cer      string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EPemArchivo Estructura de campo de Archivo
type EPemArchivo struct {
	Pem      string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//Archivo estructura de Archivos mongo
type Archivo struct {
	ID bson.ObjectId
	EKeyArchivo
	ECerArchivo
	EPemArchivo
}

//SSesion estructura de variables de sesion de Usuarios del sistema
type SSesion struct {
	Name          string
	MenuPrincipal template.HTML
	MenuUsr       template.HTML
}

//SIndex estructura de variables de index
type SIndex struct {
	SResultados bool
	SRMsj       string
	SCabecera   template.HTML
	SBody       template.HTML
	SPaginacion template.HTML
	SGrupo      template.HTML
}

//SArchivo estructura de Archivo para la vista
type SArchivo struct {
	SEstado bool
	SMsj    string
	Archivo
	SIndex
	SSesion
}
