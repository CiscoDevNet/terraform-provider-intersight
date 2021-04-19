# HyperflexBondState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.BondState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.BondState"]
**ActiveSlave** | Pointer to **string** | The current active slave. For active-active mode, this field is empty. | [optional] [readonly] 
**Mode** | Pointer to **string** | Bond mode, such as \&quot;active-backup\&quot;, \&quot;balance-slb\&quot;, \&quot;balance-tcp\&quot;. | [optional] [readonly] 
**Slaves** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHyperflexBondState

`func NewHyperflexBondState(classId string, objectType string, ) *HyperflexBondState`

NewHyperflexBondState instantiates a new HyperflexBondState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexBondStateWithDefaults

`func NewHyperflexBondStateWithDefaults() *HyperflexBondState`

NewHyperflexBondStateWithDefaults instantiates a new HyperflexBondState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexBondState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexBondState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexBondState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexBondState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexBondState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexBondState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveSlave

`func (o *HyperflexBondState) GetActiveSlave() string`

GetActiveSlave returns the ActiveSlave field if non-nil, zero value otherwise.

### GetActiveSlaveOk

`func (o *HyperflexBondState) GetActiveSlaveOk() (*string, bool)`

GetActiveSlaveOk returns a tuple with the ActiveSlave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSlave

`func (o *HyperflexBondState) SetActiveSlave(v string)`

SetActiveSlave sets ActiveSlave field to given value.

### HasActiveSlave

`func (o *HyperflexBondState) HasActiveSlave() bool`

HasActiveSlave returns a boolean if a field has been set.

### GetMode

`func (o *HyperflexBondState) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HyperflexBondState) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HyperflexBondState) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HyperflexBondState) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSlaves

`func (o *HyperflexBondState) GetSlaves() []string`

GetSlaves returns the Slaves field if non-nil, zero value otherwise.

### GetSlavesOk

`func (o *HyperflexBondState) GetSlavesOk() (*[]string, bool)`

GetSlavesOk returns a tuple with the Slaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaves

`func (o *HyperflexBondState) SetSlaves(v []string)`

SetSlaves sets Slaves field to given value.

### HasSlaves

`func (o *HyperflexBondState) HasSlaves() bool`

HasSlaves returns a boolean if a field has been set.

### SetSlavesNil

`func (o *HyperflexBondState) SetSlavesNil(b bool)`

 SetSlavesNil sets the value for Slaves to be an explicit nil

### UnsetSlaves
`func (o *HyperflexBondState) UnsetSlaves()`

UnsetSlaves ensures that no value is present for Slaves, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


