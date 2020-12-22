# AssetClientCertificateCredentialAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ClientCertificateCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ClientCertificateCredential"]
**ClientCertificate** | Pointer to **string** | PEM encoded x509 client certificate used to authenticate with the target. | [optional] 
**ClientKey** | Pointer to **string** | PEM encoded private key used to authenticate with the target. | [optional] 
**IsClientKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;clientKey&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetClientCertificateCredentialAllOf

`func NewAssetClientCertificateCredentialAllOf(classId string, objectType string, ) *AssetClientCertificateCredentialAllOf`

NewAssetClientCertificateCredentialAllOf instantiates a new AssetClientCertificateCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetClientCertificateCredentialAllOfWithDefaults

`func NewAssetClientCertificateCredentialAllOfWithDefaults() *AssetClientCertificateCredentialAllOf`

NewAssetClientCertificateCredentialAllOfWithDefaults instantiates a new AssetClientCertificateCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetClientCertificateCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetClientCertificateCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetClientCertificateCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetClientCertificateCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetClientCertificateCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetClientCertificateCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientCertificate

`func (o *AssetClientCertificateCredentialAllOf) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *AssetClientCertificateCredentialAllOf) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *AssetClientCertificateCredentialAllOf) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *AssetClientCertificateCredentialAllOf) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientKey

`func (o *AssetClientCertificateCredentialAllOf) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *AssetClientCertificateCredentialAllOf) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *AssetClientCertificateCredentialAllOf) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *AssetClientCertificateCredentialAllOf) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetIsClientKeySet

`func (o *AssetClientCertificateCredentialAllOf) GetIsClientKeySet() bool`

GetIsClientKeySet returns the IsClientKeySet field if non-nil, zero value otherwise.

### GetIsClientKeySetOk

`func (o *AssetClientCertificateCredentialAllOf) GetIsClientKeySetOk() (*bool, bool)`

GetIsClientKeySetOk returns a tuple with the IsClientKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientKeySet

`func (o *AssetClientCertificateCredentialAllOf) SetIsClientKeySet(v bool)`

SetIsClientKeySet sets IsClientKeySet field to given value.

### HasIsClientKeySet

`func (o *AssetClientCertificateCredentialAllOf) HasIsClientKeySet() bool`

HasIsClientKeySet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


