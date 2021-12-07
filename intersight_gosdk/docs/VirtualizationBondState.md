# VirtualizationBondState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.BondState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.BondState"]
**ActiveSlave** | Pointer to **string** | The current active slave. For active-active mode, this field is empty. | [optional] [readonly] 
**Mode** | Pointer to **string** | Bond mode, such as \&quot;active-backup\&quot;, \&quot;balance-slb\&quot;, \&quot;balance-tcp\&quot;. | [optional] [readonly] 
**Slaves** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVirtualizationBondState

`func NewVirtualizationBondState(classId string, objectType string, ) *VirtualizationBondState`

NewVirtualizationBondState instantiates a new VirtualizationBondState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBondStateWithDefaults

`func NewVirtualizationBondStateWithDefaults() *VirtualizationBondState`

NewVirtualizationBondStateWithDefaults instantiates a new VirtualizationBondState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBondState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBondState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBondState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBondState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBondState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBondState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveSlave

`func (o *VirtualizationBondState) GetActiveSlave() string`

GetActiveSlave returns the ActiveSlave field if non-nil, zero value otherwise.

### GetActiveSlaveOk

`func (o *VirtualizationBondState) GetActiveSlaveOk() (*string, bool)`

GetActiveSlaveOk returns a tuple with the ActiveSlave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSlave

`func (o *VirtualizationBondState) SetActiveSlave(v string)`

SetActiveSlave sets ActiveSlave field to given value.

### HasActiveSlave

`func (o *VirtualizationBondState) HasActiveSlave() bool`

HasActiveSlave returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationBondState) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationBondState) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationBondState) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationBondState) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSlaves

`func (o *VirtualizationBondState) GetSlaves() []string`

GetSlaves returns the Slaves field if non-nil, zero value otherwise.

### GetSlavesOk

`func (o *VirtualizationBondState) GetSlavesOk() (*[]string, bool)`

GetSlavesOk returns a tuple with the Slaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaves

`func (o *VirtualizationBondState) SetSlaves(v []string)`

SetSlaves sets Slaves field to given value.

### HasSlaves

`func (o *VirtualizationBondState) HasSlaves() bool`

HasSlaves returns a boolean if a field has been set.

### SetSlavesNil

`func (o *VirtualizationBondState) SetSlavesNil(b bool)`

 SetSlavesNil sets the value for Slaves to be an explicit nil

### UnsetSlaves
`func (o *VirtualizationBondState) UnsetSlaves()`

UnsetSlaves ensures that no value is present for Slaves, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


