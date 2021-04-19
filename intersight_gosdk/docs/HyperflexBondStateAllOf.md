# HyperflexBondStateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.BondState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.BondState"]
**ActiveSlave** | Pointer to **string** | The current active slave. For active-active mode, this field is empty. | [optional] [readonly] 
**Mode** | Pointer to **string** | Bond mode, such as \&quot;active-backup\&quot;, \&quot;balance-slb\&quot;, \&quot;balance-tcp\&quot;. | [optional] [readonly] 
**Slaves** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHyperflexBondStateAllOf

`func NewHyperflexBondStateAllOf(classId string, objectType string, ) *HyperflexBondStateAllOf`

NewHyperflexBondStateAllOf instantiates a new HyperflexBondStateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexBondStateAllOfWithDefaults

`func NewHyperflexBondStateAllOfWithDefaults() *HyperflexBondStateAllOf`

NewHyperflexBondStateAllOfWithDefaults instantiates a new HyperflexBondStateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexBondStateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexBondStateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexBondStateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexBondStateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexBondStateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexBondStateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveSlave

`func (o *HyperflexBondStateAllOf) GetActiveSlave() string`

GetActiveSlave returns the ActiveSlave field if non-nil, zero value otherwise.

### GetActiveSlaveOk

`func (o *HyperflexBondStateAllOf) GetActiveSlaveOk() (*string, bool)`

GetActiveSlaveOk returns a tuple with the ActiveSlave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSlave

`func (o *HyperflexBondStateAllOf) SetActiveSlave(v string)`

SetActiveSlave sets ActiveSlave field to given value.

### HasActiveSlave

`func (o *HyperflexBondStateAllOf) HasActiveSlave() bool`

HasActiveSlave returns a boolean if a field has been set.

### GetMode

`func (o *HyperflexBondStateAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HyperflexBondStateAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HyperflexBondStateAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HyperflexBondStateAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSlaves

`func (o *HyperflexBondStateAllOf) GetSlaves() []string`

GetSlaves returns the Slaves field if non-nil, zero value otherwise.

### GetSlavesOk

`func (o *HyperflexBondStateAllOf) GetSlavesOk() (*[]string, bool)`

GetSlavesOk returns a tuple with the Slaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaves

`func (o *HyperflexBondStateAllOf) SetSlaves(v []string)`

SetSlaves sets Slaves field to given value.

### HasSlaves

`func (o *HyperflexBondStateAllOf) HasSlaves() bool`

HasSlaves returns a boolean if a field has been set.

### SetSlavesNil

`func (o *HyperflexBondStateAllOf) SetSlavesNil(b bool)`

 SetSlavesNil sets the value for Slaves to be an explicit nil

### UnsetSlaves
`func (o *HyperflexBondStateAllOf) UnsetSlaves()`

UnsetSlaves ensures that no value is present for Slaves, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


