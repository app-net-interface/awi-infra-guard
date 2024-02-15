# Cloud Service Provider Connector

The connector package provides a CLI for establishing a connection
between two Cloud Service Providers.

The current implementation supports creating a connection only between
AWS and GCP.

The CSP Connector uses Bolt DB to store information about existing
connections. Every internal interaction with the database will load
or create `connections.db` file in the location from where the binary
was run.

# Usage

The CSP connector focuses mainly on 4 operations

1. connect - creates a connection between two Gateways or attempts to
    recreate it if the connection is in invalid state (for example due to
    sudden `connect` interrupt).

    ```
    ./cspConnector connect SOURCE_PROVIDER SOURCE_GW_ID SOURCE_REGION DEST_PROVIDER DEST_GW_ID DEST_REGION
    ```

    The CSP Connector requires the specification of both Gateways that
    should be connected. To do so, specify the name of the provider, the
    Gateway identifier and the region for the source and destination.

    Example:
    ```
    ./cspConnector connect aws tgw-0af95eec1e172b0fa us-west-1 gcp refactored-router us-east4
    ```

    The connection between providers is bidirectional so deciding whether
    which one is the source and which one is a destination does not matter.

    To check the details of available gateways run `list gateways` to see
    which Gateways can be used.

    The connect method creates an entry in the local database with the
    information about the status of created connection. To see those
    entries run `list connections`.

    The connect method can be extended to accept `CONNECTION_ID` argument
    as well for telling CSP Connector to recreate the connection if it
    did not complete successfully. Currently, the CSP Connector will
    not do anything for connections with state `ACTIVE` but it can be
    changed to let providers inspect if their sides of connection are not
    suddenly missing some certain resources (that could've been removed
    in the background by the Cloud Administrator or someone else).

1. disconnect - destroys the connection between two Gateways that was
    created with the CSP Connector.

    ```
    ./cspConnector disconnect CONNECTION_ID
    ```

    As opposite to `connect` method, the `disconnect` method expects only
    the information about Connection ID that should be erased. The CSP
    Connector will look up for that particular entry in the local
    database and obtain the information about both sides of the connection.

    Example:
    ```
    ./cspConnector disconnect 083b23ff-0323-4b6b-a593-4ad2d9ada516
    ```

    To check IDs of existing connections run `list connections`.

    The successful erase of connection will remove the entry from the
    database.

    The disconnect method should be improved with option to provide details
    of source and destination in case if the connection is missing in the
    local database to do an attempt of deleting resources with the use
    of provider's logic (current two supported providers AWS and GCP are
    capable of checking what resources were created by them).

1. list gateways - goes through available and authenticated providers and
    list Gateways provided by them.

1. list connections - lists all connections that are either active, are
    in the progress of creation/deletion or failed at some stage.

    As opposite to other commands, listing connections doesn't require
    access to any Cloud Provider. It entirely relies on the Bolt DB file.

    If the `connections.db` file doesn't exist in the current location,
    listing connections will create a new database file and provide empty
    list of connections.

# Configuration

Running CSP Connector requires a configuration file `csp.yaml` present
in the current location. The filepath is currently hardcoded but is
intended to be specified with an argument or environment variable.

The configuration specifies details of Logger and custom configurations
for each provider available by the CSP Connector (native ones and external
libraries provided by the user) - each provider configuration may vary and
expect different options - to check what should be configured for a given
provider check its `config.go` file.

Example configuration file:

```
logger:
  level: DEBUG
  componentLevels:
    gcpClient: ERROR
  output:
    stdout: true
    file: csp.log

providers:
  aws:
    region: "us-west-1"
  gcp:
    region: "us-east4"
    project: "MY_PROJECT"
```

# Logger

The logger is not implemented completely yet.

Currently, the logger uses trace log with no option to change. Also, it
doesn't support logging to file just yet.

The desired configuration for Logger is already defined, which allows
specifying global log level and different log levels for specified loggers.
It needs to be used by the Logger Configuration.

The CSP Connector defines a single main logger and splits that logger into
multiple subloggers for different pieces of code such as ProviderManager,
provider implementations etc. Each log entry lists field `logger` which
provides information about what logger reported that information.

When implementing a code responsible for writing to a file, the logger
should also define logic explicitely for situation when the device goes
out of space.

# Test it out

**Warning: These scripts will attempt creating resources in both AWS and**
**GCP using your local CLIs with accessible authentication. Do NOT run**
**them if you're authenticated to Providers, which you don't want to use**
**(unless you know what you're doing). These scripts are not meant to be**
**entirely bugproof so expect unexpected. It's best to check them out**
**before using them.**

You can use the bash script `demo/create_env.sh` to quickly spawn resources
in AWS and GCP representing single Gateway on both sides along with single
Virtual Machine in both providers.

The script will configure both VMs to allow SSH with the following rules:
* the AWS provider uses generated keypair that will be stored in the
    demo directory after running the script
* the GCP provider requires Service Account and assumes that you will use
    gcloud cli to access the VM

The CSP Connector and the demo script WON'T:
* attach AWS VPC to AWS Transit Gateway
* create routing rules to allow AWS reach GCP and the opposite direction
* create firewall rules allowing accessing each other

These steps are performed either by AUSM by creating Network Domain Connection
between VPCs and then creating proper App Connection or can be done manually.

# Adding new providers

The CSP Connector is being developed with a focus on making a new providers
easy. Defining a new provider means adding a new structure implementing
`provider/provider.go` interface.

Apart from extending CSP Connector codebase, the Provider Manager has a
placeholder defined for adding a logic to handle Golang Plugins. With
that approach implemented, the CSP Connector binary could be run with
additional libraries provided representing additional provider
implementations. 
