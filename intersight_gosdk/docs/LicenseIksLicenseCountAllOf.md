# LicenseIksLicenseCountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.IksLicenseCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.IksLicenseCount"]
**AdvantageCount** | Pointer to **int64** | The total number of devices claimed in the IKS Advantage tier. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseIksLicenseCountAllOf

`func NewLicenseIksLicenseCountAllOf(classId string, objectType string, ) *LicenseIksLicenseCountAllOf`

NewLicenseIksLicenseCountAllOf instantiates a new LicenseIksLicenseCountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIksLicenseCountAllOfWithDefaults

`func NewLicenseIksLicenseCountAllOfWithDefaults() *LicenseIksLicenseCountAllOf`

NewLicenseIksLicenseCountAllOfWithDefaults instantiates a new LicenseIksLicenseCountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseIksLicenseCountAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseIksLicenseCountAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseIksLicenseCountAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseIksLicenseCountAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseIksLicenseCountAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseIksLicenseCountAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdvantageCount

`func (o *LicenseIksLicenseCountAllOf) GetAdvantageCount() int64`

GetAdvantageCount returns the AdvantageCount field if non-nil, zero value otherwise.

### GetAdvantageCountOk

`func (o *LicenseIksLicenseCountAllOf) GetAdvantageCountOk() (*int64, bool)`

GetAdvantageCountOk returns a tuple with the AdvantageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvantageCount

`func (o *LicenseIksLicenseCountAllOf) SetAdvantageCount(v int64)`

SetAdvantageCount sets AdvantageCount field to given value.

### HasAdvantageCount

`func (o *LicenseIksLicenseCountAllOf) HasAdvantageCount() bool`

HasAdvantageCount returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseIksLicenseCountAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseIksLicenseCountAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseIksLicenseCountAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseIksLicenseCountAllOf) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


