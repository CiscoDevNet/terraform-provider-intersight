# NetworkFcZoneInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.FcZoneInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.FcZoneInfo"]
**UserZoneCount** | Pointer to **int64** | The number of Fibre Channel user zones defined on a Fabric Interconnect. | [optional] [readonly] 
**UserZoneLimit** | Pointer to **int64** | The maximum number of Fibre Channel user zones allowed on a Fabric Interconnect. | [optional] [readonly] 
**ZoneCount** | Pointer to **int64** | The number of Fibre Channel zones defined on a Fabric Interconnect. | [optional] [readonly] 
**ZoneLimit** | Pointer to **int64** | The maximum number of Fibre Channel zones allowed on a Fabric Interconnect. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkFcZoneInfoAllOf

`func NewNetworkFcZoneInfoAllOf(classId string, objectType string, ) *NetworkFcZoneInfoAllOf`

NewNetworkFcZoneInfoAllOf instantiates a new NetworkFcZoneInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFcZoneInfoAllOfWithDefaults

`func NewNetworkFcZoneInfoAllOfWithDefaults() *NetworkFcZoneInfoAllOf`

NewNetworkFcZoneInfoAllOfWithDefaults instantiates a new NetworkFcZoneInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkFcZoneInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkFcZoneInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkFcZoneInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkFcZoneInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkFcZoneInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkFcZoneInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUserZoneCount

`func (o *NetworkFcZoneInfoAllOf) GetUserZoneCount() int64`

GetUserZoneCount returns the UserZoneCount field if non-nil, zero value otherwise.

### GetUserZoneCountOk

`func (o *NetworkFcZoneInfoAllOf) GetUserZoneCountOk() (*int64, bool)`

GetUserZoneCountOk returns a tuple with the UserZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserZoneCount

`func (o *NetworkFcZoneInfoAllOf) SetUserZoneCount(v int64)`

SetUserZoneCount sets UserZoneCount field to given value.

### HasUserZoneCount

`func (o *NetworkFcZoneInfoAllOf) HasUserZoneCount() bool`

HasUserZoneCount returns a boolean if a field has been set.

### GetUserZoneLimit

`func (o *NetworkFcZoneInfoAllOf) GetUserZoneLimit() int64`

GetUserZoneLimit returns the UserZoneLimit field if non-nil, zero value otherwise.

### GetUserZoneLimitOk

`func (o *NetworkFcZoneInfoAllOf) GetUserZoneLimitOk() (*int64, bool)`

GetUserZoneLimitOk returns a tuple with the UserZoneLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserZoneLimit

`func (o *NetworkFcZoneInfoAllOf) SetUserZoneLimit(v int64)`

SetUserZoneLimit sets UserZoneLimit field to given value.

### HasUserZoneLimit

`func (o *NetworkFcZoneInfoAllOf) HasUserZoneLimit() bool`

HasUserZoneLimit returns a boolean if a field has been set.

### GetZoneCount

`func (o *NetworkFcZoneInfoAllOf) GetZoneCount() int64`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *NetworkFcZoneInfoAllOf) GetZoneCountOk() (*int64, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *NetworkFcZoneInfoAllOf) SetZoneCount(v int64)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *NetworkFcZoneInfoAllOf) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.

### GetZoneLimit

`func (o *NetworkFcZoneInfoAllOf) GetZoneLimit() int64`

GetZoneLimit returns the ZoneLimit field if non-nil, zero value otherwise.

### GetZoneLimitOk

`func (o *NetworkFcZoneInfoAllOf) GetZoneLimitOk() (*int64, bool)`

GetZoneLimitOk returns a tuple with the ZoneLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLimit

`func (o *NetworkFcZoneInfoAllOf) SetZoneLimit(v int64)`

SetZoneLimit sets ZoneLimit field to given value.

### HasZoneLimit

`func (o *NetworkFcZoneInfoAllOf) HasZoneLimit() bool`

HasZoneLimit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *NetworkFcZoneInfoAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *NetworkFcZoneInfoAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *NetworkFcZoneInfoAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *NetworkFcZoneInfoAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkFcZoneInfoAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkFcZoneInfoAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkFcZoneInfoAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkFcZoneInfoAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkFcZoneInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkFcZoneInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkFcZoneInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkFcZoneInfoAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


