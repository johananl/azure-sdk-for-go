//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageimportexport

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	host           string
	acceptLanguage *string
	pl             runtime.Pipeline
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
// acceptLanguage - Specifies the preferred language for the response.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationsClient(acceptLanguage *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LocationsClient{
		acceptLanguage: acceptLanguage,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns the details about a location to which you can ship the disks associated with an import or export job. A location
// is an Azure region.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-01-01
// locationName - The name of the location. For example, West US or westus.
// options - LocationsClientGetOptions contains the optional parameters for the LocationsClient.Get method.
func (client *LocationsClient) Get(ctx context.Context, locationName string, options *LocationsClientGetOptions) (LocationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, options)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LocationsClient) getCreateRequest(ctx context.Context, locationName string, options *LocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ImportExport/locations/{locationName}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if client.acceptLanguage != nil {
		req.Raw().Header["Accept-Language"] = []string{*client.acceptLanguage}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LocationsClient) getHandleResponse(resp *http.Response) (LocationsClientGetResponse, error) {
	result := LocationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Location); err != nil {
		return LocationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of locations to which you can ship the disks associated with an import or export job. A location
// is a Microsoft data center region.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-01-01
// options - LocationsClientListOptions contains the optional parameters for the LocationsClient.List method.
func (client *LocationsClient) NewListPager(options *LocationsClientListOptions) *runtime.Pager[LocationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocationsClientListResponse]{
		More: func(page LocationsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *LocationsClientListResponse) (LocationsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LocationsClient) listCreateRequest(ctx context.Context, options *LocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ImportExport/locations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if client.acceptLanguage != nil {
		req.Raw().Header["Accept-Language"] = []string{*client.acceptLanguage}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationsClient) listHandleResponse(resp *http.Response) (LocationsClientListResponse, error) {
	result := LocationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationsResponse); err != nil {
		return LocationsClientListResponse{}, err
	}
	return result, nil
}
