# NiatelemetryVpcData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.VpcData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.VpcData"]
**ConsistentStatus** | Pointer to **bool** | Indicates whether the vPC configuration is consistent between peers. | [optional] [readonly] 
**IntendedPeerName** | Pointer to **string** | Name of the switch intended to be the vPC peer. | [optional] [readonly] 
**KeepAliveStatus** | Pointer to **string** | Operational status of the vPC keepalive connection. | [optional] [readonly] 
**PeerLinkStatus** | Pointer to **string** | Operational status of the link between vPC peer switches. | [optional] [readonly] 
**PeerName** | Pointer to **string** | Name of the switch operating as the vPC peer. | [optional] [readonly] 
**PeerSwitchId** | Pointer to **string** | Identifier (strongAssetId) of the switch operating as the vPC peer. | [optional] [readonly] 
**VpcDomain** | Pointer to **string** | Identifier of the vPC domain configured on the switch. | [optional] [readonly] 
**VpcRole** | Pointer to **string** | Role assigned to the switch within the vPC domain. | [optional] [readonly] 

## Methods

### NewNiatelemetryVpcData

`func NewNiatelemetryVpcData(classId string, objectType string, ) *NiatelemetryVpcData`

NewNiatelemetryVpcData instantiates a new NiatelemetryVpcData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryVpcDataWithDefaults

`func NewNiatelemetryVpcDataWithDefaults() *NiatelemetryVpcData`

NewNiatelemetryVpcDataWithDefaults instantiates a new NiatelemetryVpcData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryVpcData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryVpcData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryVpcData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryVpcData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryVpcData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryVpcData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConsistentStatus

`func (o *NiatelemetryVpcData) GetConsistentStatus() bool`

GetConsistentStatus returns the ConsistentStatus field if non-nil, zero value otherwise.

### GetConsistentStatusOk

`func (o *NiatelemetryVpcData) GetConsistentStatusOk() (*bool, bool)`

GetConsistentStatusOk returns a tuple with the ConsistentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistentStatus

`func (o *NiatelemetryVpcData) SetConsistentStatus(v bool)`

SetConsistentStatus sets ConsistentStatus field to given value.

### HasConsistentStatus

`func (o *NiatelemetryVpcData) HasConsistentStatus() bool`

HasConsistentStatus returns a boolean if a field has been set.

### GetIntendedPeerName

`func (o *NiatelemetryVpcData) GetIntendedPeerName() string`

GetIntendedPeerName returns the IntendedPeerName field if non-nil, zero value otherwise.

### GetIntendedPeerNameOk

`func (o *NiatelemetryVpcData) GetIntendedPeerNameOk() (*string, bool)`

GetIntendedPeerNameOk returns a tuple with the IntendedPeerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPeerName

`func (o *NiatelemetryVpcData) SetIntendedPeerName(v string)`

SetIntendedPeerName sets IntendedPeerName field to given value.

### HasIntendedPeerName

`func (o *NiatelemetryVpcData) HasIntendedPeerName() bool`

HasIntendedPeerName returns a boolean if a field has been set.

### GetKeepAliveStatus

`func (o *NiatelemetryVpcData) GetKeepAliveStatus() string`

GetKeepAliveStatus returns the KeepAliveStatus field if non-nil, zero value otherwise.

### GetKeepAliveStatusOk

`func (o *NiatelemetryVpcData) GetKeepAliveStatusOk() (*string, bool)`

GetKeepAliveStatusOk returns a tuple with the KeepAliveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAliveStatus

`func (o *NiatelemetryVpcData) SetKeepAliveStatus(v string)`

SetKeepAliveStatus sets KeepAliveStatus field to given value.

### HasKeepAliveStatus

`func (o *NiatelemetryVpcData) HasKeepAliveStatus() bool`

HasKeepAliveStatus returns a boolean if a field has been set.

### GetPeerLinkStatus

`func (o *NiatelemetryVpcData) GetPeerLinkStatus() string`

GetPeerLinkStatus returns the PeerLinkStatus field if non-nil, zero value otherwise.

### GetPeerLinkStatusOk

`func (o *NiatelemetryVpcData) GetPeerLinkStatusOk() (*string, bool)`

GetPeerLinkStatusOk returns a tuple with the PeerLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerLinkStatus

`func (o *NiatelemetryVpcData) SetPeerLinkStatus(v string)`

SetPeerLinkStatus sets PeerLinkStatus field to given value.

### HasPeerLinkStatus

`func (o *NiatelemetryVpcData) HasPeerLinkStatus() bool`

HasPeerLinkStatus returns a boolean if a field has been set.

### GetPeerName

`func (o *NiatelemetryVpcData) GetPeerName() string`

GetPeerName returns the PeerName field if non-nil, zero value otherwise.

### GetPeerNameOk

`func (o *NiatelemetryVpcData) GetPeerNameOk() (*string, bool)`

GetPeerNameOk returns a tuple with the PeerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerName

`func (o *NiatelemetryVpcData) SetPeerName(v string)`

SetPeerName sets PeerName field to given value.

### HasPeerName

`func (o *NiatelemetryVpcData) HasPeerName() bool`

HasPeerName returns a boolean if a field has been set.

### GetPeerSwitchId

`func (o *NiatelemetryVpcData) GetPeerSwitchId() string`

GetPeerSwitchId returns the PeerSwitchId field if non-nil, zero value otherwise.

### GetPeerSwitchIdOk

`func (o *NiatelemetryVpcData) GetPeerSwitchIdOk() (*string, bool)`

GetPeerSwitchIdOk returns a tuple with the PeerSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSwitchId

`func (o *NiatelemetryVpcData) SetPeerSwitchId(v string)`

SetPeerSwitchId sets PeerSwitchId field to given value.

### HasPeerSwitchId

`func (o *NiatelemetryVpcData) HasPeerSwitchId() bool`

HasPeerSwitchId returns a boolean if a field has been set.

### GetVpcDomain

`func (o *NiatelemetryVpcData) GetVpcDomain() string`

GetVpcDomain returns the VpcDomain field if non-nil, zero value otherwise.

### GetVpcDomainOk

`func (o *NiatelemetryVpcData) GetVpcDomainOk() (*string, bool)`

GetVpcDomainOk returns a tuple with the VpcDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcDomain

`func (o *NiatelemetryVpcData) SetVpcDomain(v string)`

SetVpcDomain sets VpcDomain field to given value.

### HasVpcDomain

`func (o *NiatelemetryVpcData) HasVpcDomain() bool`

HasVpcDomain returns a boolean if a field has been set.

### GetVpcRole

`func (o *NiatelemetryVpcData) GetVpcRole() string`

GetVpcRole returns the VpcRole field if non-nil, zero value otherwise.

### GetVpcRoleOk

`func (o *NiatelemetryVpcData) GetVpcRoleOk() (*string, bool)`

GetVpcRoleOk returns a tuple with the VpcRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRole

`func (o *NiatelemetryVpcData) SetVpcRole(v string)`

SetVpcRole sets VpcRole field to given value.

### HasVpcRole

`func (o *NiatelemetryVpcData) HasVpcRole() bool`

HasVpcRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


