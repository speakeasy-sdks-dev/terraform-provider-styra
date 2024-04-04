// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1EnforcementConfig struct {
	// true if the policy is enforced
	Enforced bool `json:"enforced"`
	// enforcement type e.g. opa, test, mask
	Type string `json:"type"`
}

func (o *SystemsV1EnforcementConfig) GetEnforced() bool {
	if o == nil {
		return false
	}
	return o.Enforced
}

func (o *SystemsV1EnforcementConfig) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
