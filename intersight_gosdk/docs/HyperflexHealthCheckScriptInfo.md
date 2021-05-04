# HyperflexHealthCheckScriptInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HealthCheckScriptInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HealthCheckScriptInfo"]
**AggregateScriptName** | Pointer to **string** | Health check aggregate script that runs in the HyperFlex Leader Node. | It aggregates the output of all HyperFlex nodes and provides the health check result. | [optional] [readonly] 
**HyperflexVersion** | Pointer to **string** | HyperFlex Data Platform version running on the target device. | [optional] [readonly] 
**ScriptExecuteLocation** | Pointer to **string** | Location of the health check script&#39;s execution on the HyperFlex device. | [optional] [readonly] 
**ScriptInput** | Pointer to **string** | Input for the health check script execution. | [optional] [readonly] 
**ScriptName** | Pointer to **string** | Name of the health check script to be executed. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the health check script associated with the health check definition. | [optional] [readonly] 

## Methods

### NewHyperflexHealthCheckScriptInfo

`func NewHyperflexHealthCheckScriptInfo(classId string, objectType string, ) *HyperflexHealthCheckScriptInfo`

NewHyperflexHealthCheckScriptInfo instantiates a new HyperflexHealthCheckScriptInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckScriptInfoWithDefaults

`func NewHyperflexHealthCheckScriptInfoWithDefaults() *HyperflexHealthCheckScriptInfo`

NewHyperflexHealthCheckScriptInfoWithDefaults instantiates a new HyperflexHealthCheckScriptInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckScriptInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckScriptInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckScriptInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckScriptInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckScriptInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckScriptInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregateScriptName

`func (o *HyperflexHealthCheckScriptInfo) GetAggregateScriptName() string`

GetAggregateScriptName returns the AggregateScriptName field if non-nil, zero value otherwise.

### GetAggregateScriptNameOk

`func (o *HyperflexHealthCheckScriptInfo) GetAggregateScriptNameOk() (*string, bool)`

GetAggregateScriptNameOk returns a tuple with the AggregateScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateScriptName

`func (o *HyperflexHealthCheckScriptInfo) SetAggregateScriptName(v string)`

SetAggregateScriptName sets AggregateScriptName field to given value.

### HasAggregateScriptName

`func (o *HyperflexHealthCheckScriptInfo) HasAggregateScriptName() bool`

HasAggregateScriptName returns a boolean if a field has been set.

### GetHyperflexVersion

`func (o *HyperflexHealthCheckScriptInfo) GetHyperflexVersion() string`

GetHyperflexVersion returns the HyperflexVersion field if non-nil, zero value otherwise.

### GetHyperflexVersionOk

`func (o *HyperflexHealthCheckScriptInfo) GetHyperflexVersionOk() (*string, bool)`

GetHyperflexVersionOk returns a tuple with the HyperflexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexVersion

`func (o *HyperflexHealthCheckScriptInfo) SetHyperflexVersion(v string)`

SetHyperflexVersion sets HyperflexVersion field to given value.

### HasHyperflexVersion

`func (o *HyperflexHealthCheckScriptInfo) HasHyperflexVersion() bool`

HasHyperflexVersion returns a boolean if a field has been set.

### GetScriptExecuteLocation

`func (o *HyperflexHealthCheckScriptInfo) GetScriptExecuteLocation() string`

GetScriptExecuteLocation returns the ScriptExecuteLocation field if non-nil, zero value otherwise.

### GetScriptExecuteLocationOk

`func (o *HyperflexHealthCheckScriptInfo) GetScriptExecuteLocationOk() (*string, bool)`

GetScriptExecuteLocationOk returns a tuple with the ScriptExecuteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptExecuteLocation

`func (o *HyperflexHealthCheckScriptInfo) SetScriptExecuteLocation(v string)`

SetScriptExecuteLocation sets ScriptExecuteLocation field to given value.

### HasScriptExecuteLocation

`func (o *HyperflexHealthCheckScriptInfo) HasScriptExecuteLocation() bool`

HasScriptExecuteLocation returns a boolean if a field has been set.

### GetScriptInput

`func (o *HyperflexHealthCheckScriptInfo) GetScriptInput() string`

GetScriptInput returns the ScriptInput field if non-nil, zero value otherwise.

### GetScriptInputOk

`func (o *HyperflexHealthCheckScriptInfo) GetScriptInputOk() (*string, bool)`

GetScriptInputOk returns a tuple with the ScriptInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptInput

`func (o *HyperflexHealthCheckScriptInfo) SetScriptInput(v string)`

SetScriptInput sets ScriptInput field to given value.

### HasScriptInput

`func (o *HyperflexHealthCheckScriptInfo) HasScriptInput() bool`

HasScriptInput returns a boolean if a field has been set.

### GetScriptName

`func (o *HyperflexHealthCheckScriptInfo) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *HyperflexHealthCheckScriptInfo) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *HyperflexHealthCheckScriptInfo) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *HyperflexHealthCheckScriptInfo) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexHealthCheckScriptInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexHealthCheckScriptInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexHealthCheckScriptInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexHealthCheckScriptInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


