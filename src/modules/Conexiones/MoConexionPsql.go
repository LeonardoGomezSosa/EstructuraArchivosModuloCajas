package MoConexion

import (
	"database/sql"
	"fmt"

	"../Variables"
	_ "github.com/lib/pq"
)

//DataP es una estructura que contiene los datos de configuración en el archivo cfg
var DataP = MoVar.CargaSeccionCFG(MoVar.SecPsql)

//ConexionPsql abre una conexión a PostgreSql
func ConexionPsql() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DataP.Servidor, DataP.Usuario, DataP.Pass, DataP.NombreBase)
	db, err := sql.Open("postgres", dbinfo)
	return db, err
}

//IniciaSesionEspecificaPsql regresa una base de datos con su sesion especificada en el parametro para commit y rollback
func IniciaSesionEspecificaPsql() (*sql.DB, *sql.Tx, error) {
	Psql, err := ConexionPsql()
	if err != nil {
		return nil, nil, err
	}

	tx, err := Psql.Begin()
	if err != nil {
		return nil, nil, err
	}
	return Psql, tx, nil
}

//InsertaOActualizaRelacion verifica la tabla relacion e inserta el sku en caso de que no se encuentre; caso contrario actualiza la clave del sat
func InsertaOActualizaRelacion(tabla, sku, descripcion, claveSat string) error {
	var SesionPsql *sql.Tx
	var err error
	BasePsql, SesionPsql, err := IniciaSesionEspecificaPsql()
	if err != nil {
		fmt.Println("Errores al conectar con postgres: ", err)
		return err
	}
	BasePsql.Exec("set transaction isolation level serializable")

	Query := fmt.Sprintf(`SELECT "Sku" FROM "%v"  WHERE "Sku" ='%v'`, tabla, sku)
	Elemento, err := BasePsql.Query(Query)
	if err != nil {
		fmt.Println("Error al consultar el sku: ", err, Query)
		return err
	}
	var encontrado bool
	for Elemento.Next() {
		var skuEnc string
		err = Elemento.Scan(&skuEnc)
		if err != nil {
			fmt.Println("Error al consultar el sku: (2)", err)
			return err
		}
		if sku == skuEnc {
			encontrado = true
		}
	}
	if encontrado {
		Query = fmt.Sprintf(`UPDATE  public."%v"  SET  "ClaveSat" = '%v' WHERE "Sku" ='%v'`, tabla, claveSat, sku)
		_, err = SesionPsql.Exec(Query)
		if err != nil {
			fmt.Println("Ha ocurrido un error en la actualizacion", err)
			SesionPsql.Rollback()
			BasePsql.Close()
			return err
		}
	} else {
		query := fmt.Sprintf(`INSERT INTO public."%v" VALUES('%v','%v','%v')`, tabla, sku, descripcion, claveSat)
		_, errsql := SesionPsql.Exec(query)
		if errsql != nil {
			SesionPsql.Rollback()
			BasePsql.Close()
			fmt.Println("Error al insertar el producto")
			fmt.Println(query)
			return err
		}
	}
	SesionPsql.Commit()
	BasePsql.Close()
	return err
}
