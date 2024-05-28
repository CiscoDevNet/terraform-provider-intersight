# NetworkFcZoneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.FcZoneInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.FcZoneInfo"]
**UserZoneCount** | Pointer to **int64** | The number of Fibre Channel user zones defined on a Fabric Interconnect. | [optional] [readonly] 
**UserZoneLimit** | Pointer to **int64** | The maximum number of Fibre Channel user zones allowed on a Fabric Interconnect. | [optional] [readonly] 
**ZoneCount** | Pointer to **int64** | The number of Fibre Channel zones defined on a Fabric Interconnect. | [optional] [readonly] 
**ZoneLimit** | Pointer to **int64** | The maximum number of Fibre Channel zones allowed on a Fabric Interconnect. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkFcZoneInfo

`func NewNetworkFcZoneInfo(classId string, objectType string, ) *NetworkFcZoneInfo`

NewNetworkFcZoneInfo instantiates a new NetworkFcZoneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFcZoneInfoWithDefaults

`func NewNetworkFcZoneInfoWithDefaults() *NetworkFcZoneInfo`

NewNetworkFcZoneInfoWithDefaults instantiates a new NetworkFcZoneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkFcZoneInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkFcZoneInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkFcZoneInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkFcZoneInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkFcZoneInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkFcZoneInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUserZoneCount

`func (o *NetworkFcZoneInfo) GetUserZoneCount() int64`

GetUserZoneCount returns the UserZoneCount field if non-nil, zero value otherwise.

### GetUserZoneCountOk

`func (o *NetworkFcZoneInfo) GetUserZoneCountOk() (*int64, bool)`

GetUserZoneCountOk returns a tuple with the UserZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserZoneCount

`func (o *NetworkFcZoneInfo) SetUserZoneCount(v int64)`

SetUserZoneCount sets UserZoneCount field to given value.

### HasUserZoneCount

`func (o *NetworkFcZoneInfo) HasUserZoneCount() bool`

HasUserZoneCount returns a boolean if a field has been set.

### GetUserZoneLimit

`func (o *NetworkFcZoneInfo) GetUserZoneLimit() int64`

GetUserZoneLimit returns the UserZoneLimit field if non-nil, zero value otherwise.

### GetUserZoneLimitOk

`func (o *NetworkFcZoneInfo) GetUserZoneLimitOk() (*int64, bool)`

GetUserZoneLimitOk returns a tuple with the UserZoneLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserZoneLimit

`func (o *NetworkFcZoneInfo) SetUserZoneLimit(v int64)`

SetUserZoneLimit sets UserZoneLimit field to given value.

### HasUserZoneLimit

`func (o *NetworkFcZoneInfo) HasUserZoneLimit() bool`

HasUserZoneLimit returns a boolean if a field has been set.

### GetZoneCount

`func (o *NetworkFcZoneInfo) GetZoneCount() int64`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *NetworkFcZoneInfo) GetZoneCountOk() (*int64, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *NetworkFcZoneInfo) SetZoneCount(v int64)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *NetworkFcZoneInfo) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.

### GetZoneLimit

`func (o *NetworkFcZoneInfo) GetZoneLimit() int64`

GetZoneLimit returns the ZoneLimit field if non-nil, zero value otherwise.

### GetZoneLimitOk

`func (o *NetworkFcZoneInfo) GetZoneLimitOk() (*int64, bool)`

GetZoneLimitOk returns a tuple with the ZoneLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLimit

`func (o *NetworkFcZoneInfo) SetZoneLimit(v int64)`

SetZoneLimit sets ZoneLimit field to given value.

### HasZoneLimit

`func (o *NetworkFcZoneInfo) HasZoneLimit() bool`

HasZoneLimit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *NetworkFcZoneInfo) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *NetworkFcZoneInfo) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *NetworkFcZoneInfo) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *NetworkFcZoneInfo) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *NetworkFcZoneInfo) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *NetworkFcZoneInfo) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetNetworkElement

`func (o *NetworkFcZoneInfo) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkFcZoneInfo) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkFcZoneInfo) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkFcZoneInfo) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkFcZoneInfo) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkFcZoneInfo) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkFcZoneInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkFcZoneInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkFcZoneInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkFcZoneInfo) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkFcZoneInfo) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkFcZoneInfo) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


