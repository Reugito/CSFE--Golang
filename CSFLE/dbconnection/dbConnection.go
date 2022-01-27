package dbconnection

import (
	"GoLandWorkSpace/CSFLE/csfe"
	"GoLandWorkSpace/CSFLE/kms"
	"GoLandWorkSpace/CSFLE/schema"
	"GoLandWorkSpace/CSFLE/util"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	keyVaultNamespace = "encryption.__keyVault"
	uri               = "mongodb://localhost:27017"
	dbName            = "test"
	collName          = "employeeDetails"
	keyAltName        = "demo-data-key"
)

var Ctx = context.TODO()

func DbConnect() *mongo.Collection {

	preferredProvider := kms.LocalProvider(util.LocalMasterKey())

	dataKeyBase64, err := csfe.GetDataKey(keyVaultNamespace, uri, keyAltName, preferredProvider)
	if err != nil {
		log.Fatalf("problem during data key creation: %v", err)
	}

	s, err := schema.CreateJSONSchema(dataKeyBase64)
	if err != nil {
		log.Panic(err)
	}
	schemaMap := map[string]interface{}{
		dbName + "." + collName: s,
	}

	eclient, err := csfe.EncryptedClient(keyVaultNamespace, uri, schemaMap, preferredProvider)
	if err != nil {
		log.Panic(err)
	}

	//defer func() {
	//	_ = eclient.Disconnect(context.TODO())
	//}()

	return eclient.Database(dbName).Collection(collName)
}
