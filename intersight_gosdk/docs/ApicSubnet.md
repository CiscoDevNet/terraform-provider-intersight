# ApicSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.Subnet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.Subnet"]
**Dn** | Pointer to **string** | Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Ip** | Pointer to **string** | IP of an object within the Cisco Application Policy Infrastructure Controller. | [optional] 
**Name** | Pointer to **string** | Name of an object within the Cisco Application Policy Infrastructure Controller. | [optional] 
**BridgeDomain** | Pointer to [**NullableApicBridgeDomainRelationship**](ApicBridgeDomainRelationship.md) |  | [optional] 

## Methods

### NewApicSubnet

`func NewApicSubnet(classId string, objectType string, ) *ApicSubnet`

NewApicSubnet instantiates a new ApicSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicSubnetWithDefaults

`func NewApicSubnetWithDefaults() *ApicSubnet`

NewApicSubnetWithDefaults instantiates a new ApicSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicSubnet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicSubnet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicSubnet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicSubnet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicSubnet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicSubnet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *ApicSubnet) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicSubnet) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicSubnet) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicSubnet) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetIp

`func (o *ApicSubnet) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApicSubnet) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApicSubnet) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApicSubnet) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *ApicSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBridgeDomain

`func (o *ApicSubnet) GetBridgeDomain() ApicBridgeDomainRelationship`

GetBridgeDomain returns the BridgeDomain field if non-nil, zero value otherwise.

### GetBridgeDomainOk

`func (o *ApicSubnet) GetBridgeDomainOk() (*ApicBridgeDomainRelationship, bool)`

GetBridgeDomainOk returns a tuple with the BridgeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeDomain

`func (o *ApicSubnet) SetBridgeDomain(v ApicBridgeDomainRelationship)`

SetBridgeDomain sets BridgeDomain field to given value.

### HasBridgeDomain

`func (o *ApicSubnet) HasBridgeDomain() bool`

HasBridgeDomain returns a boolean if a field has been set.

### SetBridgeDomainNil

`func (o *ApicSubnet) SetBridgeDomainNil(b bool)`

 SetBridgeDomainNil sets the value for BridgeDomain to be an explicit nil

### UnsetBridgeDomain
`func (o *ApicSubnet) UnsetBridgeDomain()`

UnsetBridgeDomain ensures that no value is present for BridgeDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


