# OauthAccessTokenAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "oauth.AccessToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "oauth.AccessToken"]
**ApiType** | Pointer to **string** | Type of OAuth Api. For example, Smart-licensing-API. * &#x60;Unknown&#x60; - Unknown is the default API type. * &#x60;SmartLicensing-API&#x60; - Smart licensing API type. * &#x60;CommerceEstimate-API&#x60; - Commerce Estimate API type. | [optional] [default to "Unknown"]
**Expiry** | Pointer to **time.Time** | The date and time when the access token expires. | [optional] [readonly] 
**Issuer** | Pointer to **string** | Issuer of OAuth access token. | [optional] [readonly] 
**RefreshExpiry** | Pointer to **time.Time** | The date and time when the refresh token expires. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewOauthAccessTokenAllOf

`func NewOauthAccessTokenAllOf(classId string, objectType string, ) *OauthAccessTokenAllOf`

NewOauthAccessTokenAllOf instantiates a new OauthAccessTokenAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthAccessTokenAllOfWithDefaults

`func NewOauthAccessTokenAllOfWithDefaults() *OauthAccessTokenAllOf`

NewOauthAccessTokenAllOfWithDefaults instantiates a new OauthAccessTokenAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OauthAccessTokenAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OauthAccessTokenAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OauthAccessTokenAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OauthAccessTokenAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OauthAccessTokenAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OauthAccessTokenAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiType

`func (o *OauthAccessTokenAllOf) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *OauthAccessTokenAllOf) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *OauthAccessTokenAllOf) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *OauthAccessTokenAllOf) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetExpiry

`func (o *OauthAccessTokenAllOf) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *OauthAccessTokenAllOf) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *OauthAccessTokenAllOf) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *OauthAccessTokenAllOf) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetIssuer

`func (o *OauthAccessTokenAllOf) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OauthAccessTokenAllOf) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OauthAccessTokenAllOf) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OauthAccessTokenAllOf) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetRefreshExpiry

`func (o *OauthAccessTokenAllOf) GetRefreshExpiry() time.Time`

GetRefreshExpiry returns the RefreshExpiry field if non-nil, zero value otherwise.

### GetRefreshExpiryOk

`func (o *OauthAccessTokenAllOf) GetRefreshExpiryOk() (*time.Time, bool)`

GetRefreshExpiryOk returns a tuple with the RefreshExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpiry

`func (o *OauthAccessTokenAllOf) SetRefreshExpiry(v time.Time)`

SetRefreshExpiry sets RefreshExpiry field to given value.

### HasRefreshExpiry

`func (o *OauthAccessTokenAllOf) HasRefreshExpiry() bool`

HasRefreshExpiry returns a boolean if a field has been set.

### GetAccount

`func (o *OauthAccessTokenAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OauthAccessTokenAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OauthAccessTokenAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OauthAccessTokenAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


