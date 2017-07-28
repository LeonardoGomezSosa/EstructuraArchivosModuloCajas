package Inicio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"../../Models/Catalogo"
	"../Variables"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DataM es una estructura que contiene los datos de configuración en el archivo cfg
var DataM = MoVar.CargaSeccionCFG(MoVar.SecMongo)

//InitDatosMongo inicializalos datos en mongo
func InitDatosMongo() bool {

	//GetColectionMgo regresa una colección específica de Mongo
	host, err := mgo.Dial(DataM.Servidor)
	if err != nil {
		fmt.Println("Error al conectar con el servidor")
		return false
	}

	db := host.DB(DataM.NombreBase)
	// numUsers, err := db.C("Usuario").Find(bson.M{}).Count()
	// if numUsers == 0 {
	// 	var Usuario UsuarioModel.UsuarioMgo
	// 	Usuario.ID = bson.NewObjectId()
	// 	Usuario.Usuario = "administrador"
	// 	Usuario.Credenciales.Contraseña = "administrador"
	// 	Usuario.IsAdmin = true
	// 	db.C("Usuario").Insert(Usuario)
	// }

	num, err := db.C("Catalogo").Find(bson.M{}).Count()

	if err != nil {
		fmt.Println("Error al realizar la consulta siobre la coleccion catalogos")
		return false
	}

	if num == 0 {
		fmt.Println("Inicio la Incercion de Catalogos")
		sheetData, err := ioutil.ReadFile("Public/Resources/Sistema/Catalogos.json")
		if err != nil {
			log.Fatalln(err)
		}
		var catalogos []CatalogoModel.CatalogoMgo
		err = json.Unmarshal(sheetData, &catalogos)

		if err != nil {
			log.Fatalln(err)
		}

		for _, te := range catalogos {
			db.C("Catalogo").Insert(te)
		}
		fmt.Println("Insercion de catalogos correctamente")
	}

	Estado, err := db.C("Estados").Find(bson.M{}).Count()
	if err != nil {
		fmt.Println("Error al Realizar la consulta Sobre la Coleccion de Estados")
		return false
	}

	if Estado == 0 {
		fmt.Println("Inicio insercion de estados")
		sheetData, err := ioutil.ReadFile("Public/Resources/Sistema/Estados.json")
		if err != nil {
			log.Fatalln(err, sheetData)
		}

		var Estados []CatalogoModel.Estados
		err = json.Unmarshal(sheetData, &Estados)

		for _, est := range Estados {
			db.C("Estados").Insert(est)
		}
		fmt.Println("Estados Insertados Correctamente.")
	}

	Municipio, err := db.C("Municipios").Find(bson.M{}).Count()
	if err != nil {
		if err != nil {
			log.Fatalln(err)
		}
		return false
	}

	if Municipio == 0 {
		fmt.Println("Inicio insercion de municipios")
		sheetData, err := ioutil.ReadFile("Public/Resources/Sistema/Municipios.json")

		if err != nil {
			log.Fatalln(err)
		}

		var Municipios []CatalogoModel.Municipio
		err = json.Unmarshal(sheetData, &Municipios)

		for _, mun := range Municipios {
			db.C("Municipios").Insert(mun)
		}
		fmt.Println("Insercion de municipios correctamente")
	}

	Colonia, err := db.C("Colonias").Find(bson.M{}).Count()
	if err != nil {
		log.Fatalln(err)
		return false
	}

	if Colonia == 0 {

		fmt.Println("Inicio Insercion de Colonias")
		sheetData, err := ioutil.ReadFile("Public/Resources/Sistema/Colonias.json")
		if err != nil {
			log.Fatalln(err, sheetData)
		}

		var Colonias []CatalogoModel.Colonia
		err = json.Unmarshal(sheetData, &Colonias)

		for _, mun := range Colonias {
			db.C("Colonias").Insert(mun)
		}
		fmt.Println("Insercion de Colonias Correctamente")
	}

	return true
}
