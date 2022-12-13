# NetworkVpcPeerAllOf

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
**EtherPortChannel** | Pointer to [**EtherPortChannelRelationship**](EtherPortChannelRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVpcPeerAllOf

`func NewNetworkVpcPeerAllOf(classId string, objectType string, ) *NetworkVpcPeerAllOf`

NewNetworkVpcPeerAllOf instantiates a new NetworkVpcPeerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVpcPeerAllOfWithDefaults

`func NewNetworkVpcPeerAllOfWithDefaults() *NetworkVpcPeerAllOf`

NewNetworkVpcPeerAllOfWithDefaults instantiates a new NetworkVpcPeerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVpcPeerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVpcPeerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVpcPeerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVpcPeerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVpcPeerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVpcPeerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationalState

`func (o *NetworkVpcPeerAllOf) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *NetworkVpcPeerAllOf) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *NetworkVpcPeerAllOf) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *NetworkVpcPeerAllOf) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetPortChannel

`func (o *NetworkVpcPeerAllOf) GetPortChannel() string`

GetPortChannel returns the PortChannel field if non-nil, zero value otherwise.

### GetPortChannelOk

`func (o *NetworkVpcPeerAllOf) GetPortChannelOk() (*string, bool)`

GetPortChannelOk returns a tuple with the PortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannel

`func (o *NetworkVpcPeerAllOf) SetPortChannel(v string)`

SetPortChannel sets PortChannel field to given value.

### HasPortChannel

`func (o *NetworkVpcPeerAllOf) HasPortChannel() bool`

HasPortChannel returns a boolean if a field has been set.

### GetPortChannelId

`func (o *NetworkVpcPeerAllOf) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *NetworkVpcPeerAllOf) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *NetworkVpcPeerAllOf) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *NetworkVpcPeerAllOf) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetVpcDomainId

`func (o *NetworkVpcPeerAllOf) GetVpcDomainId() int64`

GetVpcDomainId returns the VpcDomainId field if non-nil, zero value otherwise.

### GetVpcDomainIdOk

`func (o *NetworkVpcPeerAllOf) GetVpcDomainIdOk() (*int64, bool)`

GetVpcDomainIdOk returns a tuple with the VpcDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcDomainId

`func (o *NetworkVpcPeerAllOf) SetVpcDomainId(v int64)`

SetVpcDomainId sets VpcDomainId field to given value.

### HasVpcDomainId

`func (o *NetworkVpcPeerAllOf) HasVpcDomainId() bool`

HasVpcDomainId returns a boolean if a field has been set.

### GetVpcPeerId

`func (o *NetworkVpcPeerAllOf) GetVpcPeerId() int64`

GetVpcPeerId returns the VpcPeerId field if non-nil, zero value otherwise.

### GetVpcPeerIdOk

`func (o *NetworkVpcPeerAllOf) GetVpcPeerIdOk() (*int64, bool)`

GetVpcPeerIdOk returns a tuple with the VpcPeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeerId

`func (o *NetworkVpcPeerAllOf) SetVpcPeerId(v int64)`

SetVpcPeerId sets VpcPeerId field to given value.

### HasVpcPeerId

`func (o *NetworkVpcPeerAllOf) HasVpcPeerId() bool`

HasVpcPeerId returns a boolean if a field has been set.

### GetEtherPortChannel

`func (o *NetworkVpcPeerAllOf) GetEtherPortChannel() EtherPortChannelRelationship`

GetEtherPortChannel returns the EtherPortChannel field if non-nil, zero value otherwise.

### GetEtherPortChannelOk

`func (o *NetworkVpcPeerAllOf) GetEtherPortChannelOk() (*EtherPortChannelRelationship, bool)`

GetEtherPortChannelOk returns a tuple with the EtherPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherPortChannel

`func (o *NetworkVpcPeerAllOf) SetEtherPortChannel(v EtherPortChannelRelationship)`

SetEtherPortChannel sets EtherPortChannel field to given value.

### HasEtherPortChannel

`func (o *NetworkVpcPeerAllOf) HasEtherPortChannel() bool`

HasEtherPortChannel returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVpcPeerAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVpcPeerAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVpcPeerAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVpcPeerAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVpcPeerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVpcPeerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVpcPeerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVpcPeerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


