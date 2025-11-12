# ComputePcieConnectivityPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.PcieConnectivityPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.PcieConnectivityPolicyInventory"]
**PcieZones** | Pointer to [**[]ComputePcieZone**](ComputePcieZone.md) |  | [optional] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewComputePcieConnectivityPolicyInventory

`func NewComputePcieConnectivityPolicyInventory(classId string, objectType string, ) *ComputePcieConnectivityPolicyInventory`

NewComputePcieConnectivityPolicyInventory instantiates a new ComputePcieConnectivityPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePcieConnectivityPolicyInventoryWithDefaults

`func NewComputePcieConnectivityPolicyInventoryWithDefaults() *ComputePcieConnectivityPolicyInventory`

NewComputePcieConnectivityPolicyInventoryWithDefaults instantiates a new ComputePcieConnectivityPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePcieConnectivityPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePcieConnectivityPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePcieConnectivityPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePcieConnectivityPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePcieConnectivityPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePcieConnectivityPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPcieZones

`func (o *ComputePcieConnectivityPolicyInventory) GetPcieZones() []ComputePcieZone`

GetPcieZones returns the PcieZones field if non-nil, zero value otherwise.

### GetPcieZonesOk

`func (o *ComputePcieConnectivityPolicyInventory) GetPcieZonesOk() (*[]ComputePcieZone, bool)`

GetPcieZonesOk returns a tuple with the PcieZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieZones

`func (o *ComputePcieConnectivityPolicyInventory) SetPcieZones(v []ComputePcieZone)`

SetPcieZones sets PcieZones field to given value.

### HasPcieZones

`func (o *ComputePcieConnectivityPolicyInventory) HasPcieZones() bool`

HasPcieZones returns a boolean if a field has been set.

### SetPcieZonesNil

`func (o *ComputePcieConnectivityPolicyInventory) SetPcieZonesNil(b bool)`

 SetPcieZonesNil sets the value for PcieZones to be an explicit nil

### UnsetPcieZones
`func (o *ComputePcieConnectivityPolicyInventory) UnsetPcieZones()`

UnsetPcieZones ensures that no value is present for PcieZones, not even an explicit nil
### GetTargetMo

`func (o *ComputePcieConnectivityPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *ComputePcieConnectivityPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *ComputePcieConnectivityPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *ComputePcieConnectivityPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *ComputePcieConnectivityPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *ComputePcieConnectivityPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


