# OsCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.Catalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.Catalog"]
**Name** | Pointer to **string** | The catalog name. There will be one for system and one for each user account. | [optional] 
**ConfigurationFiles** | Pointer to [**[]OsConfigurationFileRelationship**](OsConfigurationFileRelationship.md) | An array of relationships to osConfigurationFile resources. | [optional] 
**Distributions** | Pointer to [**[]OsDistributionRelationship**](OsDistributionRelationship.md) | An array of relationships to osDistribution resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewOsCatalogAllOf

`func NewOsCatalogAllOf(classId string, objectType string, ) *OsCatalogAllOf`

NewOsCatalogAllOf instantiates a new OsCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsCatalogAllOfWithDefaults

`func NewOsCatalogAllOfWithDefaults() *OsCatalogAllOf`

NewOsCatalogAllOfWithDefaults instantiates a new OsCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsCatalogAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsCatalogAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsCatalogAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsCatalogAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsCatalogAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsCatalogAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *OsCatalogAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsCatalogAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsCatalogAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsCatalogAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurationFiles

`func (o *OsCatalogAllOf) GetConfigurationFiles() []OsConfigurationFileRelationship`

GetConfigurationFiles returns the ConfigurationFiles field if non-nil, zero value otherwise.

### GetConfigurationFilesOk

`func (o *OsCatalogAllOf) GetConfigurationFilesOk() (*[]OsConfigurationFileRelationship, bool)`

GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFiles

`func (o *OsCatalogAllOf) SetConfigurationFiles(v []OsConfigurationFileRelationship)`

SetConfigurationFiles sets ConfigurationFiles field to given value.

### HasConfigurationFiles

`func (o *OsCatalogAllOf) HasConfigurationFiles() bool`

HasConfigurationFiles returns a boolean if a field has been set.

### SetConfigurationFilesNil

`func (o *OsCatalogAllOf) SetConfigurationFilesNil(b bool)`

 SetConfigurationFilesNil sets the value for ConfigurationFiles to be an explicit nil

### UnsetConfigurationFiles
`func (o *OsCatalogAllOf) UnsetConfigurationFiles()`

UnsetConfigurationFiles ensures that no value is present for ConfigurationFiles, not even an explicit nil
### GetDistributions

`func (o *OsCatalogAllOf) GetDistributions() []OsDistributionRelationship`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *OsCatalogAllOf) GetDistributionsOk() (*[]OsDistributionRelationship, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *OsCatalogAllOf) SetDistributions(v []OsDistributionRelationship)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *OsCatalogAllOf) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### SetDistributionsNil

`func (o *OsCatalogAllOf) SetDistributionsNil(b bool)`

 SetDistributionsNil sets the value for Distributions to be an explicit nil

### UnsetDistributions
`func (o *OsCatalogAllOf) UnsetDistributions()`

UnsetDistributions ensures that no value is present for Distributions, not even an explicit nil
### GetOrganization

`func (o *OsCatalogAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OsCatalogAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OsCatalogAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OsCatalogAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


