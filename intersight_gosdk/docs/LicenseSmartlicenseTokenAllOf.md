# LicenseSmartlicenseTokenAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.SmartlicenseToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.SmartlicenseToken"]
**Token** | Pointer to **string** | Smart license registration token. | [optional] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseSmartlicenseTokenAllOf

`func NewLicenseSmartlicenseTokenAllOf(classId string, objectType string, ) *LicenseSmartlicenseTokenAllOf`

NewLicenseSmartlicenseTokenAllOf instantiates a new LicenseSmartlicenseTokenAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSmartlicenseTokenAllOfWithDefaults

`func NewLicenseSmartlicenseTokenAllOfWithDefaults() *LicenseSmartlicenseTokenAllOf`

NewLicenseSmartlicenseTokenAllOfWithDefaults instantiates a new LicenseSmartlicenseTokenAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseSmartlicenseTokenAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseSmartlicenseTokenAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseSmartlicenseTokenAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseSmartlicenseTokenAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseSmartlicenseTokenAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseSmartlicenseTokenAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetToken

`func (o *LicenseSmartlicenseTokenAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LicenseSmartlicenseTokenAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LicenseSmartlicenseTokenAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LicenseSmartlicenseTokenAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseSmartlicenseTokenAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseSmartlicenseTokenAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseSmartlicenseTokenAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseSmartlicenseTokenAllOf) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


