# ComputeDownloadStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.DownloadStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.DownloadStatus"]
**DownloadMessage** | Pointer to **string** | The message from the endpoint during the download. | [optional] [readonly] 
**DownloadPercentage** | Pointer to **int64** | The percentage of the image downloaded in the endpoint. | [optional] [readonly] 
**DownloadStage** | Pointer to **string** | The image download stages. Example:downloading, flashing. | [optional] [readonly] 
**SdCardDownloadError** | Pointer to **string** | The error message from the endpoint during the SD card download. | [optional] [readonly] 
**HostOp** | Pointer to [**NullableComputeHostUtilityOperationRelationship**](ComputeHostUtilityOperationRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewComputeDownloadStatus

`func NewComputeDownloadStatus(classId string, objectType string, ) *ComputeDownloadStatus`

NewComputeDownloadStatus instantiates a new ComputeDownloadStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeDownloadStatusWithDefaults

`func NewComputeDownloadStatusWithDefaults() *ComputeDownloadStatus`

NewComputeDownloadStatusWithDefaults instantiates a new ComputeDownloadStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeDownloadStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeDownloadStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeDownloadStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeDownloadStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeDownloadStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeDownloadStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDownloadMessage

`func (o *ComputeDownloadStatus) GetDownloadMessage() string`

GetDownloadMessage returns the DownloadMessage field if non-nil, zero value otherwise.

### GetDownloadMessageOk

`func (o *ComputeDownloadStatus) GetDownloadMessageOk() (*string, bool)`

GetDownloadMessageOk returns a tuple with the DownloadMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMessage

`func (o *ComputeDownloadStatus) SetDownloadMessage(v string)`

SetDownloadMessage sets DownloadMessage field to given value.

### HasDownloadMessage

`func (o *ComputeDownloadStatus) HasDownloadMessage() bool`

HasDownloadMessage returns a boolean if a field has been set.

### GetDownloadPercentage

`func (o *ComputeDownloadStatus) GetDownloadPercentage() int64`

GetDownloadPercentage returns the DownloadPercentage field if non-nil, zero value otherwise.

### GetDownloadPercentageOk

`func (o *ComputeDownloadStatus) GetDownloadPercentageOk() (*int64, bool)`

GetDownloadPercentageOk returns a tuple with the DownloadPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentage

`func (o *ComputeDownloadStatus) SetDownloadPercentage(v int64)`

SetDownloadPercentage sets DownloadPercentage field to given value.

### HasDownloadPercentage

`func (o *ComputeDownloadStatus) HasDownloadPercentage() bool`

HasDownloadPercentage returns a boolean if a field has been set.

### GetDownloadStage

`func (o *ComputeDownloadStatus) GetDownloadStage() string`

GetDownloadStage returns the DownloadStage field if non-nil, zero value otherwise.

### GetDownloadStageOk

`func (o *ComputeDownloadStatus) GetDownloadStageOk() (*string, bool)`

GetDownloadStageOk returns a tuple with the DownloadStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStage

`func (o *ComputeDownloadStatus) SetDownloadStage(v string)`

SetDownloadStage sets DownloadStage field to given value.

### HasDownloadStage

`func (o *ComputeDownloadStatus) HasDownloadStage() bool`

HasDownloadStage returns a boolean if a field has been set.

### GetSdCardDownloadError

`func (o *ComputeDownloadStatus) GetSdCardDownloadError() string`

GetSdCardDownloadError returns the SdCardDownloadError field if non-nil, zero value otherwise.

### GetSdCardDownloadErrorOk

`func (o *ComputeDownloadStatus) GetSdCardDownloadErrorOk() (*string, bool)`

GetSdCardDownloadErrorOk returns a tuple with the SdCardDownloadError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdCardDownloadError

`func (o *ComputeDownloadStatus) SetSdCardDownloadError(v string)`

SetSdCardDownloadError sets SdCardDownloadError field to given value.

### HasSdCardDownloadError

`func (o *ComputeDownloadStatus) HasSdCardDownloadError() bool`

HasSdCardDownloadError returns a boolean if a field has been set.

### GetHostOp

`func (o *ComputeDownloadStatus) GetHostOp() ComputeHostUtilityOperationRelationship`

GetHostOp returns the HostOp field if non-nil, zero value otherwise.

### GetHostOpOk

`func (o *ComputeDownloadStatus) GetHostOpOk() (*ComputeHostUtilityOperationRelationship, bool)`

GetHostOpOk returns a tuple with the HostOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOp

`func (o *ComputeDownloadStatus) SetHostOp(v ComputeHostUtilityOperationRelationship)`

SetHostOp sets HostOp field to given value.

### HasHostOp

`func (o *ComputeDownloadStatus) HasHostOp() bool`

HasHostOp returns a boolean if a field has been set.

### SetHostOpNil

`func (o *ComputeDownloadStatus) SetHostOpNil(b bool)`

 SetHostOpNil sets the value for HostOp to be an explicit nil

### UnsetHostOp
`func (o *ComputeDownloadStatus) UnsetHostOp()`

UnsetHostOp ensures that no value is present for HostOp, not even an explicit nil
### GetWorkflow

`func (o *ComputeDownloadStatus) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ComputeDownloadStatus) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ComputeDownloadStatus) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ComputeDownloadStatus) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *ComputeDownloadStatus) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *ComputeDownloadStatus) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


