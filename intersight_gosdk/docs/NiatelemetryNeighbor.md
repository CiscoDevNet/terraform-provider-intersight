# NiatelemetryNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Neighbor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Neighbor"]
**Asn** | Pointer to **int64** | Local autonomous system number for the neighbor. | [optional] [readonly] 
**FabricName** | Pointer to **string** | Fabric name for the neighbor. | [optional] [readonly] 
**InterfaceIp** | Pointer to **string** | Interface IP address for the neighbor. | [optional] [readonly] 
**LocalIp** | Pointer to **string** | Local IP address for the neighbor. | [optional] [readonly] 
**OperationalStatus** | Pointer to **string** | Operational status of the neighbor. | [optional] [readonly] 
**OspfArea** | Pointer to **string** | OSPF area for the neighbor. | [optional] [readonly] 
**PeerAsn** | Pointer to **int64** | Peer autonomous system number for the neighbor. | [optional] [readonly] 
**PeerIp** | Pointer to **string** | Peer IP address for the neighbor. | [optional] [readonly] 
**SwitchName** | Pointer to **string** | Switch name for the neighbor. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of Layer 3 neighbor object. | [optional] [readonly] 
**VrfName** | Pointer to **string** | VRF name for the neighbor. | [optional] [readonly] 
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNeighbor

`func NewNiatelemetryNeighbor(classId string, objectType string, ) *NiatelemetryNeighbor`

NewNiatelemetryNeighbor instantiates a new NiatelemetryNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNeighborWithDefaults

`func NewNiatelemetryNeighborWithDefaults() *NiatelemetryNeighbor`

NewNiatelemetryNeighborWithDefaults instantiates a new NiatelemetryNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNeighbor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNeighbor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNeighbor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNeighbor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNeighbor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNeighbor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAsn

`func (o *NiatelemetryNeighbor) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *NiatelemetryNeighbor) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *NiatelemetryNeighbor) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *NiatelemetryNeighbor) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryNeighbor) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNeighbor) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNeighbor) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNeighbor) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *NiatelemetryNeighbor) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *NiatelemetryNeighbor) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *NiatelemetryNeighbor) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *NiatelemetryNeighbor) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetLocalIp

`func (o *NiatelemetryNeighbor) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *NiatelemetryNeighbor) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *NiatelemetryNeighbor) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *NiatelemetryNeighbor) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *NiatelemetryNeighbor) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *NiatelemetryNeighbor) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *NiatelemetryNeighbor) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *NiatelemetryNeighbor) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetOspfArea

`func (o *NiatelemetryNeighbor) GetOspfArea() string`

GetOspfArea returns the OspfArea field if non-nil, zero value otherwise.

### GetOspfAreaOk

`func (o *NiatelemetryNeighbor) GetOspfAreaOk() (*string, bool)`

GetOspfAreaOk returns a tuple with the OspfArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfArea

`func (o *NiatelemetryNeighbor) SetOspfArea(v string)`

SetOspfArea sets OspfArea field to given value.

### HasOspfArea

`func (o *NiatelemetryNeighbor) HasOspfArea() bool`

HasOspfArea returns a boolean if a field has been set.

### GetPeerAsn

`func (o *NiatelemetryNeighbor) GetPeerAsn() int64`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *NiatelemetryNeighbor) GetPeerAsnOk() (*int64, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *NiatelemetryNeighbor) SetPeerAsn(v int64)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *NiatelemetryNeighbor) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetPeerIp

`func (o *NiatelemetryNeighbor) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *NiatelemetryNeighbor) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *NiatelemetryNeighbor) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.

### HasPeerIp

`func (o *NiatelemetryNeighbor) HasPeerIp() bool`

HasPeerIp returns a boolean if a field has been set.

### GetSwitchName

`func (o *NiatelemetryNeighbor) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *NiatelemetryNeighbor) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *NiatelemetryNeighbor) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *NiatelemetryNeighbor) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetryNeighbor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetryNeighbor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetryNeighbor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetryNeighbor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVrfName

`func (o *NiatelemetryNeighbor) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *NiatelemetryNeighbor) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *NiatelemetryNeighbor) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *NiatelemetryNeighbor) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetryNeighbor) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetryNeighbor) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetryNeighbor) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetryNeighbor) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetryNeighbor) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetryNeighbor) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


