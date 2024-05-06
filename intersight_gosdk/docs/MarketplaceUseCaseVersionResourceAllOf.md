# MarketplaceUseCaseVersionResourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "marketplace.UseCaseVersionResource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "marketplace.UseCaseVersionResource"]
**ResourceId** | Pointer to **string** | A string ID for each use case. | [optional] 
**ResourceType** | Pointer to **string** | A string resource type for each use case. | [optional] 

## Methods

### NewMarketplaceUseCaseVersionResourceAllOf

`func NewMarketplaceUseCaseVersionResourceAllOf(classId string, objectType string, ) *MarketplaceUseCaseVersionResourceAllOf`

NewMarketplaceUseCaseVersionResourceAllOf instantiates a new MarketplaceUseCaseVersionResourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseVersionResourceAllOfWithDefaults

`func NewMarketplaceUseCaseVersionResourceAllOfWithDefaults() *MarketplaceUseCaseVersionResourceAllOf`

NewMarketplaceUseCaseVersionResourceAllOfWithDefaults instantiates a new MarketplaceUseCaseVersionResourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCaseVersionResourceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCaseVersionResourceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetResourceId

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *MarketplaceUseCaseVersionResourceAllOf) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *MarketplaceUseCaseVersionResourceAllOf) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *MarketplaceUseCaseVersionResourceAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *MarketplaceUseCaseVersionResourceAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *MarketplaceUseCaseVersionResourceAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


