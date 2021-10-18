# TerraformExecutorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "terraform.Executor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "terraform.Executor"]
**CloudResource** | Pointer to [**[]TerraformCloudResource**](TerraformCloudResource.md) |  | [optional] 
**Operation** | Pointer to **string** | Enum indicates what operation is being done. * &#x60;Create&#x60; - Creating a Terraform resource. * &#x60;Update&#x60; - Updating a Terraform resource. * &#x60;Delete&#x60; - Deleting a Terraform resource. | [optional] [default to "Create"]
**Output** | Pointer to **interface{}** | Terraform output of the entire execution. | [optional] 
**PlatformType** | Pointer to **string** | The Platform type used in conjunction with &#39;sourceFolderPath&#39; and &#39;sourceFolderName&#39; determines unique path for a Terraform workflow. | [optional] 
**RunState** | Pointer to [**[]TerraformRunstate**](TerraformRunstate.md) |  | [optional] 
**SourceFolderName** | Pointer to **string** | Folder Name where Terraform workflows are stored. | [optional] 
**SourceFolderPath** | Pointer to **string** | Relative folder Path where &#39;sourceFolderName&#39; is located. | [optional] 
**SourceLocation** | Pointer to **string** | Flag indicates whether workflow is internal/external. | [optional] 
**Status** | Pointer to **string** | Status of the terraform execution. | [optional] 
**Stderr** | Pointer to **interface{}** | Stderr of the terraform execution will be captured here. | [optional] 
**Stdout** | Pointer to **interface{}** | Stdout of the terraform execution will be captured here. | [optional] 
**TaskId** | Pointer to **string** | TaskId of a pontem workflow is same as the MO. | [optional] 
**Variables** | Pointer to **interface{}** | Variables needed by the terraform configuration as a JSON object. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewTerraformExecutorAllOf

`func NewTerraformExecutorAllOf(classId string, objectType string, ) *TerraformExecutorAllOf`

NewTerraformExecutorAllOf instantiates a new TerraformExecutorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformExecutorAllOfWithDefaults

`func NewTerraformExecutorAllOfWithDefaults() *TerraformExecutorAllOf`

NewTerraformExecutorAllOfWithDefaults instantiates a new TerraformExecutorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TerraformExecutorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TerraformExecutorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TerraformExecutorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TerraformExecutorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TerraformExecutorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TerraformExecutorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCloudResource

`func (o *TerraformExecutorAllOf) GetCloudResource() []TerraformCloudResource`

GetCloudResource returns the CloudResource field if non-nil, zero value otherwise.

### GetCloudResourceOk

`func (o *TerraformExecutorAllOf) GetCloudResourceOk() (*[]TerraformCloudResource, bool)`

GetCloudResourceOk returns a tuple with the CloudResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResource

`func (o *TerraformExecutorAllOf) SetCloudResource(v []TerraformCloudResource)`

SetCloudResource sets CloudResource field to given value.

### HasCloudResource

`func (o *TerraformExecutorAllOf) HasCloudResource() bool`

HasCloudResource returns a boolean if a field has been set.

### SetCloudResourceNil

`func (o *TerraformExecutorAllOf) SetCloudResourceNil(b bool)`

 SetCloudResourceNil sets the value for CloudResource to be an explicit nil

### UnsetCloudResource
`func (o *TerraformExecutorAllOf) UnsetCloudResource()`

UnsetCloudResource ensures that no value is present for CloudResource, not even an explicit nil
### GetOperation

`func (o *TerraformExecutorAllOf) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TerraformExecutorAllOf) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TerraformExecutorAllOf) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *TerraformExecutorAllOf) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOutput

`func (o *TerraformExecutorAllOf) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TerraformExecutorAllOf) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TerraformExecutorAllOf) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TerraformExecutorAllOf) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *TerraformExecutorAllOf) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *TerraformExecutorAllOf) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPlatformType

