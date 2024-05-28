# MarketplaceUseCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "marketplace.UseCase"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "marketplace.UseCase"]
**Dependencies** | Pointer to [**[]MarketplaceUseCaseDependency**](MarketplaceUseCaseDependency.md) |  | [optional] 
**Locales** | Pointer to [**[]MarketplaceUseCaseLocale**](MarketplaceUseCaseLocale.md) |  | [optional] 
**UniqueName** | Pointer to **string** | A unique identifier is used to prevent duplicates. | [optional] 
**Catalog** | Pointer to [**NullableWorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewMarketplaceUseCase

`func NewMarketplaceUseCase(classId string, objectType string, ) *MarketplaceUseCase`

NewMarketplaceUseCase instantiates a new MarketplaceUseCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseWithDefaults

`func NewMarketplaceUseCaseWithDefaults() *MarketplaceUseCase`

NewMarketplaceUseCaseWithDefaults instantiates a new MarketplaceUseCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDependencies

`func (o *MarketplaceUseCase) GetDependencies() []MarketplaceUseCaseDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *MarketplaceUseCase) GetDependenciesOk() (*[]MarketplaceUseCaseDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *MarketplaceUseCase) SetDependencies(v []MarketplaceUseCaseDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *MarketplaceUseCase) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *MarketplaceUseCase) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *MarketplaceUseCase) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetLocales

`func (o *MarketplaceUseCase) GetLocales() []MarketplaceUseCaseLocale`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *MarketplaceUseCase) GetLocalesOk() (*[]MarketplaceUseCaseLocale, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *MarketplaceUseCase) SetLocales(v []MarketplaceUseCaseLocale)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *MarketplaceUseCase) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### SetLocalesNil

`func (o *MarketplaceUseCase) SetLocalesNil(b bool)`

 SetLocalesNil sets the value for Locales to be an explicit nil

### UnsetLocales
`func (o *MarketplaceUseCase) UnsetLocales()`

UnsetLocales ensures that no value is present for Locales, not even an explicit nil
### GetUniqueName

`func (o *MarketplaceUseCase) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *MarketplaceUseCase) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *MarketplaceUseCase) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *MarketplaceUseCase) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetCatalog

`func (o *MarketplaceUseCase) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *MarketplaceUseCase) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *MarketplaceUseCase) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *MarketplaceUseCase) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *MarketplaceUseCase) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *MarketplaceUseCase) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


