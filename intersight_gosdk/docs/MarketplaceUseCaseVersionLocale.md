# MarketplaceUseCaseVersionLocale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "marketplace.UseCaseVersionLocale"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "marketplace.UseCaseVersionLocale"]
**Description** | Pointer to **string** | The string field to hold the description. | [optional] 
**Locale** | Pointer to **string** | The string field to hold the locale. | [optional] 

## Methods

### NewMarketplaceUseCaseVersionLocale

`func NewMarketplaceUseCaseVersionLocale(classId string, objectType string, ) *MarketplaceUseCaseVersionLocale`

NewMarketplaceUseCaseVersionLocale instantiates a new MarketplaceUseCaseVersionLocale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseVersionLocaleWithDefaults

`func NewMarketplaceUseCaseVersionLocaleWithDefaults() *MarketplaceUseCaseVersionLocale`

NewMarketplaceUseCaseVersionLocaleWithDefaults instantiates a new MarketplaceUseCaseVersionLocale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCaseVersionLocale) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCaseVersionLocale) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCaseVersionLocale) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCaseVersionLocale) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCaseVersionLocale) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCaseVersionLocale) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *MarketplaceUseCaseVersionLocale) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceUseCaseVersionLocale) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceUseCaseVersionLocale) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceUseCaseVersionLocale) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocale

`func (o *MarketplaceUseCaseVersionLocale) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MarketplaceUseCaseVersionLocale) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MarketplaceUseCaseVersionLocale) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MarketplaceUseCaseVersionLocale) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


