// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
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

package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getGRPCClient() (*grpc.ClientConn, error) {
	address := flag.String("url", "localhost:50052", "awi-infra-guard GRPC server address")
	flag.Parse()
	fmt.Printf("connecting to %s\n", *address)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, *address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to grpc server: %v", err)
	}
	fmt.Println("connected")
	return conn, nil
}

func main() {
	conn, err := getGRPCClient()
	if err != nil {
		panic(err)
	}
	client := infrapb.NewCloudProviderServiceClient(conn)

	response, err := client.ListInstances(context.Background(), &infrapb.ListInstancesRequest{
		Provider: "gcp",
		Labels:   map[string]string{"env": "development"},
	})
	for _, instance := range response.Instances {
		fmt.Println("instance", instance)
	}

	acClient := infrapb.NewAccessControlServiceClient(conn)
	fmt.Println("adding inbound rule to instances in development VPC with label app_type:database")
	resp, err := acClient.AddInboundAllowRuleByLabelsMatch(context.TODO(),
		&infrapb.AddInboundAllowRuleByLabelsMatchRequest{
			Provider:          "gcp",
			VpcId:             "development",
			RuleName:          "awi-rule",
			Labels:            map[string]string{"app_type": "database"},
			CidrsToAllow:      []string{"192.168.0.0/16", "10.0.0.1/32"},
			ProtocolsAndPorts: map[string]*infrapb.Ports{"icmp": {}, "tcp": {Ports: []string{"99", "100-120", "3306"}}},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("rule id", resp.GetRuleId())
	fmt.Println("matched instances", resp.GetInstances())

	k8sClient := infrapb.NewKubernetesServiceClient(conn)
	fmt.Println("listing pods and services in some cluster")
	podresp, err := k8sClient.ListPods(context.TODO(), &infrapb.ListPodsRequest{
		ClusterName: "gke_gcp-ibngctopoc-nprd-72084_us-east1-b_ml-dataset-cluster",
	})
	for _, pod := range podresp.GetPods() {
		fmt.Println("Pod:", pod)
	}
	if err != nil {
		panic(err)
	}
	servresp, err := k8sClient.ListServices(context.TODO(), &infrapb.ListServicesRequest{
		ClusterName: "gke_gcp-ibngctopoc-nprd-72084_us-east1-b_ml-dataset-cluster",
	})
	if err != nil {
		panic(err)
	}
	for _, svc := range servresp.GetServices() {
		fmt.Println("Service:", svc)
	}
}
