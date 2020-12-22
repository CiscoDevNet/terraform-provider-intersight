# HclSupportedDriverNameAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.SupportedDriverName"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.SupportedDriverName"]
**OsVendor** | Pointer to **string** | Vendor distributing the Operating System. | [optional] 
**OsVersion** | Pointer to **string** | Version of the Operating System. | [optional] 
**ProductList** | Pointer to [**[]HclProduct**](HclProduct.md) |  | [optional] 

## Methods

### NewHclSupportedDriverNameAllOf

`func NewHclSupportedDriverNameAllOf(classId string, objectType string, ) *HclSupportedDriverNameAllOf`

NewHclSupportedDriverNameAllOf instantiates a new HclSupportedDriverNameAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclSupportedDriverNameAllOfWithDefaults

`func NewHclSupportedDriverNameAllOfWithDefaults() *HclSupportedDriverNameAllOf`

NewHclSupportedDriverNameAllOfWithDefaults instantiates a new HclSupportedDriverNameAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclSupportedDriverNameAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclSupportedDriverNameAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclSupportedDriverNameAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclSupportedDriverNameAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclSupportedDriverNameAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclSupportedDriverNameAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetProductListNil

`func (o *HclSupportedDriverNameAllOf) SetProductListNil(b bool)`

 SetProductListNil sets the value for ProductList to be an explicit nil

### UnsetProductList
`func (o *HclSupportedDriverNameAllOf) UnsetProductList()`

UnsetProductList ensures that no value is present for ProductList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


