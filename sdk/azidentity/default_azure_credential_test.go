//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
)

func TestDefaultAzureCredential_GetTokenSuccess(t *testing.T) {
	env := map[string]string{azureTenantID: fakeTenantID, azureClientID: fakeClientID, azureClientSecret: secret}
	setEnvironmentVariables(t, env)
	cred, err := NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	c := cred.chain.sources[0].(*EnvironmentCredential)
	c.cred.(*ClientSecretCredential).client = fakeConfidentialClient{}
	_, err = cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{"scope"}})
	if err != nil {
		t.Fatalf("GetToken error: %v", err)
	}
}

func TestDefaultAzureCredential_ConstructorErrorHandler(t *testing.T) {
	setEnvironmentVariables(t, map[string]string{"AZURE_SDK_GO_LOGGING": "all"})
	errorMessages := []string{
		"<credential-name>: <error-message>",
		"<credential-name>: <error-message>",
	}
	err := defaultAzureCredentialConstructorErrorHandler(0, errorMessages)
	if err == nil {
		t.Fatalf("Expected an error, but received none.")
	}
	expectedError := `<credential-name>: <error-message>
	<credential-name>: <error-message>`
	if err.Error() != expectedError {
		t.Fatalf("Did not create an appropriate error message.\n\nReceived:\n%s\n\nExpected:\n%s", err.Error(), expectedError)
	}

	logMessages := []string{}
	log.SetListener(func(event log.Event, message string) {
		logMessages = append(logMessages, message)
	})

	err = defaultAzureCredentialConstructorErrorHandler(1, errorMessages)
	if err != nil {
		t.Fatal(err)
	}

	expectedLogs := `NewDefaultAzureCredential failed to initialize some credentials:
	<credential-name>: <error-message>
	<credential-name>: <error-message>`
	if len(logMessages) == 0 {
		t.Fatal("error handler logged no messages")
	}
	if logMessages[0] != expectedLogs {
		t.Fatalf("Did not receive the expected logs.\n\nReceived:\n%s\n\nExpected:\n%s", logMessages[0], expectedLogs)
	}
}

func TestDefaultAzureCredential_UserAssignedIdentity(t *testing.T) {
	for _, ID := range []ManagedIDKind{nil, ClientID("client-id")} {
		t.Run(fmt.Sprintf("%v", ID), func(t *testing.T) {
			if ID != nil {
				t.Setenv(azureClientID, ID.String())
			}
			cred, err := NewDefaultAzureCredential(nil)
			if err != nil {
				t.Fatal(err)
			}
			for _, c := range cred.chain.sources {
				if m, ok := c.(*ManagedIdentityCredential); ok {
					if actual := m.mic.id; actual != ID {
						t.Fatalf(`expected "%s", got "%v"`, ID, actual)
					}
					return
				}
			}
			t.Fatal("default chain should include ManagedIdentityCredential")
		})
	}
}

func TestDefaultAzureCredential_Workload(t *testing.T) {
	expectedAssertion := "service account token"
	tempFile := filepath.Join(t.TempDir(), "service-account-token-file")
	if err := os.WriteFile(tempFile, []byte(expectedAssertion), os.ModePerm); err != nil {
		t.Fatalf(`failed to write temporary file "%s": %v`, tempFile, err)
	}
	pred := func(req *http.Request) bool {
		if err := req.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if actual := req.PostForm["client_assertion"]; actual[0] != expectedAssertion {
			t.Fatalf(`unexpected assertion "%s"`, actual[0])
		}
		if actual := req.PostForm["client_id"]; actual[0] != fakeClientID {
			t.Fatalf(`unexpected assertion "%s"`, actual[0])
		}
		if actual := strings.Split(req.URL.Path, "/")[1]; actual != fakeTenantID {
			t.Fatalf(`unexpected tenant "%s"`, actual)
		}
		return true
	}
	srv, close := mock.NewServer(mock.WithTransformAllRequestsToTestServerUrl())
	defer close()
	srv.AppendResponse(mock.WithBody(instanceDiscoveryResponse))
	srv.AppendResponse(mock.WithBody(tenantDiscoveryResponse))
	srv.AppendResponse(mock.WithPredicate(pred), mock.WithBody(accessTokenRespSuccess))
	srv.AppendResponse()
	for k, v := range map[string]string{
		azureAuthorityHost:      cloud.AzurePublic.ActiveDirectoryAuthorityHost,
		azureClientID:           fakeClientID,
		azureFederatedTokenFile: tempFile,
		azureTenantID:           fakeTenantID,
	} {
		t.Setenv(k, v)
	}
	cred, err := NewDefaultAzureCredential(&DefaultAzureCredentialOptions{ClientOptions: policy.ClientOptions{Transport: srv}})
	if err != nil {
		t.Fatal(err)
	}
	testGetTokenSuccess(t, cred)
}
