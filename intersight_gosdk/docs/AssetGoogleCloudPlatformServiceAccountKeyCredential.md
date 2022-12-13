# AssetGoogleCloudPlatformServiceAccountKeyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.GoogleCloudPlatformServiceAccountKeyCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.GoogleCloudPlatformServiceAccountKeyCredential"]
**IsServiceAccountKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;serviceAccountKey&#39; property has been set. | [optional] [readonly] [default to false]
**KeyType** | Pointer to **string** | Google Cloud Platform (GCP) service account&#39;s key type. * &#x60;JSON&#x60; - JavaScript Object Notation (JSON) is a standard text-based format for representing structured data based on JavaScript object syntax. It is commonly used for transmitting data in web applications. * &#x60;P12&#x60; - PKCS #12 is an archive file format for storing many cryptography objects as a single file. It is commonly used to bundle a private key with its X.509 certificate or to bundle all the members of a chain of trust. | [optional] [default to "JSON"]
**ServiceAccountKey** | Pointer to **string** | Google Cloud Platform (GCP) service account&#39;s key file content in the format of the keyType specified. | [optional] 

## Methods

### NewAssetGoogleCloudPlatformServiceAccountKeyCredential

`func NewAssetGoogleCloudPlatformServiceAccountKeyCredential(classId string, objectType string, ) *AssetGoogleCloudPlatformServiceAccountKeyCredential`

NewAssetGoogleCloudPlatformServiceAccountKeyCredential instantiates a new AssetGoogleCloudPlatformServiceAccountKeyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetGoogleCloudPlatformServiceAccountKeyCredentialWithDefaults

`func NewAssetGoogleCloudPlatformServiceAccountKeyCredentialWithDefaults() *AssetGoogleCloudPlatformServiceAccountKeyCredential`

NewAssetGoogleCloudPlatformServiceAccountKeyCredentialWithDefaults instantiates a new AssetGoogleCloudPlatformServiceAccountKeyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsServiceAccountKeySet

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetIsServiceAccountKeySet() bool`

GetIsServiceAccountKeySet returns the IsServiceAccountKeySet field if non-nil, zero value otherwise.

### GetIsServiceAccountKeySetOk

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetIsServiceAccountKeySetOk() (*bool, bool)`

GetIsServiceAccountKeySetOk returns a tuple with the IsServiceAccountKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccountKeySet

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) SetIsServiceAccountKeySet(v bool)`

SetIsServiceAccountKeySet sets IsServiceAccountKeySet field to given value.

### HasIsServiceAccountKeySet

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) HasIsServiceAccountKeySet() bool`

HasIsServiceAccountKeySet returns a boolean if a field has been set.

### GetKeyType

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetServiceAccountKey

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *AssetGoogleCloudPlatformServiceAccountKeyCredential) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


