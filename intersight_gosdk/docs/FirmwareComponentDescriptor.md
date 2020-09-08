# FirmwareComponentDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandString** | Pointer to **string** | The brand string of the endpoint for which this capability information is applicable. | [optional] 
**Label** | Pointer to **string** | The label type for the component. | [optional] 
**Revision** | Pointer to **string** | The revision for the component. | [optional] 

## Methods

### NewFirmwareComponentDescriptor

`func NewFirmwareComponentDescriptor() *FirmwareComponentDescriptor`

NewFirmwareComponentDescriptor instantiates a new FirmwareComponentDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareComponentDescriptorWithDefaults

`func NewFirmwareComponentDescriptorWithDefaults() *FirmwareComponentDescriptor`

NewFirmwareComponentDescriptorWithDefaults instantiates a new FirmwareComponentDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandString

`func (o *FirmwareComponentDescriptor) GetBrandString() string`

GetBrandString returns the BrandString field if non-nil, zero value otherwise.

### GetBrandStringOk

`func (o *FirmwareComponentDescriptor) GetBrandStringOk() (*string, bool)`

GetBrandStringOk returns a tuple with the BrandString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandString

`func (o *FirmwareComponentDescriptor) SetBrandString(v string)`

SetBrandString sets BrandString field to given value.

### HasBrandString

`func (o *FirmwareComponentDescriptor) HasBrandString() bool`

HasBrandString returns a boolean if a field has been set.

### GetLabel

`func (o *FirmwareComponentDescriptor) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FirmwareComponentDescriptor) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FirmwareComponentDescriptor) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FirmwareComponentDescriptor) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRevision

`func (o *FirmwareComponentDescriptor) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FirmwareComponentDescriptor) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FirmwareComponentDescriptor) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FirmwareComponentDescriptor) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


