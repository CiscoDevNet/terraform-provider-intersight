# ResourceRackServerQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.RackServerQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.RackServerQualifier"]
**RackIdRange** | Pointer to [**[]ResourceRackIdRangeFilter**](ResourceRackIdRangeFilter.md) |  | [optional] 

## Methods

### NewResourceRackServerQualifier

`func NewResourceRackServerQualifier(classId string, objectType string, ) *ResourceRackServerQualifier`

NewResourceRackServerQualifier instantiates a new ResourceRackServerQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRackServerQualifierWithDefaults

`func NewResourceRackServerQualifierWithDefaults() *ResourceRackServerQualifier`

NewResourceRackServerQualifierWithDefaults instantiates a new ResourceRackServerQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceRackServerQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceRackServerQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceRackServerQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceRackServerQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceRackServerQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceRackServerQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRackIdRange

`func (o *ResourceRackServerQualifier) GetRackIdRange() []ResourceRackIdRangeFilter`

GetRackIdRange returns the RackIdRange field if non-nil, zero value otherwise.

### GetRackIdRangeOk

`func (o *ResourceRackServerQualifier) GetRackIdRangeOk() (*[]ResourceRackIdRangeFilter, bool)`

GetRackIdRangeOk returns a tuple with the RackIdRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackIdRange

`func (o *ResourceRackServerQualifier) SetRackIdRange(v []ResourceRackIdRangeFilter)`

SetRackIdRange sets RackIdRange field to given value.

### HasRackIdRange

`func (o *ResourceRackServerQualifier) HasRackIdRange() bool`

HasRackIdRange returns a boolean if a field has been set.

### SetRackIdRangeNil

`func (o *ResourceRackServerQualifier) SetRackIdRangeNil(b bool)`

 SetRackIdRangeNil sets the value for RackIdRange to be an explicit nil

### UnsetRackIdRange
`func (o *ResourceRackServerQualifier) UnsetRackIdRange()`

UnsetRackIdRange ensures that no value is present for RackIdRange, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


