# OsConfigurationFileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.ConfigurationFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.ConfigurationFile"]
**Description** | Pointer to **string** | Description of the OS ConfigurationFile. | [optional] 
**FileContent** | Pointer to **string** | The content of the entire configuration file is stored as value. The content can either be a static file content or a template content. The template is expected to conform to the golang template syntax. The values from os.Answers properties will be used to populate this template. | [optional] 
**Internal** | Pointer to **bool** | The internal flag is set to true when configuration file is uploaded from OS Install wizard. Internal Configuration files will not be displayed in Answer Management Page. | [optional] [default to false]
**Name** | Pointer to **string** | The name of the OS ConfigurationFile that uniquely identifies the configuration file. | [optional] 
**Placeholders** | Pointer to [**[]OsPlaceHolder**](OsPlaceHolder.md) |  | [optional] 
**Supported** | Pointer to **bool** | An internal property that is used to distinguish between the pre-canned OS configuration file entries and user provided entries. | [optional] [readonly] 
**Catalog** | Pointer to [**OsCatalogRelationship**](OsCatalogRelationship.md) |  | [optional] 
**Distributions** | Pointer to [**[]HclOperatingSystemRelationship**](HclOperatingSystemRelationship.md) | An array of relationships to hclOperatingSystem resources. | [optional] 

## Methods

### NewOsConfigurationFileAllOf

`func NewOsConfigurationFileAllOf(classId string, objectType string, ) *OsConfigurationFileAllOf`

NewOsConfigurationFileAllOf instantiates a new OsConfigurationFileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationFileAllOfWithDefaults

`func NewOsConfigurationFileAllOfWithDefaults() *OsConfigurationFileAllOf`

NewOsConfigurationFileAllOfWithDefaults instantiates a new OsConfigurationFileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsConfigurationFileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsConfigurationFileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsConfigurationFileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsConfigurationFileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsConfigurationFileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsConfigurationFileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *OsConfigurationFileAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OsConfigurationFileAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OsConfigurationFileAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OsConfigurationFileAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileContent

`func (o *OsConfigurationFileAllOf) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *OsConfigurationFileAllOf) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *OsConfigurationFileAllOf) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.

### HasFileContent

`func (o *OsConfigurationFileAllOf) HasFileContent() bool`

HasFileContent returns a boolean if a field has been set.

### GetInternal

`func (o *OsConfigurationFileAllOf) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *OsConfigurationFileAllOf) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *OsConfigurationFileAllOf) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *OsConfigurationFileAllOf) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *OsConfigurationFileAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsConfigurationFileAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsConfigurationFileAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsConfigurationFileAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaceholders

`func (o *OsConfigurationFileAllOf) GetPlaceholders() []OsPlaceHolder`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *OsConfigurationFileAllOf) GetPlaceholdersOk() (*[]OsPlaceHolder, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *OsConfigurationFileAllOf) SetPlaceholders(v []OsPlaceHolder)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *OsConfigurationFileAllOf) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### SetPlaceholdersNil

`func (o *OsConfigurationFileAllOf) SetPlaceholdersNil(b bool)`

 SetPlaceholdersNil sets the value for Placeholders to be an explicit nil

### UnsetPlaceholders
`func (o *OsConfigurationFileAllOf) UnsetPlaceholders()`

UnsetPlaceholders ensures that no value is present for Placeholders, not even an explicit nil
### GetSupported

`func (o *OsConfigurationFileAllOf) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *OsConfigurationFileAllOf) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *OsConfigurationFileAllOf) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *OsConfigurationFileAllOf) HasSupported() bool`

HasSupported returns a boolean if a field has been set.

### GetCatalog

`func (o *OsConfigurationFileAllOf) GetCatalog() OsCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *OsConfigurationFileAllOf) GetCatalogOk() (*OsCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *OsConfigurationFileAllOf) SetCatalog(v OsCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *OsConfigurationFileAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetDistributions

`func (o *OsConfigurationFileAllOf) GetDistributions() []HclOperatingSystemRelationship`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *OsConfigurationFileAllOf) GetDistributionsOk() (*[]HclOperatingSystemRelationship, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *OsConfigurationFileAllOf) SetDistributions(v []HclOperatingSystemRelationship)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *OsConfigurationFileAllOf) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### SetDistributionsNil

`func (o *OsConfigurationFileAllOf) SetDistributionsNil(b bool)`

 SetDistributionsNil sets the value for Distributions to be an explicit nil

### UnsetDistributions
`func (o *OsConfigurationFileAllOf) UnsetDistributions()`

UnsetDistributions ensures that no value is present for Distributions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


