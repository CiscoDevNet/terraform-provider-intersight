# HclHardwareCompatibilityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverIsoUrl** | Pointer to **string** | Url for the ISO with the drivers supported for the server. | [optional] 
**ErrorCode** | Pointer to **string** | Error code indicating the compatibility status. * &#x60;Success&#x60; - The input combination is valid. * &#x60;Unknown&#x60; - Unknown API request to the service. * &#x60;UnknownServer&#x60; - An invalid server model is given API requests or the server model is not present in the HCL database. * &#x60;InvalidUcsVersion&#x60; - UCS Version is not in expected format. * &#x60;ProcessorNotSupported&#x60; - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * &#x60;OSNotSupported&#x60; - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * &#x60;OSUnknown&#x60; - OS vendor or version is not known as per the HCL database. * &#x60;UCSVersionNotSupported&#x60; - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * &#x60;UcsVersionServerOSCombinationNotSupported&#x60; - Combination of UCS version, server (model and processor) and os version is not supported. * &#x60;ProductUnknown&#x60; - Product is not known as per the HCL database. * &#x60;ProductNotSupported&#x60; - Product is not supported in the given UCS version, server (model and processor) and operating system version. * &#x60;DriverNameNotSupported&#x60; - Driver protocol or name is not supported for the given product. * &#x60;FirmwareVersionNotSupported&#x60; - Firmware version not supported for the component and the server, operating system combination. * &#x60;DriverVersionNotSupported&#x60; - Driver version not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionDriverVersionCombinationNotSupported&#x60; - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionAndDriverVersionNotSupported&#x60; - Firmware and Driver version combination not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionAndDriverNameNotSupported&#x60; - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * &#x60;InternalError&#x60; - API requests to the service have either failed or timed out. * &#x60;MarshallingError&#x60; - Error in JSON marshalling. * &#x60;Exempted&#x60; - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. | [optional] [readonly] [default to "Success"]
**Id** | Pointer to **string** | Identifier of the hardware compatibility profile. | [optional] 
**OsVendor** | Pointer to **string** | Vendor of the Operating System running on the server. | [optional] 
**OsVersion** | Pointer to **string** | Version of the Operating System running on the server. | [optional] 
**ProcessorModel** | Pointer to **string** | Model of the processor present in the server. | [optional] 
**Products** | Pointer to [**[]HclProduct**](hcl.Product.md) |  | [optional] 
**ServerModel** | Pointer to **string** | Model of the server as returned by UCSM/CIMC XML API. | [optional] 
**ServerRevision** | Pointer to **string** | Revision of the server model. | [optional] 
**UcsVersion** | Pointer to **string** | Version of the UCS software. | [optional] 
**VersionType** | Pointer to **string** | Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. * &#x60;UCSM&#x60; - The server is managed by UCS Manager. * &#x60;IMC&#x60; - The server is standalone managed by CIMC. | [optional] [default to "UCSM"]

## Methods

### NewHclHardwareCompatibilityProfile

`func NewHclHardwareCompatibilityProfile() *HclHardwareCompatibilityProfile`

NewHclHardwareCompatibilityProfile instantiates a new HclHardwareCompatibilityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclHardwareCompatibilityProfileWithDefaults

`func NewHclHardwareCompatibilityProfileWithDefaults() *HclHardwareCompatibilityProfile`

NewHclHardwareCompatibilityProfileWithDefaults instantiates a new HclHardwareCompatibilityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverIsoUrl

`func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrl() string`

GetDriverIsoUrl returns the DriverIsoUrl field if non-nil, zero value otherwise.

### GetDriverIsoUrlOk

`func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrlOk() (*string, bool)`

GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverIsoUrl

`func (o *HclHardwareCompatibilityProfile) SetDriverIsoUrl(v string)`

SetDriverIsoUrl sets DriverIsoUrl field to given value.

### HasDriverIsoUrl

`func (o *HclHardwareCompatibilityProfile) HasDriverIsoUrl() bool`

HasDriverIsoUrl returns a boolean if a field has been set.

### GetErrorCode

`func (o *HclHardwareCompatibilityProfile) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclHardwareCompatibilityProfile) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclHardwareCompatibilityProfile) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclHardwareCompatibilityProfile) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetId

`func (o *HclHardwareCompatibilityProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclHardwareCompatibilityProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclHardwareCompatibilityProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclHardwareCompatibilityProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOsVendor

`func (o *HclHardwareCompatibilityProfile) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *HclHardwareCompatibilityProfile) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *HclHardwareCompatibilityProfile) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *HclHardwareCompatibilityProfile) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *HclHardwareCompatibilityProfile) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *HclHardwareCompatibilityProfile) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *HclHardwareCompatibilityProfile) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *HclHardwareCompatibilityProfile) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProcessorModel

`func (o *HclHardwareCompatibilityProfile) GetProcessorModel() string`

GetProcessorModel returns the ProcessorModel field if non-nil, zero value otherwise.

### GetProcessorModelOk

`func (o *HclHardwareCompatibilityProfile) GetProcessorModelOk() (*string, bool)`

GetProcessorModelOk returns a tuple with the ProcessorModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorModel

`func (o *HclHardwareCompatibilityProfile) SetProcessorModel(v string)`

SetProcessorModel sets ProcessorModel field to given value.

### HasProcessorModel

`func (o *HclHardwareCompatibilityProfile) HasProcessorModel() bool`

HasProcessorModel returns a boolean if a field has been set.

### GetProducts

`func (o *HclHardwareCompatibilityProfile) GetProducts() []HclProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *HclHardwareCompatibilityProfile) GetProductsOk() (*[]HclProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *HclHardwareCompatibilityProfile) SetProducts(v []HclProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *HclHardwareCompatibilityProfile) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetServerModel

`func (o *HclHardwareCompatibilityProfile) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HclHardwareCompatibilityProfile) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HclHardwareCompatibilityProfile) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HclHardwareCompatibilityProfile) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### GetServerRevision

`func (o *HclHardwareCompatibilityProfile) GetServerRevision() string`

GetServerRevision returns the ServerRevision field if non-nil, zero value otherwise.

### GetServerRevisionOk

`func (o *HclHardwareCompatibilityProfile) GetServerRevisionOk() (*string, bool)`

GetServerRevisionOk returns a tuple with the ServerRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRevision

`func (o *HclHardwareCompatibilityProfile) SetServerRevision(v string)`

SetServerRevision sets ServerRevision field to given value.

### HasServerRevision

`func (o *HclHardwareCompatibilityProfile) HasServerRevision() bool`

HasServerRevision returns a boolean if a field has been set.

### GetUcsVersion

`func (o *HclHardwareCompatibilityProfile) GetUcsVersion() string`

GetUcsVersion returns the UcsVersion field if non-nil, zero value otherwise.

### GetUcsVersionOk

`func (o *HclHardwareCompatibilityProfile) GetUcsVersionOk() (*string, bool)`

GetUcsVersionOk returns a tuple with the UcsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsVersion

`func (o *HclHardwareCompatibilityProfile) SetUcsVersion(v string)`

SetUcsVersion sets UcsVersion field to given value.

### HasUcsVersion

`func (o *HclHardwareCompatibilityProfile) HasUcsVersion() bool`

HasUcsVersion returns a boolean if a field has been set.

### GetVersionType

`func (o *HclHardwareCompatibilityProfile) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *HclHardwareCompatibilityProfile) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *HclHardwareCompatibilityProfile) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.

### HasVersionType

`func (o *HclHardwareCompatibilityProfile) HasVersionType() bool`

HasVersionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


