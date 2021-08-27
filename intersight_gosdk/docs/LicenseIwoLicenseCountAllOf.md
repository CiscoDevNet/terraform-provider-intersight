# LicenseIwoLicenseCountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.IwoLicenseCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.IwoLicenseCount"]
**VmLicenseCount** | Pointer to **int64** | The total number of devices claimed in the Intersight account. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseIwoLicenseCountAllOf

`func NewLicenseIwoLicenseCountAllOf(classId string, objectType string, ) *LicenseIwoLicenseCountAllOf`

NewLicenseIwoLicenseCountAllOf instantiates a new LicenseIwoLicenseCountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIwoLicenseCountAllOfWithDefaults

`func NewLicenseIwoLicenseCountAllOfWithDefaults() *LicenseIwoLicenseCountAllOf`

NewLicenseIwoLicenseCountAllOfWithDefaults instantiates a new LicenseIwoLicenseCountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseIwoLicenseCountAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseIwoLicenseCountAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseIwoLicenseCountAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseIwoLicenseCountAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseIwoLicenseCountAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseIwoLicenseCountAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVmLicenseCount

`func (o *LicenseIwoLicenseCountAllOf) GetVmLicenseCount() int64`

GetVmLicenseCount returns the VmLicenseCount field if non-nil, zero value otherwise.

### GetVmLicenseCountOk

`func (o *LicenseIwoLicenseCountAllOf) GetVmLicenseCountOk() (*int64, bool)`

GetVmLicenseCountOk returns a tuple with the VmLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmLicenseCount

`func (o *LicenseIwoLicenseCountAllOf) SetVmLicenseCount(v int64)`

SetVmLicenseCount sets VmLicenseCount field to given value.

### HasVmLicenseCount

`func (o *LicenseIwoLicenseCountAllOf) HasVmLicenseCount() bool`

HasVmLicenseCount returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseIwoLicenseCountAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseIwoLicenseCountAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseIwoLicenseCountAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseIwoLicenseCountAllOf) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


