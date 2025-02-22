//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// NetworkStatusClient contains the methods for the NetworkStatus group.
// Don't use this type directly, use NewNetworkStatusClient() instead.
type NetworkStatusClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNetworkStatusClient creates a new instance of NetworkStatusClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNetworkStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkStatusClient, error) {
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
	client := &NetworkStatusClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListByLocation - Gets the Connectivity Status to the external resources on which the Api Management service depends from
// inside the Cloud Service. This also returns the DNS Servers as visible to the CloudService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// locationName - Location in which the API Management service is deployed. This is one of the Azure Regions like West US,
// East US, South Central US.
// options - NetworkStatusClientListByLocationOptions contains the optional parameters for the NetworkStatusClient.ListByLocation
// method.
func (client *NetworkStatusClient) ListByLocation(ctx context.Context, resourceGroupName string, serviceName string, locationName string, options *NetworkStatusClientListByLocationOptions) (NetworkStatusClientListByLocationResponse, error) {
	req, err := client.listByLocationCreateRequest(ctx, resourceGroupName, serviceName, locationName, options)
	if err != nil {
		return NetworkStatusClientListByLocationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkStatusClientListByLocationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkStatusClientListByLocationResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByLocationHandleResponse(resp)
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *NetworkStatusClient) listByLocationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, locationName string, options *NetworkStatusClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/locations/{locationName}/networkstatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *NetworkStatusClient) listByLocationHandleResponse(resp *http.Response) (NetworkStatusClientListByLocationResponse, error) {
	result := NetworkStatusClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkStatusContract); err != nil {
		return NetworkStatusClientListByLocationResponse{}, err
	}
	return result, nil
}

// ListByService - Gets the Connectivity Status to the external resources on which the Api Management service depends from
// inside the Cloud Service. This also returns the DNS Servers as visible to the CloudService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - NetworkStatusClientListByServiceOptions contains the optional parameters for the NetworkStatusClient.ListByService
// method.
func (client *NetworkStatusClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, options *NetworkStatusClientListByServiceOptions) (NetworkStatusClientListByServiceResponse, error) {
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return NetworkStatusClientListByServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkStatusClientListByServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkStatusClientListByServiceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByServiceHandleResponse(resp)
}

// listByServiceCreateRequest creates the ListByService request.
func (client *NetworkStatusClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *NetworkStatusClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/networkstatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *NetworkStatusClient) listByServiceHandleResponse(resp *http.Response) (NetworkStatusClientListByServiceResponse, error) {
	result := NetworkStatusClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkStatusContractByLocationArray); err != nil {
		return NetworkStatusClientListByServiceResponse{}, err
	}
	return result, nil
}
