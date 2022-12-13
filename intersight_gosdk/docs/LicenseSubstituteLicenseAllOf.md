# LicenseSubstituteLicenseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.SubstituteLicense"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.SubstituteLicense"]
**SubstitutedLicense** | Pointer to **string** | The substitute license that is used. | [optional] [readonly] 
**SubstitutedQuantity** | Pointer to **int64** | The number of substitute licenses that are used. | [optional] [readonly] 
**SubstitutionType** | Pointer to **string** | The substitution from lower or from higher tier. | [optional] [readonly] 

## Methods

### NewLicenseSubstituteLicenseAllOf

`func NewLicenseSubstituteLicenseAllOf(classId string, objectType string, ) *LicenseSubstituteLicenseAllOf`

NewLicenseSubstituteLicenseAllOf instantiates a new LicenseSubstituteLicenseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSubstituteLicenseAllOfWithDefaults

`func NewLicenseSubstituteLicenseAllOfWithDefaults() *LicenseSubstituteLicenseAllOf`

NewLicenseSubstituteLicenseAllOfWithDefaults instantiates a new LicenseSubstituteLicenseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseSubstituteLicenseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseSubstituteLicenseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseSubstituteLicenseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseSubstituteLicenseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseSubstituteLicenseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseSubstituteLicenseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSubstitutedLicense

`func (o *LicenseSubstituteLicenseAllOf) GetSubstitutedLicense() string`

GetSubstitutedLicense returns the SubstitutedLicense field if non-nil, zero value otherwise.

### GetSubstitutedLicenseOk

`func (o *LicenseSubstituteLicenseAllOf) GetSubstitutedLicenseOk() (*string, bool)`

GetSubstitutedLicenseOk returns a tuple with the SubstitutedLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutedLicense

`func (o *LicenseSubstituteLicenseAllOf) SetSubstitutedLicense(v string)`

SetSubstitutedLicense sets SubstitutedLicense field to given value.

### HasSubstitutedLicense

`func (o *LicenseSubstituteLicenseAllOf) HasSubstitutedLicense() bool`

HasSubstitutedLicense returns a boolean if a field has been set.

### GetSubstitutedQuantity

`func (o *LicenseSubstituteLicenseAllOf) GetSubstitutedQuantity() int64`

GetSubstitutedQuantity returns the SubstitutedQuantity field if non-nil, zero value otherwise.

### GetSubstitutedQuantityOk

`func (o *LicenseSubstituteLicenseAllOf) GetSubstitutedQuantityOk() (*int64, bool)`

GetSubstitutedQuantityOk returns a tuple with the SubstitutedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutedQuantity

`func (o *LicenseSubstituteLicenseAllOf) SetSubstitutedQuantity(v int64)`

SetSubstitutedQuantity sets SubstitutedQuantity field to given value.

### HasSubstitutedQuantity

`func (o *LicenseSubstituteLicenseAllOf) HasSubstitutedQuantity() bool`

HasSubstitutedQuantity returns a boolean if a field has been set.

### GetSubstitutionType

`func (o *LicenseSubstituteLicenseAllOf) GetSubstitutionType() string`

GetSubstitutionType returns the SubstitutionType field if non-nil, zero value otherwise.

### GetSubstitutionTypeOk

`func (o *LicenseSubstituteLicenseAllOf) GetSubstitutionTypeOk() (*string, bool)`

GetSubstitutionTypeOk returns a tuple with the SubstitutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutionType

`func (o *LicenseSubstituteLicenseAllOf) SetSubstitutionType(v string)`

SetSubstitutionType sets SubstitutionType field to given value.

### HasSubstitutionType

`func (o *LicenseSubstituteLicenseAllOf) HasSubstitutionType() bool`

HasSubstitutionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


