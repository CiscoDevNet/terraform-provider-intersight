# AssetApiKeyCredentialAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ApiKeyCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ApiKeyCredential"]
**ApiKey** | Pointer to **string** | This a secret API key which can be used for authentication purposes for different targets like Azure Enterprise Agreement. | [optional] 
**IsApiKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;apiKey&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetApiKeyCredentialAllOf

`func NewAssetApiKeyCredentialAllOf(classId string, objectType string, ) *AssetApiKeyCredentialAllOf`

NewAssetApiKeyCredentialAllOf instantiates a new AssetApiKeyCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetApiKeyCredentialAllOfWithDefaults

`func NewAssetApiKeyCredentialAllOfWithDefaults() *AssetApiKeyCredentialAllOf`

NewAssetApiKeyCredentialAllOfWithDefaults instantiates a new AssetApiKeyCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetApiKeyCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetApiKeyCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetApiKeyCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetApiKeyCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetApiKeyCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetApiKeyCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiKey

`func (o *AssetApiKeyCredentialAllOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AssetApiKeyCredentialAllOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AssetApiKeyCredentialAllOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AssetApiKeyCredentialAllOf) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetIsApiKeySet

`func (o *AssetApiKeyCredentialAllOf) GetIsApiKeySet() bool`

GetIsApiKeySet returns the IsApiKeySet field if non-nil, zero value otherwise.

### GetIsApiKeySetOk

`func (o *AssetApiKeyCredentialAllOf) GetIsApiKeySetOk() (*bool, bool)`

GetIsApiKeySetOk returns a tuple with the IsApiKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiKeySet

`func (o *AssetApiKeyCredentialAllOf) SetIsApiKeySet(v bool)`

SetIsApiKeySet sets IsApiKeySet field to given value.

### HasIsApiKeySet

`func (o *AssetApiKeyCredentialAllOf) HasIsApiKeySet() bool`

HasIsApiKeySet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


