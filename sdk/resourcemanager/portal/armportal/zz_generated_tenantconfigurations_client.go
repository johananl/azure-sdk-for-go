//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TenantConfigurationsClient contains the methods for the TenantConfigurations group.
// Don't use this type directly, use NewTenantConfigurationsClient() instead.
type TenantConfigurationsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewTenantConfigurationsClient creates a new instance of TenantConfigurationsClient with the specified values.
func NewTenantConfigurationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *TenantConfigurationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TenantConfigurationsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Create the tenant configuration. If configuration already exists - update it. User has to be a Tenant Admin for this operation.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantConfigurationsClient) Create(ctx context.Context, configurationName ConfigurationName, tenantConfiguration Configuration, options *TenantConfigurationsCreateOptions) (TenantConfigurationsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, configurationName, tenantConfiguration, options)
	if err != nil {
		return TenantConfigurationsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantConfigurationsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TenantConfigurationsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *TenantConfigurationsClient) createCreateRequest(ctx context.Context, configurationName ConfigurationName, tenantConfiguration Configuration, options *TenantConfigurationsCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations/{configurationName}"
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, tenantConfiguration)
}

// createHandleResponse handles the Create response.
func (client *TenantConfigurationsClient) createHandleResponse(resp *http.Response) (TenantConfigurationsCreateResponse, error) {
	result := TenantConfigurationsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configuration); err != nil {
		return TenantConfigurationsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *TenantConfigurationsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete the tenant configuration. User has to be a Tenant Admin for this operation.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantConfigurationsClient) Delete(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsDeleteOptions) (TenantConfigurationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, configurationName, options)
	if err != nil {
		return TenantConfigurationsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantConfigurationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TenantConfigurationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return TenantConfigurationsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TenantConfigurationsClient) deleteCreateRequest(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations/{configurationName}"
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *TenantConfigurationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the tenant configuration.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantConfigurationsClient) Get(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsGetOptions) (TenantConfigurationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, configurationName, options)
	if err != nil {
		return TenantConfigurationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantConfigurationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return TenantConfigurationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TenantConfigurationsClient) getCreateRequest(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations/{configurationName}"
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TenantConfigurationsClient) getHandleResponse(resp *http.Response) (TenantConfigurationsGetResponse, error) {
	result := TenantConfigurationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configuration); err != nil {
		return TenantConfigurationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TenantConfigurationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets list of the tenant configurations.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantConfigurationsClient) List(options *TenantConfigurationsListOptions) *TenantConfigurationsListPager {
	return &TenantConfigurationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp TenantConfigurationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConfigurationList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *TenantConfigurationsClient) listCreateRequest(ctx context.Context, options *TenantConfigurationsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TenantConfigurationsClient) listHandleResponse(resp *http.Response) (TenantConfigurationsListResponse, error) {
	result := TenantConfigurationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationList); err != nil {
		return TenantConfigurationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *TenantConfigurationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}