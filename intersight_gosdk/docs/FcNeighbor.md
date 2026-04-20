# FcNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fc.Neighbor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fc.Neighbor"]
**PeerDeviceCapability** | Pointer to **string** | This field defines if neighbor is a switch, storage or an NPV device. * &#x60;Switch&#x60; - Switch type neighbors of an interface. * &#x60;NPV&#x60; - N Port Virtualization neighbors of an interface. * &#x60;Storage&#x60; - Storage type neighbors of an interface. | [optional] [readonly] [default to "Switch"]
**PeerInterface** | Pointer to **string** | Interface through which the relationship is established. | [optional] [readonly] 
**PeerIpAddress** | Pointer to **string** | IP address of the peer switch. | [optional] [readonly] 
**PeerPortWwn** | Pointer to **string** | World Wide Name of the neighbor port. | [optional] [readonly] 
**PeerSwitchName** | Pointer to **string** | Device Id of the neighbor switch. | [optional] [readonly] 
**PeerWwn** | Pointer to **string** | World Wide Name of the neighbor switch. | [optional] [readonly] 
**Vendor** | Pointer to **string** | Vendor name for the neighboring storage device. Available only for Storage neighbors. | [optional] [readonly] 
**Vsan** | Pointer to **int64** | VSAN associated with this neighbor port. | [optional] [readonly] 
**FcPhysicalPort** | Pointer to [**NullableFcPhysicalPortRelationship**](FcPhysicalPortRelationship.md) |  | [optional] 
**FcPortChannel** | Pointer to [**NullableFcPortChannelRelationship**](FcPortChannelRelationship.md) |  | [optional] 

## Methods

### NewFcNeighbor

`func NewFcNeighbor(classId string, objectType string, ) *FcNeighbor`

NewFcNeighbor instantiates a new FcNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcNeighborWithDefaults

`func NewFcNeighborWithDefaults() *FcNeighbor`

NewFcNeighborWithDefaults instantiates a new FcNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcNeighbor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcNeighbor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcNeighbor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcNeighbor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcNeighbor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcNeighbor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPeerDeviceCapability

`func (o *FcNeighbor) GetPeerDeviceCapability() string`

GetPeerDeviceCapability returns the PeerDeviceCapability field if non-nil, zero value otherwise.

### GetPeerDeviceCapabilityOk

`func (o *FcNeighbor) GetPeerDeviceCapabilityOk() (*string, bool)`

GetPeerDeviceCapabilityOk returns a tuple with the PeerDeviceCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDeviceCapability

`func (o *FcNeighbor) SetPeerDeviceCapability(v string)`

SetPeerDeviceCapability sets PeerDeviceCapability field to given value.

### HasPeerDeviceCapability

`func (o *FcNeighbor) HasPeerDeviceCapability() bool`

HasPeerDeviceCapability returns a boolean if a field has been set.

### GetPeerInterface

`func (o *FcNeighbor) GetPeerInterface() string`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *FcNeighbor) GetPeerInterfaceOk() (*string, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *FcNeighbor) SetPeerInterface(v string)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *FcNeighbor) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetPeerIpAddress

`func (o *FcNeighbor) GetPeerIpAddress() string`

GetPeerIpAddress returns the PeerIpAddress field if non-nil, zero value otherwise.

### GetPeerIpAddressOk

`func (o *FcNeighbor) GetPeerIpAddressOk() (*string, bool)`

GetPeerIpAddressOk returns a tuple with the PeerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIpAddress

`func (o *FcNeighbor) SetPeerIpAddress(v string)`

SetPeerIpAddress sets PeerIpAddress field to given value.

### HasPeerIpAddress

`func (o *FcNeighbor) HasPeerIpAddress() bool`

HasPeerIpAddress returns a boolean if a field has been set.

### GetPeerPortWwn

`func (o *FcNeighbor) GetPeerPortWwn() string`

GetPeerPortWwn returns the PeerPortWwn field if non-nil, zero value otherwise.

### GetPeerPortWwnOk

`func (o *FcNeighbor) GetPeerPortWwnOk() (*string, bool)`

GetPeerPortWwnOk returns a tuple with the PeerPortWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerPortWwn

`func (o *FcNeighbor) SetPeerPortWwn(v string)`

SetPeerPortWwn sets PeerPortWwn field to given value.

### HasPeerPortWwn

`func (o *FcNeighbor) HasPeerPortWwn() bool`

