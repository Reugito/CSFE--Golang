package csfe

import (
	"GoLandWorkSpace/CSFLE/kms"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EncryptedClient(keyVaultNamespace, uri string, schemaMap map[string]interface{}, provider kms.Provider) (*mongo.Client, error) {
	autoEncryptionOpts := options.AutoEncryption().
		SetKmsProviders(provider.Credentials()).
		SetKeyVaultNamespace(keyVaultNamespace).
		SetSchemaMap(schemaMap)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetAutoEncryptionOptions(autoEncryptionOpts))
	if err != nil {
		return nil, fmt.Errorf("Connect error for encrypted client: %v", err)
	}
	return client, nil
}
