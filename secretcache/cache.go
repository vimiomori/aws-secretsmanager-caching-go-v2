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

// Package secretcache provides the Cache struct for in-memory caching of secrets stored in AWS Secrets Manager
// Also exports a CacheHook, for pre-store and post-fetch processing of cached values
package secretcache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/vimiomori/aws-secretsmanager-caching-go-v2/secretsmanageriface"
)

// Cache client for AWS Secrets Manager secrets.
type Cache struct {
	lru *lruCache
	CacheConfig
	Client secretsmanageriface.SecretsManagerAPI
}

// New constructs a secret cache using functional options, uses defaults otherwise.
// Initialises a SecretsManager Client from a new session.Session.
// Initialises CacheConfig to default values.
// Initialises lru cache with a default max size.
func New(optFns ...func(*Cache)) (*Cache, error) {
	cache := &Cache{
		// Initialise default configuration
		CacheConfig: CacheConfig{
			MaxCacheSize: DefaultMaxCacheSize,
			VersionStage: DefaultVersionStage,
			CacheItemTTL: DefaultCacheItemTTL,
		},
	}

	// Iterate over options allowing user to specify alternate
	// configurations.
	for _, optFn := range optFns {
		optFn(cache)
	}

	// Initialise lru cache
	cache.lru = newLRUCache(cache.MaxCacheSize)

	// Initialise the secrets manager client
	if cache.Client == nil {
		cfg, err := config.LoadDefaultConfig(context.Background())
		if err != nil {
			return nil, err
		}

		cache.Client = secretsmanager.NewFromConfig(cfg)
	}

	return cache, nil
}

// getCachedSecret gets a cached secret for the given secret identifier.
// Returns cached secret item.
func (c *Cache) getCachedSecret(secretId string) *secretCacheItem {
	lruValue, found := c.lru.get(secretId)

	if !found {
		cacheItem := newSecretCacheItem(c.CacheConfig, c.Client, secretId)
		c.lru.putIfAbsent(secretId, &cacheItem)
		lruValue, _ = c.lru.get(secretId)
	}

	secretCacheItem, _ := lruValue.(*secretCacheItem)
	return secretCacheItem
}

// GetSecretString gets the secret string value from the cache for given secret id and a default version stage.
// Returns the secret string and an error if operation failed.
func (c *Cache) GetSecretString(ctx context.Context, secretId string) (string, error) {
	return c.GetSecretStringWithStage(ctx, secretId, DefaultVersionStage)
}

// GetSecretStringWithStage gets the secret string value from the cache for given secret id and version stage.
// Returns the secret sting and an error if operation failed.
func (c *Cache) GetSecretStringWithStage(ctx context.Context, secretId string, versionStage string) (string, error) {
	secretCacheItem := c.getCachedSecret(secretId)

	getSecretValueOutput, err := secretCacheItem.getSecretValue(ctx, versionStage)
	if err != nil {
		return "", err
	}

	if getSecretValueOutput.SecretString == nil {
		return "", &InvalidOperationError{
			baseError{
				Message: "requested secret version does not contain SecretString",
			},
		}
	}

	return *getSecretValueOutput.SecretString, nil
}

// GetSecretBinary gets the secret binary value from the cache for given secret id and a default version stage.
// Returns the secret binary and an error if operation failed.
func (c *Cache) GetSecretBinary(ctx context.Context, secretId string) ([]byte, error) {
	return c.GetSecretBinaryWithStage(ctx, secretId, DefaultVersionStage)
}

// GetSecretBinaryWithStage gets the secret binary value from the cache for given secret id and version stage.
// Returns the secret binary and an error if operation failed.
func (c *Cache) GetSecretBinaryWithStage(ctx context.Context, secretId string, versionStage string) ([]byte, error) {
	secretCacheItem := c.getCachedSecret(secretId)

	getSecretValueOutput, err := secretCacheItem.getSecretValue(ctx, versionStage)
	if err != nil {
		return nil, err
	}

	if getSecretValueOutput.SecretBinary == nil {
		return nil, &InvalidOperationError{
			baseError{
				Message: "requested secret version does not contain SecretBinary",
			},
		}
	}

	return getSecretValueOutput.SecretBinary, nil
}