`func (o *TerraformExecutorAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *TerraformExecutorAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *TerraformExecutorAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *TerraformExecutorAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRunState

`func (o *TerraformExecutorAllOf) GetRunState() []TerraformRunstate`

GetRunState returns the RunState field if non-nil, zero value otherwise.

### GetRunStateOk

`func (o *TerraformExecutorAllOf) GetRunStateOk() (*[]TerraformRunstate, bool)`

GetRunStateOk returns a tuple with the RunState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunState

`func (o *TerraformExecutorAllOf) SetRunState(v []TerraformRunstate)`

SetRunState sets RunState field to given value.

### HasRunState

`func (o *TerraformExecutorAllOf) HasRunState() bool`

HasRunState returns a boolean if a field has been set.

### SetRunStateNil

`func (o *TerraformExecutorAllOf) SetRunStateNil(b bool)`

 SetRunStateNil sets the value for RunState to be an explicit nil

### UnsetRunState
`func (o *TerraformExecutorAllOf) UnsetRunState()`

UnsetRunState ensures that no value is present for RunState, not even an explicit nil
### GetSourceFolderName

`func (o *TerraformExecutorAllOf) GetSourceFolderName() string`

GetSourceFolderName returns the SourceFolderName field if non-nil, zero value otherwise.

### GetSourceFolderNameOk

`func (o *TerraformExecutorAllOf) GetSourceFolderNameOk() (*string, bool)`

GetSourceFolderNameOk returns a tuple with the SourceFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFolderName

`func (o *TerraformExecutorAllOf) SetSourceFolderName(v string)`

SetSourceFolderName sets SourceFolderName field to given value.

### HasSourceFolderName

`func (o *TerraformExecutorAllOf) HasSourceFolderName() bool`

HasSourceFolderName returns a boolean if a field has been set.

### GetSourceFolderPath

`func (o *TerraformExecutorAllOf) GetSourceFolderPath() string`

GetSourceFolderPath returns the SourceFolderPath field if non-nil, zero value otherwise.

### GetSourceFolderPathOk

`func (o *TerraformExecutorAllOf) GetSourceFolderPathOk() (*string, bool)`

GetSourceFolderPathOk returns a tuple with the SourceFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFolderPath

`func (o *TerraformExecutorAllOf) SetSourceFolderPath(v string)`

SetSourceFolderPath sets SourceFolderPath field to given value.

### HasSourceFolderPath

`func (o *TerraformExecutorAllOf) HasSourceFolderPath() bool`

HasSourceFolderPath returns a boolean if a field has been set.

### GetSourceLocation

`func (o *TerraformExecutorAllOf) GetSourceLocation() string`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *TerraformExecutorAllOf) GetSourceLocationOk() (*string, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *TerraformExecutorAllOf) SetSourceLocation(v string)`

SetSourceLocation sets SourceLocation field to given value.

### HasSourceLocation

`func (o *TerraformExecutorAllOf) HasSourceLocation() bool`

HasSourceLocation returns a boolean if a field has been set.

### GetStatus

`func (o *TerraformExecutorAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TerraformExecutorAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TerraformExecutorAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TerraformExecutorAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStderr

`func (o *TerraformExecutorAllOf) GetStderr() interface{}`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *TerraformExecutorAllOf) GetStderrOk() (*interface{}, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *TerraformExecutorAllOf) SetStderr(v interface{})`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *TerraformExecutorAllOf) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### SetStderrNil

`func (o *TerraformExecutorAllOf) SetStderrNil(b bool)`

 SetStderrNil sets the value for Stderr to be an explicit nil

### UnsetStderr
`func (o *TerraformExecutorAllOf) UnsetStderr()`

UnsetStderr ensures that no value is present for Stderr, not even an explicit nil
### GetStdout

`func (o *TerraformExecutorAllOf) GetStdout() interface{}`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *TerraformExecutorAllOf) GetStdoutOk() (*interface{}, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *TerraformExecutorAllOf) SetStdout(v interface{})`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *TerraformExecutorAllOf) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *TerraformExecutorAllOf) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *TerraformExecutorAllOf) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil
### GetTaskId

`func (o *TerraformExecutorAllOf) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TerraformExecutorAllOf) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TerraformExecutorAllOf) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TerraformExecutorAllOf) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetVariables

`func (o *TerraformExecutorAllOf) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TerraformExecutorAllOf) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TerraformExecutorAllOf) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *TerraformExecutorAllOf) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *TerraformExecutorAllOf) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *TerraformExecutorAllOf) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetAccount

`func (o *TerraformExecutorAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TerraformExecutorAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TerraformExecutorAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TerraformExecutorAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *TerraformExecutorAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *TerraformExecutorAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *TerraformExecutorAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *TerraformExecutorAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *TerraformExecutorAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *TerraformExecutorAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *TerraformExecutorAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *TerraformExecutorAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


