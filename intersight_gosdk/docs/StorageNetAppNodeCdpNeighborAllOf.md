# StorageNetAppNodeCdpNeighborAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNodeCdpNeighbor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNodeCdpNeighbor"]
**Capabilities** | Pointer to **[]string** |  | [optional] 
**DeviceIp** | Pointer to **string** | The IP address of the CDP neighbor. | [optional] [readonly] 
**DiscoveredDevice** | Pointer to **string** | The name of the CDP neighbor. | [optional] [readonly] 
**HoldTimeRemaining** | Pointer to **int64** | The period of time for which CDP advertisements are cached. | [optional] [readonly] 
**Interface** | Pointer to **string** | The interface of the CDP neighbor. | [optional] [readonly] 
**NodeName** | Pointer to **string** | The node that owns the port for this CDP neighbor. | [optional] [readonly] 
**Platform** | Pointer to **string** | The platform of the CDP neighbor. | [optional] [readonly] 
**Port** | Pointer to **string** | The port for this CDP neighbor. | [optional] [readonly] 
**Protocol** | Pointer to **string** | The protocol used for CDP. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the operating system running on the CDP neighbor. | [optional] [readonly] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNodeCdpNeighborAllOf

`func NewStorageNetAppNodeCdpNeighborAllOf(classId string, objectType string, ) *StorageNetAppNodeCdpNeighborAllOf`

NewStorageNetAppNodeCdpNeighborAllOf instantiates a new StorageNetAppNodeCdpNeighborAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNodeCdpNeighborAllOfWithDefaults

`func NewStorageNetAppNodeCdpNeighborAllOfWithDefaults() *StorageNetAppNodeCdpNeighborAllOf`

NewStorageNetAppNodeCdpNeighborAllOfWithDefaults instantiates a new StorageNetAppNodeCdpNeighborAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapabilities

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *StorageNetAppNodeCdpNeighborAllOf) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDeviceIp

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetDeviceIp() string`

GetDeviceIp returns the DeviceIp field if non-nil, zero value otherwise.

### GetDeviceIpOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetDeviceIpOk() (*string, bool)`

GetDeviceIpOk returns a tuple with the DeviceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIp

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetDeviceIp(v string)`

SetDeviceIp sets DeviceIp field to given value.

### HasDeviceIp

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasDeviceIp() bool`

HasDeviceIp returns a boolean if a field has been set.

### GetDiscoveredDevice

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetDiscoveredDevice() string`

GetDiscoveredDevice returns the DiscoveredDevice field if non-nil, zero value otherwise.

### GetDiscoveredDeviceOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetDiscoveredDeviceOk() (*string, bool)`

GetDiscoveredDeviceOk returns a tuple with the DiscoveredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredDevice

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetDiscoveredDevice(v string)`

SetDiscoveredDevice sets DiscoveredDevice field to given value.

### HasDiscoveredDevice

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasDiscoveredDevice() bool`

HasDiscoveredDevice returns a boolean if a field has been set.

### GetHoldTimeRemaining

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetHoldTimeRemaining() int64`

GetHoldTimeRemaining returns the HoldTimeRemaining field if non-nil, zero value otherwise.

### GetHoldTimeRemainingOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetHoldTimeRemainingOk() (*int64, bool)`

GetHoldTimeRemainingOk returns a tuple with the HoldTimeRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimeRemaining

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetHoldTimeRemaining(v int64)`

SetHoldTimeRemaining sets HoldTimeRemaining field to given value.

### HasHoldTimeRemaining

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasHoldTimeRemaining() bool`

HasHoldTimeRemaining returns a boolean if a field has been set.

### GetInterface

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNodeName

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPlatform

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPort

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetVersion

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppNodeCdpNeighborAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppNodeCdpNeighborAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppNodeCdpNeighborAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


