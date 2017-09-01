package ArchivoModel

import (
	"fmt"
	"strconv"

	"../../Modules/Conexiones"
	"../../Modules/General"

	"../../Modules/Variables"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/olivere/elastic.v5"
)

//#########################< ESTRUCTURAS >##############################

//ArchivoMgo estructura de Archivos mongo
type ArchivoMgo struct {
	ID  bson.ObjectId `bson:"_id,omitempty"`
	Key string        `bson:"Usuario,omitempty"`
	Cer string        `bson:"Caja,omitempty"`
	Pem string        `bson:"Cargo,omitempty"`
}

//ArchivoElastic estructura de Archivos para insertar en Elastic
type ArchivoElastic struct {
	Key string `json:"Usuario,omitempty"`
	Cer string `json:"Caja,omitempty"`
	Pem string `json:"Cargo,omitempty"`
}

//#########################< FUNCIONES GENERALES MGO >###############################

//GetAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func GetAll() []ArchivoMgo {
	var result []ArchivoMgo
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	err = Archivos.Find(nil).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//CountAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func CountAll() int {
	var result int
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)

	if err != nil {
		fmt.Println(err)
	}
	result, err = Archivos.Find(nil).Count()
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetOne Regresa un documento específico de Mongo (Por Coleccion)
func GetOne(ID bson.ObjectId) ArchivoMgo {
	var result ArchivoMgo
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	err = Archivos.Find(bson.M{"_id": ID}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetEspecifics rsegresa un conjunto de documentos específicos de Mongo (Por Coleccion)
func GetEspecifics(Ides []bson.ObjectId) []ArchivoMgo {
	var result []ArchivoMgo
	var aux ArchivoMgo
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range Ides {
		aux = ArchivoMgo{}
		Archivos.Find(bson.M{"_id": value}).One(&aux)
		result = append(result, aux)
	}
	s.Close()
	return result
}

//GetEspecificByFields regresa un documento de Mongo especificando un campo y un determinado valor
func GetEspecificByFields(field string, valor interface{}) ArchivoMgo {
	var result ArchivoMgo
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)

	if err != nil {
		fmt.Println(err)
	}
	err = Archivos.Find(bson.M{field: valor}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetIDByField regresa un documento específico de Mongo (Por Coleccion)
func GetIDByField(field string, valor interface{}) bson.ObjectId {
	var result ArchivoMgo
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	err = Archivos.Find(bson.M{field: valor}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result.ID
}

//CargaComboArchivos regresa un combo de Archivo de mongo
func CargaComboArchivos(ID string) string {
	Archivos := GetAll()

	templ := ``

	if ID != "" {
		templ = `<option value="">--SELECCIONE--</option> `
	} else {
		templ = `<option value="" selected>--SELECCIONE--</option> `
	}

	for _, v := range Archivos {
		if ID == v.ID.Hex() {
			templ += `<option value="` + v.ID.Hex() + `" selected>  ` + v.ID.Hex() + ` </option> `
		} else {
			templ += `<option value="` + v.ID.Hex() + `">  ` + v.ID.Hex() + ` </option> `
		}

	}
	return templ
}

//GeneraTemplatesBusqueda crea templates de tabla de búsqueda
func GeneraTemplatesBusqueda(Archivos []ArchivoMgo) (string, string) {
	cuerpo := ``

	cabecera := `<tr>
			<th>#</th>
			
				<th>Key</th>					
				
				<th>Cer</th>					
				
				<th>Pem</th>					
				</tr>`

	for k, v := range Archivos {
		cuerpo += `<tr id = "` + v.ID.Hex() + `" onclick="window.location.href = '/Archivos/detalle/` + v.ID.Hex() + `';">`
		cuerpo += `<td>` + strconv.Itoa(k+1) + `</td>`
		cuerpo += `<td>` + v.Key + `</td>`

		cuerpo += `<td>` + v.Cer + `</td>`

		cuerpo += `<td>` + v.Pem + `</td>`

		cuerpo += `</tr>`
	}

	return cabecera, cuerpo
}

//########## GET NAMES ####################################

//GetNameArchivo regresa el nombre del Archivo con el ID especificado
func GetNameArchivo(ID bson.ObjectId) string {
	var result ArchivoMgo
	s, Archivos, err := MoConexion.GetColectionMgo(MoVar.ColeccionArchivo)
	if err != nil {
		fmt.Println(err)
	}
	Archivos.Find(bson.M{"_id": ID}).One(&result)

	s.Close()
	return result.ID.Hex()
}

//########################< FUNCIONES GENERALES PSQL >#############################

//######################< FUNCIONES GENERALES ELASTIC >############################

//BuscarEnElastic busca el texto solicitado en los campos solicitados
func BuscarEnElastic(texto string) *elastic.SearchResult {
	textoTilde, textoQuotes := MoGeneral.ConstruirCadenas(texto)

	queryTilde := elastic.NewQueryStringQuery(textoTilde)
	queryQuotes := elastic.NewQueryStringQuery(textoQuotes)

	queryTilde = queryTilde.Field("Key")
	queryQuotes = queryQuotes.Field("Key")

	queryTilde = queryTilde.Field("Cer")
	queryQuotes = queryQuotes.Field("Cer")

	queryTilde = queryTilde.Field("Pem")
	queryQuotes = queryQuotes.Field("Pem")

	var docs *elastic.SearchResult
	var err bool

	docs, err = MoConexion.BuscaElastic(MoVar.TipoArchivo, queryTilde)
	if err {
		fmt.Println("No Match 1st Try")
	}

	if docs.Hits.TotalHits == 0 {
		docs, err = MoConexion.BuscaElastic(MoVar.TipoArchivo, queryQuotes)
		if err {
			fmt.Println("No Match 2nd Try")
		}
	}

	return docs
}
