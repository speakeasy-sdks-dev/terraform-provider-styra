// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1ReasonMapping struct {
	// dot-separated decision property path
	Path string `json:"path"`
}

func (o *SystemsV1ReasonMapping) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}
