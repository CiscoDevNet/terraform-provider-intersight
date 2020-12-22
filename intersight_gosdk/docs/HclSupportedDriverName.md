# HclSupportedDriverName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.SupportedDriverName"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.SupportedDriverName"]
**OsVendor** | Pointer to **string** | Vendor distributing the Operating System. | [optional] 
**OsVersion** | Pointer to **string** | Version of the Operating System. | [optional] 
**ProductList** | Pointer to [**[]HclProduct**](HclProduct.md) |  | [optional] 

## Methods

### NewHclSupportedDriverName

`func NewHclSupportedDriverName(classId string, objectType string, ) *HclSupportedDriverName`

NewHclSupportedDriverName instantiates a new HclSupportedDriverName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclSupportedDriverNameWithDefaults

`func NewHclSupportedDriverNameWithDefaults() *HclSupportedDriverName`

NewHclSupportedDriverNameWithDefaults instantiates a new HclSupportedDriverName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclSupportedDriverName) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclSupportedDriverName) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclSupportedDriverName) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclSupportedDriverName) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclSupportedDriverName) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclSupportedDriverName) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOsVendor

`func (o *HclSupportedDriverName) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *HclSupportedDriverName) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *HclSupportedDriverName) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *HclSupportedDriverName) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *HclSupportedDriverName) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *HclSupportedDriverName) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *HclSupportedDriverName) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *HclSupportedDriverName) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProductList

`func (o *HclSupportedDriverName) GetProductList() []HclProduct`

GetProductList returns the ProductList field if non-nil, zero value otherwise.

### GetProductListOk

`func (o *HclSupportedDriverName) GetProductListOk() (*[]HclProduct, bool)`

GetProductListOk returns a tuple with the ProductList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductList

`func (o *HclSupportedDriverName) SetProductList(v []HclProduct)`

SetProductList sets ProductList field to given value.

### HasProductList

`func (o *HclSupportedDriverName) HasProductList() bool`

HasProductList returns a boolean if a field has been set.

### SetProductListNil

`func (o *HclSupportedDriverName) SetProductListNil(b bool)`

 SetProductListNil sets the value for ProductList to be an explicit nil

### UnsetProductList
`func (o *HclSupportedDriverName) UnsetProductList()`

UnsetProductList ensures that no value is present for ProductList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


