package ArchivoModel

import (
	"fmt"

	"../../Modules/Conexiones"
	"../../Modules/Variables"
	"gopkg.in/mgo.v2/bson"
)

//IArchivo interface con los métodos de la clase
type IArchivo interface {
	InsertaMgo() bool
	InsertaElastic() bool

	ActualizaMgo(campos []string, valores []interface{}) bool
	ActualizaElastic(campos []string, valores []interface{}) bool //Reemplaza No Actualiza

	ReemplazaMgo() bool
	ReemplazaElastic() bool

	ConsultaExistenciaByFieldMgo(field string, valor string)

	ConsultaExistenciaByIDMgo() bool
	ConsultaExistenciaByIDElastic() bool

	EliminaByIDMgo() bool
	EliminaByIDElastic() bool
}

//################################################<<METODOS DE GESTION >>################################################################

//##################################<< INSERTAR >>###################################

//InsertaMgo es un método que crea un registro en Mongo
func (p ArchivoMgo) InsertaMgo() bool {
	result := false
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}

	err = Archivos.Insert(p)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}

	s.Close()
	return result
}

//InsertaElastic es un método que crea un registro en Mongo
func (p ArchivoMgo) InsertaElastic() bool {
	var ArchivoE ArchivoElastic

	ArchivoE.Key = p.Key
	ArchivoE.Cer = p.Cer
	ArchivoE.Pem = p.Pem
	insert := MoConexion.InsertaElastic(MoVar.TipoArchivo, p.ID.Hex(), ArchivoE)
	if !insert {
		fmt.Println("Error al insertar Archivo en Elastic")
		return false
	}
	return true
}

//##########################<< UPDATE >>############################################

//ActualizaMgo es un método que encuentra y Actualiza un registro en Mongo
//IMPORTANTE --> Debe coincidir el número y orden de campos con el de valores
func (p ArchivoMgo) ActualizaMgo(campos []string, valores []interface{}) bool {
	result := false
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	var Abson bson.M
	Abson = make(map[string]interface{})
	for k, v := range campos {
		Abson[v] = valores[k]
	}
	change := bson.M{"$set": Abson}
	err = Archivos.Update(bson.M{"_id": p.ID}, change)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}
	s.Close()
	return result
}

//ActualizaElastic es un método que encuentra y Actualiza un registro en Mongo
func (p ArchivoMgo) ActualizaElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoArchivo, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar Archivo en Elastic")
		return false
	}

	if !p.InsertaElastic() {
		fmt.Println("Error al actualizar Archivo en Elastic, se perdió Referencia.")
		return false
	}

	return true
}

//##########################<< REEMPLAZA >>############################################

//ReemplazaMgo es un método que encuentra y Actualiza un registro en Mongo
func (p ArchivoMgo) ReemplazaMgo() bool {
	result := false
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	err = Archivos.Update(bson.M{"_id": p.ID}, p)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}
	s.Close()
	return result
}

//ReemplazaElastic es un método que encuentra y reemplaza un Archivo en elastic
func (p ArchivoMgo) ReemplazaElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoArchivo, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar Archivo en Elastic")
		return false
	}
	insert := MoConexion.InsertaElastic(MoVar.TipoArchivo, p.ID.Hex(), p)
	if !insert {
		fmt.Println("Error al actualizar Archivo en Elastic")
		return false
	}
	return true
}

//###########################<< CONSULTA EXISTENCIAS >>###################################

//ConsultaExistenciaByFieldMgo es un método que verifica si un registro existe en Mongo indicando un campo y un valor string
func (p ArchivoMgo) ConsultaExistenciaByFieldMgo(field string, valor string) bool {
	result := false
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	n, e := Archivos.Find(bson.M{field: valor}).Count()
	if e != nil {
		fmt.Println(e)
	}
	if n > 0 {
		result = true
	}
	s.Close()
	return result
}

//ConsultaExistenciaByIDMgo es un método que encuentra un registro en Mongo buscándolo por ID
func (p ArchivoMgo) ConsultaExistenciaByIDMgo() bool {
	result := false
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	n, e := Archivos.Find(bson.M{"_id": p.ID}).Count()
	if e != nil {
		fmt.Println(e)
	}
	if n > 0 {
		result = true
	}
	s.Close()
	return result
}

//ConsultaExistenciaByIDElastic es un método que encuentra un registro en Mongo buscándolo por ID
func (p ArchivoMgo) ConsultaExistenciaByIDElastic() bool {
	result := MoConexion.ConsultaElastic(MoVar.TipoArchivo, p.ID.Hex())
	return result
}

//##################################<< ELIMINACIONES >>#################################################

//EliminaByIDMgo es un método que elimina un registro en Mongo
func (p ArchivoMgo) EliminaByIDMgo() bool {
	result := false
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	e := Archivos.RemoveId(bson.M{"_id": p.ID})
	if e != nil {
		result = true
	} else {
		fmt.Println(e)
	}
	s.Close()
	return result
}

//EliminaByIDElastic es un método que elimina un registro en Mongo
func (p ArchivoMgo) EliminaByIDElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoArchivo, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar Archivo en Elastic")
		return false
	}
	return true
}
