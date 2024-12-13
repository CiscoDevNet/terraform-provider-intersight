# HciCapacityViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.CapacityViolation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.CapacityViolation"]
**Category** | Pointer to **string** | The category of the license that has violation. | [optional] [readonly] 
**Meter** | Pointer to **string** | The meter of the license that has violation. | [optional] [readonly] 
**Shortfall** | Pointer to **int64** | This indicates the insufficient quantity of a license. License of given quantity should be applied on the cluster for it to function properly. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the license that has violation. | [optional] [readonly] 

## Methods

### NewHciCapacityViolation

`func NewHciCapacityViolation(classId string, objectType string, ) *HciCapacityViolation`

NewHciCapacityViolation instantiates a new HciCapacityViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciCapacityViolationWithDefaults

`func NewHciCapacityViolationWithDefaults() *HciCapacityViolation`

NewHciCapacityViolationWithDefaults instantiates a new HciCapacityViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciCapacityViolation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciCapacityViolation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciCapacityViolation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciCapacityViolation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciCapacityViolation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciCapacityViolation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HciCapacityViolation) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HciCapacityViolation) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HciCapacityViolation) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HciCapacityViolation) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetMeter

`func (o *HciCapacityViolation) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *HciCapacityViolation) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *HciCapacityViolation) SetMeter(v string)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *HciCapacityViolation) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetShortfall

`func (o *HciCapacityViolation) GetShortfall() int64`

GetShortfall returns the Shortfall field if non-nil, zero value otherwise.

### GetShortfallOk

`func (o *HciCapacityViolation) GetShortfallOk() (*int64, bool)`

GetShortfallOk returns a tuple with the Shortfall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortfall

`func (o *HciCapacityViolation) SetShortfall(v int64)`

SetShortfall sets Shortfall field to given value.

### HasShortfall

`func (o *HciCapacityViolation) HasShortfall() bool`

HasShortfall returns a boolean if a field has been set.

### GetType

`func (o *HciCapacityViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciCapacityViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciCapacityViolation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciCapacityViolation) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


