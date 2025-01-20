# HciFeatureViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.FeatureViolation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.FeatureViolation"]
**AffectedEntity** | Pointer to **string** | The description of the affected entity of this feature violation. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the feature that has violation. | [optional] [readonly] 
**FeatureId** | Pointer to **string** | The id of the feature that has violation. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the feature that has violation. | [optional] [readonly] 

## Methods

### NewHciFeatureViolation

`func NewHciFeatureViolation(classId string, objectType string, ) *HciFeatureViolation`

NewHciFeatureViolation instantiates a new HciFeatureViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciFeatureViolationWithDefaults

`func NewHciFeatureViolationWithDefaults() *HciFeatureViolation`

NewHciFeatureViolationWithDefaults instantiates a new HciFeatureViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciFeatureViolation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciFeatureViolation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciFeatureViolation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciFeatureViolation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciFeatureViolation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciFeatureViolation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffectedEntity

`func (o *HciFeatureViolation) GetAffectedEntity() string`

GetAffectedEntity returns the AffectedEntity field if non-nil, zero value otherwise.

### GetAffectedEntityOk

`func (o *HciFeatureViolation) GetAffectedEntityOk() (*string, bool)`

GetAffectedEntityOk returns a tuple with the AffectedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntity

`func (o *HciFeatureViolation) SetAffectedEntity(v string)`

SetAffectedEntity sets AffectedEntity field to given value.

### HasAffectedEntity

`func (o *HciFeatureViolation) HasAffectedEntity() bool`

HasAffectedEntity returns a boolean if a field has been set.

### GetDescription

`func (o *HciFeatureViolation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HciFeatureViolation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HciFeatureViolation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HciFeatureViolation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatureId

`func (o *HciFeatureViolation) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *HciFeatureViolation) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *HciFeatureViolation) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *HciFeatureViolation) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetName

`func (o *HciFeatureViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciFeatureViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciFeatureViolation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciFeatureViolation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


