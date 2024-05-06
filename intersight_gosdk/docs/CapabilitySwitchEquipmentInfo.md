# CapabilitySwitchEquipmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchEquipmentInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchEquipmentInfo"]
**MaxFanModules** | Pointer to **int64** | Maximum number of fan modules/trays present on this switch. | [optional] [readonly] 
**MaxFansInFanModule** | Pointer to **int64** | Maximum number of fans present in a fan module/tray on this switch. | [optional] [readonly] 

## Methods

### NewCapabilitySwitchEquipmentInfo

`func NewCapabilitySwitchEquipmentInfo(classId string, objectType string, ) *CapabilitySwitchEquipmentInfo`

NewCapabilitySwitchEquipmentInfo instantiates a new CapabilitySwitchEquipmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchEquipmentInfoWithDefaults

`func NewCapabilitySwitchEquipmentInfoWithDefaults() *CapabilitySwitchEquipmentInfo`

NewCapabilitySwitchEquipmentInfoWithDefaults instantiates a new CapabilitySwitchEquipmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchEquipmentInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchEquipmentInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchEquipmentInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchEquipmentInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchEquipmentInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchEquipmentInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaxFanModules

`func (o *CapabilitySwitchEquipmentInfo) GetMaxFanModules() int64`

GetMaxFanModules returns the MaxFanModules field if non-nil, zero value otherwise.

### GetMaxFanModulesOk

`func (o *CapabilitySwitchEquipmentInfo) GetMaxFanModulesOk() (*int64, bool)`

GetMaxFanModulesOk returns a tuple with the MaxFanModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFanModules

`func (o *CapabilitySwitchEquipmentInfo) SetMaxFanModules(v int64)`

SetMaxFanModules sets MaxFanModules field to given value.

### HasMaxFanModules

`func (o *CapabilitySwitchEquipmentInfo) HasMaxFanModules() bool`

HasMaxFanModules returns a boolean if a field has been set.

### GetMaxFansInFanModule

`func (o *CapabilitySwitchEquipmentInfo) GetMaxFansInFanModule() int64`

GetMaxFansInFanModule returns the MaxFansInFanModule field if non-nil, zero value otherwise.

### GetMaxFansInFanModuleOk

`func (o *CapabilitySwitchEquipmentInfo) GetMaxFansInFanModuleOk() (*int64, bool)`

GetMaxFansInFanModuleOk returns a tuple with the MaxFansInFanModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFansInFanModule

`func (o *CapabilitySwitchEquipmentInfo) SetMaxFansInFanModule(v int64)`

SetMaxFansInFanModule sets MaxFansInFanModule field to given value.

### HasMaxFansInFanModule

`func (o *CapabilitySwitchEquipmentInfo) HasMaxFansInFanModule() bool`

HasMaxFansInFanModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


