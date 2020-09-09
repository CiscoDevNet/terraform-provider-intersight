# HclExemptedCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **string** | Reason for the exemption. | [optional] 
**Name** | Pointer to **string** | A unique descriptive name of the exemption. | [optional] 
**OsVendor** | Pointer to **string** | Vendor of the Operating System. | [optional] 
**OsVersion** | Pointer to **string** | Version of the Operating system. | [optional] 
**ProcessorName** | Pointer to **string** | Name of the processor supported for the server. | [optional] 
**ProductModels** | Pointer to **[]string** |  | [optional] 
**ProductType** | Pointer to **string** | Type of the product/adapter say PT for Pass Through controllers. | [optional] 
**ServerPid** | Pointer to **string** | Three part ID representing the server model as returned by UCSM/CIMC XML APIs. | [optional] 
**UcsVersion** | Pointer to **string** | Version of the UCS software. | [optional] 
**VersionType** | Pointer to **string** | Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. | [optional] 

## Methods

### NewHclExemptedCatalogAllOf

`func NewHclExemptedCatalogAllOf() *HclExemptedCatalogAllOf`

NewHclExemptedCatalogAllOf instantiates a new HclExemptedCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclExemptedCatalogAllOfWithDefaults

`func NewHclExemptedCatalogAllOfWithDefaults() *HclExemptedCatalogAllOf`

NewHclExemptedCatalogAllOfWithDefaults instantiates a new HclExemptedCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *HclExemptedCatalogAllOf) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *HclExemptedCatalogAllOf) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *HclExemptedCatalogAllOf) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *HclExemptedCatalogAllOf) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetName

`func (o *HclExemptedCatalogAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HclExemptedCatalogAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HclExemptedCatalogAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HclExemptedCatalogAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsVendor

`func (o *HclExemptedCatalogAllOf) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *HclExemptedCatalogAllOf) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *HclExemptedCatalogAllOf) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *HclExemptedCatalogAllOf) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *HclExemptedCatalogAllOf) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *HclExemptedCatalogAllOf) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *HclExemptedCatalogAllOf) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *HclExemptedCatalogAllOf) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProcessorName

`func (o *HclExemptedCatalogAllOf) GetProcessorName() string`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *HclExemptedCatalogAllOf) GetProcessorNameOk() (*string, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *HclExemptedCatalogAllOf) SetProcessorName(v string)`

SetProcessorName sets ProcessorName field to given value.

### HasProcessorName

`func (o *HclExemptedCatalogAllOf) HasProcessorName() bool`

HasProcessorName returns a boolean if a field has been set.

### GetProductModels

`func (o *HclExemptedCatalogAllOf) GetProductModels() []string`

GetProductModels returns the ProductModels field if non-nil, zero value otherwise.

### GetProductModelsOk

`func (o *HclExemptedCatalogAllOf) GetProductModelsOk() (*[]string, bool)`

GetProductModelsOk returns a tuple with the ProductModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductModels

`func (o *HclExemptedCatalogAllOf) SetProductModels(v []string)`

SetProductModels sets ProductModels field to given value.

### HasProductModels

`func (o *HclExemptedCatalogAllOf) HasProductModels() bool`

HasProductModels returns a boolean if a field has been set.

### GetProductType

`func (o *HclExemptedCatalogAllOf) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *HclExemptedCatalogAllOf) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *HclExemptedCatalogAllOf) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *HclExemptedCatalogAllOf) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetServerPid

`func (o *HclExemptedCatalogAllOf) GetServerPid() string`

GetServerPid returns the ServerPid field if non-nil, zero value otherwise.

### GetServerPidOk

`func (o *HclExemptedCatalogAllOf) GetServerPidOk() (*string, bool)`

GetServerPidOk returns a tuple with the ServerPid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPid

`func (o *HclExemptedCatalogAllOf) SetServerPid(v string)`

SetServerPid sets ServerPid field to given value.

### HasServerPid

`func (o *HclExemptedCatalogAllOf) HasServerPid() bool`

HasServerPid returns a boolean if a field has been set.

### GetUcsVersion

`func (o *HclExemptedCatalogAllOf) GetUcsVersion() string`

GetUcsVersion returns the UcsVersion field if non-nil, zero value otherwise.

### GetUcsVersionOk

`func (o *HclExemptedCatalogAllOf) GetUcsVersionOk() (*string, bool)`

GetUcsVersionOk returns a tuple with the UcsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsVersion

`func (o *HclExemptedCatalogAllOf) SetUcsVersion(v string)`

SetUcsVersion sets UcsVersion field to given value.

### HasUcsVersion

`func (o *HclExemptedCatalogAllOf) HasUcsVersion() bool`

HasUcsVersion returns a boolean if a field has been set.

### GetVersionType

`func (o *HclExemptedCatalogAllOf) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *HclExemptedCatalogAllOf) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *HclExemptedCatalogAllOf) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.

### HasVersionType

`func (o *HclExemptedCatalogAllOf) HasVersionType() bool`

HasVersionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


