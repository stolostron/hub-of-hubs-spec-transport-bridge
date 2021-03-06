[comment]: # ( Copyright Contributors to the Open Cluster Management project )

# Hub-of-Hubs Spec Transport Bridge

[![Go Report Card](https://goreportcard.com/badge/github.com/stolostron/hub-of-hubs-spec-transport-bridge)](https://goreportcard.com/report/github.com/stolostron/hub-of-hubs-spec-transport-bridge)
[![Go Reference](https://pkg.go.dev/badge/github.com/stolostron/hub-of-hubs-spec-transport-bridge.svg)](https://pkg.go.dev/github.com/stolostron/hub-of-hubs-spec-transport-bridge)
[![License](https://img.shields.io/github/license/stolostron/hub-of-hubs-spec-transport-bridge)](/LICENSE)

The spec transport bridge component of [Hub-of-Hubs](https://github.com/stolostron/hub-of-hubs).

Go to the [Contributing guide](CONTRIBUTING.md) to learn how to get involved.

## Getting Started

## Build and push the image to docker registry

1.  Set the `REGISTRY` environment variable to hold the name of your docker registry:
    ```
    $ export REGISTRY=...
    ```
    
1.  Set the `IMAGE_TAG` environment variable to hold the required version of the image.  
    default value is `latest`, so in that case no need to specify this variable:
    ```
    $ export IMAGE_TAG=latest
    ```
    
1.  Run make to build and push the image:
    ```
    $ make push-images
    ```

## Deploy on the hub of hubs

Set the `DATABASE_URL` according to the PostgreSQL URL format: `postgres://YourUserName:YourURLEscapedPassword@YourHostname:5432/YourDatabaseName?sslmode=verify-full&pool_max_conns=YourConnectionPoolSize`.

:exclamation: Remember to URL-escape the password, you can do it in bash:

```
python -c "import sys, urllib as ul; print ul.quote_plus(sys.argv[1])" 'YourPassword'
```

1.  Create a secret with your database url:

    ```
    kubectl create secret generic hub-of-hubs-database-transport-bridge-secret -n open-cluster-management --from-literal=url=$DATABASE_URL
    ```

1.  Set the `REGISTRY` environment variable to hold the name of your docker registry:
    ```
    $ export REGISTRY=...
    ```
    
1.  Set the `IMAGE` environment variable to hold the name of the image.

    ```
    $ export IMAGE=$REGISTRY/$(basename $(pwd)):latest
    ```

1. Set the `TRANSPORT_TYPE` environment variable to "kafka" or "sync-service" to set which transport to use.
    ```
    $ export TRANSPORT_TYPE=...
    ```
    
1.  Run the following command to deploy the `hub-of-hubs-spec-transport-bridge` to your hub of hubs cluster:  
    ```
    envsubst < deploy/hub-of-hubs-spec-transport-bridge.yaml.template | kubectl apply -f -
    ```
    
## Cleanup from the hub of hubs
    
1.  Run the following command to clean `hub-of-hubs-spec-transport-bridge` from your hub of hubs cluster:  
    ```
    envsubst < deploy/hub-of-hubs-spec-transport-bridge.yaml.template | kubectl delete -f -
    ```
