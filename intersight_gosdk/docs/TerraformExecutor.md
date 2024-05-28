# TerraformExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "terraform.Executor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "terraform.Executor"]
**CloudResource** | Pointer to [**[]TerraformCloudResource**](TerraformCloudResource.md) |  | [optional] 
**Command** | Pointer to **string** | Command to be executed during update operation. | [optional] 
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
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewTerraformExecutor

`func NewTerraformExecutor(classId string, objectType string, ) *TerraformExecutor`

NewTerraformExecutor instantiates a new TerraformExecutor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformExecutorWithDefaults

`func NewTerraformExecutorWithDefaults() *TerraformExecutor`

NewTerraformExecutorWithDefaults instantiates a new TerraformExecutor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TerraformExecutor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TerraformExecutor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TerraformExecutor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TerraformExecutor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TerraformExecutor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TerraformExecutor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCloudResource

`func (o *TerraformExecutor) GetCloudResource() []TerraformCloudResource`

GetCloudResource returns the CloudResource field if non-nil, zero value otherwise.

### GetCloudResourceOk

`func (o *TerraformExecutor) GetCloudResourceOk() (*[]TerraformCloudResource, bool)`

GetCloudResourceOk returns a tuple with the CloudResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResource

`func (o *TerraformExecutor) SetCloudResource(v []TerraformCloudResource)`

SetCloudResource sets CloudResource field to given value.

### HasCloudResource

`func (o *TerraformExecutor) HasCloudResource() bool`

HasCloudResource returns a boolean if a field has been set.

### SetCloudResourceNil

`func (o *TerraformExecutor) SetCloudResourceNil(b bool)`

 SetCloudResourceNil sets the value for CloudResource to be an explicit nil

### UnsetCloudResource
`func (o *TerraformExecutor) UnsetCloudResource()`

UnsetCloudResource ensures that no value is present for CloudResource, not even an explicit nil
### GetCommand

`func (o *TerraformExecutor) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *TerraformExecutor) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *TerraformExecutor) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *TerraformExecutor) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetOperation

`func (o *TerraformExecutor) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TerraformExecutor) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TerraformExecutor) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *TerraformExecutor) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOutput

`func (o *TerraformExecutor) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TerraformExecutor) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TerraformExecutor) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TerraformExecutor) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *TerraformExecutor) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *TerraformExecutor) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPlatformType

`func (o *TerraformExecutor) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *TerraformExecutor) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *TerraformExecutor) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *TerraformExecutor) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRunState

`func (o *TerraformExecutor) GetRunState() []TerraformRunstate`

GetRunState returns the RunState field if non-nil, zero value otherwise.

### GetRunStateOk

`func (o *TerraformExecutor) GetRunStateOk() (*[]TerraformRunstate, bool)`

GetRunStateOk returns a tuple with the RunState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunState

`func (o *TerraformExecutor) SetRunState(v []TerraformRunstate)`

SetRunState sets RunState field to given value.

### HasRunState

`func (o *TerraformExecutor) HasRunState() bool`

HasRunState returns a boolean if a field has been set.

### SetRunStateNil

`func (o *TerraformExecutor) SetRunStateNil(b bool)`

 SetRunStateNil sets the value for RunState to be an explicit nil

### UnsetRunState
`func (o *TerraformExecutor) UnsetRunState()`

UnsetRunState ensures that no value is present for RunState, not even an explicit nil
### GetSourceFolderName

`func (o *TerraformExecutor) GetSourceFolderName() string`

GetSourceFolderName returns the SourceFolderName field if non-nil, zero value otherwise.

### GetSourceFolderNameOk

`func (o *TerraformExecutor) GetSourceFolderNameOk() (*string, bool)`

GetSourceFolderNameOk returns a tuple with the SourceFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFolderName

`func (o *TerraformExecutor) SetSourceFolderName(v string)`

SetSourceFolderName sets SourceFolderName field to given value.

### HasSourceFolderName

`func (o *TerraformExecutor) HasSourceFolderName() bool`

HasSourceFolderName returns a boolean if a field has been set.

### GetSourceFolderPath

`func (o *TerraformExecutor) GetSourceFolderPath() string`

GetSourceFolderPath returns the SourceFolderPath field if non-nil, zero value otherwise.

### GetSourceFolderPathOk

`func (o *TerraformExecutor) GetSourceFolderPathOk() (*string, bool)`

GetSourceFolderPathOk returns a tuple with the SourceFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFolderPath

`func (o *TerraformExecutor) SetSourceFolderPath(v string)`

SetSourceFolderPath sets SourceFolderPath field to given value.

### HasSourceFolderPath

`func (o *TerraformExecutor) HasSourceFolderPath() bool`

HasSourceFolderPath returns a boolean if a field has been set.

### GetSourceLocation

`func (o *TerraformExecutor) GetSourceLocation() string`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *TerraformExecutor) GetSourceLocationOk() (*string, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *TerraformExecutor) SetSourceLocation(v string)`

SetSourceLocation sets SourceLocation field to given value.

### HasSourceLocation

`func (o *TerraformExecutor) HasSourceLocation() bool`

HasSourceLocation returns a boolean if a field has been set.

### GetStatus

`func (o *TerraformExecutor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TerraformExecutor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TerraformExecutor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TerraformExecutor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStderr

`func (o *TerraformExecutor) GetStderr() interface{}`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *TerraformExecutor) GetStderrOk() (*interface{}, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *TerraformExecutor) SetStderr(v interface{})`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *TerraformExecutor) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### SetStderrNil

`func (o *TerraformExecutor) SetStderrNil(b bool)`

 SetStderrNil sets the value for Stderr to be an explicit nil

### UnsetStderr
`func (o *TerraformExecutor) UnsetStderr()`

UnsetStderr ensures that no value is present for Stderr, not even an explicit nil
### GetStdout

`func (o *TerraformExecutor) GetStdout() interface{}`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *TerraformExecutor) GetStdoutOk() (*interface{}, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *TerraformExecutor) SetStdout(v interface{})`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *TerraformExecutor) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *TerraformExecutor) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *TerraformExecutor) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil
### GetTaskId

`func (o *TerraformExecutor) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TerraformExecutor) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TerraformExecutor) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TerraformExecutor) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetVariables

`func (o *TerraformExecutor) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TerraformExecutor) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TerraformExecutor) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *TerraformExecutor) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *TerraformExecutor) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *TerraformExecutor) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetAccount

`func (o *TerraformExecutor) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TerraformExecutor) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TerraformExecutor) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TerraformExecutor) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *TerraformExecutor) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TerraformExecutor) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetRegisteredDevice

`func (o *TerraformExecutor) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *TerraformExecutor) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *TerraformExecutor) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *TerraformExecutor) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *TerraformExecutor) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *TerraformExecutor) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetWorkflowInfo

`func (o *TerraformExecutor) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *TerraformExecutor) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *TerraformExecutor) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *TerraformExecutor) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.

### SetWorkflowInfoNil

`func (o *TerraformExecutor) SetWorkflowInfoNil(b bool)`

 SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil

### UnsetWorkflowInfo
`func (o *TerraformExecutor) UnsetWorkflowInfo()`

UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


