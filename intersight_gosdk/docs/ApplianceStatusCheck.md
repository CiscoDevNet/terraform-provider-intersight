# ApplianceStatusCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.StatusCheck"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.StatusCheck"]
**Code** | Pointer to **string** | Unique identifier of the status check. | [optional] 
**Result** | Pointer to **string** | Result of this status check. * &#x60;OK&#x60; - Result of the check is OK. * &#x60;Warning&#x60; - Result of the check is Warning. * &#x60;Critical&#x60; - Result of the check is Critical. * &#x60;Info&#x60; - Result of the check is low. | [optional] [default to "OK"]

## Methods

### NewApplianceStatusCheck

`func NewApplianceStatusCheck(classId string, objectType string, ) *ApplianceStatusCheck`

NewApplianceStatusCheck instantiates a new ApplianceStatusCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceStatusCheckWithDefaults

`func NewApplianceStatusCheckWithDefaults() *ApplianceStatusCheck`

NewApplianceStatusCheckWithDefaults instantiates a new ApplianceStatusCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceStatusCheck) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceStatusCheck) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceStatusCheck) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceStatusCheck) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceStatusCheck) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceStatusCheck) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCode

`func (o *ApplianceStatusCheck) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApplianceStatusCheck) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApplianceStatusCheck) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApplianceStatusCheck) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetResult

`func (o *ApplianceStatusCheck) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApplianceStatusCheck) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApplianceStatusCheck) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ApplianceStatusCheck) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


