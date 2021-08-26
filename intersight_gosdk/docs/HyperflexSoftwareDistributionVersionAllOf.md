# HyperflexSoftwareDistributionVersionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SoftwareDistributionVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SoftwareDistributionVersion"]
**Version** | Pointer to **string** | The HyperFlex Software Distribution version. | [optional] 
**DistributionComponents** | Pointer to [**[]HyperflexSoftwareDistributionComponentRelationship**](HyperflexSoftwareDistributionComponentRelationship.md) | An array of relationships to hyperflexSoftwareDistributionComponent resources. | [optional] 
**SoftwareDistributionEntry** | Pointer to [**HyperflexSoftwareDistributionEntryRelationship**](HyperflexSoftwareDistributionEntryRelationship.md) |  | [optional] 

## Methods

### NewHyperflexSoftwareDistributionVersionAllOf

`func NewHyperflexSoftwareDistributionVersionAllOf(classId string, objectType string, ) *HyperflexSoftwareDistributionVersionAllOf`

NewHyperflexSoftwareDistributionVersionAllOf instantiates a new HyperflexSoftwareDistributionVersionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSoftwareDistributionVersionAllOfWithDefaults

`func NewHyperflexSoftwareDistributionVersionAllOfWithDefaults() *HyperflexSoftwareDistributionVersionAllOf`

NewHyperflexSoftwareDistributionVersionAllOfWithDefaults instantiates a new HyperflexSoftwareDistributionVersionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSoftwareDistributionVersionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSoftwareDistributionVersionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersion

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexSoftwareDistributionVersionAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexSoftwareDistributionVersionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDistributionComponents

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetDistributionComponents() []HyperflexSoftwareDistributionComponentRelationship`

GetDistributionComponents returns the DistributionComponents field if non-nil, zero value otherwise.

### GetDistributionComponentsOk

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetDistributionComponentsOk() (*[]HyperflexSoftwareDistributionComponentRelationship, bool)`

GetDistributionComponentsOk returns a tuple with the DistributionComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionComponents

`func (o *HyperflexSoftwareDistributionVersionAllOf) SetDistributionComponents(v []HyperflexSoftwareDistributionComponentRelationship)`

SetDistributionComponents sets DistributionComponents field to given value.

### HasDistributionComponents

`func (o *HyperflexSoftwareDistributionVersionAllOf) HasDistributionComponents() bool`

HasDistributionComponents returns a boolean if a field has been set.

### SetDistributionComponentsNil

`func (o *HyperflexSoftwareDistributionVersionAllOf) SetDistributionComponentsNil(b bool)`

 SetDistributionComponentsNil sets the value for DistributionComponents to be an explicit nil

### UnsetDistributionComponents
`func (o *HyperflexSoftwareDistributionVersionAllOf) UnsetDistributionComponents()`

UnsetDistributionComponents ensures that no value is present for DistributionComponents, not even an explicit nil
### GetSoftwareDistributionEntry

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetSoftwareDistributionEntry() HyperflexSoftwareDistributionEntryRelationship`

GetSoftwareDistributionEntry returns the SoftwareDistributionEntry field if non-nil, zero value otherwise.

### GetSoftwareDistributionEntryOk

`func (o *HyperflexSoftwareDistributionVersionAllOf) GetSoftwareDistributionEntryOk() (*HyperflexSoftwareDistributionEntryRelationship, bool)`

GetSoftwareDistributionEntryOk returns a tuple with the SoftwareDistributionEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDistributionEntry

`func (o *HyperflexSoftwareDistributionVersionAllOf) SetSoftwareDistributionEntry(v HyperflexSoftwareDistributionEntryRelationship)`

SetSoftwareDistributionEntry sets SoftwareDistributionEntry field to given value.

### HasSoftwareDistributionEntry

`func (o *HyperflexSoftwareDistributionVersionAllOf) HasSoftwareDistributionEntry() bool`

HasSoftwareDistributionEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


