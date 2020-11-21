# VirtualizationProductInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.ProductInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.ProductInfo"]
**ProductName** | Pointer to **string** | Commercial product name. For example, VMware ESXi. | [optional] 
**ProductType** | Pointer to **string** | Product name provided by the vendor. For example, embeddedEsx. | [optional] 
**ProductVendor** | Pointer to **string** | Commercial vendor name. For example, VMware Inc. | [optional] 
**Version** | Pointer to **string** | Hypervisor version running on the system. | [optional] 

## Methods

### NewVirtualizationProductInfo

`func NewVirtualizationProductInfo(classId string, objectType string, ) *VirtualizationProductInfo`

NewVirtualizationProductInfo instantiates a new VirtualizationProductInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationProductInfoWithDefaults

`func NewVirtualizationProductInfoWithDefaults() *VirtualizationProductInfo`

NewVirtualizationProductInfoWithDefaults instantiates a new VirtualizationProductInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationProductInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationProductInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationProductInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationProductInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationProductInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationProductInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProductName

`func (o *VirtualizationProductInfo) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *VirtualizationProductInfo) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *VirtualizationProductInfo) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *VirtualizationProductInfo) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductType

`func (o *VirtualizationProductInfo) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *VirtualizationProductInfo) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *VirtualizationProductInfo) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *VirtualizationProductInfo) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProductVendor

`func (o *VirtualizationProductInfo) GetProductVendor() string`

GetProductVendor returns the ProductVendor field if non-nil, zero value otherwise.

### GetProductVendorOk

`func (o *VirtualizationProductInfo) GetProductVendorOk() (*string, bool)`

GetProductVendorOk returns a tuple with the ProductVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVendor

`func (o *VirtualizationProductInfo) SetProductVendor(v string)`

SetProductVendor sets ProductVendor field to given value.

### HasProductVendor

`func (o *VirtualizationProductInfo) HasProductVendor() bool`

HasProductVendor returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualizationProductInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualizationProductInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualizationProductInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualizationProductInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


