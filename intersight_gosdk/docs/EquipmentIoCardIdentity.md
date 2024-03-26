# EquipmentIoCardIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.IoCardIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.IoCardIdentity"]
**IoCardMoid** | Pointer to **string** | MO Reference to equipmentIoCard MO in inventory service. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | IOM/MUX Module ID connected to the FI. | [optional] [readonly] 
**NetworkElementMoid** | Pointer to **string** | MO Reference to networkElement MO in inventory service. | [optional] [readonly] 
**SwitchId** | Pointer to **int64** | Switch ID to which IOM is connected, ID can be either 1 or 2, equalent to A or B. | [optional] [readonly] 

## Methods

### NewEquipmentIoCardIdentity

`func NewEquipmentIoCardIdentity(classId string, objectType string, ) *EquipmentIoCardIdentity`

NewEquipmentIoCardIdentity instantiates a new EquipmentIoCardIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardIdentityWithDefaults

`func NewEquipmentIoCardIdentityWithDefaults() *EquipmentIoCardIdentity`

NewEquipmentIoCardIdentityWithDefaults instantiates a new EquipmentIoCardIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIoCardIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIoCardIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIoCardIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIoCardIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIoCardIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIoCardIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIoCardMoid

`func (o *EquipmentIoCardIdentity) GetIoCardMoid() string`

GetIoCardMoid returns the IoCardMoid field if non-nil, zero value otherwise.

### GetIoCardMoidOk

`func (o *EquipmentIoCardIdentity) GetIoCardMoidOk() (*string, bool)`

GetIoCardMoidOk returns a tuple with the IoCardMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardMoid

`func (o *EquipmentIoCardIdentity) SetIoCardMoid(v string)`

SetIoCardMoid sets IoCardMoid field to given value.

### HasIoCardMoid

`func (o *EquipmentIoCardIdentity) HasIoCardMoid() bool`

HasIoCardMoid returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentIoCardIdentity) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentIoCardIdentity) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentIoCardIdentity) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentIoCardIdentity) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetNetworkElementMoid

`func (o *EquipmentIoCardIdentity) GetNetworkElementMoid() string`

GetNetworkElementMoid returns the NetworkElementMoid field if non-nil, zero value otherwise.

### GetNetworkElementMoidOk

`func (o *EquipmentIoCardIdentity) GetNetworkElementMoidOk() (*string, bool)`

GetNetworkElementMoidOk returns a tuple with the NetworkElementMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElementMoid

`func (o *EquipmentIoCardIdentity) SetNetworkElementMoid(v string)`

SetNetworkElementMoid sets NetworkElementMoid field to given value.

### HasNetworkElementMoid

`func (o *EquipmentIoCardIdentity) HasNetworkElementMoid() bool`

HasNetworkElementMoid returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentIoCardIdentity) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentIoCardIdentity) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentIoCardIdentity) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentIoCardIdentity) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


