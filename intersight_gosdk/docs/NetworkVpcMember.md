# NetworkVpcMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.VpcMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.VpcMember"]
**OperationalState** | Pointer to **string** | Operational state of the virtual port channel. | [optional] [readonly] 
**PortChannel** | Pointer to **string** | Name of the virtual port channel. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel identity of the virtual port channel. | [optional] [readonly] 
**VpcDomainId** | Pointer to **int64** | Identity of the virtual port channel. | [optional] [readonly] 
**VpcMemberId** | Pointer to **int64** | Identity of the virtual port channel. | [optional] [readonly] 
**EtherPortChannel** | Pointer to [**EtherPortChannelRelationship**](EtherPortChannelRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVpcMember

`func NewNetworkVpcMember(classId string, objectType string, ) *NetworkVpcMember`

NewNetworkVpcMember instantiates a new NetworkVpcMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVpcMemberWithDefaults

`func NewNetworkVpcMemberWithDefaults() *NetworkVpcMember`

NewNetworkVpcMemberWithDefaults instantiates a new NetworkVpcMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVpcMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVpcMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVpcMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVpcMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVpcMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVpcMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationalState

`func (o *NetworkVpcMember) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *NetworkVpcMember) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *NetworkVpcMember) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *NetworkVpcMember) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetPortChannel

`func (o *NetworkVpcMember) GetPortChannel() string`

GetPortChannel returns the PortChannel field if non-nil, zero value otherwise.

### GetPortChannelOk

`func (o *NetworkVpcMember) GetPortChannelOk() (*string, bool)`

GetPortChannelOk returns a tuple with the PortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannel

`func (o *NetworkVpcMember) SetPortChannel(v string)`

SetPortChannel sets PortChannel field to given value.

### HasPortChannel

`func (o *NetworkVpcMember) HasPortChannel() bool`

HasPortChannel returns a boolean if a field has been set.

### GetPortChannelId

`func (o *NetworkVpcMember) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *NetworkVpcMember) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *NetworkVpcMember) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *NetworkVpcMember) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetVpcDomainId

`func (o *NetworkVpcMember) GetVpcDomainId() int64`

GetVpcDomainId returns the VpcDomainId field if non-nil, zero value otherwise.

### GetVpcDomainIdOk

`func (o *NetworkVpcMember) GetVpcDomainIdOk() (*int64, bool)`

GetVpcDomainIdOk returns a tuple with the VpcDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcDomainId

`func (o *NetworkVpcMember) SetVpcDomainId(v int64)`

SetVpcDomainId sets VpcDomainId field to given value.

### HasVpcDomainId

`func (o *NetworkVpcMember) HasVpcDomainId() bool`

HasVpcDomainId returns a boolean if a field has been set.

### GetVpcMemberId

`func (o *NetworkVpcMember) GetVpcMemberId() int64`

GetVpcMemberId returns the VpcMemberId field if non-nil, zero value otherwise.

### GetVpcMemberIdOk

`func (o *NetworkVpcMember) GetVpcMemberIdOk() (*int64, bool)`

GetVpcMemberIdOk returns a tuple with the VpcMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcMemberId

`func (o *NetworkVpcMember) SetVpcMemberId(v int64)`

SetVpcMemberId sets VpcMemberId field to given value.

### HasVpcMemberId

`func (o *NetworkVpcMember) HasVpcMemberId() bool`

HasVpcMemberId returns a boolean if a field has been set.

### GetEtherPortChannel

`func (o *NetworkVpcMember) GetEtherPortChannel() EtherPortChannelRelationship`

GetEtherPortChannel returns the EtherPortChannel field if non-nil, zero value otherwise.

### GetEtherPortChannelOk

`func (o *NetworkVpcMember) GetEtherPortChannelOk() (*EtherPortChannelRelationship, bool)`

GetEtherPortChannelOk returns a tuple with the EtherPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherPortChannel

`func (o *NetworkVpcMember) SetEtherPortChannel(v EtherPortChannelRelationship)`

SetEtherPortChannel sets EtherPortChannel field to given value.

### HasEtherPortChannel

`func (o *NetworkVpcMember) HasEtherPortChannel() bool`

HasEtherPortChannel returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVpcMember) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVpcMember) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVpcMember) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVpcMember) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVpcMember) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVpcMember) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVpcMember) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVpcMember) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


