# FcNeighborAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fc.Neighbor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fc.Neighbor"]
**PeerDeviceCapability** | Pointer to **string** | This field defines if neighbor is a switch or an NPV device. * &#x60;Switch&#x60; - Switch type neighbors of an interface. * &#x60;NPV&#x60; - N Port Virtualization neighbors of an interface. | [optional] [readonly] [default to "Switch"]
**PeerInterface** | Pointer to **string** | Interface through which the relationship is established. | [optional] [readonly] 
**PeerIpAddress** | Pointer to **string** | IP address of the peer switch. | [optional] [readonly] 
**PeerSwitchName** | Pointer to **string** | Device Id of the neighbor switch. | [optional] [readonly] 
**PeerWwn** | Pointer to **string** | World Wide Name of the neighbor switch. | [optional] [readonly] 
**FcPhysicalPort** | Pointer to [**FcPhysicalPortRelationship**](FcPhysicalPortRelationship.md) |  | [optional] 
**FcPortChannel** | Pointer to [**FcPortChannelRelationship**](FcPortChannelRelationship.md) |  | [optional] 

## Methods

### NewFcNeighborAllOf

`func NewFcNeighborAllOf(classId string, objectType string, ) *FcNeighborAllOf`

NewFcNeighborAllOf instantiates a new FcNeighborAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcNeighborAllOfWithDefaults

`func NewFcNeighborAllOfWithDefaults() *FcNeighborAllOf`

NewFcNeighborAllOfWithDefaults instantiates a new FcNeighborAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcNeighborAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcNeighborAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcNeighborAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcNeighborAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcNeighborAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcNeighborAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPeerDeviceCapability

`func (o *FcNeighborAllOf) GetPeerDeviceCapability() string`

GetPeerDeviceCapability returns the PeerDeviceCapability field if non-nil, zero value otherwise.

### GetPeerDeviceCapabilityOk

`func (o *FcNeighborAllOf) GetPeerDeviceCapabilityOk() (*string, bool)`

GetPeerDeviceCapabilityOk returns a tuple with the PeerDeviceCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDeviceCapability

`func (o *FcNeighborAllOf) SetPeerDeviceCapability(v string)`

SetPeerDeviceCapability sets PeerDeviceCapability field to given value.

### HasPeerDeviceCapability

`func (o *FcNeighborAllOf) HasPeerDeviceCapability() bool`

HasPeerDeviceCapability returns a boolean if a field has been set.

### GetPeerInterface

`func (o *FcNeighborAllOf) GetPeerInterface() string`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *FcNeighborAllOf) GetPeerInterfaceOk() (*string, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *FcNeighborAllOf) SetPeerInterface(v string)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *FcNeighborAllOf) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetPeerIpAddress

`func (o *FcNeighborAllOf) GetPeerIpAddress() string`

GetPeerIpAddress returns the PeerIpAddress field if non-nil, zero value otherwise.

### GetPeerIpAddressOk

`func (o *FcNeighborAllOf) GetPeerIpAddressOk() (*string, bool)`

GetPeerIpAddressOk returns a tuple with the PeerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIpAddress

`func (o *FcNeighborAllOf) SetPeerIpAddress(v string)`

SetPeerIpAddress sets PeerIpAddress field to given value.

### HasPeerIpAddress

`func (o *FcNeighborAllOf) HasPeerIpAddress() bool`

HasPeerIpAddress returns a boolean if a field has been set.

### GetPeerSwitchName

`func (o *FcNeighborAllOf) GetPeerSwitchName() string`

GetPeerSwitchName returns the PeerSwitchName field if non-nil, zero value otherwise.

### GetPeerSwitchNameOk

`func (o *FcNeighborAllOf) GetPeerSwitchNameOk() (*string, bool)`

GetPeerSwitchNameOk returns a tuple with the PeerSwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSwitchName

`func (o *FcNeighborAllOf) SetPeerSwitchName(v string)`

SetPeerSwitchName sets PeerSwitchName field to given value.

### HasPeerSwitchName

`func (o *FcNeighborAllOf) HasPeerSwitchName() bool`

HasPeerSwitchName returns a boolean if a field has been set.

### GetPeerWwn

`func (o *FcNeighborAllOf) GetPeerWwn() string`

GetPeerWwn returns the PeerWwn field if non-nil, zero value otherwise.

### GetPeerWwnOk

`func (o *FcNeighborAllOf) GetPeerWwnOk() (*string, bool)`

GetPeerWwnOk returns a tuple with the PeerWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerWwn

`func (o *FcNeighborAllOf) SetPeerWwn(v string)`

SetPeerWwn sets PeerWwn field to given value.

### HasPeerWwn

`func (o *FcNeighborAllOf) HasPeerWwn() bool`

HasPeerWwn returns a boolean if a field has been set.

### GetFcPhysicalPort

`func (o *FcNeighborAllOf) GetFcPhysicalPort() FcPhysicalPortRelationship`

GetFcPhysicalPort returns the FcPhysicalPort field if non-nil, zero value otherwise.

### GetFcPhysicalPortOk

`func (o *FcNeighborAllOf) GetFcPhysicalPortOk() (*FcPhysicalPortRelationship, bool)`

GetFcPhysicalPortOk returns a tuple with the FcPhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPhysicalPort

`func (o *FcNeighborAllOf) SetFcPhysicalPort(v FcPhysicalPortRelationship)`

SetFcPhysicalPort sets FcPhysicalPort field to given value.

### HasFcPhysicalPort

`func (o *FcNeighborAllOf) HasFcPhysicalPort() bool`

HasFcPhysicalPort returns a boolean if a field has been set.

### GetFcPortChannel

`func (o *FcNeighborAllOf) GetFcPortChannel() FcPortChannelRelationship`

GetFcPortChannel returns the FcPortChannel field if non-nil, zero value otherwise.

### GetFcPortChannelOk

`func (o *FcNeighborAllOf) GetFcPortChannelOk() (*FcPortChannelRelationship, bool)`

GetFcPortChannelOk returns a tuple with the FcPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPortChannel

`func (o *FcNeighborAllOf) SetFcPortChannel(v FcPortChannelRelationship)`

SetFcPortChannel sets FcPortChannel field to given value.

### HasFcPortChannel

`func (o *FcNeighborAllOf) HasFcPortChannel() bool`

HasFcPortChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


