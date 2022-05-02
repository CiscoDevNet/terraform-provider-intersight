# StorageNetAppNodeCdpNeighbor

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

### NewStorageNetAppNodeCdpNeighbor

`func NewStorageNetAppNodeCdpNeighbor(classId string, objectType string, ) *StorageNetAppNodeCdpNeighbor`

NewStorageNetAppNodeCdpNeighbor instantiates a new StorageNetAppNodeCdpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNodeCdpNeighborWithDefaults

`func NewStorageNetAppNodeCdpNeighborWithDefaults() *StorageNetAppNodeCdpNeighbor`

NewStorageNetAppNodeCdpNeighborWithDefaults instantiates a new StorageNetAppNodeCdpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNodeCdpNeighbor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNodeCdpNeighbor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNodeCdpNeighbor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNodeCdpNeighbor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNodeCdpNeighbor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNodeCdpNeighbor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapabilities

`func (o *StorageNetAppNodeCdpNeighbor) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *StorageNetAppNodeCdpNeighbor) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *StorageNetAppNodeCdpNeighbor) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *StorageNetAppNodeCdpNeighbor) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *StorageNetAppNodeCdpNeighbor) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *StorageNetAppNodeCdpNeighbor) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDeviceIp

`func (o *StorageNetAppNodeCdpNeighbor) GetDeviceIp() string`

GetDeviceIp returns the DeviceIp field if non-nil, zero value otherwise.

### GetDeviceIpOk

`func (o *StorageNetAppNodeCdpNeighbor) GetDeviceIpOk() (*string, bool)`

GetDeviceIpOk returns a tuple with the DeviceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIp

`func (o *StorageNetAppNodeCdpNeighbor) SetDeviceIp(v string)`

SetDeviceIp sets DeviceIp field to given value.

### HasDeviceIp

`func (o *StorageNetAppNodeCdpNeighbor) HasDeviceIp() bool`

HasDeviceIp returns a boolean if a field has been set.

### GetDiscoveredDevice

`func (o *StorageNetAppNodeCdpNeighbor) GetDiscoveredDevice() string`

GetDiscoveredDevice returns the DiscoveredDevice field if non-nil, zero value otherwise.

### GetDiscoveredDeviceOk

`func (o *StorageNetAppNodeCdpNeighbor) GetDiscoveredDeviceOk() (*string, bool)`

GetDiscoveredDeviceOk returns a tuple with the DiscoveredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredDevice

`func (o *StorageNetAppNodeCdpNeighbor) SetDiscoveredDevice(v string)`

SetDiscoveredDevice sets DiscoveredDevice field to given value.

### HasDiscoveredDevice

`func (o *StorageNetAppNodeCdpNeighbor) HasDiscoveredDevice() bool`

HasDiscoveredDevice returns a boolean if a field has been set.

### GetHoldTimeRemaining

`func (o *StorageNetAppNodeCdpNeighbor) GetHoldTimeRemaining() int64`

GetHoldTimeRemaining returns the HoldTimeRemaining field if non-nil, zero value otherwise.

### GetHoldTimeRemainingOk

`func (o *StorageNetAppNodeCdpNeighbor) GetHoldTimeRemainingOk() (*int64, bool)`

GetHoldTimeRemainingOk returns a tuple with the HoldTimeRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimeRemaining

`func (o *StorageNetAppNodeCdpNeighbor) SetHoldTimeRemaining(v int64)`

SetHoldTimeRemaining sets HoldTimeRemaining field to given value.

### HasHoldTimeRemaining

`func (o *StorageNetAppNodeCdpNeighbor) HasHoldTimeRemaining() bool`

HasHoldTimeRemaining returns a boolean if a field has been set.

### GetInterface

`func (o *StorageNetAppNodeCdpNeighbor) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *StorageNetAppNodeCdpNeighbor) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *StorageNetAppNodeCdpNeighbor) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *StorageNetAppNodeCdpNeighbor) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNodeName

`func (o *StorageNetAppNodeCdpNeighbor) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *StorageNetAppNodeCdpNeighbor) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *StorageNetAppNodeCdpNeighbor) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *StorageNetAppNodeCdpNeighbor) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPlatform

`func (o *StorageNetAppNodeCdpNeighbor) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *StorageNetAppNodeCdpNeighbor) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *StorageNetAppNodeCdpNeighbor) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *StorageNetAppNodeCdpNeighbor) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPort

`func (o *StorageNetAppNodeCdpNeighbor) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StorageNetAppNodeCdpNeighbor) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StorageNetAppNodeCdpNeighbor) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *StorageNetAppNodeCdpNeighbor) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *StorageNetAppNodeCdpNeighbor) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageNetAppNodeCdpNeighbor) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageNetAppNodeCdpNeighbor) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageNetAppNodeCdpNeighbor) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetVersion

`func (o *StorageNetAppNodeCdpNeighbor) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageNetAppNodeCdpNeighbor) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageNetAppNodeCdpNeighbor) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageNetAppNodeCdpNeighbor) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppNodeCdpNeighbor) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppNodeCdpNeighbor) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppNodeCdpNeighbor) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppNodeCdpNeighbor) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


