# FabricIpv4FlowKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.Ipv4FlowKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.Ipv4FlowKey"]
**DestinationIpAddress** | Pointer to **bool** | IPv4 Packet destination address. | [optional] [default to false]
**DestinationPort** | Pointer to **bool** | IPv4 Packet destination port field. | [optional] [default to false]
**Protocol** | Pointer to **bool** | IPv4 packet protocol field. | [optional] [default to false]
**SourceIpAddress** | Pointer to **bool** | IPv4 Packet source address. | [optional] [default to false]
**SourcePort** | Pointer to **bool** | IPv4 Packet source port field. | [optional] [default to false]
**Tos** | Pointer to **bool** | IPv4 Packet type of service. | [optional] [default to false]

## Methods

### NewFabricIpv4FlowKey

`func NewFabricIpv4FlowKey(classId string, objectType string, ) *FabricIpv4FlowKey`

NewFabricIpv4FlowKey instantiates a new FabricIpv4FlowKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricIpv4FlowKeyWithDefaults

`func NewFabricIpv4FlowKeyWithDefaults() *FabricIpv4FlowKey`

NewFabricIpv4FlowKeyWithDefaults instantiates a new FabricIpv4FlowKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricIpv4FlowKey) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricIpv4FlowKey) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricIpv4FlowKey) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricIpv4FlowKey) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricIpv4FlowKey) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricIpv4FlowKey) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationIpAddress

`func (o *FabricIpv4FlowKey) GetDestinationIpAddress() bool`

GetDestinationIpAddress returns the DestinationIpAddress field if non-nil, zero value otherwise.

### GetDestinationIpAddressOk

`func (o *FabricIpv4FlowKey) GetDestinationIpAddressOk() (*bool, bool)`

GetDestinationIpAddressOk returns a tuple with the DestinationIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpAddress

`func (o *FabricIpv4FlowKey) SetDestinationIpAddress(v bool)`

SetDestinationIpAddress sets DestinationIpAddress field to given value.

### HasDestinationIpAddress

`func (o *FabricIpv4FlowKey) HasDestinationIpAddress() bool`

HasDestinationIpAddress returns a boolean if a field has been set.

### GetDestinationPort

`func (o *FabricIpv4FlowKey) GetDestinationPort() bool`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *FabricIpv4FlowKey) GetDestinationPortOk() (*bool, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *FabricIpv4FlowKey) SetDestinationPort(v bool)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *FabricIpv4FlowKey) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetProtocol

`func (o *FabricIpv4FlowKey) GetProtocol() bool`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FabricIpv4FlowKey) GetProtocolOk() (*bool, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FabricIpv4FlowKey) SetProtocol(v bool)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FabricIpv4FlowKey) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceIpAddress

`func (o *FabricIpv4FlowKey) GetSourceIpAddress() bool`

GetSourceIpAddress returns the SourceIpAddress field if non-nil, zero value otherwise.

### GetSourceIpAddressOk

`func (o *FabricIpv4FlowKey) GetSourceIpAddressOk() (*bool, bool)`

GetSourceIpAddressOk returns a tuple with the SourceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddress

`func (o *FabricIpv4FlowKey) SetSourceIpAddress(v bool)`

SetSourceIpAddress sets SourceIpAddress field to given value.

### HasSourceIpAddress

`func (o *FabricIpv4FlowKey) HasSourceIpAddress() bool`

HasSourceIpAddress returns a boolean if a field has been set.

### GetSourcePort

`func (o *FabricIpv4FlowKey) GetSourcePort() bool`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *FabricIpv4FlowKey) GetSourcePortOk() (*bool, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *FabricIpv4FlowKey) SetSourcePort(v bool)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *FabricIpv4FlowKey) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetTos

`func (o *FabricIpv4FlowKey) GetTos() bool`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *FabricIpv4FlowKey) GetTosOk() (*bool, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *FabricIpv4FlowKey) SetTos(v bool)`

SetTos sets Tos field to given value.

### HasTos

`func (o *FabricIpv4FlowKey) HasTos() bool`

HasTos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


