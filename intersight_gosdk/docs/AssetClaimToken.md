# AssetClaimToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ClaimToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ClaimToken"]
**Claim** | Pointer to **string** | The claim generated to configure the device to connect and claim to the user account. | [optional] 
**ExpirationDate** | Pointer to **time.Time** | The expiry time of the claim token, after this time any generated claims from this token are no longer valid. | [optional] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewAssetClaimToken

`func NewAssetClaimToken(classId string, objectType string, ) *AssetClaimToken`

NewAssetClaimToken instantiates a new AssetClaimToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetClaimTokenWithDefaults

`func NewAssetClaimTokenWithDefaults() *AssetClaimToken`

NewAssetClaimTokenWithDefaults instantiates a new AssetClaimToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetClaimToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetClaimToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetClaimToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetClaimToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetClaimToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetClaimToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClaim

`func (o *AssetClaimToken) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *AssetClaimToken) GetClaimOk() (*string, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *AssetClaimToken) SetClaim(v string)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *AssetClaimToken) HasClaim() bool`

HasClaim returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AssetClaimToken) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AssetClaimToken) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AssetClaimToken) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AssetClaimToken) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetUser

`func (o *AssetClaimToken) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AssetClaimToken) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AssetClaimToken) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *AssetClaimToken) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AssetClaimToken) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AssetClaimToken) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


