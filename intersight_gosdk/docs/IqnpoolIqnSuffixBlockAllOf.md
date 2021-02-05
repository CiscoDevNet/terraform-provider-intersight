# IqnpoolIqnSuffixBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.IqnSuffixBlock"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.IqnSuffixBlock"]
**From** | Pointer to **int64** | The first suffix number in the block. | [optional] [readonly] 
**Suffix** | Pointer to **string** | The suffix for this bock of IQNs. | [optional] 
**To** | Pointer to **int64** | The last suffix number in the block. | [optional] [readonly] 

## Methods

### NewIqnpoolIqnSuffixBlockAllOf

`func NewIqnpoolIqnSuffixBlockAllOf(classId string, objectType string, ) *IqnpoolIqnSuffixBlockAllOf`

NewIqnpoolIqnSuffixBlockAllOf instantiates a new IqnpoolIqnSuffixBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolIqnSuffixBlockAllOfWithDefaults

`func NewIqnpoolIqnSuffixBlockAllOfWithDefaults() *IqnpoolIqnSuffixBlockAllOf`

NewIqnpoolIqnSuffixBlockAllOfWithDefaults instantiates a new IqnpoolIqnSuffixBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolIqnSuffixBlockAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolIqnSuffixBlockAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolIqnSuffixBlockAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolIqnSuffixBlockAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolIqnSuffixBlockAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolIqnSuffixBlockAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFrom

`func (o *IqnpoolIqnSuffixBlockAllOf) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IqnpoolIqnSuffixBlockAllOf) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IqnpoolIqnSuffixBlockAllOf) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *IqnpoolIqnSuffixBlockAllOf) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSuffix

`func (o *IqnpoolIqnSuffixBlockAllOf) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *IqnpoolIqnSuffixBlockAllOf) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *IqnpoolIqnSuffixBlockAllOf) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *IqnpoolIqnSuffixBlockAllOf) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetTo

`func (o *IqnpoolIqnSuffixBlockAllOf) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *IqnpoolIqnSuffixBlockAllOf) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *IqnpoolIqnSuffixBlockAllOf) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *IqnpoolIqnSuffixBlockAllOf) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


