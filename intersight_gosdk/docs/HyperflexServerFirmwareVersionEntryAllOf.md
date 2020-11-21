# HyperflexServerFirmwareVersionEntryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ServerFirmwareVersionEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ServerFirmwareVersionEntry"]
**Constraint** | Pointer to [**NullableHyperflexAppSettingConstraint**](hyperflex.AppSettingConstraint.md) |  | [optional] 
**Label** | Pointer to **string** | The display name for server firmware bundle version in UI. | [optional] 

## Methods

### NewHyperflexServerFirmwareVersionEntryAllOf

`func NewHyperflexServerFirmwareVersionEntryAllOf(classId string, objectType string, ) *HyperflexServerFirmwareVersionEntryAllOf`

NewHyperflexServerFirmwareVersionEntryAllOf instantiates a new HyperflexServerFirmwareVersionEntryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults

`func NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults() *HyperflexServerFirmwareVersionEntryAllOf`

NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults instantiates a new HyperflexServerFirmwareVersionEntryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraint

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetConstraint() HyperflexAppSettingConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetConstraintOk() (*HyperflexAppSettingConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetConstraint(v HyperflexAppSettingConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *HyperflexServerFirmwareVersionEntryAllOf) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### SetConstraintNil

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetConstraintNil(b bool)`

 SetConstraintNil sets the value for Constraint to be an explicit nil

### UnsetConstraint
`func (o *HyperflexServerFirmwareVersionEntryAllOf) UnsetConstraint()`

UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
### GetLabel

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *HyperflexServerFirmwareVersionEntryAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


