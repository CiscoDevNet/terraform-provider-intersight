# AssetNewRelicCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.NewRelicCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.NewRelicCredential"]
**GraphQlCredential** | Pointer to [**AssetApiKeyCredential**](AssetApiKeyCredential.md) |  | [optional] 
**RestApiCredential** | Pointer to [**AssetApiKeyCredential**](AssetApiKeyCredential.md) |  | [optional] 

## Methods

### NewAssetNewRelicCredential

`func NewAssetNewRelicCredential(classId string, objectType string, ) *AssetNewRelicCredential`

NewAssetNewRelicCredential instantiates a new AssetNewRelicCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetNewRelicCredentialWithDefaults

`func NewAssetNewRelicCredentialWithDefaults() *AssetNewRelicCredential`

NewAssetNewRelicCredentialWithDefaults instantiates a new AssetNewRelicCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetNewRelicCredential) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetNewRelicCredential) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetNewRelicCredential) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetNewRelicCredential) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetNewRelicCredential) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetNewRelicCredential) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGraphQlCredential

`func (o *AssetNewRelicCredential) GetGraphQlCredential() AssetApiKeyCredential`

GetGraphQlCredential returns the GraphQlCredential field if non-nil, zero value otherwise.

### GetGraphQlCredentialOk

`func (o *AssetNewRelicCredential) GetGraphQlCredentialOk() (*AssetApiKeyCredential, bool)`

GetGraphQlCredentialOk returns a tuple with the GraphQlCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphQlCredential

`func (o *AssetNewRelicCredential) SetGraphQlCredential(v AssetApiKeyCredential)`

SetGraphQlCredential sets GraphQlCredential field to given value.

### HasGraphQlCredential

`func (o *AssetNewRelicCredential) HasGraphQlCredential() bool`

HasGraphQlCredential returns a boolean if a field has been set.

### GetRestApiCredential

`func (o *AssetNewRelicCredential) GetRestApiCredential() AssetApiKeyCredential`

GetRestApiCredential returns the RestApiCredential field if non-nil, zero value otherwise.

### GetRestApiCredentialOk

`func (o *AssetNewRelicCredential) GetRestApiCredentialOk() (*AssetApiKeyCredential, bool)`

GetRestApiCredentialOk returns a tuple with the RestApiCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestApiCredential

`func (o *AssetNewRelicCredential) SetRestApiCredential(v AssetApiKeyCredential)`

SetRestApiCredential sets RestApiCredential field to given value.

### HasRestApiCredential

`func (o *AssetNewRelicCredential) HasRestApiCredential() bool`

HasRestApiCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


