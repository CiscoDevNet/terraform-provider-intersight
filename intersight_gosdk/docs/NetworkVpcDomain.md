# NetworkVpcDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.VpcDomain"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.VpcDomain"]
**AutoRecoveryStatus** | Pointer to **string** | Auto recovery status of the virtual port channel domain. | [optional] [readonly] 
**ConsistencyStatus** | Pointer to **string** | Consistency status of the virtual port channel domain. | [optional] [readonly] 
**DualActiveExcludedVlans** | Pointer to **int64** | Dual Active Excluded VLANs of the virtual port channel domain. | [optional] [readonly] 
**KeepAliveStatus** | Pointer to **string** | Keep alive status of the virtual port channel domain. | [optional] [readonly] 
**PeerStatus** | Pointer to **string** | Peer status of the virtual port channel domain. | [optional] [readonly] 
**Role** | Pointer to **string** | Role of the virtual port channel domain. | [optional] [readonly] 
**VpcDomainId** | Pointer to **int64** | Identity of the virtual port channel domain. | [optional] [readonly] 
**VpcsConfiguredCount** | Pointer to **int64** | Number of VPCs configured on the virtual port channel domain. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVpcDomain

`func NewNetworkVpcDomain(classId string, objectType string, ) *NetworkVpcDomain`

NewNetworkVpcDomain instantiates a new NetworkVpcDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVpcDomainWithDefaults

`func NewNetworkVpcDomainWithDefaults() *NetworkVpcDomain`

NewNetworkVpcDomainWithDefaults instantiates a new NetworkVpcDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVpcDomain) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVpcDomain) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVpcDomain) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVpcDomain) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVpcDomain) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVpcDomain) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoRecoveryStatus

`func (o *NetworkVpcDomain) GetAutoRecoveryStatus() string`

GetAutoRecoveryStatus returns the AutoRecoveryStatus field if non-nil, zero value otherwise.

### GetAutoRecoveryStatusOk

`func (o *NetworkVpcDomain) GetAutoRecoveryStatusOk() (*string, bool)`

GetAutoRecoveryStatusOk returns a tuple with the AutoRecoveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoveryStatus

`func (o *NetworkVpcDomain) SetAutoRecoveryStatus(v string)`

SetAutoRecoveryStatus sets AutoRecoveryStatus field to given value.

### HasAutoRecoveryStatus

`func (o *NetworkVpcDomain) HasAutoRecoveryStatus() bool`

HasAutoRecoveryStatus returns a boolean if a field has been set.

### GetConsistencyStatus

`func (o *NetworkVpcDomain) GetConsistencyStatus() string`

GetConsistencyStatus returns the ConsistencyStatus field if non-nil, zero value otherwise.

### GetConsistencyStatusOk

`func (o *NetworkVpcDomain) GetConsistencyStatusOk() (*string, bool)`

GetConsistencyStatusOk returns a tuple with the ConsistencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyStatus

`func (o *NetworkVpcDomain) SetConsistencyStatus(v string)`

SetConsistencyStatus sets ConsistencyStatus field to given value.

### HasConsistencyStatus

`func (o *NetworkVpcDomain) HasConsistencyStatus() bool`

HasConsistencyStatus returns a boolean if a field has been set.

### GetDualActiveExcludedVlans

`func (o *NetworkVpcDomain) GetDualActiveExcludedVlans() int64`

GetDualActiveExcludedVlans returns the DualActiveExcludedVlans field if non-nil, zero value otherwise.

### GetDualActiveExcludedVlansOk

`func (o *NetworkVpcDomain) GetDualActiveExcludedVlansOk() (*int64, bool)`

GetDualActiveExcludedVlansOk returns a tuple with the DualActiveExcludedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualActiveExcludedVlans

`func (o *NetworkVpcDomain) SetDualActiveExcludedVlans(v int64)`

SetDualActiveExcludedVlans sets DualActiveExcludedVlans field to given value.

### HasDualActiveExcludedVlans

`func (o *NetworkVpcDomain) HasDualActiveExcludedVlans() bool`

HasDualActiveExcludedVlans returns a boolean if a field has been set.

### GetKeepAliveStatus

`func (o *NetworkVpcDomain) GetKeepAliveStatus() string`

GetKeepAliveStatus returns the KeepAliveStatus field if non-nil, zero value otherwise.

### GetKeepAliveStatusOk

`func (o *NetworkVpcDomain) GetKeepAliveStatusOk() (*string, bool)`

GetKeepAliveStatusOk returns a tuple with the KeepAliveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAliveStatus

`func (o *NetworkVpcDomain) SetKeepAliveStatus(v string)`

SetKeepAliveStatus sets KeepAliveStatus field to given value.

### HasKeepAliveStatus

`func (o *NetworkVpcDomain) HasKeepAliveStatus() bool`

HasKeepAliveStatus returns a boolean if a field has been set.

### GetPeerStatus

`func (o *NetworkVpcDomain) GetPeerStatus() string`

GetPeerStatus returns the PeerStatus field if non-nil, zero value otherwise.

### GetPeerStatusOk

`func (o *NetworkVpcDomain) GetPeerStatusOk() (*string, bool)`

GetPeerStatusOk returns a tuple with the PeerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerStatus

`func (o *NetworkVpcDomain) SetPeerStatus(v string)`

SetPeerStatus sets PeerStatus field to given value.

### HasPeerStatus

`func (o *NetworkVpcDomain) HasPeerStatus() bool`

HasPeerStatus returns a boolean if a field has been set.

### GetRole

`func (o *NetworkVpcDomain) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NetworkVpcDomain) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NetworkVpcDomain) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NetworkVpcDomain) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetVpcDomainId

`func (o *NetworkVpcDomain) GetVpcDomainId() int64`

GetVpcDomainId returns the VpcDomainId field if non-nil, zero value otherwise.

### GetVpcDomainIdOk

`func (o *NetworkVpcDomain) GetVpcDomainIdOk() (*int64, bool)`

GetVpcDomainIdOk returns a tuple with the VpcDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcDomainId

`func (o *NetworkVpcDomain) SetVpcDomainId(v int64)`

SetVpcDomainId sets VpcDomainId field to given value.

### HasVpcDomainId

`func (o *NetworkVpcDomain) HasVpcDomainId() bool`

HasVpcDomainId returns a boolean if a field has been set.

### GetVpcsConfiguredCount

`func (o *NetworkVpcDomain) GetVpcsConfiguredCount() int64`

GetVpcsConfiguredCount returns the VpcsConfiguredCount field if non-nil, zero value otherwise.

### GetVpcsConfiguredCountOk

`func (o *NetworkVpcDomain) GetVpcsConfiguredCountOk() (*int64, bool)`

GetVpcsConfiguredCountOk returns a tuple with the VpcsConfiguredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcsConfiguredCount

`func (o *NetworkVpcDomain) SetVpcsConfiguredCount(v int64)`

SetVpcsConfiguredCount sets VpcsConfiguredCount field to given value.

### HasVpcsConfiguredCount

`func (o *NetworkVpcDomain) HasVpcsConfiguredCount() bool`

HasVpcsConfiguredCount returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVpcDomain) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVpcDomain) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVpcDomain) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVpcDomain) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkVpcDomain) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkVpcDomain) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkVpcDomain) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVpcDomain) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVpcDomain) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVpcDomain) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkVpcDomain) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkVpcDomain) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


