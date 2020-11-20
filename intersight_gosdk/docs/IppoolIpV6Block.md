# IppoolIpV6Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.IpV6Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.IpV6Block"]
**From** | Pointer to **string** | First IPv6 address of the block. | [optional] [readonly] 
**To** | Pointer to **string** | Last IPv6 address of the block. | [optional] [readonly] 

## Methods

### NewIppoolIpV6Block

`func NewIppoolIpV6Block(classId string, objectType string, ) *IppoolIpV6Block`

NewIppoolIpV6Block instantiates a new IppoolIpV6Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpV6BlockWithDefaults

`func NewIppoolIpV6BlockWithDefaults() *IppoolIpV6Block`

NewIppoolIpV6BlockWithDefaults instantiates a new IppoolIpV6Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolIpV6Block) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolIpV6Block) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolIpV6Block) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolIpV6Block) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolIpV6Block) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolIpV6Block) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFrom

`func (o *IppoolIpV6Block) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IppoolIpV6Block) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IppoolIpV6Block) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *IppoolIpV6Block) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *IppoolIpV6Block) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *IppoolIpV6Block) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *IppoolIpV6Block) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *IppoolIpV6Block) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


