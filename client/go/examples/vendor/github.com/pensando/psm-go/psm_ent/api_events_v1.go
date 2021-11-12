/*
 * Events API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// EventsV1ApiService EventsV1Api service
type EventsV1ApiService service

type ApiEventsGetGetEventRequest struct {
	ctx _context.Context
	ApiService *EventsV1ApiService
	uUID string
}


func (r ApiEventsGetGetEventRequest) Execute() (EventsEvent, *_nethttp.Response, error) {
	return r.ApiService.GetGetEventExecute(r)
}

/*
 * GetGetEvent Get specific event
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uUID
 * @return ApiEventsGetGetEventRequest
 */
func (a *EventsV1ApiService) GetGetEvent(ctx _context.Context, uUID string) ApiEventsGetGetEventRequest {
	return ApiEventsGetGetEventRequest{
		ApiService: a,
		ctx: ctx,
		uUID: uUID,
	}
}

/*
 * Execute executes the request
 * @return EventsEvent
 */
func (a *EventsV1ApiService) GetGetEventExecute(r ApiEventsGetGetEventRequest) (EventsEvent, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EventsEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsV1ApiService.GetGetEvent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/v1/events/{UUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"UUID"+"}", _neturl.PathEscape(parameterToString(r.uUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.GetGetEventExecute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEventsGetGetEvents1Request struct {
	ctx _context.Context
	ApiService *EventsV1ApiService
	oName *string
	oTenant *string
	oNamespace *string
	oGenerationId *string
	oResourceVersion *string
	oUuid *string
	oCreationTime *time.Time
	oModTime *time.Time
	oSelfLink *string
	labelSelector *string
	fieldSelector *string
	fieldChangeSelector *[]string
	from *int32
	maxResults *int32
	sortOrder *string
	metaOnly *bool
}

func (r ApiEventsGetGetEvents1Request) OName(oName string) ApiEventsGetGetEvents1Request {
	r.oName = &oName
	return r
}
func (r ApiEventsGetGetEvents1Request) OTenant(oTenant string) ApiEventsGetGetEvents1Request {
	r.oTenant = &oTenant
	return r
}
func (r ApiEventsGetGetEvents1Request) ONamespace(oNamespace string) ApiEventsGetGetEvents1Request {
	r.oNamespace = &oNamespace
	return r
}
func (r ApiEventsGetGetEvents1Request) OGenerationId(oGenerationId string) ApiEventsGetGetEvents1Request {
	r.oGenerationId = &oGenerationId
	return r
}
func (r ApiEventsGetGetEvents1Request) OResourceVersion(oResourceVersion string) ApiEventsGetGetEvents1Request {
	r.oResourceVersion = &oResourceVersion
	return r
}
func (r ApiEventsGetGetEvents1Request) OUuid(oUuid string) ApiEventsGetGetEvents1Request {
	r.oUuid = &oUuid
	return r
}
func (r ApiEventsGetGetEvents1Request) OCreationTime(oCreationTime time.Time) ApiEventsGetGetEvents1Request {
	r.oCreationTime = &oCreationTime
	return r
}
func (r ApiEventsGetGetEvents1Request) OModTime(oModTime time.Time) ApiEventsGetGetEvents1Request {
	r.oModTime = &oModTime
	return r
}
func (r ApiEventsGetGetEvents1Request) OSelfLink(oSelfLink string) ApiEventsGetGetEvents1Request {
	r.oSelfLink = &oSelfLink
	return r
}
func (r ApiEventsGetGetEvents1Request) LabelSelector(labelSelector string) ApiEventsGetGetEvents1Request {
	r.labelSelector = &labelSelector
	return r
}
func (r ApiEventsGetGetEvents1Request) FieldSelector(fieldSelector string) ApiEventsGetGetEvents1Request {
	r.fieldSelector = &fieldSelector
	return r
}
func (r ApiEventsGetGetEvents1Request) FieldChangeSelector(fieldChangeSelector []string) ApiEventsGetGetEvents1Request {
	r.fieldChangeSelector = &fieldChangeSelector
	return r
}
func (r ApiEventsGetGetEvents1Request) From(from int32) ApiEventsGetGetEvents1Request {
	r.from = &from
	return r
}
func (r ApiEventsGetGetEvents1Request) MaxResults(maxResults int32) ApiEventsGetGetEvents1Request {
	r.maxResults = &maxResults
	return r
}
func (r ApiEventsGetGetEvents1Request) SortOrder(sortOrder string) ApiEventsGetGetEvents1Request {
	r.sortOrder = &sortOrder
	return r
}
func (r ApiEventsGetGetEvents1Request) MetaOnly(metaOnly bool) ApiEventsGetGetEvents1Request {
	r.metaOnly = &metaOnly
	return r
}

func (r ApiEventsGetGetEvents1Request) Execute() (EventsEventList, *_nethttp.Response, error) {
	return r.ApiService.GetGetEvents1Execute(r)
}

/*
 * GetGetEvents1 Method for GetGetEvents1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEventsGetGetEvents1Request
 */
func (a *EventsV1ApiService) GetGetEvents1(ctx _context.Context) ApiEventsGetGetEvents1Request {
	return ApiEventsGetGetEvents1Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return EventsEventList
 */
func (a *EventsV1ApiService) GetGetEvents1Execute(r ApiEventsGetGetEvents1Request) (EventsEventList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EventsEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsV1ApiService.GetGetEvents1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/v1/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.oName != nil {
		localVarQueryParams.Add("O.name", parameterToString(*r.oName, ""))
	}
	if r.oTenant != nil {
		localVarQueryParams.Add("O.tenant", parameterToString(*r.oTenant, ""))
	}
	if r.oNamespace != nil {
		localVarQueryParams.Add("O.namespace", parameterToString(*r.oNamespace, ""))
	}
	if r.oGenerationId != nil {
		localVarQueryParams.Add("O.generation-id", parameterToString(*r.oGenerationId, ""))
	}
	if r.oResourceVersion != nil {
		localVarQueryParams.Add("O.resource-version", parameterToString(*r.oResourceVersion, ""))
	}
	if r.oUuid != nil {
		localVarQueryParams.Add("O.uuid", parameterToString(*r.oUuid, ""))
	}
	if r.oCreationTime != nil {
		localVarQueryParams.Add("O.creation-time", parameterToString(*r.oCreationTime, ""))
	}
	if r.oModTime != nil {
		localVarQueryParams.Add("O.mod-time", parameterToString(*r.oModTime, ""))
	}
	if r.oSelfLink != nil {
		localVarQueryParams.Add("O.self-link", parameterToString(*r.oSelfLink, ""))
	}
	if r.labelSelector != nil {
		localVarQueryParams.Add("label-selector", parameterToString(*r.labelSelector, ""))
	}
	if r.fieldSelector != nil {
		localVarQueryParams.Add("field-selector", parameterToString(*r.fieldSelector, ""))
	}
	if r.fieldChangeSelector != nil {
		localVarQueryParams.Add("field-change-selector", parameterToString(*r.fieldChangeSelector, "csv"))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("max-results", parameterToString(*r.maxResults, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sort-order", parameterToString(*r.sortOrder, ""))
	}
	if r.metaOnly != nil {
		localVarQueryParams.Add("meta-only", parameterToString(*r.metaOnly, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.GetGetEvents1Execute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEventsPostGetEventsRequest struct {
	ctx _context.Context
	ApiService *EventsV1ApiService
}


func (r ApiEventsPostGetEventsRequest) Execute() (EventsEventList, *_nethttp.Response, error) {
	return r.ApiService.PostGetEventsExecute(r)
}

/*
 * PostGetEvents Method for PostGetEvents
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEventsPostGetEventsRequest
 */
func (a *EventsV1ApiService) PostGetEvents(ctx _context.Context) ApiEventsPostGetEventsRequest {
	return ApiEventsPostGetEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return EventsEventList
 */
func (a *EventsV1ApiService) PostGetEventsExecute(r ApiEventsPostGetEventsRequest) (EventsEventList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EventsEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsV1ApiService.PostGetEvents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/v1/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.PostGetEventsExecute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}