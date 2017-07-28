package Inicio

// import (
// 	"context"
// 	"fmt"

// 	elastic "gopkg.in/olivere/elastic.v5"

// 	//"../../Modulos/Variables"
// )

// //DataE es una estructura que contiene los datos de configuración en el archivo cfg
// var DataE = MoVar.CargaSeccionCFG(MoVar.SecElastic)

// //ctx Contexto
// var ctx = context.Background()

// //InitDatosElastic funcion que inicializa los datos de elastic
// func InitDatosElastic() bool {

// 	client, err := elastic.NewClient(elastic.SetURL(DataE.BaseURL))
// 	if err != nil {
// 		fmt.Println("Error al crear un nuevo cliente", err)
// 		return false
// 	}
// 	defer client.Stop()
// 	exists, err := client.IndexExists(MoVar.IndexElastic).Do(ctx)
// 	if err != nil {
// 		fmt.Println("Error al verificar index de elastic", err)
// 		return false
// 	}

// 	if !exists {

// 		fmt.Println("\n El Indice: %s no existia se ha Creado ", MoVar.IndexElastic)

// 		// Create an index
// 		_, err = client.CreateIndex(MoVar.IndexElastic).Do(ctx)
// 		if err != nil {
// 			return false
// 		}

// 		return false
// 	}
// 	return true

// }
