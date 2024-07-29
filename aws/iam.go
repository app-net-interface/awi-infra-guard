// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package aws

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go"
)

type Auth struct {
	iam *iam.Client
}

// Policy
type PolicyDocument struct {
	Version   string
	Statement []PolicyStatement
}

// PolicyStatement defines a statement in a policy document.
type PolicyStatement struct {
	Effect    string
	Action    []string
	Principal map[string]string `json:",omitempty"`
	Resource  *string           `json:",omitempty"`
}

func (auth *Auth) ListPolicies(maxPolicies int32) ([]types.Policy, error) {
	var policies []types.Policy
	result, err := auth.iam.ListPolicies(context.TODO(), &iam.ListPoliciesInput{
		MaxItems: aws.Int32(maxPolicies),
	})
	if err != nil {
		log.Printf("Couldn't list policies. Here's why: %v\n", err)
	} else {
		policies = result.Policies
	}
	return policies, err
}

func (auth *Auth) CreatePolicy(policyName string, actions []string,
	resourceArn string) (*types.Policy, error) {
	var policy *types.Policy
	policyDoc := PolicyDocument{
		Version: "2012-10-17",
		Statement: []PolicyStatement{{
			Effect:   "Allow",
			Action:   actions,
			Resource: aws.String(resourceArn),
		}},
	}
	policyBytes, err := json.Marshal(policyDoc)
	if err != nil {
		log.Printf("Couldn't create policy document for %v. Here's why: %v\n", resourceArn, err)
		return nil, err
	}
	result, err := auth.iam.CreatePolicy(context.TODO(), &iam.CreatePolicyInput{
		PolicyDocument: aws.String(string(policyBytes)),
		PolicyName:     aws.String(policyName),
	})
	if err != nil {
		log.Printf("Couldn't create policy %v. Here's why: %v\n", policyName, err)
	} else {
		policy = result.Policy
	}
	return policy, err
}

func (auth *Auth) GetPolicy(policyArn string) (*types.Policy, error) {
	var policy *types.Policy
	result, err := auth.iam.GetPolicy(context.TODO(), &iam.GetPolicyInput{
		PolicyArn: aws.String(policyArn),
	})
	if err != nil {
		log.Printf("Couldn't get policy %v. Here's why: %v\n", policyArn, err)
	} else {
		policy = result.Policy
	}
	return policy, err
}

func (auth *Auth) DeletePolicy(policyArn string) error {
	_, err := auth.iam.DeletePolicy(context.TODO(), &iam.DeletePolicyInput{
		PolicyArn: aws.String(policyArn),
	})
	if err != nil {
		log.Printf("Couldn't delete policy %v. Here's why: %v\n", policyArn, err)
	}
	return err
}

//Roles

func (auth *Auth) ListRoles(maxRoles int32) ([]types.Role, error) {
	var roles []types.Role
	result, err := auth.iam.ListRoles(context.TODO(),
		&iam.ListRolesInput{MaxItems: aws.Int32(maxRoles)},
	)
	if err != nil {
		log.Printf("Couldn't list roles. Here's why: %v\n", err)
	} else {
		roles = result.Roles
	}
	return roles, err
}

func (auth *Auth) CreateRole(roleName string, trustedUserArn string) (*types.Role, error) {
	var role *types.Role
	trustPolicy := PolicyDocument{
		Version: "2012-10-17",
		Statement: []PolicyStatement{{
			Effect:    "Allow",
			Principal: map[string]string{"AWS": trustedUserArn},
			Action:    []string{"sts:AssumeRole"},
		}},
	}
	policyBytes, err := json.Marshal(trustPolicy)
	if err != nil {
		log.Printf("Couldn't create trust policy for %v. Here's why: %v\n", trustedUserArn, err)
		return nil, err
	}
	result, err := auth.iam.CreateRole(context.TODO(), &iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(string(policyBytes)),
		RoleName:                 aws.String(roleName),
	})
	if err != nil {
		log.Printf("Couldn't create role %v. Here's why: %v\n", roleName, err)
	} else {
		role = result.Role
	}
	return role, err
}

