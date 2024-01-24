// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You
// may not use this file except in compliance with the License. A copy of
// the License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is
// distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF
// ANY KIND, either express or implied. See the License for the specific
// language governing permissions and limitations under the License.

package secretcache_test

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// A struct to be used in unit tests as a mock Client
type mockSecretsManagerClient struct {
	*secretsmanager.Client
	MockedGetResult         *secretsmanager.GetSecretValueOutput
	MockedDescribeResult    *secretsmanager.DescribeSecretOutput
	GetSecretValueErr       error
	DescribeSecretErr       error
	GetSecretValueCallCount int
	DescribeSecretCallCount int
}

// Initialises a mock Client with dummy outputs for GetSecretValue and DescribeSecret APIs
func newMockedClientWithDummyResults() (mockSecretsManagerClient, string, string) {
	createDate := time.Now().Add(-time.Hour * 12) // 12 hours ago
	versionId := getStrPtr("very-random-uuid")
	otherVersionId := getStrPtr("other-random-uuid")
	versionStages := []string{"hello", "versionStage-42", "AWSCURRENT"}
	otherVersionStages := []string{"AWSPREVIOUS"}
	versionIdsToStages := make(map[string][]string)
	versionIdsToStages[*versionId] = versionStages
	versionIdsToStages[*otherVersionId] = otherVersionStages
	secretId := getStrPtr("dummy-secret-name")
	secretString := getStrPtr("my secret string")

	mockedGetResult := secretsmanager.GetSecretValueOutput{
		ARN:           getStrPtr("dummy-arn"),
		CreatedDate:   &createDate,
		Name:          secretId,
		SecretString:  secretString,
		VersionId:     versionId,
		VersionStages: versionStages,
	}

	mockedDescribeResult := secretsmanager.DescribeSecretOutput{
		ARN:                getStrPtr("dummy-arn"),
		Name:               secretId,
		Description:        getStrPtr("my dummy description"),
		VersionIdsToStages: versionIdsToStages,
	}

	return mockSecretsManagerClient{
		MockedDescribeResult: &mockedDescribeResult,
		MockedGetResult:      &mockedGetResult,
	}, *secretId, *secretString
}

// Overrides the interface method to return dummy result.
func (m *mockSecretsManagerClient) GetSecretValue(context.Context, *secretsmanager.GetSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	m.GetSecretValueCallCount++

	if m.GetSecretValueErr != nil {
		return nil, m.GetSecretValueErr
	}

	return m.MockedGetResult, nil
}

// Overrides the interface method to return dummy result.
func (m *mockSecretsManagerClient) DescribeSecret(context.Context, *secretsmanager.DescribeSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error) {
	m.DescribeSecretCallCount++

	if m.DescribeSecretErr != nil {
		return nil, m.DescribeSecretErr
	}

	return m.MockedDescribeResult, nil
}

// Helper function to get a string pointer for input string.
func getStrPtr(str string) *string {
	return &str
}
