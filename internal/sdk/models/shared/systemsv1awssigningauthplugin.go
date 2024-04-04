// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1AwsSigningAuthPlugin struct {
	EnvironmentCredentials interface{}                            `json:"environment_credentials,omitempty"`
	MetadataCredentials    *SystemsV1AwsMetadataCredentialService `json:"metadata_credentials,omitempty"`
	ProfileCredentials     *SystemsV1AwsProfileCredentialService  `json:"profile_credentials,omitempty"`
	// the AWS service to sign requests with, eg execute-api or s3 (default: s3)
	Service                *string                                   `json:"service,omitempty"`
	WebIdentityCredentials *SystemsV1AwsWebIdentityCredentialService `json:"web_identity_credentials,omitempty"`
}

func (o *SystemsV1AwsSigningAuthPlugin) GetEnvironmentCredentials() interface{} {
	if o == nil {
		return nil
	}
	return o.EnvironmentCredentials
}

func (o *SystemsV1AwsSigningAuthPlugin) GetMetadataCredentials() *SystemsV1AwsMetadataCredentialService {
	if o == nil {
		return nil
	}
	return o.MetadataCredentials
}

func (o *SystemsV1AwsSigningAuthPlugin) GetProfileCredentials() *SystemsV1AwsProfileCredentialService {
	if o == nil {
		return nil
	}
	return o.ProfileCredentials
}

func (o *SystemsV1AwsSigningAuthPlugin) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *SystemsV1AwsSigningAuthPlugin) GetWebIdentityCredentials() *SystemsV1AwsWebIdentityCredentialService {
	if o == nil {
		return nil
	}
	return o.WebIdentityCredentials
}
