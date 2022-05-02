# WorkflowPowerShellApiAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.PowerShellApi"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.PowerShellApi"]
**Depth** | Pointer to **int64** | The response of a PowerShell script is an object, since PowerShell is an Object based language. Each object can contain multiple objects as properties, each of which in turn can contain multiple objects and so on and so forth. The depth field specifies how many levels of contained objects are included in the JSON representation. | [optional] 
**OperationTimeout** | Pointer to **string** | The timeout in seconds for the execution of the script against the given endpoint. | [optional] 
**PowerShellResponseSpec** | Pointer to **interface{}** | The grammar specification to parse the response and extract the required values. | [optional] 

## Methods

### NewWorkflowPowerShellApiAllOf

`func NewWorkflowPowerShellApiAllOf(classId string, objectType string, ) *WorkflowPowerShellApiAllOf`

NewWorkflowPowerShellApiAllOf instantiates a new WorkflowPowerShellApiAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPowerShellApiAllOfWithDefaults

`func NewWorkflowPowerShellApiAllOfWithDefaults() *WorkflowPowerShellApiAllOf`

NewWorkflowPowerShellApiAllOfWithDefaults instantiates a new WorkflowPowerShellApiAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowPowerShellApiAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowPowerShellApiAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowPowerShellApiAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowPowerShellApiAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowPowerShellApiAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowPowerShellApiAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDepth

`func (o *WorkflowPowerShellApiAllOf) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WorkflowPowerShellApiAllOf) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WorkflowPowerShellApiAllOf) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *WorkflowPowerShellApiAllOf) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetOperationTimeout

`func (o *WorkflowPowerShellApiAllOf) GetOperationTimeout() string`

GetOperationTimeout returns the OperationTimeout field if non-nil, zero value otherwise.

### GetOperationTimeoutOk

`func (o *WorkflowPowerShellApiAllOf) GetOperationTimeoutOk() (*string, bool)`

GetOperationTimeoutOk returns a tuple with the OperationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTimeout

`func (o *WorkflowPowerShellApiAllOf) SetOperationTimeout(v string)`

SetOperationTimeout sets OperationTimeout field to given value.

### HasOperationTimeout

`func (o *WorkflowPowerShellApiAllOf) HasOperationTimeout() bool`

HasOperationTimeout returns a boolean if a field has been set.

### GetPowerShellResponseSpec

`func (o *WorkflowPowerShellApiAllOf) GetPowerShellResponseSpec() interface{}`

GetPowerShellResponseSpec returns the PowerShellResponseSpec field if non-nil, zero value otherwise.

### GetPowerShellResponseSpecOk

`func (o *WorkflowPowerShellApiAllOf) GetPowerShellResponseSpecOk() (*interface{}, bool)`

GetPowerShellResponseSpecOk returns a tuple with the PowerShellResponseSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShellResponseSpec

`func (o *WorkflowPowerShellApiAllOf) SetPowerShellResponseSpec(v interface{})`

SetPowerShellResponseSpec sets PowerShellResponseSpec field to given value.

### HasPowerShellResponseSpec

`func (o *WorkflowPowerShellApiAllOf) HasPowerShellResponseSpec() bool`

HasPowerShellResponseSpec returns a boolean if a field has been set.

### SetPowerShellResponseSpecNil

`func (o *WorkflowPowerShellApiAllOf) SetPowerShellResponseSpecNil(b bool)`

 SetPowerShellResponseSpecNil sets the value for PowerShellResponseSpec to be an explicit nil

### UnsetPowerShellResponseSpec
`func (o *WorkflowPowerShellApiAllOf) UnsetPowerShellResponseSpec()`

UnsetPowerShellResponseSpec ensures that no value is present for PowerShellResponseSpec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


