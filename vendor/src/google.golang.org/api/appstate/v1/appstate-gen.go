// Package appstate provides access to the Google App State API.
//
// See https://developers.google.com/games/services/web/api/states
//
// Usage example:
//
//   import "google.golang.org/api/appstate/v1"
//   ...
//   appstateService, err := appstate.New(oauthHttpClient)
package appstate // import "google.golang.org/api/appstate/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/internal"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = internal.MarshalJSON

const apiId = "appstate:v1"
const apiName = "appstate"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/appstate/v1/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data for this application
	AppstateScope = "https://www.googleapis.com/auth/appstate"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.States = NewStatesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	States *StatesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewStatesService(s *Service) *StatesService {
	rs := &StatesService{s: s}
	return rs
}

type StatesService struct {
	s *Service
}

// GetResponse: This is a JSON template for an app state resource.
type GetResponse struct {
	// CurrentStateVersion: The current app state version.
	CurrentStateVersion string `json:"currentStateVersion,omitempty"`

	// Data: The requested data.
	Data string `json:"data,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#getResponse.
	Kind string `json:"kind,omitempty"`

	// StateKey: The key for the data.
	StateKey int64 `json:"stateKey,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CurrentStateVersion")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GetResponse) MarshalJSON() ([]byte, error) {
	type noMethod GetResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ListResponse: This is a JSON template to convert a list-response for
// app state.
type ListResponse struct {
	// Items: The app state data.
	Items []*GetResponse `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#listResponse.
	Kind string `json:"kind,omitempty"`

	// MaximumKeyCount: The maximum number of keys allowed for this user.
	MaximumKeyCount int64 `json:"maximumKeyCount,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// UpdateRequest: This is a JSON template for a requests which update
// app state
type UpdateRequest struct {
	// Data: The new app state data that your application is trying to
	// update with.
	Data string `json:"data,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#updateRequest.
	Kind string `json:"kind,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Data") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod UpdateRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// WriteResult: This is a JSON template for an app state write result.
type WriteResult struct {
	// CurrentStateVersion: The version of the data for this key on the
	// server.
	CurrentStateVersion string `json:"currentStateVersion,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#writeResult.
	Kind string `json:"kind,omitempty"`

	// StateKey: The written key.
	StateKey int64 `json:"stateKey,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CurrentStateVersion")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WriteResult) MarshalJSON() ([]byte, error) {
	type noMethod WriteResult
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// method id "appstate.states.clear":

type StatesClearCall struct {
	s        *Service
	stateKey int64
	opt_     map[string]interface{}
	ctx_     context.Context
}

// Clear: Clears (sets to empty) the data for the passed key if and only
// if the passed version matches the currently stored version. This
// method results in a conflict error on version mismatch.
func (r *StatesService) Clear(stateKey int64) *StatesClearCall {
	c := &StatesClearCall{s: r.s, opt_: make(map[string]interface{})}
	c.stateKey = stateKey
	return c
}

// CurrentDataVersion sets the optional parameter "currentDataVersion":
// The version of the data to be cleared. Version strings are returned
// by the server.
func (c *StatesClearCall) CurrentDataVersion(currentDataVersion string) *StatesClearCall {
	c.opt_["currentDataVersion"] = currentDataVersion
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesClearCall) Fields(s ...googleapi.Field) *StatesClearCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *StatesClearCall) Context(ctx context.Context) *StatesClearCall {
	c.ctx_ = ctx
	return c
}

func (c *StatesClearCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["currentDataVersion"]; ok {
		params.Set("currentDataVersion", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "states/{stateKey}/clear")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "appstate.states.clear" call.
// Exactly one of *WriteResult or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *WriteResult.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *StatesClearCall) Do() (*WriteResult, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WriteResult{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Clears (sets to empty) the data for the passed key if and only if the passed version matches the currently stored version. This method results in a conflict error on version mismatch.",
	//   "httpMethod": "POST",
	//   "id": "appstate.states.clear",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "currentDataVersion": {
	//       "description": "The version of the data to be cleared. Version strings are returned by the server.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}/clear",
	//   "response": {
	//     "$ref": "WriteResult"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.delete":

type StatesDeleteCall struct {
	s        *Service
	stateKey int64
	opt_     map[string]interface{}
	ctx_     context.Context
}

// Delete: Deletes a key and the data associated with it. The key is
// removed and no longer counts against the key quota. Note that since
// this method is not safe in the face of concurrent modifications, it
// should only be used for development and testing purposes. Invoking
// this method in shipping code can result in data loss and data
// corruption.
func (r *StatesService) Delete(stateKey int64) *StatesDeleteCall {
	c := &StatesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.stateKey = stateKey
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesDeleteCall) Fields(s ...googleapi.Field) *StatesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *StatesDeleteCall) Context(ctx context.Context) *StatesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *StatesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "states/{stateKey}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "appstate.states.delete" call.
func (c *StatesDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a key and the data associated with it. The key is removed and no longer counts against the key quota. Note that since this method is not safe in the face of concurrent modifications, it should only be used for development and testing purposes. Invoking this method in shipping code can result in data loss and data corruption.",
	//   "httpMethod": "DELETE",
	//   "id": "appstate.states.delete",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.get":

type StatesGetCall struct {
	s        *Service
	stateKey int64
	opt_     map[string]interface{}
	ctx_     context.Context
}

// Get: Retrieves the data corresponding to the passed key. If the key
// does not exist on the server, an HTTP 404 will be returned.
func (r *StatesService) Get(stateKey int64) *StatesGetCall {
	c := &StatesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.stateKey = stateKey
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesGetCall) Fields(s ...googleapi.Field) *StatesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *StatesGetCall) IfNoneMatch(entityTag string) *StatesGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *StatesGetCall) Context(ctx context.Context) *StatesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *StatesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "states/{stateKey}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "appstate.states.get" call.
// Exactly one of *GetResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *GetResponse.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *StatesGetCall) Do() (*GetResponse, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GetResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the data corresponding to the passed key. If the key does not exist on the server, an HTTP 404 will be returned.",
	//   "httpMethod": "GET",
	//   "id": "appstate.states.get",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}",
	//   "response": {
	//     "$ref": "GetResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.list":

type StatesListCall struct {
	s    *Service
	opt_ map[string]interface{}
	ctx_ context.Context
}

// List: Lists all the states keys, and optionally the state data.
func (r *StatesService) List() *StatesListCall {
	c := &StatesListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// IncludeData sets the optional parameter "includeData": Whether to
// include the full data in addition to the version number
func (c *StatesListCall) IncludeData(includeData bool) *StatesListCall {
	c.opt_["includeData"] = includeData
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesListCall) Fields(s ...googleapi.Field) *StatesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *StatesListCall) IfNoneMatch(entityTag string) *StatesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *StatesListCall) Context(ctx context.Context) *StatesListCall {
	c.ctx_ = ctx
	return c
}

