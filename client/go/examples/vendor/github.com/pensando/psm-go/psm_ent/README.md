# Go API client for psm_ent

Service name



## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./psm_ent"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*WorkloadV1Api* | [**AbortMigration**](docs/WorkloadV1Api.md#abortmigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/AbortMigration | Abort Workload Migration operation
*WorkloadV1Api* | [**AbortMigration1**](docs/WorkloadV1Api.md#abortmigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/AbortMigration | Abort Workload Migration operation
*WorkloadV1Api* | [**AddWorkload**](docs/WorkloadV1Api.md#addworkload) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads | Create Workload object
*WorkloadV1Api* | [**AddWorkload1**](docs/WorkloadV1Api.md#addworkload1) | **Post** /configs/workload/v1/workloads | Create Workload object
*WorkloadV1Api* | [**DeleteWorkload**](docs/WorkloadV1Api.md#deleteworkload) | **Delete** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name} | Delete Workload object
*WorkloadV1Api* | [**DeleteWorkload1**](docs/WorkloadV1Api.md#deleteworkload1) | **Delete** /configs/workload/v1/workloads/{O.Name} | Delete Workload object
*WorkloadV1Api* | [**FinalSyncMigration**](docs/WorkloadV1Api.md#finalsyncmigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/FinalSyncMigration | Initiates the final sync for the Workload Migration operation
*WorkloadV1Api* | [**FinalSyncMigration1**](docs/WorkloadV1Api.md#finalsyncmigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/FinalSyncMigration | Initiates the final sync for the Workload Migration operation
*WorkloadV1Api* | [**FinishMigration**](docs/WorkloadV1Api.md#finishmigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/FinishMigration | Finish Workload Migration operation
*WorkloadV1Api* | [**FinishMigration1**](docs/WorkloadV1Api.md#finishmigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/FinishMigration | Finish Workload Migration operation
*WorkloadV1Api* | [**GetEndpoint**](docs/WorkloadV1Api.md#getendpoint) | **Get** /configs/workload/v1/tenant/{O.Tenant}/endpoints/{O.Name} | Get Endpoint object
*WorkloadV1Api* | [**GetEndpoint1**](docs/WorkloadV1Api.md#getendpoint1) | **Get** /configs/workload/v1/endpoints/{O.Name} | Get Endpoint object
*WorkloadV1Api* | [**GetWorkload**](docs/WorkloadV1Api.md#getworkload) | **Get** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name} | Get Workload object
*WorkloadV1Api* | [**GetWorkload1**](docs/WorkloadV1Api.md#getworkload1) | **Get** /configs/workload/v1/workloads/{O.Name} | Get Workload object
*WorkloadV1Api* | [**LabelWorkload**](docs/WorkloadV1Api.md#labelworkload) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/label | Label Workload object
*WorkloadV1Api* | [**LabelWorkload1**](docs/WorkloadV1Api.md#labelworkload1) | **Post** /configs/workload/v1/workloads/{O.Name}/label | Label Workload object
*WorkloadV1Api* | [**ListEndpoint**](docs/WorkloadV1Api.md#listendpoint) | **Get** /configs/workload/v1/tenant/{O.Tenant}/endpoints | List Endpoint objects
*WorkloadV1Api* | [**ListEndpoint1**](docs/WorkloadV1Api.md#listendpoint1) | **Get** /configs/workload/v1/endpoints | List Endpoint objects
*WorkloadV1Api* | [**ListWorkload**](docs/WorkloadV1Api.md#listworkload) | **Get** /configs/workload/v1/tenant/{O.Tenant}/workloads | List Workload objects
*WorkloadV1Api* | [**ListWorkload1**](docs/WorkloadV1Api.md#listworkload1) | **Get** /configs/workload/v1/workloads | List Workload objects
*WorkloadV1Api* | [**StartMigration**](docs/WorkloadV1Api.md#startmigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/StartMigration | Start Workload Migration operation
*WorkloadV1Api* | [**StartMigration1**](docs/WorkloadV1Api.md#startmigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/StartMigration | Start Workload Migration operation
*WorkloadV1Api* | [**UpdateWorkload**](docs/WorkloadV1Api.md#updateworkload) | **Put** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name} | Update Workload object
*WorkloadV1Api* | [**UpdateWorkload1**](docs/WorkloadV1Api.md#updateworkload1) | **Put** /configs/workload/v1/workloads/{O.Name} | Update Workload object
*WorkloadV1Api* | [**WatchEndpoint**](docs/WorkloadV1Api.md#watchendpoint) | **Get** /configs/workload/v1/watch/tenant/{O.Tenant}/endpoints | Watch Endpoint objects. Supports WebSockets or HTTP long poll
*WorkloadV1Api* | [**WatchEndpoint1**](docs/WorkloadV1Api.md#watchendpoint1) | **Get** /configs/workload/v1/watch/endpoints | Watch Endpoint objects. Supports WebSockets or HTTP long poll
*WorkloadV1Api* | [**WatchWorkload**](docs/WorkloadV1Api.md#watchworkload) | **Get** /configs/workload/v1/watch/tenant/{O.Tenant}/workloads | Watch Workload objects. Supports WebSockets or HTTP long poll
*WorkloadV1Api* | [**WatchWorkload1**](docs/WorkloadV1Api.md#watchworkload1) | **Get** /configs/workload/v1/watch/workloads | Watch Workload objects. Supports WebSockets or HTTP long poll


## Documentation For Models

 - [ApiAggWatchOptions](docs/ApiAggWatchOptions.md)
 - [ApiKindWatchOptions](docs/ApiKindWatchOptions.md)
 - [ApiLabel](docs/ApiLabel.md)
 - [ApiListMeta](docs/ApiListMeta.md)
 - [ApiListWatchOptions](docs/ApiListWatchOptions.md)
 - [ApiObjectMeta](docs/ApiObjectMeta.md)
 - [ApiObjectRef](docs/ApiObjectRef.md)
 - [ApiStatus](docs/ApiStatus.md)
 - [ApiStatusResult](docs/ApiStatusResult.md)
 - [ApiTimestamp](docs/ApiTimestamp.md)
 - [ApiTypeMeta](docs/ApiTypeMeta.md)
 - [ApiWatchControl](docs/ApiWatchControl.md)
 - [ApiWatchEvent](docs/ApiWatchEvent.md)
 - [ApiWatchEventList](docs/ApiWatchEventList.md)
 - [ClusterIPConfig](docs/ClusterIPConfig.md)
 - [GoogleprotobufAny](docs/GoogleprotobufAny.md)
 - [SecurityDSCStatus](docs/SecurityDSCStatus.md)
 - [SecurityPropagationStatus](docs/SecurityPropagationStatus.md)
 - [WorkloadAutoMsgEndpointWatchHelper](docs/WorkloadAutoMsgEndpointWatchHelper.md)
 - [WorkloadAutoMsgEndpointWatchHelperWatchEvent](docs/WorkloadAutoMsgEndpointWatchHelperWatchEvent.md)
 - [WorkloadAutoMsgWorkloadWatchHelper](docs/WorkloadAutoMsgWorkloadWatchHelper.md)
 - [WorkloadAutoMsgWorkloadWatchHelperWatchEvent](docs/WorkloadAutoMsgWorkloadWatchHelperWatchEvent.md)
 - [WorkloadEndpoint](docs/WorkloadEndpoint.md)
 - [WorkloadEndpointList](docs/WorkloadEndpointList.md)
 - [WorkloadEndpointMigrationStatus](docs/WorkloadEndpointMigrationStatus.md)
 - [WorkloadEndpointSpec](docs/WorkloadEndpointSpec.md)
 - [WorkloadEndpointStatus](docs/WorkloadEndpointStatus.md)
 - [WorkloadMigrationSource](docs/WorkloadMigrationSource.md)
 - [WorkloadWorkload](docs/WorkloadWorkload.md)
 - [WorkloadWorkloadIntfSpec](docs/WorkloadWorkloadIntfSpec.md)
 - [WorkloadWorkloadIntfStatus](docs/WorkloadWorkloadIntfStatus.md)
 - [WorkloadWorkloadList](docs/WorkloadWorkloadList.md)
 - [WorkloadWorkloadMigrationStatus](docs/WorkloadWorkloadMigrationStatus.md)
 - [WorkloadWorkloadSpec](docs/WorkloadWorkloadSpec.md)
 - [WorkloadWorkloadStatus](docs/WorkloadWorkloadStatus.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


