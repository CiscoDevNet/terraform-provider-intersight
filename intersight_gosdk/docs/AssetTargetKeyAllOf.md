# AssetTargetKeyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.TargetKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.TargetKey"]
**IsPrivateKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;privateKey&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetTargetKeyAllOf

`func NewAssetTargetKeyAllOf(classId string, objectType string, ) *AssetTargetKeyAllOf`

NewAssetTargetKeyAllOf instantiates a new AssetTargetKeyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTargetKeyAllOfWithDefaults

`func NewAssetTargetKeyAllOfWithDefaults() *AssetTargetKeyAllOf`

NewAssetTargetKeyAllOfWithDefaults instantiates a new AssetTargetKeyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetTargetKeyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetTargetKeyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetTargetKeyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetTargetKeyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetTargetKeyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetTargetKeyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPrivateKeySet

`func (o *AssetTargetKeyAllOf) GetIsPrivateKeySet() bool`

GetIsPrivateKeySet returns the IsPrivateKeySet field if non-nil, zero value otherwise.

### GetIsPrivateKeySetOk

`func (o *AssetTargetKeyAllOf) GetIsPrivateKeySetOk() (*bool, bool)`

GetIsPrivateKeySetOk returns a tuple with the IsPrivateKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateKeySet

`func (o *AssetTargetKeyAllOf) SetIsPrivateKeySet(v bool)`

SetIsPrivateKeySet sets IsPrivateKeySet field to given value.

### HasIsPrivateKeySet

`func (o *AssetTargetKeyAllOf) HasIsPrivateKeySet() bool`

HasIsPrivateKeySet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


