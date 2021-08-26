# PortInterfaceBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**OperState** | Pointer to **string** | Operational state of an Interface. | [optional] 
**AcknowledgedPeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](EtherPhysicalPortBaseRelationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](EtherPhysicalPortBaseRelationship.md) |  | [optional] 

## Methods

### NewPortInterfaceBase

`func NewPortInterfaceBase(classId string, objectType string, ) *PortInterfaceBase`

NewPortInterfaceBase instantiates a new PortInterfaceBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortInterfaceBaseWithDefaults

`func NewPortInterfaceBaseWithDefaults() *PortInterfaceBase`

NewPortInterfaceBaseWithDefaults instantiates a new PortInterfaceBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PortInterfaceBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PortInterfaceBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PortInterfaceBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PortInterfaceBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PortInterfaceBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PortInterfaceBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperState

`func (o *PortInterfaceBase) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PortInterfaceBase) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PortInterfaceBase) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PortInterfaceBase) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *PortInterfaceBase) GetAcknowledgedPeerInterface() EtherPhysicalPortBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *PortInterfaceBase) GetAcknowledgedPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *PortInterfaceBase) SetAcknowledgedPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *PortInterfaceBase) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *PortInterfaceBase) GetPeerInterface() EtherPhysicalPortBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *PortInterfaceBase) GetPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *PortInterfaceBase) SetPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *PortInterfaceBase) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


