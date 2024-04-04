// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Styra/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type CreateStackResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	StacksV1StacksPostResponse *shared.StacksV1StacksPostResponse
}

func (o *CreateStackResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateStackResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateStackResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateStackResponse) GetStacksV1StacksPostResponse() *shared.StacksV1StacksPostResponse {
	if o == nil {
		return nil
	}
	return o.StacksV1StacksPostResponse
}
