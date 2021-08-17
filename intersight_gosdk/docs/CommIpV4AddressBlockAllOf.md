# CommIpV4AddressBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "comm.IpV4AddressBlock"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "comm.IpV4AddressBlock"]
**EndAddress** | Pointer to **string** | The end address of the IPv4 block. | [optional] 
**StartAddress** | Pointer to **string** | The start address of the IPv4 block. | [optional] 

## Methods

### NewCommIpV4AddressBlockAllOf

`func NewCommIpV4AddressBlockAllOf(classId string, objectType string, ) *CommIpV4AddressBlockAllOf`

NewCommIpV4AddressBlockAllOf instantiates a new CommIpV4AddressBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommIpV4AddressBlockAllOfWithDefaults

`func NewCommIpV4AddressBlockAllOfWithDefaults() *CommIpV4AddressBlockAllOf`

NewCommIpV4AddressBlockAllOfWithDefaults instantiates a new CommIpV4AddressBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommIpV4AddressBlockAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommIpV4AddressBlockAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommIpV4AddressBlockAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommIpV4AddressBlockAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommIpV4AddressBlockAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommIpV4AddressBlockAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndAddress

`func (o *CommIpV4AddressBlockAllOf) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *CommIpV4AddressBlockAllOf) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *CommIpV4AddressBlockAllOf) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *CommIpV4AddressBlockAllOf) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetStartAddress

`func (o *CommIpV4AddressBlockAllOf) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *CommIpV4AddressBlockAllOf) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *CommIpV4AddressBlockAllOf) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *CommIpV4AddressBlockAllOf) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


