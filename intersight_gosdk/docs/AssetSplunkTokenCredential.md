# AssetSplunkTokenCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.SplunkTokenCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.SplunkTokenCredential"]
**IsTokenSet** | Pointer to **bool** | Indicates whether the value of the &#39;token&#39; property has been set. | [optional] [readonly] [default to false]
**Token** | Pointer to **string** | The Splunk token used to authenticate with a managed target. | [optional] 

## Methods

### NewAssetSplunkTokenCredential

`func NewAssetSplunkTokenCredential(classId string, objectType string, ) *AssetSplunkTokenCredential`

NewAssetSplunkTokenCredential instantiates a new AssetSplunkTokenCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetSplunkTokenCredentialWithDefaults

`func NewAssetSplunkTokenCredentialWithDefaults() *AssetSplunkTokenCredential`

NewAssetSplunkTokenCredentialWithDefaults instantiates a new AssetSplunkTokenCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetSplunkTokenCredential) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetSplunkTokenCredential) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetSplunkTokenCredential) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetSplunkTokenCredential) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetSplunkTokenCredential) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetSplunkTokenCredential) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsTokenSet

`func (o *AssetSplunkTokenCredential) GetIsTokenSet() bool`

GetIsTokenSet returns the IsTokenSet field if non-nil, zero value otherwise.

### GetIsTokenSetOk

`func (o *AssetSplunkTokenCredential) GetIsTokenSetOk() (*bool, bool)`

GetIsTokenSetOk returns a tuple with the IsTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenSet

`func (o *AssetSplunkTokenCredential) SetIsTokenSet(v bool)`

SetIsTokenSet sets IsTokenSet field to given value.

### HasIsTokenSet

`func (o *AssetSplunkTokenCredential) HasIsTokenSet() bool`

HasIsTokenSet returns a boolean if a field has been set.

### GetToken

`func (o *AssetSplunkTokenCredential) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AssetSplunkTokenCredential) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AssetSplunkTokenCredential) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AssetSplunkTokenCredential) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


