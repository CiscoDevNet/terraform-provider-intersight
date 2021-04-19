# AssetClaimSignatureAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Signature** | Pointer to **string** | The result of signing the deviceId appended with the timeStamp fields with the devices private key. | [optional] 
**TimeStamp** | Pointer to **time.Time** | The time at which the signature was generated. Date is accurate to Intersights clock. Used to expire the signature. | [optional] 

## Methods

### NewAssetClaimSignatureAllOf

`func NewAssetClaimSignatureAllOf(classId string, objectType string, ) *AssetClaimSignatureAllOf`

NewAssetClaimSignatureAllOf instantiates a new AssetClaimSignatureAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetClaimSignatureAllOfWithDefaults

`func NewAssetClaimSignatureAllOfWithDefaults() *AssetClaimSignatureAllOf`

NewAssetClaimSignatureAllOfWithDefaults instantiates a new AssetClaimSignatureAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetClaimSignatureAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetClaimSignatureAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetClaimSignatureAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetClaimSignatureAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetClaimSignatureAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetClaimSignatureAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSignature

`func (o *AssetClaimSignatureAllOf) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AssetClaimSignatureAllOf) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AssetClaimSignatureAllOf) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AssetClaimSignatureAllOf) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTimeStamp

`func (o *AssetClaimSignatureAllOf) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AssetClaimSignatureAllOf) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AssetClaimSignatureAllOf) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *AssetClaimSignatureAllOf) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


