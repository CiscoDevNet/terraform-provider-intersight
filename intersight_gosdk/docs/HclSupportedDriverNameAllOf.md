# HclSupportedDriverNameAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsVendor** | Pointer to **string** | Vendor distributing the Operating System. | [optional] 
**OsVersion** | Pointer to **string** | Version of the Operating System. | [optional] 
**ProductList** | Pointer to [**[]HclProduct**](hcl.Product.md) |  | [optional] 

## Methods

### NewHclSupportedDriverNameAllOf

`func NewHclSupportedDriverNameAllOf() *HclSupportedDriverNameAllOf`

NewHclSupportedDriverNameAllOf instantiates a new HclSupportedDriverNameAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclSupportedDriverNameAllOfWithDefaults

`func NewHclSupportedDriverNameAllOfWithDefaults() *HclSupportedDriverNameAllOf`

NewHclSupportedDriverNameAllOfWithDefaults instantiates a new HclSupportedDriverNameAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsVendor

`func (o *HclSupportedDriverNameAllOf) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *HclSupportedDriverNameAllOf) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *HclSupportedDriverNameAllOf) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *HclSupportedDriverNameAllOf) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *HclSupportedDriverNameAllOf) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *HclSupportedDriverNameAllOf) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *HclSupportedDriverNameAllOf) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *HclSupportedDriverNameAllOf) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProductList

`func (o *HclSupportedDriverNameAllOf) GetProductList() []HclProduct`

GetProductList returns the ProductList field if non-nil, zero value otherwise.

### GetProductListOk

`func (o *HclSupportedDriverNameAllOf) GetProductListOk() (*[]HclProduct, bool)`

GetProductListOk returns a tuple with the ProductList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductList

`func (o *HclSupportedDriverNameAllOf) SetProductList(v []HclProduct)`

SetProductList sets ProductList field to given value.

### HasProductList

`func (o *HclSupportedDriverNameAllOf) HasProductList() bool`

HasProductList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


