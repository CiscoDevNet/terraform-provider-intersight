# IppoolIpV4Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.IpV4Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.IpV4Block"]
**From** | Pointer to **string** | First IPv4 address of the block. | [optional] [readonly] 
**To** | Pointer to **string** | Last IPv4 address of the block. | [optional] [readonly] 

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


