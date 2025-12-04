# SdaaciConnectionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdaaci.ConnectionDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdaaci.ConnectionDetail"]
**Description** | Pointer to **string** | Description of this connection between two peers. | [optional] 
**IpPool** | Pointer to **string** | Id of the ip pool configured for this connection. | [optional] 
**Layer3HandoffId** | Pointer to **string** | Id of layer 3 handoff configured between a border node and a border leaf. | [optional] 
**PeerAinterface** | Pointer to **string** | Interface id configured on Peer A. | [optional] 
**PeerAipAddress** | Pointer to **string** | The IP Address of the device used as the local peer. | [optional] 
**PeerAtype** | Pointer to **string** | Type of device used as Peer A for this peer connection. | [optional] 
**PeerBinterface** | Pointer to **string** | Interface id configured on Peer B. | [optional] 
**PeerBipAddress** | Pointer to **string** | The IP Address of the device used as the remote peer. | [optional] 
**PeerBtype** | Pointer to **string** | Type of device used as Peer B for this peer connection. | [optional] 
**Peera** | Pointer to **string** | First peer of the connection. | [optional] 
**Peerb** | Pointer to **string** | Second Peer of the connection. | [optional] 
**RouterId** | Pointer to **string** | Router id defined for this peer connection. | [optional] 
**Status** | Pointer to **string** | Connection status between the peers. * &#x60;NotConnected&#x60; - Connection Status NotConnected. * &#x60;Connected&#x60; - Connection Status Connected. | [optional] [default to "NotConnected"]
**Connection** | Pointer to [**NullableSdaaciConnectionRelationship**](SdaaciConnectionRelationship.md) |  | [optional] 

## Methods

### NewSdaaciConnectionDetail

`func NewSdaaciConnectionDetail(classId string, objectType string, ) *SdaaciConnectionDetail`

NewSdaaciConnectionDetail instantiates a new SdaaciConnectionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdaaciConnectionDetailWithDefaults

`func NewSdaaciConnectionDetailWithDefaults() *SdaaciConnectionDetail`

NewSdaaciConnectionDetailWithDefaults instantiates a new SdaaciConnectionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdaaciConnectionDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdaaciConnectionDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdaaciConnectionDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdaaciConnectionDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdaaciConnectionDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdaaciConnectionDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *SdaaciConnectionDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdaaciConnectionDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdaaciConnectionDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdaaciConnectionDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpPool

`func (o *SdaaciConnectionDetail) GetIpPool() string`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *SdaaciConnectionDetail) GetIpPoolOk() (*string, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *SdaaciConnectionDetail) SetIpPool(v string)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *SdaaciConnectionDetail) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetLayer3HandoffId

`func (o *SdaaciConnectionDetail) GetLayer3HandoffId() string`

GetLayer3HandoffId returns the Layer3HandoffId field if non-nil, zero value otherwise.

### GetLayer3HandoffIdOk

`func (o *SdaaciConnectionDetail) GetLayer3HandoffIdOk() (*string, bool)`

GetLayer3HandoffIdOk returns a tuple with the Layer3HandoffId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer3HandoffId

`func (o *SdaaciConnectionDetail) SetLayer3HandoffId(v string)`

SetLayer3HandoffId sets Layer3HandoffId field to given value.

### HasLayer3HandoffId

`func (o *SdaaciConnectionDetail) HasLayer3HandoffId() bool`

HasLayer3HandoffId returns a boolean if a field has been set.

### GetPeerAinterface

`func (o *SdaaciConnectionDetail) GetPeerAinterface() string`

GetPeerAinterface returns the PeerAinterface field if non-nil, zero value otherwise.

### GetPeerAinterfaceOk

`func (o *SdaaciConnectionDetail) GetPeerAinterfaceOk() (*string, bool)`

GetPeerAinterfaceOk returns a tuple with the PeerAinterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAinterface

`func (o *SdaaciConnectionDetail) SetPeerAinterface(v string)`

SetPeerAinterface sets PeerAinterface field to given value.

### HasPeerAinterface

`func (o *SdaaciConnectionDetail) HasPeerAinterface() bool`

HasPeerAinterface returns a boolean if a field has been set.

### GetPeerAipAddress

`func (o *SdaaciConnectionDetail) GetPeerAipAddress() string`

GetPeerAipAddress returns the PeerAipAddress field if non-nil, zero value otherwise.

### GetPeerAipAddressOk

`func (o *SdaaciConnectionDetail) GetPeerAipAddressOk() (*string, bool)`

GetPeerAipAddressOk returns a tuple with the PeerAipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAipAddress

`func (o *SdaaciConnectionDetail) SetPeerAipAddress(v string)`

SetPeerAipAddress sets PeerAipAddress field to given value.

### HasPeerAipAddress

`func (o *SdaaciConnectionDetail) HasPeerAipAddress() bool`

HasPeerAipAddress returns a boolean if a field has been set.

### GetPeerAtype