func (c *StatesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["includeData"]; ok {
		params.Set("includeData", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "states")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "appstate.states.list" call.
// Exactly one of *ListResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ListResponse.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *StatesListCall) Do() (*ListResponse, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the states keys, and optionally the state data.",
	//   "httpMethod": "GET",
	//   "id": "appstate.states.list",
	//   "parameters": {
	//     "includeData": {
	//       "default": "false",
	//       "description": "Whether to include the full data in addition to the version number",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "states",
	//   "response": {
	//     "$ref": "ListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.update":

type StatesUpdateCall struct {
	s             *Service
	stateKey      int64
	updaterequest *UpdateRequest
	opt_          map[string]interface{}
	ctx_          context.Context
}

// Update: Update the data associated with the input key if and only if
// the passed version matches the currently stored version. This method
// is safe in the face of concurrent writes. Maximum per-key size is
// 128KB.
func (r *StatesService) Update(stateKey int64, updaterequest *UpdateRequest) *StatesUpdateCall {
	c := &StatesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.stateKey = stateKey
	c.updaterequest = updaterequest
	return c
}

// CurrentStateVersion sets the optional parameter
// "currentStateVersion": The version of the app state your application
// is attempting to update. If this does not match the current version,
// this method will return a conflict error. If there is no data stored
// on the server for this key, the update will succeed irrespective of
// the value of this parameter.
func (c *StatesUpdateCall) CurrentStateVersion(currentStateVersion string) *StatesUpdateCall {
	c.opt_["currentStateVersion"] = currentStateVersion
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesUpdateCall) Fields(s ...googleapi.Field) *StatesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *StatesUpdateCall) Context(ctx context.Context) *StatesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *StatesUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["currentStateVersion"]; ok {
		params.Set("currentStateVersion", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "states/{stateKey}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "appstate.states.update" call.
// Exactly one of *WriteResult or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *WriteResult.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *StatesUpdateCall) Do() (*WriteResult, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WriteResult{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the data associated with the input key if and only if the passed version matches the currently stored version. This method is safe in the face of concurrent writes. Maximum per-key size is 128KB.",
	//   "httpMethod": "PUT",
	//   "id": "appstate.states.update",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "currentStateVersion": {
	//       "description": "The version of the app state your application is attempting to update. If this does not match the current version, this method will return a conflict error. If there is no data stored on the server for this key, the update will succeed irrespective of the value of this parameter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}",
	//   "request": {
	//     "$ref": "UpdateRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteResult"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}
