package csfe

import (
	"GoLandWorkSpace/CSFLE/kms"
	"context"
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

// GetDataKey creates a new data key and returns the base64 encoding to be used
// in schema configuration for automatic encryption

func GetDataKey(keyVaultNamespace, uri, keyAltName string, provider kms.Provider) (string, error) {

	clientEnctriptionOpts := options.ClientEncryption().SetKeyVaultNamespace(keyVaultNamespace).
		SetKmsProviders(provider.Credentials())

	keyVaultClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return "", fmt.Errorf("client encryption connect error %v", err)
	}

	clientEnc, err := mongo.NewClientEncryption(keyVaultClient, clientEnctriptionOpts)
	if err != nil {
		return "", fmt.Errorf("NewClientEncryption error %v", err)
	}
	defer func() {
		_ = clientEnc.Close(context.TODO())
	}()

	keyVault := strings.Split(keyVaultNamespace, ".")

	db := keyVault[0]
	coll := keyVault[1]

	var dataKey bson.M

	err = keyVaultClient.Database(db).Collection(coll).FindOne(context.TODO(), bson.M{"keyAltNames": keyAltName}).
		Decode(&dataKey)

	if err == mongo.ErrNoDocuments {
		dataKeyOpts := options.DataKey().
			SetKeyAltNames([]string{keyAltName}).
			SetKeyAltNames([]string{keyAltName})
		dataKeyID, err := clientEnc.CreateDataKey(context.TODO(), provider.Name(), dataKeyOpts)
		if err != nil {
			return "", fmt.Errorf("create data key error %v", err)
		}

		return base64.StdEncoding.EncodeToString(dataKeyID.Data), nil
	}
	if err != nil {
		return "", fmt.Errorf("error encountered while attempting to find key")
	}

	return base64.StdEncoding.EncodeToString(dataKey["_id"].(primitive.Binary).Data), nil

}
