// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1KeyConfig struct {
	// name of the signing algorithm
	Algorithm *string `json:"algorithm,omitempty"`
	// PEM encoded public key to use for signature verification
	Key *string `json:"key,omitempty"`
	// PEM encoded private key to use for signing
	PrivateKey *string `json:"private_key,omitempty"`
	// scope to use for bundle signature verification
	Scope *string `json:"scope,omitempty"`
}

func (o *SystemsV1KeyConfig) GetAlgorithm() *string {
	if o == nil {
		return nil
	}
	return o.Algorithm
}

func (o *SystemsV1KeyConfig) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *SystemsV1KeyConfig) GetPrivateKey() *string {
	if o == nil {
		return nil
	}
	return o.PrivateKey
}

func (o *SystemsV1KeyConfig) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}
