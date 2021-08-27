# HyperflexSoftwareDistributionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SoftwareDistributionVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SoftwareDistributionVersion"]
**Version** | Pointer to **string** | The HyperFlex Software Distribution version. | [optional] 
**DistributionComponents** | Pointer to [**[]HyperflexSoftwareDistributionComponentRelationship**](HyperflexSoftwareDistributionComponentRelationship.md) | An array of relationships to hyperflexSoftwareDistributionComponent resources. | [optional] 
**SoftwareDistributionEntry** | Pointer to [**HyperflexSoftwareDistributionEntryRelationship**](HyperflexSoftwareDistributionEntryRelationship.md) |  | [optional] 

## Methods

### NewHyperflexSoftwareDistributionVersion

`func NewHyperflexSoftwareDistributionVersion(classId string, objectType string, ) *HyperflexSoftwareDistributionVersion`

NewHyperflexSoftwareDistributionVersion instantiates a new HyperflexSoftwareDistributionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSoftwareDistributionVersionWithDefaults

`func NewHyperflexSoftwareDistributionVersionWithDefaults() *HyperflexSoftwareDistributionVersion`

NewHyperflexSoftwareDistributionVersionWithDefaults instantiates a new HyperflexSoftwareDistributionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSoftwareDistributionVersion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSoftwareDistributionVersion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSoftwareDistributionVersion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSoftwareDistributionVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSoftwareDistributionVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSoftwareDistributionVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersion

`func (o *HyperflexSoftwareDistributionVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexSoftwareDistributionVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexSoftwareDistributionVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexSoftwareDistributionVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDistributionComponents

`func (o *HyperflexSoftwareDistributionVersion) GetDistributionComponents() []HyperflexSoftwareDistributionComponentRelationship`

GetDistributionComponents returns the DistributionComponents field if non-nil, zero value otherwise.

### GetDistributionComponentsOk

`func (o *HyperflexSoftwareDistributionVersion) GetDistributionComponentsOk() (*[]HyperflexSoftwareDistributionComponentRelationship, bool)`

GetDistributionComponentsOk returns a tuple with the DistributionComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionComponents

`func (o *HyperflexSoftwareDistributionVersion) SetDistributionComponents(v []HyperflexSoftwareDistributionComponentRelationship)`

SetDistributionComponents sets DistributionComponents field to given value.

### HasDistributionComponents

`func (o *HyperflexSoftwareDistributionVersion) HasDistributionComponents() bool`

HasDistributionComponents returns a boolean if a field has been set.

### SetDistributionComponentsNil

`func (o *HyperflexSoftwareDistributionVersion) SetDistributionComponentsNil(b bool)`

 SetDistributionComponentsNil sets the value for DistributionComponents to be an explicit nil

### UnsetDistributionComponents
`func (o *HyperflexSoftwareDistributionVersion) UnsetDistributionComponents()`

UnsetDistributionComponents ensures that no value is present for DistributionComponents, not even an explicit nil
### GetSoftwareDistributionEntry

`func (o *HyperflexSoftwareDistributionVersion) GetSoftwareDistributionEntry() HyperflexSoftwareDistributionEntryRelationship`

GetSoftwareDistributionEntry returns the SoftwareDistributionEntry field if non-nil, zero value otherwise.

### GetSoftwareDistributionEntryOk

`func (o *HyperflexSoftwareDistributionVersion) GetSoftwareDistributionEntryOk() (*HyperflexSoftwareDistributionEntryRelationship, bool)`

GetSoftwareDistributionEntryOk returns a tuple with the SoftwareDistributionEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDistributionEntry

`func (o *HyperflexSoftwareDistributionVersion) SetSoftwareDistributionEntry(v HyperflexSoftwareDistributionEntryRelationship)`

SetSoftwareDistributionEntry sets SoftwareDistributionEntry field to given value.

### HasSoftwareDistributionEntry

`func (o *HyperflexSoftwareDistributionVersion) HasSoftwareDistributionEntry() bool`

HasSoftwareDistributionEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


