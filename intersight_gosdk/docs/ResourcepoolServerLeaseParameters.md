# ResourcepoolServerLeaseParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.ServerLeaseParameters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.ServerLeaseParameters"]
**LifeCycleStates** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResourcepoolServerLeaseParameters

`func NewResourcepoolServerLeaseParameters(classId string, objectType string, ) *ResourcepoolServerLeaseParameters`

NewResourcepoolServerLeaseParameters instantiates a new ResourcepoolServerLeaseParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolServerLeaseParametersWithDefaults

`func NewResourcepoolServerLeaseParametersWithDefaults() *ResourcepoolServerLeaseParameters`

NewResourcepoolServerLeaseParametersWithDefaults instantiates a new ResourcepoolServerLeaseParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolServerLeaseParameters) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolServerLeaseParameters) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolServerLeaseParameters) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolServerLeaseParameters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolServerLeaseParameters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolServerLeaseParameters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLifeCycleStates

`func (o *ResourcepoolServerLeaseParameters) GetLifeCycleStates() []string`

GetLifeCycleStates returns the LifeCycleStates field if non-nil, zero value otherwise.

### GetLifeCycleStatesOk

`func (o *ResourcepoolServerLeaseParameters) GetLifeCycleStatesOk() (*[]string, bool)`

GetLifeCycleStatesOk returns a tuple with the LifeCycleStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStates

`func (o *ResourcepoolServerLeaseParameters) SetLifeCycleStates(v []string)`

SetLifeCycleStates sets LifeCycleStates field to given value.

### HasLifeCycleStates

`func (o *ResourcepoolServerLeaseParameters) HasLifeCycleStates() bool`

HasLifeCycleStates returns a boolean if a field has been set.

### SetLifeCycleStatesNil

`func (o *ResourcepoolServerLeaseParameters) SetLifeCycleStatesNil(b bool)`

 SetLifeCycleStatesNil sets the value for LifeCycleStates to be an explicit nil

### UnsetLifeCycleStates
`func (o *ResourcepoolServerLeaseParameters) UnsetLifeCycleStates()`

UnsetLifeCycleStates ensures that no value is present for LifeCycleStates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


