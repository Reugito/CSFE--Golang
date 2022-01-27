package kms

import (
	"github.com/fatih/structs"
)

// Provider is a common interface for each kms provider (aws, azure, gcp, local)
type Provider interface {
	Name() string
	Credentials() map[string]map[string]interface{}
	DataKeyOpts() interface{}
}

// gcpKMSCredentials used to access this kms provider
// See https://github.com/mongodb/specifications/blob/master/source/client-side-encryption/client-side-encryption.rst#kmsproviders
type gcpKMSCredentials struct {
	Email      string `structs:"email"`
	PrivateKey string `structs:"privateKey"`
	Endpoint   string `structs:"endpoint,omitempty"` // optional, defaults to oauth2.googleapis.com
}

// gcpKMSDataKeyOpts are the data key options used for this kms provider.
// See https://github.com/mongodb/specifications/blob/master/source/client-side-encryption/client-side-encryption.rst#datakeyopts
type gcpKMSDataKeyOpts struct {
	ProjectID  string `bson:"projectId"`
	Location   string `bson:"location"`
	KeyRing    string `bson:"keyRing"`
	KeyName    string `bson:"keyName"`
	KeyVersion string `bson:"keyVersion,omitempty"` // optional, defaults to the key's primary version
	Endpoint   string `bson:"endpoint,omitempty"`   // optional, defaults to cloudkms.googleapis.com
}

// GCP holds the credentials and master key information to use this kms
// Get an instance of this with the kms.GCPProvider() method
type GCP struct {
	credentials gcpKMSCredentials `structs:"gcp"`
	dataKeyOpts gcpKMSDataKeyOpts
	name        string
}

// Name is the name of this provider
func (g *GCP) Name() string {
	return g.name
}

// Credentials are the credentials for this provider returned in the format necessary
// to immediately pass to the driver
func (g *GCP) Credentials() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{"gcp": structs.Map(g.credentials)}
}

// DataKeyOpts are the data key options for this provider returned in the format necessary
// to immediately pass to the driver
func (g *GCP) DataKeyOpts() interface{} {
	return g.dataKeyOpts
}

// Local holds the credentials and master key information to use this kms
// Get an instance of this with the kms.LocalProvider() method
type Local struct {
	name string
	key  []byte
}

// LocalProvider returns information for using the local kms.
// Not for production
func LocalProvider(key []byte) *Local {
	return &Local{name: "local", key: key}
}

// Name is the name of this provider
func (l *Local) Name() string {
	return l.name
}

// Credentials are the credentials for this provider returned in the format necessary
// to immediately pass to the driver
func (l *Local) Credentials() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"local": {
			"key": l.key,
		},
	}
}

// DataKeyOpts are the data key options for this provider returned in the format necessary
// to immediately pass to the driver
func (l *Local) DataKeyOpts() interface{} {
	return nil
}