HasPeerPortWwn returns a boolean if a field has been set.

### GetPeerSwitchName

`func (o *FcNeighbor) GetPeerSwitchName() string`

GetPeerSwitchName returns the PeerSwitchName field if non-nil, zero value otherwise.

### GetPeerSwitchNameOk

`func (o *FcNeighbor) GetPeerSwitchNameOk() (*string, bool)`

GetPeerSwitchNameOk returns a tuple with the PeerSwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSwitchName

`func (o *FcNeighbor) SetPeerSwitchName(v string)`

SetPeerSwitchName sets PeerSwitchName field to given value.

### HasPeerSwitchName

`func (o *FcNeighbor) HasPeerSwitchName() bool`

HasPeerSwitchName returns a boolean if a field has been set.

### GetPeerWwn

`func (o *FcNeighbor) GetPeerWwn() string`

GetPeerWwn returns the PeerWwn field if non-nil, zero value otherwise.

### GetPeerWwnOk

`func (o *FcNeighbor) GetPeerWwnOk() (*string, bool)`

GetPeerWwnOk returns a tuple with the PeerWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerWwn

`func (o *FcNeighbor) SetPeerWwn(v string)`

SetPeerWwn sets PeerWwn field to given value.

### HasPeerWwn

`func (o *FcNeighbor) HasPeerWwn() bool`

HasPeerWwn returns a boolean if a field has been set.

### GetVendor

`func (o *FcNeighbor) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FcNeighbor) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FcNeighbor) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *FcNeighbor) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVsan

`func (o *FcNeighbor) GetVsan() int64`

GetVsan returns the Vsan field if non-nil, zero value otherwise.

### GetVsanOk

`func (o *FcNeighbor) GetVsanOk() (*int64, bool)`

GetVsanOk returns a tuple with the Vsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsan

`func (o *FcNeighbor) SetVsan(v int64)`

SetVsan sets Vsan field to given value.

### HasVsan

`func (o *FcNeighbor) HasVsan() bool`

HasVsan returns a boolean if a field has been set.

### GetFcPhysicalPort

`func (o *FcNeighbor) GetFcPhysicalPort() FcPhysicalPortRelationship`

GetFcPhysicalPort returns the FcPhysicalPort field if non-nil, zero value otherwise.

### GetFcPhysicalPortOk

`func (o *FcNeighbor) GetFcPhysicalPortOk() (*FcPhysicalPortRelationship, bool)`

GetFcPhysicalPortOk returns a tuple with the FcPhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPhysicalPort

`func (o *FcNeighbor) SetFcPhysicalPort(v FcPhysicalPortRelationship)`

SetFcPhysicalPort sets FcPhysicalPort field to given value.

### HasFcPhysicalPort

`func (o *FcNeighbor) HasFcPhysicalPort() bool`

HasFcPhysicalPort returns a boolean if a field has been set.

### SetFcPhysicalPortNil

`func (o *FcNeighbor) SetFcPhysicalPortNil(b bool)`

 SetFcPhysicalPortNil sets the value for FcPhysicalPort to be an explicit nil

### UnsetFcPhysicalPort
`func (o *FcNeighbor) UnsetFcPhysicalPort()`

UnsetFcPhysicalPort ensures that no value is present for FcPhysicalPort, not even an explicit nil
### GetFcPortChannel

`func (o *FcNeighbor) GetFcPortChannel() FcPortChannelRelationship`

GetFcPortChannel returns the FcPortChannel field if non-nil, zero value otherwise.

### GetFcPortChannelOk

`func (o *FcNeighbor) GetFcPortChannelOk() (*FcPortChannelRelationship, bool)`

GetFcPortChannelOk returns a tuple with the FcPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPortChannel

`func (o *FcNeighbor) SetFcPortChannel(v FcPortChannelRelationship)`

SetFcPortChannel sets FcPortChannel field to given value.

### HasFcPortChannel

`func (o *FcNeighbor) HasFcPortChannel() bool`

HasFcPortChannel returns a boolean if a field has been set.

### SetFcPortChannelNil

`func (o *FcNeighbor) SetFcPortChannelNil(b bool)`

 SetFcPortChannelNil sets the value for FcPortChannel to be an explicit nil

### UnsetFcPortChannel
`func (o *FcNeighbor) UnsetFcPortChannel()`

UnsetFcPortChannel ensures that no value is present for FcPortChannel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


