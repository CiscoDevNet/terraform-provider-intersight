# VirtualizationProductInfoAllOf

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

### NewVirtualizationProductInfoAllOf

`func NewVirtualizationProductInfoAllOf(classId string, objectType string, ) *VirtualizationProductInfoAllOf`

NewVirtualizationProductInfoAllOf instantiates a new VirtualizationProductInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationProductInfoAllOfWithDefaults

`func NewVirtualizationProductInfoAllOfWithDefaults() *VirtualizationProductInfoAllOf`

NewVirtualizationProductInfoAllOfWithDefaults instantiates a new VirtualizationProductInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationProductInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationProductInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationProductInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationProductInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationProductInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationProductInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProductName

`func (o *VirtualizationProductInfoAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *VirtualizationProductInfoAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *VirtualizationProductInfoAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *VirtualizationProductInfoAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductType

`func (o *VirtualizationProductInfoAllOf) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *VirtualizationProductInfoAllOf) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *VirtualizationProductInfoAllOf) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *VirtualizationProductInfoAllOf) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProductVendor

`func (o *VirtualizationProductInfoAllOf) GetProductVendor() string`

GetProductVendor returns the ProductVendor field if non-nil, zero value otherwise.

### GetProductVendorOk

`func (o *VirtualizationProductInfoAllOf) GetProductVendorOk() (*string, bool)`

GetProductVendorOk returns a tuple with the ProductVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVendor

`func (o *VirtualizationProductInfoAllOf) SetProductVendor(v string)`

SetProductVendor sets ProductVendor field to given value.

### HasProductVendor

`func (o *VirtualizationProductInfoAllOf) HasProductVendor() bool`

HasProductVendor returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualizationProductInfoAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualizationProductInfoAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualizationProductInfoAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualizationProductInfoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


