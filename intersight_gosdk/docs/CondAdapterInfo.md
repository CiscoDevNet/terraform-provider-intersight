# CondAdapterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**DriverName** | Pointer to **string** | It specifies the name of the driver supported by the adapter. | [optional] 
**DriverVersion** | Pointer to **string** | It specifies the version of the driver installed for the adapter. | [optional] 
**Firmware** | Pointer to **string** | It specifies the firmware version installed on the adapter. | [optional] 
**HclReason** | Pointer to **string** | The reason for the Cisco HCL validation status, more useful when status is NotListed. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed because the given server has no operating system driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper operating system information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given operating system vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and operating system vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, operating system vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, operating system vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, operating system vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, operating system vendor and version, component model, component firmware and driver version. | [optional] [readonly] [default to "Missing-Os-Driver-Info"]
**HclStatus** | Pointer to **string** | The Cisco HCL validation status of the adapter. * &#x60;Incomplete&#x60; - This means we do not have operating system information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper operating system information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that in component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet or the configurations provided are insufficient for HCL status assessment. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. * &#x60;Not-Applicable&#x60; - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. | [optional] [readonly] [default to "Incomplete"]
**Model** | Pointer to **string** | It specifies the model name or Product ID (PID) of the adapter. | [optional] 
**Type** | Pointer to **string** | It specifies the type of the adapter, such as Network Adapter, Storage Controller, or GPU. * &#x60;Unknown&#x60; - It indicates no adapter type. it is used when component type enum is not defined. * &#x60;NetworkAdapter&#x60; - Indicates network adapter. * &#x60;StorageController&#x60; - Indicates Storage controller. * &#x60;GPU&#x60; - It refers to Graphics Processing Unit (GPU) adapters, which are used for rendering graphics and performing parallel processing tasks. * &#x60;SSD&#x60; - It refers NVME solid state drives (SSD) drives. | [optional] [default to "Unknown"]
**Vendor** | Pointer to **string** | It specifies the vendor or manufacturer of the adapter. | [optional] 

## Methods

### NewCondAdapterInfo

`func NewCondAdapterInfo(classId string, objectType string, ) *CondAdapterInfo`

NewCondAdapterInfo instantiates a new CondAdapterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAdapterInfoWithDefaults

`func NewCondAdapterInfoWithDefaults() *CondAdapterInfo`

NewCondAdapterInfoWithDefaults instantiates a new CondAdapterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAdapterInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAdapterInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAdapterInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAdapterInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAdapterInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAdapterInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriverName

`func (o *CondAdapterInfo) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *CondAdapterInfo) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *CondAdapterInfo) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *CondAdapterInfo) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetDriverVersion

`func (o *CondAdapterInfo) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *CondAdapterInfo) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *CondAdapterInfo) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *CondAdapterInfo) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetFirmware

`func (o *CondAdapterInfo) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *CondAdapterInfo) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *CondAdapterInfo) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *CondAdapterInfo) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetHclReason

`func (o *CondAdapterInfo) GetHclReason() string`

GetHclReason returns the HclReason field if non-nil, zero value otherwise.

### GetHclReasonOk

`func (o *CondAdapterInfo) GetHclReasonOk() (*string, bool)`

GetHclReasonOk returns a tuple with the HclReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclReason

`func (o *CondAdapterInfo) SetHclReason(v string)`

SetHclReason sets HclReason field to given value.

### HasHclReason

`func (o *CondAdapterInfo) HasHclReason() bool`

HasHclReason returns a boolean if a field has been set.

### GetHclStatus

`func (o *CondAdapterInfo) GetHclStatus() string`

GetHclStatus returns the HclStatus field if non-nil, zero value otherwise.

### GetHclStatusOk

`func (o *CondAdapterInfo) GetHclStatusOk() (*string, bool)`

GetHclStatusOk returns a tuple with the HclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatus

`func (o *CondAdapterInfo) SetHclStatus(v string)`

SetHclStatus sets HclStatus field to given value.

### HasHclStatus

`func (o *CondAdapterInfo) HasHclStatus() bool`

HasHclStatus returns a boolean if a field has been set.

### GetModel

`func (o *CondAdapterInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CondAdapterInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CondAdapterInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CondAdapterInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *CondAdapterInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CondAdapterInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CondAdapterInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CondAdapterInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *CondAdapterInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CondAdapterInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CondAdapterInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CondAdapterInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


