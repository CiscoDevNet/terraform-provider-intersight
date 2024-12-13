# HciExpiredLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ExpiredLicense"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ExpiredLicense"]
**Category** | Pointer to **string** | The category of the license that has expired. | [optional] [readonly] 
**ExpiryDate** | Pointer to **time.Time** | The expiry date of the license that has expired. | [optional] [readonly] 
**ExtId** | Pointer to **string** | The id of the license that has expired. | [optional] [readonly] 
**LicenseId** | Pointer to **string** | The id of the license that has expired. | [optional] [readonly] 
**Meter** | Pointer to **string** | The meter of the license. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the license that has expired. | [optional] [readonly] 
**SubCategory** | Pointer to **string** | The subCategory of the license that has expired. | [optional] [readonly] 
**UsedQuantity** | Pointer to **int64** | The used quantity of the license. | [optional] [readonly] 

## Methods

### NewHciExpiredLicense

`func NewHciExpiredLicense(classId string, objectType string, ) *HciExpiredLicense`

NewHciExpiredLicense instantiates a new HciExpiredLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciExpiredLicenseWithDefaults

`func NewHciExpiredLicenseWithDefaults() *HciExpiredLicense`

NewHciExpiredLicenseWithDefaults instantiates a new HciExpiredLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciExpiredLicense) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciExpiredLicense) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciExpiredLicense) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciExpiredLicense) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciExpiredLicense) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciExpiredLicense) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HciExpiredLicense) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HciExpiredLicense) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HciExpiredLicense) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HciExpiredLicense) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetExpiryDate

`func (o *HciExpiredLicense) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *HciExpiredLicense) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *HciExpiredLicense) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *HciExpiredLicense) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetExtId

`func (o *HciExpiredLicense) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *HciExpiredLicense) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *HciExpiredLicense) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *HciExpiredLicense) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetLicenseId

`func (o *HciExpiredLicense) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *HciExpiredLicense) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *HciExpiredLicense) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *HciExpiredLicense) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetMeter

`func (o *HciExpiredLicense) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *HciExpiredLicense) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *HciExpiredLicense) SetMeter(v string)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *HciExpiredLicense) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetName

`func (o *HciExpiredLicense) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciExpiredLicense) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciExpiredLicense) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciExpiredLicense) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubCategory

`func (o *HciExpiredLicense) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *HciExpiredLicense) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *HciExpiredLicense) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *HciExpiredLicense) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetUsedQuantity

`func (o *HciExpiredLicense) GetUsedQuantity() int64`

GetUsedQuantity returns the UsedQuantity field if non-nil, zero value otherwise.

### GetUsedQuantityOk

`func (o *HciExpiredLicense) GetUsedQuantityOk() (*int64, bool)`

GetUsedQuantityOk returns a tuple with the UsedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuantity

`func (o *HciExpiredLicense) SetUsedQuantity(v int64)`

SetUsedQuantity sets UsedQuantity field to given value.

### HasUsedQuantity

`func (o *HciExpiredLicense) HasUsedQuantity() bool`

HasUsedQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


