//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

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

// ScriptExecutionHistoryClient contains the methods for the ScriptExecutionHistory group.
// Don't use this type directly, use NewScriptExecutionHistoryClient() instead.
type ScriptExecutionHistoryClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewScriptExecutionHistoryClient creates a new instance of ScriptExecutionHistoryClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewScriptExecutionHistoryClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScriptExecutionHistoryClient, error) {
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
	client := &ScriptExecutionHistoryClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByClusterPager - Lists all scripts' execution history for the specified cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// options - ScriptExecutionHistoryClientListByClusterOptions contains the optional parameters for the ScriptExecutionHistoryClient.ListByCluster
// method.
func (client *ScriptExecutionHistoryClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *ScriptExecutionHistoryClientListByClusterOptions) *runtime.Pager[ScriptExecutionHistoryClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScriptExecutionHistoryClientListByClusterResponse]{
		More: func(page ScriptExecutionHistoryClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScriptExecutionHistoryClientListByClusterResponse) (ScriptExecutionHistoryClientListByClusterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ScriptExecutionHistoryClientListByClusterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ScriptExecutionHistoryClientListByClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ScriptExecutionHistoryClientListByClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterHandleResponse(resp)
		},
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *ScriptExecutionHistoryClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ScriptExecutionHistoryClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *ScriptExecutionHistoryClient) listByClusterHandleResponse(resp *http.Response) (ScriptExecutionHistoryClientListByClusterResponse, error) {
	result := ScriptExecutionHistoryClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptActionExecutionHistoryList); err != nil {
		return ScriptExecutionHistoryClientListByClusterResponse{}, err
	}
	return result, nil
}

// Promote - Promotes the specified ad-hoc script execution to a persisted script.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// scriptExecutionID - The script execution Id
// options - ScriptExecutionHistoryClientPromoteOptions contains the optional parameters for the ScriptExecutionHistoryClient.Promote
// method.
func (client *ScriptExecutionHistoryClient) Promote(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string, options *ScriptExecutionHistoryClientPromoteOptions) (ScriptExecutionHistoryClientPromoteResponse, error) {
	req, err := client.promoteCreateRequest(ctx, resourceGroupName, clusterName, scriptExecutionID, options)
	if err != nil {
		return ScriptExecutionHistoryClientPromoteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptExecutionHistoryClientPromoteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScriptExecutionHistoryClientPromoteResponse{}, runtime.NewResponseError(resp)
	}
	return ScriptExecutionHistoryClientPromoteResponse{}, nil
}

// promoteCreateRequest creates the Promote request.
func (client *ScriptExecutionHistoryClient) promoteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string, options *ScriptExecutionHistoryClientPromoteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory/{scriptExecutionId}/promote"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if scriptExecutionID == "" {
		return nil, errors.New("parameter scriptExecutionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionId}", url.PathEscape(scriptExecutionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
