# TamFnSeverity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.FnSeverity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.FnSeverity"]
**Level** | Pointer to **string** | Severity level associated with a field notice. * &#x60;na&#x60; - &lt; For field notices, where Severity or Impact rating is not applicable. * &#x60;critical&#x60; - &lt; If applicable, customers may experience downtime or might be at risk of affecting services. Urgent attention and response action are desirable. * &#x60;high&#x60; - &lt; If applicable, it may degrade or might have a potential degradation, Attention and response based on deployment scenarios. * &#x60;medium&#x60; - &lt; If applicable, it may affect system performance or functionality but are less like to pose an immediate risk to the network. Actions to be taken in due course of time based on deployment scenarios. * &#x60;low&#x60; - &lt; If applicable, Unlikely to expect disrupted functionality. Minor in nature and can be addressed as per convenience. | [optional] [default to "na"]

## Methods

### NewTamFnSeverity

`func NewTamFnSeverity(classId string, objectType string, ) *TamFnSeverity`

NewTamFnSeverity instantiates a new TamFnSeverity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamFnSeverityWithDefaults

`func NewTamFnSeverityWithDefaults() *TamFnSeverity`

NewTamFnSeverityWithDefaults instantiates a new TamFnSeverity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamFnSeverity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamFnSeverity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamFnSeverity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamFnSeverity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamFnSeverity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamFnSeverity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLevel

`func (o *TamFnSeverity) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TamFnSeverity) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TamFnSeverity) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *TamFnSeverity) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


