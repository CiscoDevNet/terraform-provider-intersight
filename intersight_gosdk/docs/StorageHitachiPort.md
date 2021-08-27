# StorageHitachiPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiPort"]
**FabricMode** | Pointer to **bool** | Fabric mode of the port. true, Set. false, Not set. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | IPv4 address of Hitachi Port. | [optional] [readonly] 
**Ipv6GlobalAddress** | Pointer to **string** | IPv6 global address value. | [optional] [readonly] 
**Ipv6LinkLocalAddress** | Pointer to **string** | IPv6 link local address value. | [optional] [readonly] 
**IsIpv6Enable** | Pointer to **bool** | IPv6 mode. | [optional] [readonly] 
**LoopId** | Pointer to **string** | The value set for the port loop ID (AL_PA). | [optional] [readonly] 
**PortConnection** | Pointer to **string** | Topology setting for the port. | [optional] [readonly] 
**PortLunSecurity** | Pointer to **bool** | LUN security setting for the port. | [optional] [readonly] 
**ShortportId** | Pointer to **string** | Port ID (short) of the port. | [optional] [readonly] 
**TcpMtu** | Pointer to **int64** | Value of MTU for iSCSI communication. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiPort

`func NewStorageHitachiPort(classId string, objectType string, ) *StorageHitachiPort`

NewStorageHitachiPort instantiates a new StorageHitachiPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiPortWithDefaults

`func NewStorageHitachiPortWithDefaults() *StorageHitachiPort`

NewStorageHitachiPortWithDefaults instantiates a new StorageHitachiPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFabricMode

`func (o *StorageHitachiPort) GetFabricMode() bool`

GetFabricMode returns the FabricMode field if non-nil, zero value otherwise.

### GetFabricModeOk

`func (o *StorageHitachiPort) GetFabricModeOk() (*bool, bool)`

GetFabricModeOk returns a tuple with the FabricMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricMode

`func (o *StorageHitachiPort) SetFabricMode(v bool)`

SetFabricMode sets FabricMode field to given value.

### HasFabricMode

`func (o *StorageHitachiPort) HasFabricMode() bool`

HasFabricMode returns a boolean if a field has been set.

### GetIpv4Address

