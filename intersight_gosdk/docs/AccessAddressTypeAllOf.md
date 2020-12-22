# AccessAddressTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "access.AddressType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "access.AddressType"]
**EnableIpV4** | Pointer to **bool** | This flag enables the use of IPv4 address for end-point access. | [optional] [default to true]
**EnableIpV6** | Pointer to **bool** | This flag enables the use of IPv6 address for end-point access. | [optional] [default to false]

## Methods

### NewAccessAddressTypeAllOf

`func NewAccessAddressTypeAllOf(classId string, objectType string, ) *AccessAddressTypeAllOf`

NewAccessAddressTypeAllOf instantiates a new AccessAddressTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessAddressTypeAllOfWithDefaults

`func NewAccessAddressTypeAllOfWithDefaults() *AccessAddressTypeAllOf`

NewAccessAddressTypeAllOfWithDefaults instantiates a new AccessAddressTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AccessAddressTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AccessAddressTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AccessAddressTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AccessAddressTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessAddressTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessAddressTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableIpV4

`func (o *AccessAddressTypeAllOf) GetEnableIpV4() bool`

GetEnableIpV4 returns the EnableIpV4 field if non-nil, zero value otherwise.

### GetEnableIpV4Ok

`func (o *AccessAddressTypeAllOf) GetEnableIpV4Ok() (*bool, bool)`

GetEnableIpV4Ok returns a tuple with the EnableIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpV4

`func (o *AccessAddressTypeAllOf) SetEnableIpV4(v bool)`

SetEnableIpV4 sets EnableIpV4 field to given value.

### HasEnableIpV4

`func (o *AccessAddressTypeAllOf) HasEnableIpV4() bool`

HasEnableIpV4 returns a boolean if a field has been set.

### GetEnableIpV6

`func (o *AccessAddressTypeAllOf) GetEnableIpV6() bool`

GetEnableIpV6 returns the EnableIpV6 field if non-nil, zero value otherwise.

### GetEnableIpV6Ok

`func (o *AccessAddressTypeAllOf) GetEnableIpV6Ok() (*bool, bool)`

GetEnableIpV6Ok returns a tuple with the EnableIpV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpV6

`func (o *AccessAddressTypeAllOf) SetEnableIpV6(v bool)`

SetEnableIpV6 sets EnableIpV6 field to given value.

### HasEnableIpV6

`func (o *AccessAddressTypeAllOf) HasEnableIpV6() bool`

HasEnableIpV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


