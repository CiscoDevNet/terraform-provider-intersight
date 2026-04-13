# FabricL2FlowKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.L2FlowKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.L2FlowKey"]
**DestinationMac** | Pointer to **bool** | L2 Packet destination address. | [optional] [default to false]
**EtherType** | Pointer to **bool** | L2 Packet ether type field. | [optional] [default to false]
**SourceMac** | Pointer to **bool** | L2 Packet source address. | [optional] [default to false]

## Methods

### NewFabricL2FlowKey

`func NewFabricL2FlowKey(classId string, objectType string, ) *FabricL2FlowKey`

NewFabricL2FlowKey instantiates a new FabricL2FlowKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricL2FlowKeyWithDefaults

`func NewFabricL2FlowKeyWithDefaults() *FabricL2FlowKey`

NewFabricL2FlowKeyWithDefaults instantiates a new FabricL2FlowKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricL2FlowKey) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricL2FlowKey) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricL2FlowKey) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricL2FlowKey) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricL2FlowKey) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricL2FlowKey) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationMac

`func (o *FabricL2FlowKey) GetDestinationMac() bool`

GetDestinationMac returns the DestinationMac field if non-nil, zero value otherwise.

### GetDestinationMacOk

`func (o *FabricL2FlowKey) GetDestinationMacOk() (*bool, bool)`

GetDestinationMacOk returns a tuple with the DestinationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMac

`func (o *FabricL2FlowKey) SetDestinationMac(v bool)`

SetDestinationMac sets DestinationMac field to given value.

### HasDestinationMac

`func (o *FabricL2FlowKey) HasDestinationMac() bool`

HasDestinationMac returns a boolean if a field has been set.

### GetEtherType

`func (o *FabricL2FlowKey) GetEtherType() bool`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *FabricL2FlowKey) GetEtherTypeOk() (*bool, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *FabricL2FlowKey) SetEtherType(v bool)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *FabricL2FlowKey) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetSourceMac

`func (o *FabricL2FlowKey) GetSourceMac() bool`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *FabricL2FlowKey) GetSourceMacOk() (*bool, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *FabricL2FlowKey) SetSourceMac(v bool)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *FabricL2FlowKey) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


