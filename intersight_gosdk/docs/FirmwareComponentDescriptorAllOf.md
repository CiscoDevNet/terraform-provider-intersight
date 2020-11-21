# FirmwareComponentDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**BrandString** | Pointer to **string** | The brand string of the endpoint for which this capability information is applicable. | [optional] 
**Label** | Pointer to **string** | The label type for the component. | [optional] 
**Revision** | Pointer to **string** | The revision for the component. | [optional] 

## Methods

### NewFirmwareComponentDescriptorAllOf

`func NewFirmwareComponentDescriptorAllOf(classId string, objectType string, ) *FirmwareComponentDescriptorAllOf`

NewFirmwareComponentDescriptorAllOf instantiates a new FirmwareComponentDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareComponentDescriptorAllOfWithDefaults

`func NewFirmwareComponentDescriptorAllOfWithDefaults() *FirmwareComponentDescriptorAllOf`

NewFirmwareComponentDescriptorAllOfWithDefaults instantiates a new FirmwareComponentDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareComponentDescriptorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareComponentDescriptorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareComponentDescriptorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareComponentDescriptorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareComponentDescriptorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareComponentDescriptorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBrandString

`func (o *FirmwareComponentDescriptorAllOf) GetBrandString() string`

GetBrandString returns the BrandString field if non-nil, zero value otherwise.

### GetBrandStringOk

`func (o *FirmwareComponentDescriptorAllOf) GetBrandStringOk() (*string, bool)`

GetBrandStringOk returns a tuple with the BrandString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandString

`func (o *FirmwareComponentDescriptorAllOf) SetBrandString(v string)`

SetBrandString sets BrandString field to given value.

### HasBrandString

`func (o *FirmwareComponentDescriptorAllOf) HasBrandString() bool`

HasBrandString returns a boolean if a field has been set.

### GetLabel

`func (o *FirmwareComponentDescriptorAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FirmwareComponentDescriptorAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FirmwareComponentDescriptorAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FirmwareComponentDescriptorAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRevision

`func (o *FirmwareComponentDescriptorAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FirmwareComponentDescriptorAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FirmwareComponentDescriptorAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FirmwareComponentDescriptorAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


