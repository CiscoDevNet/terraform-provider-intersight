# FabricIpv6FlowKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.Ipv6FlowKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.Ipv6FlowKey"]
**DestinationIpAddress** | Pointer to **bool** | IPv6 Packet destination address. | [optional] [default to false]
**DestinationPort** | Pointer to **bool** | IPv6 Packet destination port field. | [optional] [default to false]
**Protocol** | Pointer to **bool** | IPv6 Packet protocol field. | [optional] [default to false]
**SourceIpAddress** | Pointer to **bool** | IPv6 Packet source address. | [optional] [default to false]
**SourcePort** | Pointer to **bool** | IPv6 Packet source port field. | [optional] [default to false]

## Methods

### NewFabricIpv6FlowKey

`func NewFabricIpv6FlowKey(classId string, objectType string, ) *FabricIpv6FlowKey`

NewFabricIpv6FlowKey instantiates a new FabricIpv6FlowKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricIpv6FlowKeyWithDefaults

`func NewFabricIpv6FlowKeyWithDefaults() *FabricIpv6FlowKey`

NewFabricIpv6FlowKeyWithDefaults instantiates a new FabricIpv6FlowKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricIpv6FlowKey) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricIpv6FlowKey) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricIpv6FlowKey) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricIpv6FlowKey) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricIpv6FlowKey) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricIpv6FlowKey) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationIpAddress

`func (o *FabricIpv6FlowKey) GetDestinationIpAddress() bool`

GetDestinationIpAddress returns the DestinationIpAddress field if non-nil, zero value otherwise.

### GetDestinationIpAddressOk

`func (o *FabricIpv6FlowKey) GetDestinationIpAddressOk() (*bool, bool)`

GetDestinationIpAddressOk returns a tuple with the DestinationIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpAddress

`func (o *FabricIpv6FlowKey) SetDestinationIpAddress(v bool)`

SetDestinationIpAddress sets DestinationIpAddress field to given value.

### HasDestinationIpAddress

`func (o *FabricIpv6FlowKey) HasDestinationIpAddress() bool`

HasDestinationIpAddress returns a boolean if a field has been set.

### GetDestinationPort

`func (o *FabricIpv6FlowKey) GetDestinationPort() bool`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *FabricIpv6FlowKey) GetDestinationPortOk() (*bool, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *FabricIpv6FlowKey) SetDestinationPort(v bool)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *FabricIpv6FlowKey) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetProtocol

`func (o *FabricIpv6FlowKey) GetProtocol() bool`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FabricIpv6FlowKey) GetProtocolOk() (*bool, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FabricIpv6FlowKey) SetProtocol(v bool)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FabricIpv6FlowKey) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceIpAddress

`func (o *FabricIpv6FlowKey) GetSourceIpAddress() bool`

GetSourceIpAddress returns the SourceIpAddress field if non-nil, zero value otherwise.

### GetSourceIpAddressOk

`func (o *FabricIpv6FlowKey) GetSourceIpAddressOk() (*bool, bool)`

GetSourceIpAddressOk returns a tuple with the SourceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddress

`func (o *FabricIpv6FlowKey) SetSourceIpAddress(v bool)`

SetSourceIpAddress sets SourceIpAddress field to given value.

### HasSourceIpAddress

`func (o *FabricIpv6FlowKey) HasSourceIpAddress() bool`

HasSourceIpAddress returns a boolean if a field has been set.

### GetSourcePort

`func (o *FabricIpv6FlowKey) GetSourcePort() bool`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *FabricIpv6FlowKey) GetSourcePortOk() (*bool, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *FabricIpv6FlowKey) SetSourcePort(v bool)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *FabricIpv6FlowKey) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


