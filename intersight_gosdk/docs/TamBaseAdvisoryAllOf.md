# TamBaseAdvisoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Description** | Pointer to **string** | Brief description of the advisory details. | [optional] 
**Name** | Pointer to **string** | A user defined name for the Intersight Advisory. | [optional] 
**Severity** | Pointer to [**NullableTamSeverity**](TamSeverity.md) |  | [optional] 
**State** | Pointer to **string** | Current state of the advisory. * &#x60;ready&#x60; - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created. * &#x60;evaluating&#x60; - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation. | [optional] [default to "ready"]

## Methods

### NewTamBaseAdvisoryAllOf

`func NewTamBaseAdvisoryAllOf(classId string, objectType string, ) *TamBaseAdvisoryAllOf`

NewTamBaseAdvisoryAllOf instantiates a new TamBaseAdvisoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamBaseAdvisoryAllOfWithDefaults

`func NewTamBaseAdvisoryAllOfWithDefaults() *TamBaseAdvisoryAllOf`

NewTamBaseAdvisoryAllOfWithDefaults instantiates a new TamBaseAdvisoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamBaseAdvisoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamBaseAdvisoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamBaseAdvisoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamBaseAdvisoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamBaseAdvisoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamBaseAdvisoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *TamBaseAdvisoryAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TamBaseAdvisoryAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TamBaseAdvisoryAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TamBaseAdvisoryAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *TamBaseAdvisoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamBaseAdvisoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamBaseAdvisoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamBaseAdvisoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *TamBaseAdvisoryAllOf) GetSeverity() TamSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TamBaseAdvisoryAllOf) GetSeverityOk() (*TamSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TamBaseAdvisoryAllOf) SetSeverity(v TamSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TamBaseAdvisoryAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *TamBaseAdvisoryAllOf) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *TamBaseAdvisoryAllOf) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetState

`func (o *TamBaseAdvisoryAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TamBaseAdvisoryAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TamBaseAdvisoryAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TamBaseAdvisoryAllOf) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


