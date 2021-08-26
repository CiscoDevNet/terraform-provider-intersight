# EquipmentFru

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Fru"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Fru"]
**Action** | Pointer to **string** | This field identifies the action performed on a component. * &#x60;None&#x60; - No action performed on the FRU. * &#x60;Inserted&#x60; - A new FRU is inserted or added. * &#x60;Removed&#x60; - The previous FRU is removed. * &#x60;Replaced&#x60; - The previous FRU is replaced with a new FRU. * &#x60;ReplacedWithAlarm&#x60; - The previous FRU is replaced with a new FRU and a alarm is raised. | [optional] [default to "None"]
**CurrentFru** | Pointer to [**EquipmentBaseRelationship**](EquipmentBaseRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentFru

`func NewEquipmentFru(classId string, objectType string, ) *EquipmentFru`

NewEquipmentFru instantiates a new EquipmentFru object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFruWithDefaults

`func NewEquipmentFruWithDefaults() *EquipmentFru`

NewEquipmentFruWithDefaults instantiates a new EquipmentFru object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentFru) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentFru) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentFru) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentFru) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentFru) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentFru) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *EquipmentFru) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EquipmentFru) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EquipmentFru) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EquipmentFru) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCurrentFru

`func (o *EquipmentFru) GetCurrentFru() EquipmentBaseRelationship`

GetCurrentFru returns the CurrentFru field if non-nil, zero value otherwise.

### GetCurrentFruOk

`func (o *EquipmentFru) GetCurrentFruOk() (*EquipmentBaseRelationship, bool)`

GetCurrentFruOk returns a tuple with the CurrentFru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFru

`func (o *EquipmentFru) SetCurrentFru(v EquipmentBaseRelationship)`

SetCurrentFru sets CurrentFru field to given value.

### HasCurrentFru

`func (o *EquipmentFru) HasCurrentFru() bool`

HasCurrentFru returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentFru) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentFru) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentFru) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentFru) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentFru) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentFru) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentFru) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentFru) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


