# MarketplaceUseCaseDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "marketplace.UseCaseDependency"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "marketplace.UseCaseDependency"]
**Name** | Pointer to **string** | The string field to hold the key name. | [optional] 
**Version** | Pointer to **string** | The string field to hold the value. | [optional] 

## Methods

### NewMarketplaceUseCaseDependency

`func NewMarketplaceUseCaseDependency(classId string, objectType string, ) *MarketplaceUseCaseDependency`

NewMarketplaceUseCaseDependency instantiates a new MarketplaceUseCaseDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceUseCaseDependencyWithDefaults

`func NewMarketplaceUseCaseDependencyWithDefaults() *MarketplaceUseCaseDependency`

NewMarketplaceUseCaseDependencyWithDefaults instantiates a new MarketplaceUseCaseDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MarketplaceUseCaseDependency) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MarketplaceUseCaseDependency) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MarketplaceUseCaseDependency) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MarketplaceUseCaseDependency) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MarketplaceUseCaseDependency) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MarketplaceUseCaseDependency) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *MarketplaceUseCaseDependency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceUseCaseDependency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceUseCaseDependency) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceUseCaseDependency) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *MarketplaceUseCaseDependency) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketplaceUseCaseDependency) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketplaceUseCaseDependency) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketplaceUseCaseDependency) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


