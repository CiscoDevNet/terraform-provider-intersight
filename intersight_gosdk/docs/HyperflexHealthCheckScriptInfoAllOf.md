# HyperflexHealthCheckScriptInfoAllOf

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

### NewHyperflexHealthCheckScriptInfoAllOf

`func NewHyperflexHealthCheckScriptInfoAllOf(classId string, objectType string, ) *HyperflexHealthCheckScriptInfoAllOf`

NewHyperflexHealthCheckScriptInfoAllOf instantiates a new HyperflexHealthCheckScriptInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckScriptInfoAllOfWithDefaults

`func NewHyperflexHealthCheckScriptInfoAllOfWithDefaults() *HyperflexHealthCheckScriptInfoAllOf`

NewHyperflexHealthCheckScriptInfoAllOfWithDefaults instantiates a new HyperflexHealthCheckScriptInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregateScriptName

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetAggregateScriptName() string`

GetAggregateScriptName returns the AggregateScriptName field if non-nil, zero value otherwise.

### GetAggregateScriptNameOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetAggregateScriptNameOk() (*string, bool)`

GetAggregateScriptNameOk returns a tuple with the AggregateScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateScriptName

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetAggregateScriptName(v string)`

SetAggregateScriptName sets AggregateScriptName field to given value.

### HasAggregateScriptName

`func (o *HyperflexHealthCheckScriptInfoAllOf) HasAggregateScriptName() bool`

HasAggregateScriptName returns a boolean if a field has been set.

### GetHyperflexVersion

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetHyperflexVersion() string`

GetHyperflexVersion returns the HyperflexVersion field if non-nil, zero value otherwise.

### GetHyperflexVersionOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetHyperflexVersionOk() (*string, bool)`

GetHyperflexVersionOk returns a tuple with the HyperflexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexVersion

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetHyperflexVersion(v string)`

SetHyperflexVersion sets HyperflexVersion field to given value.

### HasHyperflexVersion

`func (o *HyperflexHealthCheckScriptInfoAllOf) HasHyperflexVersion() bool`

HasHyperflexVersion returns a boolean if a field has been set.

### GetScriptExecuteLocation

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptExecuteLocation() string`

GetScriptExecuteLocation returns the ScriptExecuteLocation field if non-nil, zero value otherwise.

### GetScriptExecuteLocationOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptExecuteLocationOk() (*string, bool)`

GetScriptExecuteLocationOk returns a tuple with the ScriptExecuteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptExecuteLocation

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetScriptExecuteLocation(v string)`

SetScriptExecuteLocation sets ScriptExecuteLocation field to given value.

### HasScriptExecuteLocation

`func (o *HyperflexHealthCheckScriptInfoAllOf) HasScriptExecuteLocation() bool`

HasScriptExecuteLocation returns a boolean if a field has been set.

### GetScriptInput

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptInput() string`

GetScriptInput returns the ScriptInput field if non-nil, zero value otherwise.

### GetScriptInputOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptInputOk() (*string, bool)`

GetScriptInputOk returns a tuple with the ScriptInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptInput

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetScriptInput(v string)`

SetScriptInput sets ScriptInput field to given value.

### HasScriptInput

`func (o *HyperflexHealthCheckScriptInfoAllOf) HasScriptInput() bool`

HasScriptInput returns a boolean if a field has been set.

### GetScriptName

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *HyperflexHealthCheckScriptInfoAllOf) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexHealthCheckScriptInfoAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexHealthCheckScriptInfoAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexHealthCheckScriptInfoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


