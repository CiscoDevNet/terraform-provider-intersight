# CapabilityEquipmentPhysicalDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.EquipmentPhysicalDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.EquipmentPhysicalDef"]
**Depth** | Pointer to **float32** | Depth information for a Switch/Fabric-Interconnect. | [optional] 
**Height** | Pointer to **float32** | Height information for a Switch/Fabric-Interconnect. | [optional] 
**MaxPower** | Pointer to **int64** | Max Power information for a Switch/Fabric-Interconnect. | [optional] 
**MinPower** | Pointer to **int64** | Min Power information for a Switch/Fabric-Interconnect. | [optional] 
**NominalPower** | Pointer to **int64** | Nominal Power information for a Switch/Fabric-Interconnect. | [optional] 
**Weight** | Pointer to **float32** | Weight information for a Switch/Fabric-Interconnect. | [optional] 
**Width** | Pointer to **float32** | Width information for a Switch/Fabric-Interconnect. | [optional] 

## Methods

### NewCapabilityEquipmentPhysicalDefAllOf

`func NewCapabilityEquipmentPhysicalDefAllOf(classId string, objectType string, ) *CapabilityEquipmentPhysicalDefAllOf`

NewCapabilityEquipmentPhysicalDefAllOf instantiates a new CapabilityEquipmentPhysicalDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityEquipmentPhysicalDefAllOfWithDefaults

`func NewCapabilityEquipmentPhysicalDefAllOfWithDefaults() *CapabilityEquipmentPhysicalDefAllOf`

NewCapabilityEquipmentPhysicalDefAllOfWithDefaults instantiates a new CapabilityEquipmentPhysicalDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDepth

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetDepth() float32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetDepthOk() (*float32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetDepth(v float32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *CapabilityEquipmentPhysicalDefAllOf) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetHeight

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *CapabilityEquipmentPhysicalDefAllOf) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMaxPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetMaxPower() int64`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetMaxPowerOk() (*int64, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetMaxPower(v int64)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetMinPower() int64`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetMinPowerOk() (*int64, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetMinPower(v int64)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetNominalPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetNominalPower() int64`

GetNominalPower returns the NominalPower field if non-nil, zero value otherwise.

### GetNominalPowerOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetNominalPowerOk() (*int64, bool)`

GetNominalPowerOk returns a tuple with the NominalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominalPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetNominalPower(v int64)`

SetNominalPower sets NominalPower field to given value.

### HasNominalPower

`func (o *CapabilityEquipmentPhysicalDefAllOf) HasNominalPower() bool`

HasNominalPower returns a boolean if a field has been set.

### GetWeight

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CapabilityEquipmentPhysicalDefAllOf) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWidth

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CapabilityEquipmentPhysicalDefAllOf) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CapabilityEquipmentPhysicalDefAllOf) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CapabilityEquipmentPhysicalDefAllOf) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


