# TamEolSeverityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.EolSeverity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.EolSeverity"]
**Level** | Pointer to **string** | Severity level associated with an end-of-life (EOL) milestone advisory. * &#x60;info&#x60; - The end-of-life (EOL) milestone is at info level. * &#x60;critical&#x60; - The end-of-life (EOL) milestone is at critical level. It usually hints &#39;red&#39; in a color-map. * &#x60;high&#x60; - The end-of-life (EOL) milestone is at high level. * &#x60;medium&#x60; - The end-of-life (EOL) milestone is at medium level. | [optional] [default to "info"]

## Methods

### NewTamEolSeverityAllOf

`func NewTamEolSeverityAllOf(classId string, objectType string, ) *TamEolSeverityAllOf`

NewTamEolSeverityAllOf instantiates a new TamEolSeverityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamEolSeverityAllOfWithDefaults

`func NewTamEolSeverityAllOfWithDefaults() *TamEolSeverityAllOf`

NewTamEolSeverityAllOfWithDefaults instantiates a new TamEolSeverityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamEolSeverityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamEolSeverityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamEolSeverityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamEolSeverityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamEolSeverityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamEolSeverityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLevel

`func (o *TamEolSeverityAllOf) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TamEolSeverityAllOf) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TamEolSeverityAllOf) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *TamEolSeverityAllOf) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