func (auth *Auth) GetRole(roleName string) (*types.Role, error) {
	var role *types.Role
	result, err := auth.iam.GetRole(context.TODO(),
		&iam.GetRoleInput{RoleName: aws.String(roleName)})
	if err != nil {
		log.Printf("Couldn't get role %v. Here's why: %v\n", roleName, err)
	} else {
		role = result.Role
	}
	return role, err
}

func (auth *Auth) CreateServiceLinkedRole(serviceName string, description string) (*types.Role, error) {
	var role *types.Role
	result, err := auth.iam.CreateServiceLinkedRole(context.TODO(), &iam.CreateServiceLinkedRoleInput{
		AWSServiceName: aws.String(serviceName),
		Description:    aws.String(description),
	})
	if err != nil {
		log.Printf("Couldn't create service-linked role %v. Here's why: %v\n", serviceName, err)
	} else {
		role = result.Role
	}
	return role, err
}

func (auth *Auth) DeleteServiceLinkedRole(roleName string) error {
	_, err := auth.iam.DeleteServiceLinkedRole(context.TODO(), &iam.DeleteServiceLinkedRoleInput{
		RoleName: aws.String(roleName)},
	)
	if err != nil {
		log.Printf("Couldn't delete service-linked role %v. Here's why: %v\n", roleName, err)
	}
	return err
}

func (auth *Auth) AttachRolePolicy(policyArn string, roleName string) error {
	_, err := auth.iam.AttachRolePolicy(context.TODO(), &iam.AttachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(roleName),
	})
	if err != nil {
		log.Printf("Couldn't attach policy %v to role %v. Here's why: %v\n", policyArn, roleName, err)
	}
	return err
}

func (auth *Auth) ListAttachedRolePolicies(roleName string) ([]types.AttachedPolicy, error) {
	var policies []types.AttachedPolicy
	result, err := auth.iam.ListAttachedRolePolicies(context.TODO(), &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(roleName),
	})
	if err != nil {
		log.Printf("Couldn't list attached policies for role %v. Here's why: %v\n", roleName, err)
	} else {
		policies = result.AttachedPolicies
	}
	return policies, err
}

func (auth *Auth) DetachRolePolicy(roleName string, policyArn string) error {
	_, err := auth.iam.DetachRolePolicy(context.TODO(), &iam.DetachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(roleName),
	})
	if err != nil {
		log.Printf("Couldn't detach policy from role %v. Here's why: %v\n", roleName, err)
	}
	return err
}

func (auth *Auth) ListRolePolicies(roleName string) ([]string, error) {
	var policies []string
	result, err := auth.iam.ListRolePolicies(context.TODO(), &iam.ListRolePoliciesInput{
		RoleName: aws.String(roleName),
	})
	if err != nil {
		log.Printf("Couldn't list policies for role %v. Here's why: %v\n", roleName, err)
	} else {
		policies = result.PolicyNames
	}
	return policies, err
}

func (auth *Auth) DeleteRole(roleName string) error {
	_, err := auth.iam.DeleteRole(context.TODO(), &iam.DeleteRoleInput{
		RoleName: aws.String(roleName),
	})
	if err != nil {
		log.Printf("Couldn't delete role %v. Here's why: %v\n", roleName, err)
	}
	return err
}

