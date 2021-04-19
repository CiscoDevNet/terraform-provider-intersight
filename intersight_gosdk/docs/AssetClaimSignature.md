# AssetClaimSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Signature** | Pointer to **string** | The result of signing the deviceId appended with the timeStamp fields with the devices private key. | [optional] 
**TimeStamp** | Pointer to **time.Time** | The time at which the signature was generated. Date is accurate to Intersights clock. Used to expire the signature. | [optional] 

## Methods

### NewAssetClaimSignature

`func NewAssetClaimSignature(classId string, objectType string, ) *AssetClaimSignature`

NewAssetClaimSignature instantiates a new AssetClaimSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetClaimSignatureWithDefaults

`func NewAssetClaimSignatureWithDefaults() *AssetClaimSignature`

NewAssetClaimSignatureWithDefaults instantiates a new AssetClaimSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetClaimSignature) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetClaimSignature) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetClaimSignature) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetClaimSignature) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetClaimSignature) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetClaimSignature) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSignature

`func (o *AssetClaimSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AssetClaimSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AssetClaimSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AssetClaimSignature) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTimeStamp

`func (o *AssetClaimSignature) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AssetClaimSignature) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AssetClaimSignature) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *AssetClaimSignature) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


