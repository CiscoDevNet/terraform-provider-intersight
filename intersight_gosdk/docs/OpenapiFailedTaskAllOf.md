# OpenapiFailedTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.FailedTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.FailedTask"]
**Name** | Pointer to **string** | Name of the task. | [optional] 
**Path** | Pointer to **string** | REST API path of the task. | [optional] 
**Reason** | Pointer to **string** | Indicates the reason for task creation failure. | [optional] 

## Methods

### NewOpenapiFailedTaskAllOf

`func NewOpenapiFailedTaskAllOf(classId string, objectType string, ) *OpenapiFailedTaskAllOf`

NewOpenapiFailedTaskAllOf instantiates a new OpenapiFailedTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiFailedTaskAllOfWithDefaults

`func NewOpenapiFailedTaskAllOfWithDefaults() *OpenapiFailedTaskAllOf`

NewOpenapiFailedTaskAllOfWithDefaults instantiates a new OpenapiFailedTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiFailedTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiFailedTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiFailedTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiFailedTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiFailedTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiFailedTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *OpenapiFailedTaskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenapiFailedTaskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenapiFailedTaskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenapiFailedTaskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *OpenapiFailedTaskAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OpenapiFailedTaskAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OpenapiFailedTaskAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OpenapiFailedTaskAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReason

`func (o *OpenapiFailedTaskAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OpenapiFailedTaskAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OpenapiFailedTaskAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OpenapiFailedTaskAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