// AssumeRole assumes the specified role and returns the credentials.
func (auth *Auth) AssumeRole(ctx context.Context, cfg aws.Config, roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	svc := sts.NewFromConfig(cfg)
	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
	}
	result, err := svc.AssumeRole(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Users

func (auth *Auth) ListUsers(maxUsers int32) ([]types.User, error) {
	var users []types.User
	result, err := auth.iam.ListUsers(context.TODO(), &iam.ListUsersInput{
		MaxItems: aws.Int32(maxUsers),
	})
	if err != nil {
		log.Printf("Couldn't list users. Here's why: %v\n", err)
	} else {
		users = result.Users
	}
	return users, err
}

func (auth *Auth) GetUser(userName string) (*types.User, error) {
	var user *types.User
	result, err := auth.iam.GetUser(context.TODO(), &iam.GetUserInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NoSuchEntityException:
				log.Printf("User %v does not exist.\n", userName)
				err = nil
			default:
				log.Printf("Couldn't get user %v. Here's why: %v\n", userName, err)
			}
		}
	} else {
		user = result.User
	}
	return user, err
}

func (auth *Auth) CreateUser(userName string) (*types.User, error) {
	var user *types.User
	result, err := auth.iam.CreateUser(context.TODO(), &iam.CreateUserInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't create user %v. Here's why: %v\n", userName, err)
	} else {
		user = result.User
	}
	return user, err
}

func (auth *Auth) CreateUserPolicy(userName string, policyName string, actions []string,
	roleArn string) error {
	policyDoc := PolicyDocument{
		Version: "2012-10-17",
		Statement: []PolicyStatement{{
			Effect:   "Allow",
			Action:   actions,
			Resource: aws.String(roleArn),
		}},
	}
	policyBytes, err := json.Marshal(policyDoc)
	if err != nil {
		log.Printf("Couldn't create policy document for %v. Here's why: %v\n", roleArn, err)
		return err
	}
	_, err = auth.iam.PutUserPolicy(context.TODO(), &iam.PutUserPolicyInput{
		PolicyDocument: aws.String(string(policyBytes)),
		PolicyName:     aws.String(policyName),
		UserName:       aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't create policy for user %v. Here's why: %v\n", userName, err)
	}
	return err
}

func (auth *Auth) ListUserPolicies(userName string) ([]string, error) {
	var policies []string
	result, err := auth.iam.ListUserPolicies(context.TODO(), &iam.ListUserPoliciesInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't list policies for user %v. Here's why: %v\n", userName, err)
	} else {
		policies = result.PolicyNames
	}
	return policies, err
}

func (auth *Auth) DeleteUserPolicy(userName string, policyName string) error {
	_, err := auth.iam.DeleteUserPolicy(context.TODO(), &iam.DeleteUserPolicyInput{
		PolicyName: aws.String(policyName),
		UserName:   aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't delete policy from user %v. Here's why: %v\n", userName, err)
	}
	return err
}

func (auth *Auth) DeleteUser(userName string) error {
	_, err := auth.iam.DeleteUser(context.TODO(), &iam.DeleteUserInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't delete user %v. Here's why: %v\n", userName, err)
	}
	return err
}

func (auth *Auth) CreateAccessKeyPair(userName string) (*types.AccessKey, error) {
	var key *types.AccessKey
	result, err := auth.iam.CreateAccessKey(context.TODO(), &iam.CreateAccessKeyInput{
		UserName: aws.String(userName)})
	if err != nil {
		log.Printf("Couldn't create access key pair for user %v. Here's why: %v\n", userName, err)
	} else {
		key = result.AccessKey
	}
	return key, err
}

func (auth *Auth) DeleteAccessKey(userName string, keyId string) error {
	_, err := auth.iam.DeleteAccessKey(context.TODO(), &iam.DeleteAccessKeyInput{
		AccessKeyId: aws.String(keyId),
		UserName:    aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't delete access key %v. Here's why: %v\n", keyId, err)
	}
	return err
}

func (auth *Auth) ListAccessKeys(userName string) ([]types.AccessKeyMetadata, error) {
	var keys []types.AccessKeyMetadata
	result, err := auth.iam.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't list access keys for user %v. Here's why: %v\n", userName, err)
	} else {
		keys = result.AccessKeyMetadata
	}
	return keys, err
}
