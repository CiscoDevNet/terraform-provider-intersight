# IppoolIpV4Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.IpV4Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.IpV4Block"]
**From** | Pointer to **string** | First IPv4 address of the block. | [optional] 
**IdMappingPolicyMoid** | Pointer to **string** | The managed object ID of the ID mapping policy. | [optional] 
**IpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**To** | Pointer to **string** | Last IPv4 address of the block. | [optional] 

## Methods

### NewIppoolIpV4Block

`func NewIppoolIpV4Block(classId string, objectType string, ) *IppoolIpV4Block`

NewIppoolIpV4Block instantiates a new IppoolIpV4Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpV4BlockWithDefaults

`func NewIppoolIpV4BlockWithDefaults() *IppoolIpV4Block`

NewIppoolIpV4BlockWithDefaults instantiates a new IppoolIpV4Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolIpV4Block) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolIpV4Block) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolIpV4Block) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolIpV4Block) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolIpV4Block) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolIpV4Block) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFrom

`func (o *IppoolIpV4Block) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IppoolIpV4Block) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IppoolIpV4Block) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *IppoolIpV4Block) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIdMappingPolicyMoid

`func (o *IppoolIpV4Block) GetIdMappingPolicyMoid() string`

GetIdMappingPolicyMoid returns the IdMappingPolicyMoid field if non-nil, zero value otherwise.

### GetIdMappingPolicyMoidOk

`func (o *IppoolIpV4Block) GetIdMappingPolicyMoidOk() (*string, bool)`

GetIdMappingPolicyMoidOk returns a tuple with the IdMappingPolicyMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMappingPolicyMoid

`func (o *IppoolIpV4Block) SetIdMappingPolicyMoid(v string)`

SetIdMappingPolicyMoid sets IdMappingPolicyMoid field to given value.

### HasIdMappingPolicyMoid

`func (o *IppoolIpV4Block) HasIdMappingPolicyMoid() bool`

HasIdMappingPolicyMoid returns a boolean if a field has been set.

### GetIpV4Config

`func (o *IppoolIpV4Block) GetIpV4Config() IppoolIpV4Config`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *IppoolIpV4Block) GetIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *IppoolIpV4Block) SetIpV4Config(v IppoolIpV4Config)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *IppoolIpV4Block) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### SetIpV4ConfigNil

`func (o *IppoolIpV4Block) SetIpV4ConfigNil(b bool)`

 SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil

### UnsetIpV4Config
`func (o *IppoolIpV4Block) UnsetIpV4Config()`

UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
### GetTo

`func (o *IppoolIpV4Block) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *IppoolIpV4Block) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *IppoolIpV4Block) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *IppoolIpV4Block) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


