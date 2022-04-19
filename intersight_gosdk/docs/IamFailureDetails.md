# IamFailureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.FailureDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.FailureDetails"]
**FailureReason** | Pointer to **string** | Reason for the failure during verification. | [optional] [readonly] 
**FailureTime** | Pointer to **time.Time** | Timestamp of the failure during verification. | [optional] [readonly] 

## Methods

### NewIamFailureDetails

`func NewIamFailureDetails(classId string, objectType string, ) *IamFailureDetails`

NewIamFailureDetails instantiates a new IamFailureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamFailureDetailsWithDefaults

`func NewIamFailureDetailsWithDefaults() *IamFailureDetails`

NewIamFailureDetailsWithDefaults instantiates a new IamFailureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamFailureDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamFailureDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamFailureDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamFailureDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamFailureDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamFailureDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailureReason

`func (o *IamFailureDetails) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *IamFailureDetails) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *IamFailureDetails) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *IamFailureDetails) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFailureTime

`func (o *IamFailureDetails) GetFailureTime() time.Time`

GetFailureTime returns the FailureTime field if non-nil, zero value otherwise.

### GetFailureTimeOk

`func (o *IamFailureDetails) GetFailureTimeOk() (*time.Time, bool)`

GetFailureTimeOk returns a tuple with the FailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureTime

`func (o *IamFailureDetails) SetFailureTime(v time.Time)`

SetFailureTime sets FailureTime field to given value.

### HasFailureTime

`func (o *IamFailureDetails) HasFailureTime() bool`

HasFailureTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


