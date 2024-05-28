# CapabilityHsuIsoFileSupportMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.HsuIsoFileSupportMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.HsuIsoFileSupportMeta"]
**DefaultFileName** | Pointer to **string** | Name of the symbolic link to the actual iso file in the default case. | [optional] [readonly] 
**DefaultMinVersion** | Pointer to **string** | Firmware version from which the HSU capability is present in the default case. | [optional] [readonly] 
**ModelSpecificConstraint** | Pointer to [**[]CapabilityHsuIsoModelSpecificConstraint**](CapabilityHsuIsoModelSpecificConstraint.md) |  | [optional] 
**Series** | Pointer to **string** | Series name for the capability group. | [optional] [readonly] 

## Methods

### NewCapabilityHsuIsoFileSupportMeta

`func NewCapabilityHsuIsoFileSupportMeta(classId string, objectType string, ) *CapabilityHsuIsoFileSupportMeta`

NewCapabilityHsuIsoFileSupportMeta instantiates a new CapabilityHsuIsoFileSupportMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityHsuIsoFileSupportMetaWithDefaults

`func NewCapabilityHsuIsoFileSupportMetaWithDefaults() *CapabilityHsuIsoFileSupportMeta`

NewCapabilityHsuIsoFileSupportMetaWithDefaults instantiates a new CapabilityHsuIsoFileSupportMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityHsuIsoFileSupportMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityHsuIsoFileSupportMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityHsuIsoFileSupportMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityHsuIsoFileSupportMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityHsuIsoFileSupportMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityHsuIsoFileSupportMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultFileName

`func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultFileName() string`

GetDefaultFileName returns the DefaultFileName field if non-nil, zero value otherwise.

### GetDefaultFileNameOk

`func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultFileNameOk() (*string, bool)`

GetDefaultFileNameOk returns a tuple with the DefaultFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFileName

`func (o *CapabilityHsuIsoFileSupportMeta) SetDefaultFileName(v string)`

SetDefaultFileName sets DefaultFileName field to given value.

### HasDefaultFileName

`func (o *CapabilityHsuIsoFileSupportMeta) HasDefaultFileName() bool`

HasDefaultFileName returns a boolean if a field has been set.

### GetDefaultMinVersion

`func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultMinVersion() string`

GetDefaultMinVersion returns the DefaultMinVersion field if non-nil, zero value otherwise.

### GetDefaultMinVersionOk

`func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultMinVersionOk() (*string, bool)`

GetDefaultMinVersionOk returns a tuple with the DefaultMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMinVersion

`func (o *CapabilityHsuIsoFileSupportMeta) SetDefaultMinVersion(v string)`

SetDefaultMinVersion sets DefaultMinVersion field to given value.

### HasDefaultMinVersion

`func (o *CapabilityHsuIsoFileSupportMeta) HasDefaultMinVersion() bool`

HasDefaultMinVersion returns a boolean if a field has been set.

### GetModelSpecificConstraint

`func (o *CapabilityHsuIsoFileSupportMeta) GetModelSpecificConstraint() []CapabilityHsuIsoModelSpecificConstraint`

GetModelSpecificConstraint returns the ModelSpecificConstraint field if non-nil, zero value otherwise.

### GetModelSpecificConstraintOk

`func (o *CapabilityHsuIsoFileSupportMeta) GetModelSpecificConstraintOk() (*[]CapabilityHsuIsoModelSpecificConstraint, bool)`

GetModelSpecificConstraintOk returns a tuple with the ModelSpecificConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelSpecificConstraint

`func (o *CapabilityHsuIsoFileSupportMeta) SetModelSpecificConstraint(v []CapabilityHsuIsoModelSpecificConstraint)`

SetModelSpecificConstraint sets ModelSpecificConstraint field to given value.

### HasModelSpecificConstraint

`func (o *CapabilityHsuIsoFileSupportMeta) HasModelSpecificConstraint() bool`

HasModelSpecificConstraint returns a boolean if a field has been set.

### SetModelSpecificConstraintNil

`func (o *CapabilityHsuIsoFileSupportMeta) SetModelSpecificConstraintNil(b bool)`

 SetModelSpecificConstraintNil sets the value for ModelSpecificConstraint to be an explicit nil

### UnsetModelSpecificConstraint
`func (o *CapabilityHsuIsoFileSupportMeta) UnsetModelSpecificConstraint()`

UnsetModelSpecificConstraint ensures that no value is present for ModelSpecificConstraint, not even an explicit nil
### GetSeries

`func (o *CapabilityHsuIsoFileSupportMeta) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CapabilityHsuIsoFileSupportMeta) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CapabilityHsuIsoFileSupportMeta) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CapabilityHsuIsoFileSupportMeta) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