`func (o *SdaaciConnectionDetail) GetPeerAtype() string`

GetPeerAtype returns the PeerAtype field if non-nil, zero value otherwise.

### GetPeerAtypeOk

`func (o *SdaaciConnectionDetail) GetPeerAtypeOk() (*string, bool)`

GetPeerAtypeOk returns a tuple with the PeerAtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAtype

`func (o *SdaaciConnectionDetail) SetPeerAtype(v string)`

SetPeerAtype sets PeerAtype field to given value.

### HasPeerAtype

`func (o *SdaaciConnectionDetail) HasPeerAtype() bool`

HasPeerAtype returns a boolean if a field has been set.

### GetPeerBinterface

`func (o *SdaaciConnectionDetail) GetPeerBinterface() string`

GetPeerBinterface returns the PeerBinterface field if non-nil, zero value otherwise.

### GetPeerBinterfaceOk

`func (o *SdaaciConnectionDetail) GetPeerBinterfaceOk() (*string, bool)`

GetPeerBinterfaceOk returns a tuple with the PeerBinterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerBinterface

`func (o *SdaaciConnectionDetail) SetPeerBinterface(v string)`

SetPeerBinterface sets PeerBinterface field to given value.

### HasPeerBinterface

`func (o *SdaaciConnectionDetail) HasPeerBinterface() bool`

HasPeerBinterface returns a boolean if a field has been set.

### GetPeerBipAddress

`func (o *SdaaciConnectionDetail) GetPeerBipAddress() string`

GetPeerBipAddress returns the PeerBipAddress field if non-nil, zero value otherwise.

### GetPeerBipAddressOk

`func (o *SdaaciConnectionDetail) GetPeerBipAddressOk() (*string, bool)`

GetPeerBipAddressOk returns a tuple with the PeerBipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerBipAddress

`func (o *SdaaciConnectionDetail) SetPeerBipAddress(v string)`

SetPeerBipAddress sets PeerBipAddress field to given value.

### HasPeerBipAddress

`func (o *SdaaciConnectionDetail) HasPeerBipAddress() bool`

HasPeerBipAddress returns a boolean if a field has been set.

### GetPeerBtype

`func (o *SdaaciConnectionDetail) GetPeerBtype() string`

GetPeerBtype returns the PeerBtype field if non-nil, zero value otherwise.

### GetPeerBtypeOk

`func (o *SdaaciConnectionDetail) GetPeerBtypeOk() (*string, bool)`

GetPeerBtypeOk returns a tuple with the PeerBtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerBtype

`func (o *SdaaciConnectionDetail) SetPeerBtype(v string)`

SetPeerBtype sets PeerBtype field to given value.

### HasPeerBtype

`func (o *SdaaciConnectionDetail) HasPeerBtype() bool`

HasPeerBtype returns a boolean if a field has been set.

### GetPeera

`func (o *SdaaciConnectionDetail) GetPeera() string`

GetPeera returns the Peera field if non-nil, zero value otherwise.

### GetPeeraOk

`func (o *SdaaciConnectionDetail) GetPeeraOk() (*string, bool)`

GetPeeraOk returns a tuple with the Peera field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeera

`func (o *SdaaciConnectionDetail) SetPeera(v string)`

SetPeera sets Peera field to given value.

### HasPeera

`func (o *SdaaciConnectionDetail) HasPeera() bool`

HasPeera returns a boolean if a field has been set.

### GetPeerb

`func (o *SdaaciConnectionDetail) GetPeerb() string`

GetPeerb returns the Peerb field if non-nil, zero value otherwise.

### GetPeerbOk

`func (o *SdaaciConnectionDetail) GetPeerbOk() (*string, bool)`

GetPeerbOk returns a tuple with the Peerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerb

`func (o *SdaaciConnectionDetail) SetPeerb(v string)`

SetPeerb sets Peerb field to given value.

### HasPeerb

`func (o *SdaaciConnectionDetail) HasPeerb() bool`

HasPeerb returns a boolean if a field has been set.

### GetRouterId

`func (o *SdaaciConnectionDetail) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *SdaaciConnectionDetail) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *SdaaciConnectionDetail) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *SdaaciConnectionDetail) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetStatus

`func (o *SdaaciConnectionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdaaciConnectionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdaaciConnectionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdaaciConnectionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnection

`func (o *SdaaciConnectionDetail) GetConnection() SdaaciConnectionRelationship`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *SdaaciConnectionDetail) GetConnectionOk() (*SdaaciConnectionRelationship, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *SdaaciConnectionDetail) SetConnection(v SdaaciConnectionRelationship)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *SdaaciConnectionDetail) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### SetConnectionNil

`func (o *SdaaciConnectionDetail) SetConnectionNil(b bool)`

 SetConnectionNil sets the value for Connection to be an explicit nil

### UnsetConnection
`func (o *SdaaciConnectionDetail) UnsetConnection()`

UnsetConnection ensures that no value is present for Connection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


