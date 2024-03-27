# MarketplaceUseCaseVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "marketplace.UseCaseVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "marketplace.UseCaseVersion"]
**Locales** | Pointer to [**[]MarketplaceUseCaseVersionLocale**](MarketplaceUseCaseVersionLocale.md) |  | [optional] 
**Resources** | Pointer to [**[]MarketplaceUseCaseVersionResources**](MarketplaceUseCaseVersionResources.md) |  | [optional] 
**Version** | Pointer to **string** | A string version for each use case | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**UseCase** | Pointer to [**MarketplaceUseCaseRelationship**](MarketplaceUseCaseRelationship.md) |  | [optional] 

## Methods

### NewMarketplaceUseCaseVersion

`func NewMarketplaceUseCaseVersion(classId string, objectType string, ) *MarketplaceUseCaseVersion`

NewMarketplaceUseCaseVersion instantiates a new MarketplaceUseCaseVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseVersionWithDefaults

`func NewMarketplaceUseCaseVersionWithDefaults() *MarketplaceUseCaseVersion`

NewMarketplaceUseCaseVersionWithDefaults instantiates a new MarketplaceUseCaseVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCaseVersion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCaseVersion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCaseVersion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCaseVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCaseVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCaseVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocales

`func (o *MarketplaceUseCaseVersion) GetLocales() []MarketplaceUseCaseVersionLocale`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *MarketplaceUseCaseVersion) GetLocalesOk() (*[]MarketplaceUseCaseVersionLocale, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *MarketplaceUseCaseVersion) SetLocales(v []MarketplaceUseCaseVersionLocale)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *MarketplaceUseCaseVersion) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### SetLocalesNil

`func (o *MarketplaceUseCaseVersion) SetLocalesNil(b bool)`

 SetLocalesNil sets the value for Locales to be an explicit nil

### UnsetLocales
`func (o *MarketplaceUseCaseVersion) UnsetLocales()`

UnsetLocales ensures that no value is present for Locales, not even an explicit nil
### GetResources

`func (o *MarketplaceUseCaseVersion) GetResources() []MarketplaceUseCaseVersionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MarketplaceUseCaseVersion) GetResourcesOk() (*[]MarketplaceUseCaseVersionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MarketplaceUseCaseVersion) SetResources(v []MarketplaceUseCaseVersionResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MarketplaceUseCaseVersion) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *MarketplaceUseCaseVersion) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *MarketplaceUseCaseVersion) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetVersion

`func (o *MarketplaceUseCaseVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketplaceUseCaseVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketplaceUseCaseVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketplaceUseCaseVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrganization

`func (o *MarketplaceUseCaseVersion) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MarketplaceUseCaseVersion) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MarketplaceUseCaseVersion) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MarketplaceUseCaseVersion) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetUseCase

`func (o *MarketplaceUseCaseVersion) GetUseCase() MarketplaceUseCaseRelationship`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *MarketplaceUseCaseVersion) GetUseCaseOk() (*MarketplaceUseCaseRelationship, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *MarketplaceUseCaseVersion) SetUseCase(v MarketplaceUseCaseRelationship)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *MarketplaceUseCaseVersion) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