`func (o *StorageHitachiPort) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *StorageHitachiPort) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *StorageHitachiPort) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *StorageHitachiPort) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6GlobalAddress

`func (o *StorageHitachiPort) GetIpv6GlobalAddress() string`

GetIpv6GlobalAddress returns the Ipv6GlobalAddress field if non-nil, zero value otherwise.

### GetIpv6GlobalAddressOk

`func (o *StorageHitachiPort) GetIpv6GlobalAddressOk() (*string, bool)`

GetIpv6GlobalAddressOk returns a tuple with the Ipv6GlobalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6GlobalAddress

`func (o *StorageHitachiPort) SetIpv6GlobalAddress(v string)`

SetIpv6GlobalAddress sets Ipv6GlobalAddress field to given value.

### HasIpv6GlobalAddress

`func (o *StorageHitachiPort) HasIpv6GlobalAddress() bool`

HasIpv6GlobalAddress returns a boolean if a field has been set.

### GetIpv6LinkLocalAddress

`func (o *StorageHitachiPort) GetIpv6LinkLocalAddress() string`

GetIpv6LinkLocalAddress returns the Ipv6LinkLocalAddress field if non-nil, zero value otherwise.

### GetIpv6LinkLocalAddressOk

`func (o *StorageHitachiPort) GetIpv6LinkLocalAddressOk() (*string, bool)`

GetIpv6LinkLocalAddressOk returns a tuple with the Ipv6LinkLocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6LinkLocalAddress

`func (o *StorageHitachiPort) SetIpv6LinkLocalAddress(v string)`

SetIpv6LinkLocalAddress sets Ipv6LinkLocalAddress field to given value.

### HasIpv6LinkLocalAddress

`func (o *StorageHitachiPort) HasIpv6LinkLocalAddress() bool`

HasIpv6LinkLocalAddress returns a boolean if a field has been set.

### GetIsIpv6Enable

`func (o *StorageHitachiPort) GetIsIpv6Enable() bool`

GetIsIpv6Enable returns the IsIpv6Enable field if non-nil, zero value otherwise.

### GetIsIpv6EnableOk

`func (o *StorageHitachiPort) GetIsIpv6EnableOk() (*bool, bool)`

GetIsIpv6EnableOk returns a tuple with the IsIpv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIpv6Enable

`func (o *StorageHitachiPort) SetIsIpv6Enable(v bool)`

SetIsIpv6Enable sets IsIpv6Enable field to given value.

### HasIsIpv6Enable

`func (o *StorageHitachiPort) HasIsIpv6Enable() bool`

HasIsIpv6Enable returns a boolean if a field has been set.

### GetLoopId

`func (o *StorageHitachiPort) GetLoopId() string`

GetLoopId returns the LoopId field if non-nil, zero value otherwise.

### GetLoopIdOk

`func (o *StorageHitachiPort) GetLoopIdOk() (*string, bool)`

GetLoopIdOk returns a tuple with the LoopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopId

`func (o *StorageHitachiPort) SetLoopId(v string)`

SetLoopId sets LoopId field to given value.

### HasLoopId

`func (o *StorageHitachiPort) HasLoopId() bool`

HasLoopId returns a boolean if a field has been set.

### GetPortConnection

`func (o *StorageHitachiPort) GetPortConnection() string`

GetPortConnection returns the PortConnection field if non-nil, zero value otherwise.

### GetPortConnectionOk

`func (o *StorageHitachiPort) GetPortConnectionOk() (*string, bool)`

GetPortConnectionOk returns a tuple with the PortConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConnection

`func (o *StorageHitachiPort) SetPortConnection(v string)`

SetPortConnection sets PortConnection field to given value.

### HasPortConnection

`func (o *StorageHitachiPort) HasPortConnection() bool`

HasPortConnection returns a boolean if a field has been set.

### GetPortLunSecurity

`func (o *StorageHitachiPort) GetPortLunSecurity() bool`

GetPortLunSecurity returns the PortLunSecurity field if non-nil, zero value otherwise.

### GetPortLunSecurityOk

`func (o *StorageHitachiPort) GetPortLunSecurityOk() (*bool, bool)`

GetPortLunSecurityOk returns a tuple with the PortLunSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLunSecurity

`func (o *StorageHitachiPort) SetPortLunSecurity(v bool)`

SetPortLunSecurity sets PortLunSecurity field to given value.

### HasPortLunSecurity

`func (o *StorageHitachiPort) HasPortLunSecurity() bool`

HasPortLunSecurity returns a boolean if a field has been set.

### GetShortportId

`func (o *StorageHitachiPort) GetShortportId() string`

GetShortportId returns the ShortportId field if non-nil, zero value otherwise.

### GetShortportIdOk

`func (o *StorageHitachiPort) GetShortportIdOk() (*string, bool)`

GetShortportIdOk returns a tuple with the ShortportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortportId

`func (o *StorageHitachiPort) SetShortportId(v string)`

SetShortportId sets ShortportId field to given value.

### HasShortportId

`func (o *StorageHitachiPort) HasShortportId() bool`

HasShortportId returns a boolean if a field has been set.

### GetTcpMtu

`func (o *StorageHitachiPort) GetTcpMtu() int64`

GetTcpMtu returns the TcpMtu field if non-nil, zero value otherwise.

### GetTcpMtuOk

`func (o *StorageHitachiPort) GetTcpMtuOk() (*int64, bool)`

GetTcpMtuOk returns a tuple with the TcpMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMtu

`func (o *StorageHitachiPort) SetTcpMtu(v int64)`

SetTcpMtu sets TcpMtu field to given value.

### HasTcpMtu

`func (o *StorageHitachiPort) HasTcpMtu() bool`

HasTcpMtu returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiPort) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiPort) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiPort) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiPort) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


