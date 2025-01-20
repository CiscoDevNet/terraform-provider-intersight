# HciServiceViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ServiceViolation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ServiceViolation"]
**StartDate** | Pointer to **time.Time** | The start date of the violation. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the service violation. The possible values are: &#39;LICENSE_EXPIRED&#39;, &#39;FEATURE_VIOLATION&#39;, &#39;CAPACITY_VIOLATION&#39;. | [optional] [readonly] 

## Methods

### NewHciServiceViolation

`func NewHciServiceViolation(classId string, objectType string, ) *HciServiceViolation`

NewHciServiceViolation instantiates a new HciServiceViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciServiceViolationWithDefaults

`func NewHciServiceViolationWithDefaults() *HciServiceViolation`

NewHciServiceViolationWithDefaults instantiates a new HciServiceViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciServiceViolation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciServiceViolation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciServiceViolation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciServiceViolation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciServiceViolation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciServiceViolation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStartDate

`func (o *HciServiceViolation) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *HciServiceViolation) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *HciServiceViolation) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *HciServiceViolation) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetType

`func (o *HciServiceViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciServiceViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciServiceViolation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciServiceViolation) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


