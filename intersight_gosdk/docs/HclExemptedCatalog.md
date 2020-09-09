# HclExemptedCatalog

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

### NewHclExemptedCatalog

`func NewHclExemptedCatalog() *HclExemptedCatalog`

NewHclExemptedCatalog instantiates a new HclExemptedCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclExemptedCatalogWithDefaults

`func NewHclExemptedCatalogWithDefaults() *HclExemptedCatalog`

NewHclExemptedCatalogWithDefaults instantiates a new HclExemptedCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *HclExemptedCatalog) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *HclExemptedCatalog) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *HclExemptedCatalog) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *HclExemptedCatalog) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetName

`func (o *HclExemptedCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HclExemptedCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HclExemptedCatalog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HclExemptedCatalog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsVendor

`func (o *HclExemptedCatalog) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *HclExemptedCatalog) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *HclExemptedCatalog) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *HclExemptedCatalog) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *HclExemptedCatalog) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *HclExemptedCatalog) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *HclExemptedCatalog) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *HclExemptedCatalog) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProcessorName

`func (o *HclExemptedCatalog) GetProcessorName() string`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *HclExemptedCatalog) GetProcessorNameOk() (*string, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *HclExemptedCatalog) SetProcessorName(v string)`

SetProcessorName sets ProcessorName field to given value.

### HasProcessorName

`func (o *HclExemptedCatalog) HasProcessorName() bool`

HasProcessorName returns a boolean if a field has been set.

### GetProductModels

`func (o *HclExemptedCatalog) GetProductModels() []string`

GetProductModels returns the ProductModels field if non-nil, zero value otherwise.

### GetProductModelsOk

`func (o *HclExemptedCatalog) GetProductModelsOk() (*[]string, bool)`

GetProductModelsOk returns a tuple with the ProductModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductModels

`func (o *HclExemptedCatalog) SetProductModels(v []string)`

SetProductModels sets ProductModels field to given value.

### HasProductModels

`func (o *HclExemptedCatalog) HasProductModels() bool`

HasProductModels returns a boolean if a field has been set.

### GetProductType

`func (o *HclExemptedCatalog) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *HclExemptedCatalog) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *HclExemptedCatalog) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *HclExemptedCatalog) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetServerPid

`func (o *HclExemptedCatalog) GetServerPid() string`

GetServerPid returns the ServerPid field if non-nil, zero value otherwise.

### GetServerPidOk

`func (o *HclExemptedCatalog) GetServerPidOk() (*string, bool)`

GetServerPidOk returns a tuple with the ServerPid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPid

`func (o *HclExemptedCatalog) SetServerPid(v string)`

SetServerPid sets ServerPid field to given value.

### HasServerPid

`func (o *HclExemptedCatalog) HasServerPid() bool`

HasServerPid returns a boolean if a field has been set.

### GetUcsVersion

`func (o *HclExemptedCatalog) GetUcsVersion() string`

GetUcsVersion returns the UcsVersion field if non-nil, zero value otherwise.

### GetUcsVersionOk

`func (o *HclExemptedCatalog) GetUcsVersionOk() (*string, bool)`

GetUcsVersionOk returns a tuple with the UcsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsVersion

`func (o *HclExemptedCatalog) SetUcsVersion(v string)`

SetUcsVersion sets UcsVersion field to given value.

### HasUcsVersion

`func (o *HclExemptedCatalog) HasUcsVersion() bool`

HasUcsVersion returns a boolean if a field has been set.

### GetVersionType

`func (o *HclExemptedCatalog) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *HclExemptedCatalog) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *HclExemptedCatalog) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.

### HasVersionType

`func (o *HclExemptedCatalog) HasVersionType() bool`

HasVersionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


