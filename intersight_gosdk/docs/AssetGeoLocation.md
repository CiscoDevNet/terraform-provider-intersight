# AssetGeoLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.GeoLocation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.GeoLocation"]
**Address** | Pointer to [**NullableCommPhysicalAddress**](CommPhysicalAddress.md) |  | [optional] 
**Coordinates** | Pointer to [**NullableCommGeoPoint**](CommGeoPoint.md) |  | [optional] 
**Name** | Pointer to **string** | A user provided name for the location. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewAssetGeoLocation

`func NewAssetGeoLocation(classId string, objectType string, ) *AssetGeoLocation`

NewAssetGeoLocation instantiates a new AssetGeoLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetGeoLocationWithDefaults

`func NewAssetGeoLocationWithDefaults() *AssetGeoLocation`

NewAssetGeoLocationWithDefaults instantiates a new AssetGeoLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetGeoLocation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetGeoLocation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetGeoLocation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetGeoLocation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetGeoLocation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetGeoLocation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *AssetGeoLocation) GetAddress() CommPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AssetGeoLocation) GetAddressOk() (*CommPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AssetGeoLocation) SetAddress(v CommPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AssetGeoLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *AssetGeoLocation) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *AssetGeoLocation) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCoordinates

`func (o *AssetGeoLocation) GetCoordinates() CommGeoPoint`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *AssetGeoLocation) GetCoordinatesOk() (*CommGeoPoint, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *AssetGeoLocation) SetCoordinates(v CommGeoPoint)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *AssetGeoLocation) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### SetCoordinatesNil

`func (o *AssetGeoLocation) SetCoordinatesNil(b bool)`

 SetCoordinatesNil sets the value for Coordinates to be an explicit nil

### UnsetCoordinates
`func (o *AssetGeoLocation) UnsetCoordinates()`

UnsetCoordinates ensures that no value is present for Coordinates, not even an explicit nil
### GetName

`func (o *AssetGeoLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetGeoLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetGeoLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetGeoLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *AssetGeoLocation) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetGeoLocation) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetGeoLocation) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetGeoLocation) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *AssetGeoLocation) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AssetGeoLocation) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


