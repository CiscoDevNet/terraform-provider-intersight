# LicenseIwoLicenseCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.IwoLicenseCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.IwoLicenseCount"]
**VmLicenseCount** | Pointer to **int64** | The total number of devices claimed in the Intersight account. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseIwoLicenseCount

`func NewLicenseIwoLicenseCount(classId string, objectType string, ) *LicenseIwoLicenseCount`

NewLicenseIwoLicenseCount instantiates a new LicenseIwoLicenseCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIwoLicenseCountWithDefaults

`func NewLicenseIwoLicenseCountWithDefaults() *LicenseIwoLicenseCount`

NewLicenseIwoLicenseCountWithDefaults instantiates a new LicenseIwoLicenseCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseIwoLicenseCount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseIwoLicenseCount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseIwoLicenseCount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseIwoLicenseCount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseIwoLicenseCount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseIwoLicenseCount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVmLicenseCount

`func (o *LicenseIwoLicenseCount) GetVmLicenseCount() int64`

GetVmLicenseCount returns the VmLicenseCount field if non-nil, zero value otherwise.

### GetVmLicenseCountOk

`func (o *LicenseIwoLicenseCount) GetVmLicenseCountOk() (*int64, bool)`

GetVmLicenseCountOk returns a tuple with the VmLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmLicenseCount

`func (o *LicenseIwoLicenseCount) SetVmLicenseCount(v int64)`

SetVmLicenseCount sets VmLicenseCount field to given value.

### HasVmLicenseCount

`func (o *LicenseIwoLicenseCount) HasVmLicenseCount() bool`

HasVmLicenseCount returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseIwoLicenseCount) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseIwoLicenseCount) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseIwoLicenseCount) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseIwoLicenseCount) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


