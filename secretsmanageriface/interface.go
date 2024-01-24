package secretsmanageriface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretsManagerAPI interface {
	BatchGetSecretValue(context.Context, *secretsmanager.BatchGetSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.BatchGetSecretValueOutput, error)
	CancelRotateSecret(context.Context, *secretsmanager.CancelRotateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.CancelRotateSecretOutput, error)
	CreateSecret(context.Context, *secretsmanager.CreateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error)
	DeleteResourcePolicy(context.Context, *secretsmanager.DeleteResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.DeleteResourcePolicyOutput, error)
	DeleteSecret(context.Context, *secretsmanager.DeleteSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.DeleteSecretOutput, error)
	DescribeSecret(context.Context, *secretsmanager.DescribeSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error)
	GetRandomPassword(context.Context, *secretsmanager.GetRandomPasswordInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetRandomPasswordOutput, error)
	GetResourcePolicy(context.Context, *secretsmanager.GetResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error)
	GetSecretValue(context.Context, *secretsmanager.GetSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
	ListSecretVersionIds(context.Context, *secretsmanager.ListSecretVersionIdsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecrets(context.Context, *secretsmanager.ListSecretsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)
	PutResourcePolicy(context.Context, *secretsmanager.PutResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.PutResourcePolicyOutput, error)
	PutSecretValue(context.Context, *secretsmanager.PutSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.PutSecretValueOutput, error)
	RemoveRegionsFromReplication(context.Context, *secretsmanager.RemoveRegionsFromReplicationInput, ...func(*secretsmanager.Options)) (*secretsmanager.RemoveRegionsFromReplicationOutput, error)
	ReplicateSecretToRegions(context.Context, *secretsmanager.ReplicateSecretToRegionsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ReplicateSecretToRegionsOutput, error)
	RestoreSecret(context.Context, *secretsmanager.RestoreSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.RestoreSecretOutput, error)
	RotateSecret(context.Context, *secretsmanager.RotateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.RotateSecretOutput, error)
	StopReplicationToReplica(context.Context, *secretsmanager.StopReplicationToReplicaInput, ...func(*secretsmanager.Options)) (*secretsmanager.StopReplicationToReplicaOutput, error)
	TagResource(context.Context, *secretsmanager.TagResourceInput, ...func(*secretsmanager.Options)) (*secretsmanager.TagResourceOutput, error)
	UntagResource(context.Context, *secretsmanager.UntagResourceInput, ...func(*secretsmanager.Options)) (*secretsmanager.UntagResourceOutput, error)
	UpdateSecret(context.Context, *secretsmanager.UpdateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error)
	UpdateSecretVersionStage(context.Context, *secretsmanager.UpdateSecretVersionStageInput, ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretVersionStageOutput, error)
	ValidateResourcePolicy(context.Context, *secretsmanager.ValidateResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.ValidateResourcePolicyOutput, error)
}

var _ SecretsManagerAPI = (*secretsmanager.Client)(nil)
