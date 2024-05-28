# NetworkVpcPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.VpcPeer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.VpcPeer"]
**OperationalState** | Pointer to **string** | Operational state of the virtual port channel. | [optional] [readonly] 
**PortChannel** | Pointer to **string** | Name of the virtual port channel. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel identity of the virtual port channel. | [optional] [readonly] 
**VpcDomainId** | Pointer to **int64** | Identity of the virtual port channel. | [optional] [readonly] 
**VpcPeerId** | Pointer to **int64** | Identity of the virtual port channel. | [optional] [readonly] 
**EtherPortChannel** | Pointer to [**NullableEtherPortChannelRelationship**](EtherPortChannelRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVpcPeer

`func NewNetworkVpcPeer(classId string, objectType string, ) *NetworkVpcPeer`

NewNetworkVpcPeer instantiates a new NetworkVpcPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVpcPeerWithDefaults

`func NewNetworkVpcPeerWithDefaults() *NetworkVpcPeer`

NewNetworkVpcPeerWithDefaults instantiates a new NetworkVpcPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVpcPeer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVpcPeer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVpcPeer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVpcPeer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVpcPeer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVpcPeer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationalState

`func (o *NetworkVpcPeer) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *NetworkVpcPeer) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *NetworkVpcPeer) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *NetworkVpcPeer) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetPortChannel

`func (o *NetworkVpcPeer) GetPortChannel() string`

GetPortChannel returns the PortChannel field if non-nil, zero value otherwise.

### GetPortChannelOk

`func (o *NetworkVpcPeer) GetPortChannelOk() (*string, bool)`

GetPortChannelOk returns a tuple with the PortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannel

`func (o *NetworkVpcPeer) SetPortChannel(v string)`

SetPortChannel sets PortChannel field to given value.

### HasPortChannel

`func (o *NetworkVpcPeer) HasPortChannel() bool`

HasPortChannel returns a boolean if a field has been set.

### GetPortChannelId

`func (o *NetworkVpcPeer) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *NetworkVpcPeer) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *NetworkVpcPeer) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *NetworkVpcPeer) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetVpcDomainId

`func (o *NetworkVpcPeer) GetVpcDomainId() int64`

GetVpcDomainId returns the VpcDomainId field if non-nil, zero value otherwise.

### GetVpcDomainIdOk

`func (o *NetworkVpcPeer) GetVpcDomainIdOk() (*int64, bool)`

GetVpcDomainIdOk returns a tuple with the VpcDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcDomainId

`func (o *NetworkVpcPeer) SetVpcDomainId(v int64)`

SetVpcDomainId sets VpcDomainId field to given value.

### HasVpcDomainId

`func (o *NetworkVpcPeer) HasVpcDomainId() bool`

HasVpcDomainId returns a boolean if a field has been set.

### GetVpcPeerId

`func (o *NetworkVpcPeer) GetVpcPeerId() int64`

GetVpcPeerId returns the VpcPeerId field if non-nil, zero value otherwise.

### GetVpcPeerIdOk

`func (o *NetworkVpcPeer) GetVpcPeerIdOk() (*int64, bool)`

GetVpcPeerIdOk returns a tuple with the VpcPeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeerId

`func (o *NetworkVpcPeer) SetVpcPeerId(v int64)`

SetVpcPeerId sets VpcPeerId field to given value.

### HasVpcPeerId

`func (o *NetworkVpcPeer) HasVpcPeerId() bool`

HasVpcPeerId returns a boolean if a field has been set.

### GetEtherPortChannel

`func (o *NetworkVpcPeer) GetEtherPortChannel() EtherPortChannelRelationship`

GetEtherPortChannel returns the EtherPortChannel field if non-nil, zero value otherwise.

### GetEtherPortChannelOk

`func (o *NetworkVpcPeer) GetEtherPortChannelOk() (*EtherPortChannelRelationship, bool)`

GetEtherPortChannelOk returns a tuple with the EtherPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherPortChannel

`func (o *NetworkVpcPeer) SetEtherPortChannel(v EtherPortChannelRelationship)`

SetEtherPortChannel sets EtherPortChannel field to given value.

### HasEtherPortChannel

`func (o *NetworkVpcPeer) HasEtherPortChannel() bool`

HasEtherPortChannel returns a boolean if a field has been set.

### SetEtherPortChannelNil

`func (o *NetworkVpcPeer) SetEtherPortChannelNil(b bool)`

 SetEtherPortChannelNil sets the value for EtherPortChannel to be an explicit nil

### UnsetEtherPortChannel
`func (o *NetworkVpcPeer) UnsetEtherPortChannel()`

UnsetEtherPortChannel ensures that no value is present for EtherPortChannel, not even an explicit nil
### GetNetworkElement

`func (o *NetworkVpcPeer) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVpcPeer) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVpcPeer) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVpcPeer) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkVpcPeer) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkVpcPeer) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkVpcPeer) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVpcPeer) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVpcPeer) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVpcPeer) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkVpcPeer) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkVpcPeer) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


