# MarketplaceUseCaseVersionAllOf

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

### NewMarketplaceUseCaseVersionAllOf

`func NewMarketplaceUseCaseVersionAllOf(classId string, objectType string, ) *MarketplaceUseCaseVersionAllOf`

NewMarketplaceUseCaseVersionAllOf instantiates a new MarketplaceUseCaseVersionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseVersionAllOfWithDefaults

`func NewMarketplaceUseCaseVersionAllOfWithDefaults() *MarketplaceUseCaseVersionAllOf`

NewMarketplaceUseCaseVersionAllOfWithDefaults instantiates a new MarketplaceUseCaseVersionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCaseVersionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCaseVersionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCaseVersionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCaseVersionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCaseVersionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCaseVersionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocales

`func (o *MarketplaceUseCaseVersionAllOf) GetLocales() []MarketplaceUseCaseVersionLocale`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *MarketplaceUseCaseVersionAllOf) GetLocalesOk() (*[]MarketplaceUseCaseVersionLocale, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *MarketplaceUseCaseVersionAllOf) SetLocales(v []MarketplaceUseCaseVersionLocale)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *MarketplaceUseCaseVersionAllOf) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### SetLocalesNil

`func (o *MarketplaceUseCaseVersionAllOf) SetLocalesNil(b bool)`

 SetLocalesNil sets the value for Locales to be an explicit nil

### UnsetLocales
`func (o *MarketplaceUseCaseVersionAllOf) UnsetLocales()`

UnsetLocales ensures that no value is present for Locales, not even an explicit nil
### GetResources

`func (o *MarketplaceUseCaseVersionAllOf) GetResources() []MarketplaceUseCaseVersionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MarketplaceUseCaseVersionAllOf) GetResourcesOk() (*[]MarketplaceUseCaseVersionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MarketplaceUseCaseVersionAllOf) SetResources(v []MarketplaceUseCaseVersionResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MarketplaceUseCaseVersionAllOf) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *MarketplaceUseCaseVersionAllOf) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *MarketplaceUseCaseVersionAllOf) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetVersion

`func (o *MarketplaceUseCaseVersionAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketplaceUseCaseVersionAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketplaceUseCaseVersionAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketplaceUseCaseVersionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrganization

`func (o *MarketplaceUseCaseVersionAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MarketplaceUseCaseVersionAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MarketplaceUseCaseVersionAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MarketplaceUseCaseVersionAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetUseCase

`func (o *MarketplaceUseCaseVersionAllOf) GetUseCase() MarketplaceUseCaseRelationship`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *MarketplaceUseCaseVersionAllOf) GetUseCaseOk() (*MarketplaceUseCaseRelationship, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *MarketplaceUseCaseVersionAllOf) SetUseCase(v MarketplaceUseCaseRelationship)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *MarketplaceUseCaseVersionAllOf) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


