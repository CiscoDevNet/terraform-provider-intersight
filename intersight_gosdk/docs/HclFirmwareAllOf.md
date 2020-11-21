# HclFirmwareAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.Firmware"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.Firmware"]
**DriverName** | Pointer to **string** | Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3. | [optional] 
**DriverVersion** | Pointer to **string** | Version of the Driver supported. | [optional] 
**ErrorCode** | Pointer to **string** | Error code for the support status. * &#x60;Success&#x60; - The input combination is valid. * &#x60;Unknown&#x60; - Unknown API request to the service. * &#x60;UnknownServer&#x60; - An invalid server model is given API requests or the server model is not present in the HCL database. * &#x60;InvalidUcsVersion&#x60; - UCS Version is not in expected format. * &#x60;ProcessorNotSupported&#x60; - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * &#x60;OSNotSupported&#x60; - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * &#x60;OSUnknown&#x60; - OS vendor or version is not known as per the HCL database. * &#x60;UCSVersionNotSupported&#x60; - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * &#x60;UcsVersionServerOSCombinationNotSupported&#x60; - Combination of UCS version, server (model and processor) and os version is not supported. * &#x60;ProductUnknown&#x60; - Product is not known as per the HCL database. * &#x60;ProductNotSupported&#x60; - Product is not supported in the given UCS version, server (model and processor) and operating system version. * &#x60;DriverNameNotSupported&#x60; - Driver protocol or name is not supported for the given product. * &#x60;FirmwareVersionNotSupported&#x60; - Firmware version not supported for the component and the server, operating system combination. * &#x60;DriverVersionNotSupported&#x60; - Driver version not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionDriverVersionCombinationNotSupported&#x60; - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionAndDriverVersionNotSupported&#x60; - Firmware and Driver version combination not supported for the component and the server, operating system combination. * &#x60;FirmwareVersionAndDriverNameNotSupported&#x60; - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * &#x60;InternalError&#x60; - API requests to the service have either failed or timed out. * &#x60;MarshallingError&#x60; - Error in JSON marshalling. * &#x60;Exempted&#x60; - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. | [optional] [readonly] [default to "Success"]
**FirmwareVersion** | Pointer to **string** | Firmware version of the product/adapter supported. | [optional] 
**Id** | Pointer to **string** | Identifier of the firmware. | [optional] 
**LatestDriver** | Pointer to **bool** | True if the driver is latest recommended driver. | [optional] [readonly] 
**LatestFirmware** | Pointer to **bool** | True if the firmware is latest recommended firmware. | [optional] [readonly] 

## Methods

### NewHclFirmwareAllOf

`func NewHclFirmwareAllOf(classId string, objectType string, ) *HclFirmwareAllOf`

NewHclFirmwareAllOf instantiates a new HclFirmwareAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclFirmwareAllOfWithDefaults

`func NewHclFirmwareAllOfWithDefaults() *HclFirmwareAllOf`

NewHclFirmwareAllOfWithDefaults instantiates a new HclFirmwareAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclFirmwareAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclFirmwareAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclFirmwareAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclFirmwareAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclFirmwareAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclFirmwareAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriverName

`func (o *HclFirmwareAllOf) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *HclFirmwareAllOf) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *HclFirmwareAllOf) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *HclFirmwareAllOf) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetDriverVersion

`func (o *HclFirmwareAllOf) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *HclFirmwareAllOf) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *HclFirmwareAllOf) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *HclFirmwareAllOf) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetErrorCode

`func (o *HclFirmwareAllOf) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclFirmwareAllOf) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclFirmwareAllOf) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclFirmwareAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *HclFirmwareAllOf) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *HclFirmwareAllOf) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *HclFirmwareAllOf) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *HclFirmwareAllOf) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetId

`func (o *HclFirmwareAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclFirmwareAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclFirmwareAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclFirmwareAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestDriver

`func (o *HclFirmwareAllOf) GetLatestDriver() bool`

GetLatestDriver returns the LatestDriver field if non-nil, zero value otherwise.

### GetLatestDriverOk

`func (o *HclFirmwareAllOf) GetLatestDriverOk() (*bool, bool)`

GetLatestDriverOk returns a tuple with the LatestDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDriver

`func (o *HclFirmwareAllOf) SetLatestDriver(v bool)`

SetLatestDriver sets LatestDriver field to given value.

### HasLatestDriver

`func (o *HclFirmwareAllOf) HasLatestDriver() bool`

HasLatestDriver returns a boolean if a field has been set.

### GetLatestFirmware

`func (o *HclFirmwareAllOf) GetLatestFirmware() bool`

GetLatestFirmware returns the LatestFirmware field if non-nil, zero value otherwise.

### GetLatestFirmwareOk

`func (o *HclFirmwareAllOf) GetLatestFirmwareOk() (*bool, bool)`

GetLatestFirmwareOk returns a tuple with the LatestFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFirmware

`func (o *HclFirmwareAllOf) SetLatestFirmware(v bool)`

SetLatestFirmware sets LatestFirmware field to given value.

### HasLatestFirmware

`func (o *HclFirmwareAllOf) HasLatestFirmware() bool`

HasLatestFirmware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


