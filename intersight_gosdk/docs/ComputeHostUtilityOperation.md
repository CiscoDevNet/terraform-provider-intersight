# ComputeHostUtilityOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.HostUtilityOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.HostUtilityOperation"]
**HostOpConfig** | Pointer to [**NullableComputeHostUtilityOperationConfguration**](ComputeHostUtilityOperationConfguration.md) |  | [optional] 
**HostUtilityOperationMode** | Pointer to **string** | Host utility operation need to be performed in the endpoint. * &#x60;None&#x60; - Host utility mode of the operation is set to none by default. * &#x60;SecureErase&#x60; - EU LOT-9 secure data cleanup on the server components. * &#x60;SecureEraseWithDecommission&#x60; - EU LOT-9 secure data cleanup on the server components and do decommission. * &#x60;Scrub&#x60; - Quick cleanup on storage and BIOS. | [optional] [default to "None"]
**HostUtilityOperationStatus** | Pointer to **string** | Task status of the host utility operation. * &#x60;Initiated&#x60; - This status indicates that host utility operation request is initiated. * &#x60;InProgress&#x60; - The operation status indicates that host utility operation is in-progress after the basic validations. * &#x60;CompletedOk&#x60; - The operation status indicates that host utility operation is completed successfully with no error or warning. * &#x60;CompletedError&#x60; - The operation status indicates that host utility operation is completed with error. * &#x60;CompletedWarning&#x60; - The operation status indicates that host utility operation is completed with warning. * &#x60;Aborted&#x60; - The operation status indicates that host utility operation is terminated or aborted. * &#x60;Invalidated&#x60; - The operation status indicates that host utility operation is invalid due to validation failure. | [optional] [readonly] [default to "Initiated"]
**DownloadStatus** | Pointer to [**NullableComputeDownloadStatusRelationship**](ComputeDownloadStatusRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewComputeHostUtilityOperation

`func NewComputeHostUtilityOperation(classId string, objectType string, ) *ComputeHostUtilityOperation`

NewComputeHostUtilityOperation instantiates a new ComputeHostUtilityOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeHostUtilityOperationWithDefaults

`func NewComputeHostUtilityOperationWithDefaults() *ComputeHostUtilityOperation`

NewComputeHostUtilityOperationWithDefaults instantiates a new ComputeHostUtilityOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeHostUtilityOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeHostUtilityOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeHostUtilityOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeHostUtilityOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeHostUtilityOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeHostUtilityOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostOpConfig

`func (o *ComputeHostUtilityOperation) GetHostOpConfig() ComputeHostUtilityOperationConfguration`

GetHostOpConfig returns the HostOpConfig field if non-nil, zero value otherwise.

### GetHostOpConfigOk

`func (o *ComputeHostUtilityOperation) GetHostOpConfigOk() (*ComputeHostUtilityOperationConfguration, bool)`

GetHostOpConfigOk returns a tuple with the HostOpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOpConfig

`func (o *ComputeHostUtilityOperation) SetHostOpConfig(v ComputeHostUtilityOperationConfguration)`

SetHostOpConfig sets HostOpConfig field to given value.

### HasHostOpConfig

`func (o *ComputeHostUtilityOperation) HasHostOpConfig() bool`

HasHostOpConfig returns a boolean if a field has been set.

### SetHostOpConfigNil

`func (o *ComputeHostUtilityOperation) SetHostOpConfigNil(b bool)`

 SetHostOpConfigNil sets the value for HostOpConfig to be an explicit nil

### UnsetHostOpConfig
`func (o *ComputeHostUtilityOperation) UnsetHostOpConfig()`

UnsetHostOpConfig ensures that no value is present for HostOpConfig, not even an explicit nil
### GetHostUtilityOperationMode

`func (o *ComputeHostUtilityOperation) GetHostUtilityOperationMode() string`

GetHostUtilityOperationMode returns the HostUtilityOperationMode field if non-nil, zero value otherwise.

### GetHostUtilityOperationModeOk

`func (o *ComputeHostUtilityOperation) GetHostUtilityOperationModeOk() (*string, bool)`

GetHostUtilityOperationModeOk returns a tuple with the HostUtilityOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUtilityOperationMode

`func (o *ComputeHostUtilityOperation) SetHostUtilityOperationMode(v string)`

SetHostUtilityOperationMode sets HostUtilityOperationMode field to given value.

### HasHostUtilityOperationMode

`func (o *ComputeHostUtilityOperation) HasHostUtilityOperationMode() bool`

HasHostUtilityOperationMode returns a boolean if a field has been set.

### GetHostUtilityOperationStatus

`func (o *ComputeHostUtilityOperation) GetHostUtilityOperationStatus() string`

GetHostUtilityOperationStatus returns the HostUtilityOperationStatus field if non-nil, zero value otherwise.

### GetHostUtilityOperationStatusOk

`func (o *ComputeHostUtilityOperation) GetHostUtilityOperationStatusOk() (*string, bool)`

GetHostUtilityOperationStatusOk returns a tuple with the HostUtilityOperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUtilityOperationStatus

`func (o *ComputeHostUtilityOperation) SetHostUtilityOperationStatus(v string)`

SetHostUtilityOperationStatus sets HostUtilityOperationStatus field to given value.

### HasHostUtilityOperationStatus

`func (o *ComputeHostUtilityOperation) HasHostUtilityOperationStatus() bool`

HasHostUtilityOperationStatus returns a boolean if a field has been set.

### GetDownloadStatus

`func (o *ComputeHostUtilityOperation) GetDownloadStatus() ComputeDownloadStatusRelationship`

GetDownloadStatus returns the DownloadStatus field if non-nil, zero value otherwise.

### GetDownloadStatusOk

`func (o *ComputeHostUtilityOperation) GetDownloadStatusOk() (*ComputeDownloadStatusRelationship, bool)`

GetDownloadStatusOk returns a tuple with the DownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStatus

`func (o *ComputeHostUtilityOperation) SetDownloadStatus(v ComputeDownloadStatusRelationship)`

SetDownloadStatus sets DownloadStatus field to given value.

### HasDownloadStatus

`func (o *ComputeHostUtilityOperation) HasDownloadStatus() bool`

HasDownloadStatus returns a boolean if a field has been set.

### SetDownloadStatusNil

`func (o *ComputeHostUtilityOperation) SetDownloadStatusNil(b bool)`

 SetDownloadStatusNil sets the value for DownloadStatus to be an explicit nil

### UnsetDownloadStatus
`func (o *ComputeHostUtilityOperation) UnsetDownloadStatus()`

UnsetDownloadStatus ensures that no value is present for DownloadStatus, not even an explicit nil
### GetServer

`func (o *ComputeHostUtilityOperation) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ComputeHostUtilityOperation) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ComputeHostUtilityOperation) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ComputeHostUtilityOperation) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *ComputeHostUtilityOperation) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *ComputeHostUtilityOperation) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetWorkflow

`func (o *ComputeHostUtilityOperation) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ComputeHostUtilityOperation) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ComputeHostUtilityOperation) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ComputeHostUtilityOperation) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *ComputeHostUtilityOperation) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *ComputeHostUtilityOperation) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


