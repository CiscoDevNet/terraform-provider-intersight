# IppoolShadowBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.ShadowBlock"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.ShadowBlock"]
**IpType** | Pointer to **string** | Type of this IP addresses blocks. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [default to "IPv4"]
**IpV4Block** | Pointer to [**IppoolIpV4Block**](IppoolIpV4Block.md) |  | [optional] 
**IpV6Block** | Pointer to [**IppoolIpV6Block**](IppoolIpV6Block.md) |  | [optional] 
**Pool** | Pointer to [**IppoolShadowPoolRelationship**](IppoolShadowPoolRelationship.md) |  | [optional] 

## Methods

### NewIppoolShadowBlockAllOf

`func NewIppoolShadowBlockAllOf(classId string, objectType string, ) *IppoolShadowBlockAllOf`

NewIppoolShadowBlockAllOf instantiates a new IppoolShadowBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolShadowBlockAllOfWithDefaults

`func NewIppoolShadowBlockAllOfWithDefaults() *IppoolShadowBlockAllOf`

NewIppoolShadowBlockAllOfWithDefaults instantiates a new IppoolShadowBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolShadowBlockAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolShadowBlockAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolShadowBlockAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolShadowBlockAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolShadowBlockAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolShadowBlockAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolShadowBlockAllOf) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolShadowBlockAllOf) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolShadowBlockAllOf) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolShadowBlockAllOf) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpV4Block

`func (o *IppoolShadowBlockAllOf) GetIpV4Block() IppoolIpV4Block`

GetIpV4Block returns the IpV4Block field if non-nil, zero value otherwise.

### GetIpV4BlockOk

`func (o *IppoolShadowBlockAllOf) GetIpV4BlockOk() (*IppoolIpV4Block, bool)`

GetIpV4BlockOk returns a tuple with the IpV4Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Block

`func (o *IppoolShadowBlockAllOf) SetIpV4Block(v IppoolIpV4Block)`

SetIpV4Block sets IpV4Block field to given value.

### HasIpV4Block

`func (o *IppoolShadowBlockAllOf) HasIpV4Block() bool`

HasIpV4Block returns a boolean if a field has been set.

### GetIpV6Block

`func (o *IppoolShadowBlockAllOf) GetIpV6Block() IppoolIpV6Block`

GetIpV6Block returns the IpV6Block field if non-nil, zero value otherwise.

### GetIpV6BlockOk

`func (o *IppoolShadowBlockAllOf) GetIpV6BlockOk() (*IppoolIpV6Block, bool)`

GetIpV6BlockOk returns a tuple with the IpV6Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Block

`func (o *IppoolShadowBlockAllOf) SetIpV6Block(v IppoolIpV6Block)`

SetIpV6Block sets IpV6Block field to given value.

### HasIpV6Block

`func (o *IppoolShadowBlockAllOf) HasIpV6Block() bool`

HasIpV6Block returns a boolean if a field has been set.

### GetPool

`func (o *IppoolShadowBlockAllOf) GetPool() IppoolShadowPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolShadowBlockAllOf) GetPoolOk() (*IppoolShadowPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolShadowBlockAllOf) SetPool(v IppoolShadowPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolShadowBlockAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


