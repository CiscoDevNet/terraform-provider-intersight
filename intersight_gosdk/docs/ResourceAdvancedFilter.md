# ResourceAdvancedFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ConditionType** | Pointer to **string** | Filter condition to extract the resource. * &#x60;RANGE&#x60; - Condition matches with range over value. | [optional] [default to "RANGE"]

## Methods

### NewResourceAdvancedFilter

`func NewResourceAdvancedFilter(classId string, objectType string, ) *ResourceAdvancedFilter`

NewResourceAdvancedFilter instantiates a new ResourceAdvancedFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAdvancedFilterWithDefaults

`func NewResourceAdvancedFilterWithDefaults() *ResourceAdvancedFilter`

NewResourceAdvancedFilterWithDefaults instantiates a new ResourceAdvancedFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceAdvancedFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceAdvancedFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceAdvancedFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceAdvancedFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceAdvancedFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceAdvancedFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConditionType

`func (o *ResourceAdvancedFilter) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ResourceAdvancedFilter) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ResourceAdvancedFilter) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *ResourceAdvancedFilter) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


