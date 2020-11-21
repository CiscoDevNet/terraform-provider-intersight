# UuidpoolUuidBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "uuidpool.UuidBlock"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "uuidpool.UuidBlock"]
**From** | Pointer to **string** | Starting UUID suffix of the block must be in hexadecimal format xxxx-xxxxxxxxxxxx. | [optional] 
**To** | Pointer to **string** | Starting UUID suffix of the block must be in hexadecimal format xxxx-xxxxxxxxxxxx. | [optional] 

## Methods

### NewUuidpoolUuidBlockAllOf

`func NewUuidpoolUuidBlockAllOf(classId string, objectType string, ) *UuidpoolUuidBlockAllOf`

NewUuidpoolUuidBlockAllOf instantiates a new UuidpoolUuidBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolUuidBlockAllOfWithDefaults

`func NewUuidpoolUuidBlockAllOfWithDefaults() *UuidpoolUuidBlockAllOf`

NewUuidpoolUuidBlockAllOfWithDefaults instantiates a new UuidpoolUuidBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UuidpoolUuidBlockAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UuidpoolUuidBlockAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UuidpoolUuidBlockAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UuidpoolUuidBlockAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UuidpoolUuidBlockAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UuidpoolUuidBlockAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFrom

`func (o *UuidpoolUuidBlockAllOf) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *UuidpoolUuidBlockAllOf) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *UuidpoolUuidBlockAllOf) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *UuidpoolUuidBlockAllOf) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *UuidpoolUuidBlockAllOf) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UuidpoolUuidBlockAllOf) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UuidpoolUuidBlockAllOf) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *UuidpoolUuidBlockAllOf) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


