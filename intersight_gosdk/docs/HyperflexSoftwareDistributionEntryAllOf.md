# HyperflexSoftwareDistributionEntryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SoftwareDistributionEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SoftwareDistributionEntry"]
**DistributionType** | Pointer to **string** | The HyperFlex Software Distribution type. | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 
**DistributionVersions** | Pointer to [**[]HyperflexSoftwareDistributionVersionRelationship**](HyperflexSoftwareDistributionVersionRelationship.md) | An array of relationships to hyperflexSoftwareDistributionVersion resources. | [optional] 

## Methods

### NewHyperflexSoftwareDistributionEntryAllOf

`func NewHyperflexSoftwareDistributionEntryAllOf(classId string, objectType string, ) *HyperflexSoftwareDistributionEntryAllOf`

NewHyperflexSoftwareDistributionEntryAllOf instantiates a new HyperflexSoftwareDistributionEntryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSoftwareDistributionEntryAllOfWithDefaults

`func NewHyperflexSoftwareDistributionEntryAllOfWithDefaults() *HyperflexSoftwareDistributionEntryAllOf`

NewHyperflexSoftwareDistributionEntryAllOfWithDefaults instantiates a new HyperflexSoftwareDistributionEntryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSoftwareDistributionEntryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSoftwareDistributionEntryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDistributionType

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetDistributionType() string`

GetDistributionType returns the DistributionType field if non-nil, zero value otherwise.

### GetDistributionTypeOk

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetDistributionTypeOk() (*string, bool)`

GetDistributionTypeOk returns a tuple with the DistributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionType

`func (o *HyperflexSoftwareDistributionEntryAllOf) SetDistributionType(v string)`

SetDistributionType sets DistributionType field to given value.

### HasDistributionType

`func (o *HyperflexSoftwareDistributionEntryAllOf) HasDistributionType() bool`

HasDistributionType returns a boolean if a field has been set.

### GetAppCatalog

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexSoftwareDistributionEntryAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexSoftwareDistributionEntryAllOf) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.

### GetDistributionVersions

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetDistributionVersions() []HyperflexSoftwareDistributionVersionRelationship`

GetDistributionVersions returns the DistributionVersions field if non-nil, zero value otherwise.

### GetDistributionVersionsOk

`func (o *HyperflexSoftwareDistributionEntryAllOf) GetDistributionVersionsOk() (*[]HyperflexSoftwareDistributionVersionRelationship, bool)`

GetDistributionVersionsOk returns a tuple with the DistributionVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionVersions

`func (o *HyperflexSoftwareDistributionEntryAllOf) SetDistributionVersions(v []HyperflexSoftwareDistributionVersionRelationship)`

SetDistributionVersions sets DistributionVersions field to given value.

### HasDistributionVersions

`func (o *HyperflexSoftwareDistributionEntryAllOf) HasDistributionVersions() bool`

HasDistributionVersions returns a boolean if a field has been set.

### SetDistributionVersionsNil

`func (o *HyperflexSoftwareDistributionEntryAllOf) SetDistributionVersionsNil(b bool)`

 SetDistributionVersionsNil sets the value for DistributionVersions to be an explicit nil

### UnsetDistributionVersions
`func (o *HyperflexSoftwareDistributionEntryAllOf) UnsetDistributionVersions()`

UnsetDistributionVersions ensures that no value is present for DistributionVersions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


