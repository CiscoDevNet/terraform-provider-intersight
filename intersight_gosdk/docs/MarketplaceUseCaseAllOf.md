# MarketplaceUseCaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "marketplace.UseCase"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "marketplace.UseCase"]
**Dependencies** | Pointer to [**[]MarketplaceUseCaseDependency**](MarketplaceUseCaseDependency.md) |  | [optional] 
**Locales** | Pointer to [**[]MarketplaceUseCaseLocale**](MarketplaceUseCaseLocale.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewMarketplaceUseCaseAllOf

`func NewMarketplaceUseCaseAllOf(classId string, objectType string, ) *MarketplaceUseCaseAllOf`

NewMarketplaceUseCaseAllOf instantiates a new MarketplaceUseCaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseAllOfWithDefaults

`func NewMarketplaceUseCaseAllOfWithDefaults() *MarketplaceUseCaseAllOf`

NewMarketplaceUseCaseAllOfWithDefaults instantiates a new MarketplaceUseCaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDependencies

`func (o *MarketplaceUseCaseAllOf) GetDependencies() []MarketplaceUseCaseDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *MarketplaceUseCaseAllOf) GetDependenciesOk() (*[]MarketplaceUseCaseDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *MarketplaceUseCaseAllOf) SetDependencies(v []MarketplaceUseCaseDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *MarketplaceUseCaseAllOf) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *MarketplaceUseCaseAllOf) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *MarketplaceUseCaseAllOf) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetLocales

`func (o *MarketplaceUseCaseAllOf) GetLocales() []MarketplaceUseCaseLocale`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *MarketplaceUseCaseAllOf) GetLocalesOk() (*[]MarketplaceUseCaseLocale, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *MarketplaceUseCaseAllOf) SetLocales(v []MarketplaceUseCaseLocale)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *MarketplaceUseCaseAllOf) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### SetLocalesNil

`func (o *MarketplaceUseCaseAllOf) SetLocalesNil(b bool)`

 SetLocalesNil sets the value for Locales to be an explicit nil

### UnsetLocales
`func (o *MarketplaceUseCaseAllOf) UnsetLocales()`

UnsetLocales ensures that no value is present for Locales, not even an explicit nil
### GetOrganization

`func (o *MarketplaceUseCaseAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MarketplaceUseCaseAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MarketplaceUseCaseAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MarketplaceUseCaseAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


