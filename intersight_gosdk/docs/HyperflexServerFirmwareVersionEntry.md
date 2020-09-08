# HyperflexServerFirmwareVersionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraint** | Pointer to [**HyperflexAppSettingConstraint**](hyperflex.AppSettingConstraint.md) |  | [optional] 
**Label** | Pointer to **string** | The display name for server firmware bundle version in UI. | [optional] 

## Methods

### NewHyperflexServerFirmwareVersionEntry

`func NewHyperflexServerFirmwareVersionEntry() *HyperflexServerFirmwareVersionEntry`

NewHyperflexServerFirmwareVersionEntry instantiates a new HyperflexServerFirmwareVersionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerFirmwareVersionEntryWithDefaults

`func NewHyperflexServerFirmwareVersionEntryWithDefaults() *HyperflexServerFirmwareVersionEntry`

NewHyperflexServerFirmwareVersionEntryWithDefaults instantiates a new HyperflexServerFirmwareVersionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraint

`func (o *HyperflexServerFirmwareVersionEntry) GetConstraint() HyperflexAppSettingConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *HyperflexServerFirmwareVersionEntry) GetConstraintOk() (*HyperflexAppSettingConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *HyperflexServerFirmwareVersionEntry) SetConstraint(v HyperflexAppSettingConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *HyperflexServerFirmwareVersionEntry) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetLabel

`func (o *HyperflexServerFirmwareVersionEntry) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HyperflexServerFirmwareVersionEntry) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HyperflexServerFirmwareVersionEntry) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *HyperflexServerFirmwareVersionEntry) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


