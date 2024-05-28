# EquipmentDeviceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.DeviceSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.DeviceSummary"]
**Dn** | Pointer to **string** | The distinguished name that unambiguously identifies an object in the system. | [optional] [readonly] 
**Model** | Pointer to **string** | This field identifies the model of the given component. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field identifies the serial number of the given component. | [optional] [readonly] 
**SourceObjectType** | Pointer to **string** | The source object type of the given component. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**CustomPermissionResources** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**EquipmentFex** | Pointer to [**NullableEquipmentFexRelationship**](EquipmentFexRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**InventoryParent** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentDeviceSummary

`func NewEquipmentDeviceSummary(classId string, objectType string, ) *EquipmentDeviceSummary`

NewEquipmentDeviceSummary instantiates a new EquipmentDeviceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentDeviceSummaryWithDefaults

`func NewEquipmentDeviceSummaryWithDefaults() *EquipmentDeviceSummary`

NewEquipmentDeviceSummaryWithDefaults instantiates a new EquipmentDeviceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentDeviceSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentDeviceSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentDeviceSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentDeviceSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentDeviceSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentDeviceSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *EquipmentDeviceSummary) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EquipmentDeviceSummary) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EquipmentDeviceSummary) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EquipmentDeviceSummary) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentDeviceSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentDeviceSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentDeviceSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentDeviceSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentDeviceSummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentDeviceSummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentDeviceSummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentDeviceSummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSourceObjectType

`func (o *EquipmentDeviceSummary) GetSourceObjectType() string`

GetSourceObjectType returns the SourceObjectType field if non-nil, zero value otherwise.

### GetSourceObjectTypeOk

`func (o *EquipmentDeviceSummary) GetSourceObjectTypeOk() (*string, bool)`

GetSourceObjectTypeOk returns a tuple with the SourceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectType

`func (o *EquipmentDeviceSummary) SetSourceObjectType(v string)`

SetSourceObjectType sets SourceObjectType field to given value.

### HasSourceObjectType

`func (o *EquipmentDeviceSummary) HasSourceObjectType() bool`

HasSourceObjectType returns a boolean if a field has been set.

### GetComputeBlade

`func (o *EquipmentDeviceSummary) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *EquipmentDeviceSummary) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *EquipmentDeviceSummary) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *EquipmentDeviceSummary) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### SetComputeBladeNil

`func (o *EquipmentDeviceSummary) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *EquipmentDeviceSummary) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
### GetComputeRackUnit

`func (o *EquipmentDeviceSummary) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentDeviceSummary) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentDeviceSummary) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentDeviceSummary) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### SetComputeRackUnitNil

`func (o *EquipmentDeviceSummary) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *EquipmentDeviceSummary) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
### GetCustomPermissionResources

`func (o *EquipmentDeviceSummary) GetCustomPermissionResources() []MoBaseMoRelationship`

GetCustomPermissionResources returns the CustomPermissionResources field if non-nil, zero value otherwise.

### GetCustomPermissionResourcesOk

`func (o *EquipmentDeviceSummary) GetCustomPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetCustomPermissionResourcesOk returns a tuple with the CustomPermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPermissionResources

`func (o *EquipmentDeviceSummary) SetCustomPermissionResources(v []MoBaseMoRelationship)`

SetCustomPermissionResources sets CustomPermissionResources field to given value.

### HasCustomPermissionResources

`func (o *EquipmentDeviceSummary) HasCustomPermissionResources() bool`

HasCustomPermissionResources returns a boolean if a field has been set.

### SetCustomPermissionResourcesNil

`func (o *EquipmentDeviceSummary) SetCustomPermissionResourcesNil(b bool)`

 SetCustomPermissionResourcesNil sets the value for CustomPermissionResources to be an explicit nil

### UnsetCustomPermissionResources
`func (o *EquipmentDeviceSummary) UnsetCustomPermissionResources()`

UnsetCustomPermissionResources ensures that no value is present for CustomPermissionResources, not even an explicit nil
### GetEquipmentChassis

`func (o *EquipmentDeviceSummary) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentDeviceSummary) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentDeviceSummary) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentDeviceSummary) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *EquipmentDeviceSummary) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *EquipmentDeviceSummary) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetEquipmentFex

`func (o *EquipmentDeviceSummary) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentDeviceSummary) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentDeviceSummary) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentDeviceSummary) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

### SetEquipmentFexNil

`func (o *EquipmentDeviceSummary) SetEquipmentFexNil(b bool)`

 SetEquipmentFexNil sets the value for EquipmentFex to be an explicit nil

### UnsetEquipmentFex
`func (o *EquipmentDeviceSummary) UnsetEquipmentFex()`

UnsetEquipmentFex ensures that no value is present for EquipmentFex, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentDeviceSummary) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentDeviceSummary) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentDeviceSummary) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentDeviceSummary) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *EquipmentDeviceSummary) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *EquipmentDeviceSummary) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetInventoryParent

`func (o *EquipmentDeviceSummary) GetInventoryParent() MoBaseMoRelationship`

GetInventoryParent returns the InventoryParent field if non-nil, zero value otherwise.

### GetInventoryParentOk

`func (o *EquipmentDeviceSummary) GetInventoryParentOk() (*MoBaseMoRelationship, bool)`

GetInventoryParentOk returns a tuple with the InventoryParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryParent

`func (o *EquipmentDeviceSummary) SetInventoryParent(v MoBaseMoRelationship)`

SetInventoryParent sets InventoryParent field to given value.

### HasInventoryParent

`func (o *EquipmentDeviceSummary) HasInventoryParent() bool`

HasInventoryParent returns a boolean if a field has been set.

### SetInventoryParentNil

`func (o *EquipmentDeviceSummary) SetInventoryParentNil(b bool)`

 SetInventoryParentNil sets the value for InventoryParent to be an explicit nil

### UnsetInventoryParent
`func (o *EquipmentDeviceSummary) UnsetInventoryParent()`

UnsetInventoryParent ensures that no value is present for InventoryParent, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentDeviceSummary) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentDeviceSummary) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentDeviceSummary) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentDeviceSummary) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentDeviceSummary) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentDeviceSummary) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


