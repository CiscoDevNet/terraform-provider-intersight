# FabricSpanSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SpanSession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SpanSession"]
**DestPorts** | Pointer to [**[]FabricAbstractSpanDestPortRelationship**](FabricAbstractSpanDestPortRelationship.md) | An array of relationships to fabricAbstractSpanDestPort resources. | [optional] 
**SourcePortChannels** | Pointer to [**[]FabricAbstractSpanSourcePortChannelRelationship**](FabricAbstractSpanSourcePortChannelRelationship.md) | An array of relationships to fabricAbstractSpanSourcePortChannel resources. | [optional] 
**SourcePorts** | Pointer to [**[]FabricAbstractSpanSourcePortRelationship**](FabricAbstractSpanSourcePortRelationship.md) | An array of relationships to fabricAbstractSpanSourcePort resources. | [optional] 
**SourceVirtualIfs** | Pointer to [**[]FabricAbstractSpanSourceVirtualIfRelationship**](FabricAbstractSpanSourceVirtualIfRelationship.md) | An array of relationships to fabricAbstractSpanSourceVirtualIf resources. | [optional] 
**SourceVlans** | Pointer to [**[]FabricSpanSourceVlanRelationship**](FabricSpanSourceVlanRelationship.md) | An array of relationships to fabricSpanSourceVlan resources. | [optional] 

## Methods

### NewFabricSpanSession

`func NewFabricSpanSession(classId string, objectType string, ) *FabricSpanSession`

NewFabricSpanSession instantiates a new FabricSpanSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSpanSessionWithDefaults

`func NewFabricSpanSessionWithDefaults() *FabricSpanSession`

NewFabricSpanSessionWithDefaults instantiates a new FabricSpanSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSpanSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSpanSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSpanSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSpanSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSpanSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSpanSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestPorts

`func (o *FabricSpanSession) GetDestPorts() []FabricAbstractSpanDestPortRelationship`

GetDestPorts returns the DestPorts field if non-nil, zero value otherwise.

### GetDestPortsOk

`func (o *FabricSpanSession) GetDestPortsOk() (*[]FabricAbstractSpanDestPortRelationship, bool)`

GetDestPortsOk returns a tuple with the DestPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPorts

`func (o *FabricSpanSession) SetDestPorts(v []FabricAbstractSpanDestPortRelationship)`

SetDestPorts sets DestPorts field to given value.

### HasDestPorts

`func (o *FabricSpanSession) HasDestPorts() bool`

HasDestPorts returns a boolean if a field has been set.

### SetDestPortsNil

`func (o *FabricSpanSession) SetDestPortsNil(b bool)`

 SetDestPortsNil sets the value for DestPorts to be an explicit nil

### UnsetDestPorts
`func (o *FabricSpanSession) UnsetDestPorts()`

UnsetDestPorts ensures that no value is present for DestPorts, not even an explicit nil
### GetSourcePortChannels

`func (o *FabricSpanSession) GetSourcePortChannels() []FabricAbstractSpanSourcePortChannelRelationship`

GetSourcePortChannels returns the SourcePortChannels field if non-nil, zero value otherwise.

### GetSourcePortChannelsOk

`func (o *FabricSpanSession) GetSourcePortChannelsOk() (*[]FabricAbstractSpanSourcePortChannelRelationship, bool)`

GetSourcePortChannelsOk returns a tuple with the SourcePortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortChannels

`func (o *FabricSpanSession) SetSourcePortChannels(v []FabricAbstractSpanSourcePortChannelRelationship)`

SetSourcePortChannels sets SourcePortChannels field to given value.

### HasSourcePortChannels

`func (o *FabricSpanSession) HasSourcePortChannels() bool`

HasSourcePortChannels returns a boolean if a field has been set.

### SetSourcePortChannelsNil

`func (o *FabricSpanSession) SetSourcePortChannelsNil(b bool)`

 SetSourcePortChannelsNil sets the value for SourcePortChannels to be an explicit nil

### UnsetSourcePortChannels
`func (o *FabricSpanSession) UnsetSourcePortChannels()`

UnsetSourcePortChannels ensures that no value is present for SourcePortChannels, not even an explicit nil
### GetSourcePorts

`func (o *FabricSpanSession) GetSourcePorts() []FabricAbstractSpanSourcePortRelationship`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *FabricSpanSession) GetSourcePortsOk() (*[]FabricAbstractSpanSourcePortRelationship, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *FabricSpanSession) SetSourcePorts(v []FabricAbstractSpanSourcePortRelationship)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *FabricSpanSession) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.

### SetSourcePortsNil

`func (o *FabricSpanSession) SetSourcePortsNil(b bool)`

 SetSourcePortsNil sets the value for SourcePorts to be an explicit nil

### UnsetSourcePorts
`func (o *FabricSpanSession) UnsetSourcePorts()`

UnsetSourcePorts ensures that no value is present for SourcePorts, not even an explicit nil
### GetSourceVirtualIfs

`func (o *FabricSpanSession) GetSourceVirtualIfs() []FabricAbstractSpanSourceVirtualIfRelationship`

GetSourceVirtualIfs returns the SourceVirtualIfs field if non-nil, zero value otherwise.

### GetSourceVirtualIfsOk

`func (o *FabricSpanSession) GetSourceVirtualIfsOk() (*[]FabricAbstractSpanSourceVirtualIfRelationship, bool)`

GetSourceVirtualIfsOk returns a tuple with the SourceVirtualIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVirtualIfs

`func (o *FabricSpanSession) SetSourceVirtualIfs(v []FabricAbstractSpanSourceVirtualIfRelationship)`

SetSourceVirtualIfs sets SourceVirtualIfs field to given value.

### HasSourceVirtualIfs

`func (o *FabricSpanSession) HasSourceVirtualIfs() bool`

HasSourceVirtualIfs returns a boolean if a field has been set.

### SetSourceVirtualIfsNil

`func (o *FabricSpanSession) SetSourceVirtualIfsNil(b bool)`

 SetSourceVirtualIfsNil sets the value for SourceVirtualIfs to be an explicit nil

### UnsetSourceVirtualIfs
`func (o *FabricSpanSession) UnsetSourceVirtualIfs()`

UnsetSourceVirtualIfs ensures that no value is present for SourceVirtualIfs, not even an explicit nil
### GetSourceVlans

`func (o *FabricSpanSession) GetSourceVlans() []FabricSpanSourceVlanRelationship`

GetSourceVlans returns the SourceVlans field if non-nil, zero value otherwise.

### GetSourceVlansOk

`func (o *FabricSpanSession) GetSourceVlansOk() (*[]FabricSpanSourceVlanRelationship, bool)`

GetSourceVlansOk returns a tuple with the SourceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVlans

`func (o *FabricSpanSession) SetSourceVlans(v []FabricSpanSourceVlanRelationship)`

SetSourceVlans sets SourceVlans field to given value.

### HasSourceVlans

`func (o *FabricSpanSession) HasSourceVlans() bool`

HasSourceVlans returns a boolean if a field has been set.

### SetSourceVlansNil

`func (o *FabricSpanSession) SetSourceVlansNil(b bool)`

 SetSourceVlansNil sets the value for SourceVlans to be an explicit nil

### UnsetSourceVlans
`func (o *FabricSpanSession) UnsetSourceVlans()`

UnsetSourceVlans ensures that no value is present for SourceVlans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


