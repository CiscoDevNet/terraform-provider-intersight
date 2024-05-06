# MarketplaceUseCaseLocale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "marketplace.UseCaseLocale"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "marketplace.UseCaseLocale"]
**Automations** | Pointer to [**[]MarketplaceUseCaseAutomation**](MarketplaceUseCaseAutomation.md) |  | [optional] 
**Contents** | Pointer to **string** | The string field to hold the contents value. | [optional] 
**Description** | Pointer to **string** | The string field to hold the description value. | [optional] 
**Icon** | Pointer to **string** | A base64-encoded image for the use case. | [optional] 
**Locale** | Pointer to **string** | The string field to hold the locale. | [optional] 
**Name** | Pointer to **string** | The string field to hold the name value. | [optional] 
**Summary** | Pointer to **string** | The string field to hold the summary value. | [optional] 

## Methods

### NewMarketplaceUseCaseLocale

`func NewMarketplaceUseCaseLocale(classId string, objectType string, ) *MarketplaceUseCaseLocale`

NewMarketplaceUseCaseLocale instantiates a new MarketplaceUseCaseLocale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseLocaleWithDefaults

`func NewMarketplaceUseCaseLocaleWithDefaults() *MarketplaceUseCaseLocale`

NewMarketplaceUseCaseLocaleWithDefaults instantiates a new MarketplaceUseCaseLocale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCaseLocale) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCaseLocale) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCaseLocale) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCaseLocale) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCaseLocale) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCaseLocale) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutomations

`func (o *MarketplaceUseCaseLocale) GetAutomations() []MarketplaceUseCaseAutomation`

GetAutomations returns the Automations field if non-nil, zero value otherwise.

### GetAutomationsOk

`func (o *MarketplaceUseCaseLocale) GetAutomationsOk() (*[]MarketplaceUseCaseAutomation, bool)`

GetAutomationsOk returns a tuple with the Automations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomations

`func (o *MarketplaceUseCaseLocale) SetAutomations(v []MarketplaceUseCaseAutomation)`

SetAutomations sets Automations field to given value.

### HasAutomations

`func (o *MarketplaceUseCaseLocale) HasAutomations() bool`

HasAutomations returns a boolean if a field has been set.

### SetAutomationsNil

`func (o *MarketplaceUseCaseLocale) SetAutomationsNil(b bool)`

 SetAutomationsNil sets the value for Automations to be an explicit nil

### UnsetAutomations
`func (o *MarketplaceUseCaseLocale) UnsetAutomations()`

UnsetAutomations ensures that no value is present for Automations, not even an explicit nil
### GetContents

`func (o *MarketplaceUseCaseLocale) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *MarketplaceUseCaseLocale) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *MarketplaceUseCaseLocale) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *MarketplaceUseCaseLocale) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplaceUseCaseLocale) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceUseCaseLocale) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceUseCaseLocale) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceUseCaseLocale) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *MarketplaceUseCaseLocale) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *MarketplaceUseCaseLocale) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *MarketplaceUseCaseLocale) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *MarketplaceUseCaseLocale) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLocale

`func (o *MarketplaceUseCaseLocale) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MarketplaceUseCaseLocale) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MarketplaceUseCaseLocale) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MarketplaceUseCaseLocale) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *MarketplaceUseCaseLocale) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceUseCaseLocale) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceUseCaseLocale) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceUseCaseLocale) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *MarketplaceUseCaseLocale) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *MarketplaceUseCaseLocale) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *MarketplaceUseCaseLocale) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *MarketplaceUseCaseLocale) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


