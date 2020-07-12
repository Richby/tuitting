package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD*/
var MongoCN = ConectarBD()
var clienteOtions = options.Client().ApplyURI("mongodb+srv://dbtuitting:ISCricardo19@tuitting.umcyb.gcp.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar la BD */
func ConectarBD() *mongo.Client {
	//Verificamos si hay conexión a la BD, de lo contrario, cachamos el error
	client, err := mongo.Connect(context.TODO(), clienteOtions)
	if err != nil {
		log.Fatal(err.Error())
		return client //se hace un retorno de variable aunque este vacia
	}

	//Conexion para verificar si el cliente de BD esta disponible
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client //se hace un retorno de variable aunque este vacia
	}
	log.Println("Conexión exitosa a la BD")
	return client
}

/*ChequeoConexion es el Ping a la BD*/
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
