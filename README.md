# awi-infra-guard

SDK/GRPC service to fetch infrastructure resource information and push updates to multiple infrastructure provider such as AWS, GCP, AZURE, VMWare and ACI.

## Supported infrastructure providers

Currently supported providers:

- AWS
- Google Cloud Platform (GCP).

## Kubernetes support

Kubernetes clusters operations are supported. Optionally, clusters information can be provided in kube config file present in
HOME/.kube/config. EKS and GKE clusters should be discovered automatically.

## awi-infra-guard as a library or as a service

awi-infra-guard can be used an imported Go library or as a standalone GRPC service.

### Credentials configuration

#### AWS credentials

Setup .aws/configuration file in your home directory or specify environment variables based on instruction from AWS
[guide](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials).

Multiple accounts are supported, they can be configured using profiles in credentials file, instructions can be found in
"Specifying profiles" section in [guide](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials).

#### GCP credentials

Setup application default credentials based on instructions from GCP [guide](https://cloud.google.com/docs/authentication/application-default-credentials).
Multiple projects are supported, for instructions how to specify them check "awi-infra-guard as a library" and "awi-infra-guard as a service"
sections.

### awi-infra-guard as a library

To use awi-infra-guard as a library import github.com/app-net-interface/awi-infra-guard package:

```sh
go get github.com/app-net-interface/awi-infra-guard@develop
```

Initialize provider strategy and use it for calling requests as in an example below:

```go
package main

import (
    "context"
    "fmt"

    "github.com/sirupsen/logrus"
    "github.com/app-net-interface/awi-infra-guard/provider"
)

func main() {
    ctx := context.Background()
    providerStrategy := provider.NewRealProviderStrategy(ctx, logrus.New(), "")

    awsProvider, err := providerStrategy.GetProvider(context.TODO(), "aws")
    if err != nil {
        panic(err)
    }
    instances, err := awsProvider.ListInstances(context.TODO(), &infrapb.ListInstancesRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println("Instances in AWS:")
    for _, instance := range instances {
        fmt.Println(instance.VPCID, instance.Name)
    }

    gcpProvider, err := providerStrategy.GetProvider(context.TODO(), "gcp")
    if err != nil {
        panic(err)
    }
    instances, err = gcpProvider.ListInstances(context.TODO(), &infrapb.ListInstancesRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println("Instances in GCP:")
    for _, instance := range instances {
        fmt.Println(instance.VPCID, instance.Name)
    }
}
```

### awi-infra-guard as a service

To run awi-infra-guard as a separate service you can start it using `make run` command.

Example:

```sh
$ make run
go run main.go
INFO[0000] server listening at [::]:50052
```

You can connect to this server using [grpc_cli tool](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md).
Example:

```sh
$ grpc_cli call localhost:50052 ListInstances "provider: 'aws', vpc_id: 'vpc-04a1eaad3aa81310f'"
connecting to localhost:50052
instances {
            "labels": [
                {
                    "key": "CreatedBy",
                    "value": "terraform"
                },
                {
                    "key": "Name",
                    "value": "sdwan-vmanage-00"
                },
                {
                    "key": "owner",
                    "value": "natal"
                },
                {
                    "key": "project",
                    "value": "xxx"
                },
                {
                    "key": "ApplicationName",
                    "value": "Cisco SD-WAN Control Plane"
                }
            ],
            "id": "i-007755280666fdca5",
            "name": "sdwan-vmanage-00",
            "publicIP": "",
            "privateIP": "10.128.0.32",
            "subnetID": "subnet-xxx",
            "project": "",
            "vpcId": "vpc-xxx",
            "region": "us-west-2",
            "zone": "us-west-2a",
            "provider": "AWS",
            "account_id": "xxxxxx",
            "state": "running",
            "type": "c5.4xlarge",
            "last_sync_time": "2024-06-14T17:25:30Z",
            "self_link": "https://us-west-2.console.aws.amazon.com/xxxx"
}

Rpc succeeded with OK status

$ grpc_cli call localhost:50052 ListClusters ""
connecting to localhost:50052
clusters {
  name: "gke-demo-cluster"
}
clusters {
  name: "eks-awi-demo"
}
clusters {
  name: "kind-awi"
}

$ grpc_cli call localhost:50052 ListPods "cluster_name: 'eks-awi-demo'"
connecting to localhost:50052
pods {
  cluster: "eks-awi-demo"
  namespace: "kube-system"
  name: "coredns-6ff9c46cd8-m8lwv"
  labels {
    key: "eks.amazonaws.com/component"
    value: "coredns"
  }
  labels {
    key: "k8s-app"
    value: "kube-dns"
  }
  labels {
    key: "pod-template-hash"
    value: "6ff9c46cd8"
  }
}
pods {
  cluster: "eks-awi-demo"
  namespace: "kube-system"
  name: "coredns-6ff9c46cd8-s4b95"
  labels {
    key: "eks.amazonaws.com/component"
    value: "coredns"
  }
  labels {
    key: "k8s-app"
    value: "kube-dns"
  }
  labels {
    key: "pod-template-hash"
    value: "6ff9c46cd8"
  }
}
Rpc succeeded with OK status
```

Example Go client usage can be found in example/client directory:

```sh
$ cd example/client
$ go run main.go
connecting to localhost:50052
connected
instance ID:"4894037167304189131" Name:"development-dashboard-1" PublicIP:"35.212.252.162" PrivateIP:"10.150.0.2" SubnetID:"development-subnet-1" VPCID:"development"
instance ID:"8825713928722555929" Name:"development-database-1" PublicIP:"35.212.129.188" PrivateIP:"10.150.0.3" SubnetID:"development-subnet-1" VPCID:"development"
instance ID:"7411617185127835047" Name:"development-database-2" PublicIP:"35.212.176.237" PrivateIP:"10.150.0.4" SubnetID:"development-subnet-1" VPCID:"development"
instance ID:"258418092159915173" Name:"development-database-3" PublicIP:"35.212.218.134" PrivateIP:"10.150.0.7" SubnetID:"development-subnet-1" VPCID:"development"
adding inbound rule to instances in development VPC with label app_type:database
rule id 3114023319057261683
matched instances IDs [8825713928722555929 7411617185127835047 258418092159915173]
```

## Docker instructions

### Building and pushing image

To build your image:

```sh
make docker-build IMG=<your-repo>/<name>
```

To push it to your repository:

```sh
make docker-push IMG=<your-repo>/<name>
```

> ℹ️ Info: You can also do both steps at once with
> `make docker-build docker-push IMG=<your-repo>/<name>`

### Running docker image

The awi-infra-guard accepts following files:

- `/root/config/config.yaml` - the configuration file
- `/root/.aws/credentials` - the credentials for AWS
- `/app/gcp-key/gcp-key.json` - the credentials for GCP
- `/root/.kube/config` - configuration and credentials for k8s cluster

In order tp configure and gain access for different providers for awi-infra-guard
one need to mount these files while starting container.

## Contributing

Thank you for interest in contributing! Please refer to our
[contributing guide](CONTRIBUTING.md).

## License

awi-infra-guard is released under the Apache 2.0 license. See
[LICENSE](./LICENSE).

awi-infra-guard is also made possible thanks to
[third party open source projects](NOTICE).
