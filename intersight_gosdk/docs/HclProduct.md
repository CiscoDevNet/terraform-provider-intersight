# HclProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.Product"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.Product"]
**DriverNames** | Pointer to **[]string** |  | [optional] 
**ErrorCode** | Pointer to **string** | Error code indicating the support status. * &#x60;Success&#x60; - The input combination is valid. * &#x60;Unknown&#x60; - Unknown API request to the service. * &#x60;UnknownServer&#x60; - An invalid server model is given API requests or the server model is not present in the HCL database. * &#x60;InvalidUcsVersion&#x60; - UCS Version is not in expected format. * &#x60;ProcessorNotSupported&#x60; - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * &#x60;OSNotSupported&#x60; - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * &#x60;OSUnknown&#x60; - OS vendor or version is not known as per the HCL database. * &#x60;UCSVersionNotSupported&#x60; - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * &#x60;UcsVersionServerOSCombinationNotSupported&#x60; - Combination of UCS version, server (model and processor) and os version is not supported. * &#x60;ProductUnknown&#x60; - Product is not known as per the HCL database. * &#x60;ProductNotSupported&#x60; - Product is not supported in the given UCS version, server (model and processor) and operating system version. * &#x60;DriverNameNotSupported&#x60; - Driver protocol or name is not supported for the given product. * &#x60;FirmwareVersionNotSupported&#x60; - Firmware version not supported for the component and the server, operating system combination. * &#x60;DriverVersionNotSupported&#x60; - Driver version not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionDriverVersionCombinationNotSupported&#x60; - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionAndDriverVersionNotSupported&#x60; - Firmware and Driver version combination not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionAndDriverNameNotSupported&#x60; - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * &#x60;InternalError&#x60; - API requests to the service have either failed or timed out. * &#x60;MarshallingError&#x60; - Error in JSON marshalling. * &#x60;Exempted&#x60; - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. | [optional] [readonly] [default to "Success"]
**Firmwares** | Pointer to [**[]HclFirmware**](HclFirmware.md) |  | [optional] 
**Id** | Pointer to **string** | Identifier of the product. | [optional] 
**Model** | Pointer to **string** | Model/PID of the product/adapter. | [optional] 
**Revision** | Pointer to **string** | Revision of the adapter model. | [optional] 
**Type** | Pointer to **string** | Type of the product/adapter say OCP, PT, GPU. * &#x60;&#x60; - Default type of the Product. * &#x60;Adapter&#x60; - Represents network adapter cards. * &#x60;StorageController&#x60; - Represents storage controllers. * &#x60;GPU&#x60; - Represents graphics cards. | [optional] [default to ""]
**Vendor** | Pointer to **string** | Vendor of the product or adapter. | [optional] 

## Methods

### NewHclProduct

`func NewHclProduct(classId string, objectType string, ) *HclProduct`

NewHclProduct instantiates a new HclProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclProductWithDefaults

`func NewHclProductWithDefaults() *HclProduct`

NewHclProductWithDefaults instantiates a new HclProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclProduct) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclProduct) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclProduct) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclProduct) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclProduct) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclProduct) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriverNames

`func (o *HclProduct) GetDriverNames() []string`

GetDriverNames returns the DriverNames field if non-nil, zero value otherwise.

### GetDriverNamesOk

`func (o *HclProduct) GetDriverNamesOk() (*[]string, bool)`

GetDriverNamesOk returns a tuple with the DriverNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverNames

`func (o *HclProduct) SetDriverNames(v []string)`

SetDriverNames sets DriverNames field to given value.

### HasDriverNames

`func (o *HclProduct) HasDriverNames() bool`

HasDriverNames returns a boolean if a field has been set.

### SetDriverNamesNil

`func (o *HclProduct) SetDriverNamesNil(b bool)`

 SetDriverNamesNil sets the value for DriverNames to be an explicit nil

### UnsetDriverNames
`func (o *HclProduct) UnsetDriverNames()`

UnsetDriverNames ensures that no value is present for DriverNames, not even an explicit nil
### GetErrorCode

`func (o *HclProduct) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclProduct) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclProduct) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclProduct) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFirmwares

`func (o *HclProduct) GetFirmwares() []HclFirmware`

GetFirmwares returns the Firmwares field if non-nil, zero value otherwise.

### GetFirmwaresOk

`func (o *HclProduct) GetFirmwaresOk() (*[]HclFirmware, bool)`

GetFirmwaresOk returns a tuple with the Firmwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwares

`func (o *HclProduct) SetFirmwares(v []HclFirmware)`

SetFirmwares sets Firmwares field to given value.

### HasFirmwares

`func (o *HclProduct) HasFirmwares() bool`

HasFirmwares returns a boolean if a field has been set.

### SetFirmwaresNil

`func (o *HclProduct) SetFirmwaresNil(b bool)`

 SetFirmwaresNil sets the value for Firmwares to be an explicit nil

### UnsetFirmwares
`func (o *HclProduct) UnsetFirmwares()`

UnsetFirmwares ensures that no value is present for Firmwares, not even an explicit nil
### GetId

`func (o *HclProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *HclProduct) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HclProduct) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HclProduct) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HclProduct) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *HclProduct) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HclProduct) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HclProduct) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *HclProduct) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetType

`func (o *HclProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HclProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HclProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HclProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *HclProduct) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HclProduct) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HclProduct) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HclProduct) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


