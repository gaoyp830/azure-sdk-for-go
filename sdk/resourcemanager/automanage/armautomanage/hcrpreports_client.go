//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomanage

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

// HCRPReportsClient contains the methods for the HCRPReports group.
// Don't use this type directly, use NewHCRPReportsClient() instead.
type HCRPReportsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHCRPReportsClient creates a new instance of HCRPReportsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHCRPReportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HCRPReportsClient, error) {
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
	client := &HCRPReportsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get information about a report associated with a configuration profile assignment run
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-04
// resourceGroupName - The name of the resource group. The name is case insensitive.
// machineName - The name of the Arc machine.
// configurationProfileAssignmentName - The configuration profile assignment name.
// reportName - The report name.
// options - HCRPReportsClientGetOptions contains the optional parameters for the HCRPReportsClient.Get method.
func (client *HCRPReportsClient) Get(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string, reportName string, options *HCRPReportsClientGetOptions) (HCRPReportsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, configurationProfileAssignmentName, reportName, options)
	if err != nil {
		return HCRPReportsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HCRPReportsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HCRPReportsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HCRPReportsClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string, reportName string, options *HCRPReportsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports/{reportName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HCRPReportsClient) getHandleResponse(resp *http.Response) (HCRPReportsClientGetResponse, error) {
	result := HCRPReportsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Report); err != nil {
		return HCRPReportsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByConfigurationProfileAssignmentsPager - Retrieve a list of reports within a given configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-04
// resourceGroupName - The name of the resource group. The name is case insensitive.
// machineName - The name of the Arc machine.
// configurationProfileAssignmentName - The configuration profile assignment name.
// options - HCRPReportsClientListByConfigurationProfileAssignmentsOptions contains the optional parameters for the HCRPReportsClient.ListByConfigurationProfileAssignments
// method.
func (client *HCRPReportsClient) NewListByConfigurationProfileAssignmentsPager(resourceGroupName string, machineName string, configurationProfileAssignmentName string, options *HCRPReportsClientListByConfigurationProfileAssignmentsOptions) *runtime.Pager[HCRPReportsClientListByConfigurationProfileAssignmentsResponse] {
	return runtime.NewPager(runtime.PagingHandler[HCRPReportsClientListByConfigurationProfileAssignmentsResponse]{
		More: func(page HCRPReportsClientListByConfigurationProfileAssignmentsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *HCRPReportsClientListByConfigurationProfileAssignmentsResponse) (HCRPReportsClientListByConfigurationProfileAssignmentsResponse, error) {
			req, err := client.listByConfigurationProfileAssignmentsCreateRequest(ctx, resourceGroupName, machineName, configurationProfileAssignmentName, options)
			if err != nil {
				return HCRPReportsClientListByConfigurationProfileAssignmentsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HCRPReportsClientListByConfigurationProfileAssignmentsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HCRPReportsClientListByConfigurationProfileAssignmentsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByConfigurationProfileAssignmentsHandleResponse(resp)
		},
	})
}

// listByConfigurationProfileAssignmentsCreateRequest creates the ListByConfigurationProfileAssignments request.
func (client *HCRPReportsClient) listByConfigurationProfileAssignmentsCreateRequest(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string, options *HCRPReportsClientListByConfigurationProfileAssignmentsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByConfigurationProfileAssignmentsHandleResponse handles the ListByConfigurationProfileAssignments response.
func (client *HCRPReportsClient) listByConfigurationProfileAssignmentsHandleResponse(resp *http.Response) (HCRPReportsClientListByConfigurationProfileAssignmentsResponse, error) {
	result := HCRPReportsClientListByConfigurationProfileAssignmentsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportList); err != nil {
		return HCRPReportsClientListByConfigurationProfileAssignmentsResponse{}, err
	}
	return result, nil
}
