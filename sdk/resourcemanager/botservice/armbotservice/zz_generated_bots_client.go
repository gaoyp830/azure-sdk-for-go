//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice

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

// BotsClient contains the methods for the Bots group.
// Don't use this type directly, use NewBotsClient() instead.
type BotsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBotsClient creates a new instance of BotsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBotsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BotsClient, error) {
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
	client := &BotsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Creates a Bot Service. Bot Service is a resource group wide resource type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the Bot resource group in the user subscription.
// resourceName - The name of the Bot resource.
// parameters - The parameters to provide for the created bot.
// options - BotsClientCreateOptions contains the optional parameters for the BotsClient.Create method.
func (client *BotsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot, options *BotsClientCreateOptions) (BotsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return BotsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BotsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *BotsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot, options *BotsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *BotsClient) createHandleResponse(resp *http.Response) (BotsClientCreateResponse, error) {
	result := BotsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Bot); err != nil {
		return BotsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Bot Service from the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the Bot resource group in the user subscription.
// resourceName - The name of the Bot resource.
// options - BotsClientDeleteOptions contains the optional parameters for the BotsClient.Delete method.
func (client *BotsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, options *BotsClientDeleteOptions) (BotsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return BotsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return BotsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return BotsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *BotsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a BotService specified by the parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the Bot resource group in the user subscription.
// resourceName - The name of the Bot resource.
// options - BotsClientGetOptions contains the optional parameters for the BotsClient.Get method.
func (client *BotsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *BotsClientGetOptions) (BotsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return BotsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *BotsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BotsClient) getHandleResponse(resp *http.Response) (BotsClientGetResponse, error) {
	result := BotsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Bot); err != nil {
		return BotsClientGetResponse{}, err
	}
	return result, nil
}

// GetCheckNameAvailability - Check whether a bot name is available.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// parameters - The request body parameters to provide for the check name availability request
// options - BotsClientGetCheckNameAvailabilityOptions contains the optional parameters for the BotsClient.GetCheckNameAvailability
// method.
func (client *BotsClient) GetCheckNameAvailability(ctx context.Context, parameters CheckNameAvailabilityRequestBody, options *BotsClientGetCheckNameAvailabilityOptions) (BotsClientGetCheckNameAvailabilityResponse, error) {
	req, err := client.getCheckNameAvailabilityCreateRequest(ctx, parameters, options)
	if err != nil {
		return BotsClientGetCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotsClientGetCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotsClientGetCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.getCheckNameAvailabilityHandleResponse(resp)
}

// getCheckNameAvailabilityCreateRequest creates the GetCheckNameAvailability request.
func (client *BotsClient) getCheckNameAvailabilityCreateRequest(ctx context.Context, parameters CheckNameAvailabilityRequestBody, options *BotsClientGetCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BotService/checkNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// getCheckNameAvailabilityHandleResponse handles the GetCheckNameAvailability response.
func (client *BotsClient) getCheckNameAvailabilityHandleResponse(resp *http.Response) (BotsClientGetCheckNameAvailabilityResponse, error) {
	result := BotsClientGetCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponseBody); err != nil {
		return BotsClientGetCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns all the resources of a particular type belonging to a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// options - BotsClientListOptions contains the optional parameters for the BotsClient.List method.
func (client *BotsClient) NewListPager(options *BotsClientListOptions) *runtime.Pager[BotsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BotsClientListResponse]{
		More: func(page BotsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BotsClientListResponse) (BotsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BotsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BotsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BotsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BotsClient) listCreateRequest(ctx context.Context, options *BotsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BotService/botServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BotsClient) listHandleResponse(resp *http.Response) (BotsClientListResponse, error) {
	result := BotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotResponseList); err != nil {
		return BotsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns all the resources of a particular type belonging to a resource group
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the Bot resource group in the user subscription.
// options - BotsClientListByResourceGroupOptions contains the optional parameters for the BotsClient.ListByResourceGroup
// method.
func (client *BotsClient) NewListByResourceGroupPager(resourceGroupName string, options *BotsClientListByResourceGroupOptions) *runtime.Pager[BotsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[BotsClientListByResourceGroupResponse]{
		More: func(page BotsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BotsClientListByResourceGroupResponse) (BotsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BotsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BotsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BotsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *BotsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *BotsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *BotsClient) listByResourceGroupHandleResponse(resp *http.Response) (BotsClientListByResourceGroupResponse, error) {
	result := BotsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotResponseList); err != nil {
		return BotsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the Bot resource group in the user subscription.
// resourceName - The name of the Bot resource.
// parameters - The parameters to provide for the created bot.
// options - BotsClientUpdateOptions contains the optional parameters for the BotsClient.Update method.
func (client *BotsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot, options *BotsClientUpdateOptions) (BotsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return BotsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BotsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *BotsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot, options *BotsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *BotsClient) updateHandleResponse(resp *http.Response) (BotsClientUpdateResponse, error) {
	result := BotsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Bot); err != nil {
		return BotsClientUpdateResponse{}, err
	}
	return result, nil
}
