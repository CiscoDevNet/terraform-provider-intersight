# CondHclStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentStatus** | Pointer to **string** | The overall status for the components found in the HCL. This will provide the HCL validation status for all the components. It can be one of the following. \&quot;Validated\&quot; - all the components hardware/software profiles are listed in the HCL. \&quot;Not-Listed\&quot; - one or more components hardware/software profiles are not listed in the HCL \&quot;Incomplete\&quot; - the components are not evaluated as the server&#39;s software/hardware profiles are not listed in the HCL. \&quot;Not-Evaluated\&quot; - The components are not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**HardwareStatus** | Pointer to **string** | The server model, processor and firmware are considered as part of the hardware profile for the server. This will provide the HCL validation status for the hardware profile. For the failure reason see the serverReason property. The status can be one of the following \&quot;Validated\&quot; - The server model, processor and firmware combination is listed in the HCL \&quot;Not-Listed\&quot; - The server model, processor and firmware combination is not listed in the HCL. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**HclFirmwareVersion** | Pointer to **string** | The current CIMC version for the server normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclModel** | Pointer to **string** | The managed object&#39;s model to validate normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclOsVendor** | Pointer to **string** | The OS Vendor for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclOsVersion** | Pointer to **string** | The OS Version for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclProcessor** | Pointer to **string** | The managed object&#39;s processor to validate if applicable normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**InvFirmwareVersion** | Pointer to **string** | The current CIMC version for the server as received from inventory. It is empty if we are missing this information. | [optional] 
**InvModel** | Pointer to **string** | The managed object&#39;s model to validate as received from the inventory. It is empty if we are missing this information. | [optional] 
**InvOsVendor** | Pointer to **string** | The OS Vendor for the managed object to validate as received from inventory. It is empty if we are missing this information. | [optional] 
**InvOsVersion** | Pointer to **string** | The OS Version for the managed object to validate as received from inventory. It is empty if we are missing this information. | [optional] 
**InvProcessor** | Pointer to **string** | The managed object&#39;s processor to validate if applicable as received from inventory. It is empty if we are missing this information. | [optional] 
**Reason** | Pointer to **string** | The reason for the HCL status. It will be one of the following \&quot;Missing-Os-Info\&quot; - we are missing os information in the inventory from the device connector \&quot;Incompatible-Components\&quot; - we have 1 or more components with \&quot;Not-Validated\&quot; status \&quot;Compatible\&quot; - all the components have \&quot;Validated\&quot; status. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Missing-Os-Info&#x60; - This means the HclStatus for the sever failed HCL validation because we have missing os information. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Components&#x60; - This means the HclStatus for the sever failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails. * &#x60;Compatible&#x60; - This means the HclStatus for the sever has passed HCL validation for all of its related components. * &#x60;Not-Evaluated&#x60; - This means the HclStatus for the sever has not been evaluated because it is exempted. | [optional] [default to "Missing-Os-Info"]
**ServerReason** | Pointer to **string** | The reason generated by the server&#39;s HCL validation. For HCL the evaluation can be seen in three logical stages 1. Evaluate the server&#39;s hardware status 2. Evaluate the server&#39;s software status 3. Evaluate the server&#39;s components (each component has its own hardware/software evaluation) The evaluation does not proceed to the next stage until the previous stage is evaluated. Therefore there can be only one validation reason. \&quot;Incompatible-Server\&quot; - the server model is not listed in the HCL. \&quot;Incompatible-Processor\&quot; - the server model and processor combination is not listed in HCL. \&quot;Incompatible-Firmware\&quot; - the server model, processor and server firmware is not listed in HCL. \&quot;Missing-Os-Info\&quot; - the os vendor and version is not listed in HCL with the HW profile. \&quot;Incompatible-Os-Info\&quot; - the os vendor and version is not listed in HCL with the HW profile. \&quot;Incompatible-Components\&quot; - there is one or more components with \&quot;Not-Validated\&quot; status \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). \&quot;Compatible\&quot; - the server and all its components are validated. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information. * &#x60;Incompatible-Server&#x60; - The validation failed for this server because the server&#39;s model was not listed in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not listed for the given server model. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has returned a service error or unrecognized result. * &#x60;Not-Evaluated&#x60; - This means the HclStatus for the sever has not been evaluated because it is exempted. * &#x60;Incompatible-Components&#x60; - The validation has failed for this server because one or more components have \&quot;Not-Listed\&quot; status. * &#x60;Compatible&#x60; - The validation has passed for this server&#39;s model, processor, OS vendor and version. | [optional] [default to "Missing-Os-Driver-Info"]
**SoftwareStatus** | Pointer to **string** | The OS vendor and version are considered part of the software profile for the server. This will provide the HCL validation status for the software profile. For the failure reason see the serverReason property. The status can be be one of the following \&quot;Validated\&quot; - The os vendor/version is listed in the HCL for the server model, processor and firmware \&quot;Not-Listed\&quot; - The os vendor/version is not listed in the HCL for the server model, processor and firmware \&quot;Incomplete\&quot; - The inventory is missing os vendor/version and HCL validation was not performed. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**Status** | Pointer to **string** | The HCL compatibility status of the managed object. The status can be one of the following \&quot;Incomplete\&quot; - there is no enough information to evaluate against the HCL data \&quot;Validated\&quot; - all components have been evaluated against the HCL and they all have \&quot;Validated\&quot; status \&quot;Not-Listed\&quot; - all components have been evaluated against the HCL and one or more have \&quot;Not-Listed\&quot; status. \&quot;Not-Evaluated\&quot; - server is not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**Details** | Pointer to [**[]CondHclStatusDetailRelationship**](cond.HclStatusDetail.Relationship.md) | An array of relationships to condHclStatusDetail resources. | [optional] [readonly] 
**ManagedObject** | Pointer to [**InventoryBaseRelationship**](inventory.Base.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewCondHclStatus

`func NewCondHclStatus() *CondHclStatus`

NewCondHclStatus instantiates a new CondHclStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondHclStatusWithDefaults

`func NewCondHclStatusWithDefaults() *CondHclStatus`

NewCondHclStatusWithDefaults instantiates a new CondHclStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentStatus

`func (o *CondHclStatus) GetComponentStatus() string`

GetComponentStatus returns the ComponentStatus field if non-nil, zero value otherwise.

### GetComponentStatusOk

`func (o *CondHclStatus) GetComponentStatusOk() (*string, bool)`

GetComponentStatusOk returns a tuple with the ComponentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentStatus

`func (o *CondHclStatus) SetComponentStatus(v string)`

SetComponentStatus sets ComponentStatus field to given value.

### HasComponentStatus

`func (o *CondHclStatus) HasComponentStatus() bool`

HasComponentStatus returns a boolean if a field has been set.

### GetHardwareStatus

`func (o *CondHclStatus) GetHardwareStatus() string`

GetHardwareStatus returns the HardwareStatus field if non-nil, zero value otherwise.

### GetHardwareStatusOk

`func (o *CondHclStatus) GetHardwareStatusOk() (*string, bool)`

GetHardwareStatusOk returns a tuple with the HardwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareStatus

`func (o *CondHclStatus) SetHardwareStatus(v string)`

SetHardwareStatus sets HardwareStatus field to given value.

### HasHardwareStatus

`func (o *CondHclStatus) HasHardwareStatus() bool`

HasHardwareStatus returns a boolean if a field has been set.

### GetHclFirmwareVersion

`func (o *CondHclStatus) GetHclFirmwareVersion() string`

GetHclFirmwareVersion returns the HclFirmwareVersion field if non-nil, zero value otherwise.

### GetHclFirmwareVersionOk

`func (o *CondHclStatus) GetHclFirmwareVersionOk() (*string, bool)`

GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclFirmwareVersion

`func (o *CondHclStatus) SetHclFirmwareVersion(v string)`

SetHclFirmwareVersion sets HclFirmwareVersion field to given value.

### HasHclFirmwareVersion

`func (o *CondHclStatus) HasHclFirmwareVersion() bool`

HasHclFirmwareVersion returns a boolean if a field has been set.

### GetHclModel

`func (o *CondHclStatus) GetHclModel() string`

GetHclModel returns the HclModel field if non-nil, zero value otherwise.

### GetHclModelOk

`func (o *CondHclStatus) GetHclModelOk() (*string, bool)`

GetHclModelOk returns a tuple with the HclModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclModel

`func (o *CondHclStatus) SetHclModel(v string)`

SetHclModel sets HclModel field to given value.

### HasHclModel

`func (o *CondHclStatus) HasHclModel() bool`

HasHclModel returns a boolean if a field has been set.

### GetHclOsVendor

`func (o *CondHclStatus) GetHclOsVendor() string`

GetHclOsVendor returns the HclOsVendor field if non-nil, zero value otherwise.

### GetHclOsVendorOk

`func (o *CondHclStatus) GetHclOsVendorOk() (*string, bool)`

GetHclOsVendorOk returns a tuple with the HclOsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclOsVendor

`func (o *CondHclStatus) SetHclOsVendor(v string)`

SetHclOsVendor sets HclOsVendor field to given value.

### HasHclOsVendor

`func (o *CondHclStatus) HasHclOsVendor() bool`

HasHclOsVendor returns a boolean if a field has been set.

### GetHclOsVersion

`func (o *CondHclStatus) GetHclOsVersion() string`

GetHclOsVersion returns the HclOsVersion field if non-nil, zero value otherwise.

### GetHclOsVersionOk

`func (o *CondHclStatus) GetHclOsVersionOk() (*string, bool)`

GetHclOsVersionOk returns a tuple with the HclOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclOsVersion

`func (o *CondHclStatus) SetHclOsVersion(v string)`

SetHclOsVersion sets HclOsVersion field to given value.

### HasHclOsVersion

`func (o *CondHclStatus) HasHclOsVersion() bool`

HasHclOsVersion returns a boolean if a field has been set.

### GetHclProcessor

`func (o *CondHclStatus) GetHclProcessor() string`

GetHclProcessor returns the HclProcessor field if non-nil, zero value otherwise.

### GetHclProcessorOk

`func (o *CondHclStatus) GetHclProcessorOk() (*string, bool)`

GetHclProcessorOk returns a tuple with the HclProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclProcessor

`func (o *CondHclStatus) SetHclProcessor(v string)`

SetHclProcessor sets HclProcessor field to given value.

### HasHclProcessor

`func (o *CondHclStatus) HasHclProcessor() bool`

HasHclProcessor returns a boolean if a field has been set.

### GetInvFirmwareVersion

`func (o *CondHclStatus) GetInvFirmwareVersion() string`

GetInvFirmwareVersion returns the InvFirmwareVersion field if non-nil, zero value otherwise.

### GetInvFirmwareVersionOk

`func (o *CondHclStatus) GetInvFirmwareVersionOk() (*string, bool)`

GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvFirmwareVersion

`func (o *CondHclStatus) SetInvFirmwareVersion(v string)`

SetInvFirmwareVersion sets InvFirmwareVersion field to given value.

### HasInvFirmwareVersion

`func (o *CondHclStatus) HasInvFirmwareVersion() bool`

HasInvFirmwareVersion returns a boolean if a field has been set.

### GetInvModel

`func (o *CondHclStatus) GetInvModel() string`

GetInvModel returns the InvModel field if non-nil, zero value otherwise.

### GetInvModelOk

`func (o *CondHclStatus) GetInvModelOk() (*string, bool)`

GetInvModelOk returns a tuple with the InvModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvModel

`func (o *CondHclStatus) SetInvModel(v string)`

SetInvModel sets InvModel field to given value.

### HasInvModel

`func (o *CondHclStatus) HasInvModel() bool`

HasInvModel returns a boolean if a field has been set.

### GetInvOsVendor

`func (o *CondHclStatus) GetInvOsVendor() string`

GetInvOsVendor returns the InvOsVendor field if non-nil, zero value otherwise.

### GetInvOsVendorOk

`func (o *CondHclStatus) GetInvOsVendorOk() (*string, bool)`

GetInvOsVendorOk returns a tuple with the InvOsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvOsVendor

`func (o *CondHclStatus) SetInvOsVendor(v string)`

SetInvOsVendor sets InvOsVendor field to given value.

### HasInvOsVendor

`func (o *CondHclStatus) HasInvOsVendor() bool`

HasInvOsVendor returns a boolean if a field has been set.

### GetInvOsVersion

`func (o *CondHclStatus) GetInvOsVersion() string`

GetInvOsVersion returns the InvOsVersion field if non-nil, zero value otherwise.

### GetInvOsVersionOk

`func (o *CondHclStatus) GetInvOsVersionOk() (*string, bool)`

GetInvOsVersionOk returns a tuple with the InvOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvOsVersion

`func (o *CondHclStatus) SetInvOsVersion(v string)`

SetInvOsVersion sets InvOsVersion field to given value.

### HasInvOsVersion

`func (o *CondHclStatus) HasInvOsVersion() bool`

HasInvOsVersion returns a boolean if a field has been set.

### GetInvProcessor

`func (o *CondHclStatus) GetInvProcessor() string`

GetInvProcessor returns the InvProcessor field if non-nil, zero value otherwise.

### GetInvProcessorOk

`func (o *CondHclStatus) GetInvProcessorOk() (*string, bool)`

GetInvProcessorOk returns a tuple with the InvProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvProcessor

`func (o *CondHclStatus) SetInvProcessor(v string)`

SetInvProcessor sets InvProcessor field to given value.

### HasInvProcessor

`func (o *CondHclStatus) HasInvProcessor() bool`

HasInvProcessor returns a boolean if a field has been set.

### GetReason

`func (o *CondHclStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CondHclStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CondHclStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CondHclStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetServerReason

`func (o *CondHclStatus) GetServerReason() string`

GetServerReason returns the ServerReason field if non-nil, zero value otherwise.

### GetServerReasonOk

`func (o *CondHclStatus) GetServerReasonOk() (*string, bool)`

GetServerReasonOk returns a tuple with the ServerReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerReason

`func (o *CondHclStatus) SetServerReason(v string)`

SetServerReason sets ServerReason field to given value.

### HasServerReason

`func (o *CondHclStatus) HasServerReason() bool`

HasServerReason returns a boolean if a field has been set.

### GetSoftwareStatus

`func (o *CondHclStatus) GetSoftwareStatus() string`

GetSoftwareStatus returns the SoftwareStatus field if non-nil, zero value otherwise.

### GetSoftwareStatusOk

`func (o *CondHclStatus) GetSoftwareStatusOk() (*string, bool)`

GetSoftwareStatusOk returns a tuple with the SoftwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareStatus

`func (o *CondHclStatus) SetSoftwareStatus(v string)`

SetSoftwareStatus sets SoftwareStatus field to given value.

### HasSoftwareStatus

`func (o *CondHclStatus) HasSoftwareStatus() bool`

HasSoftwareStatus returns a boolean if a field has been set.

### GetStatus

`func (o *CondHclStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondHclStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondHclStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondHclStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetails

`func (o *CondHclStatus) GetDetails() []CondHclStatusDetailRelationship`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CondHclStatus) GetDetailsOk() (*[]CondHclStatusDetailRelationship, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CondHclStatus) SetDetails(v []CondHclStatusDetailRelationship)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CondHclStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CondHclStatus) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CondHclStatus) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetManagedObject

`func (o *CondHclStatus) GetManagedObject() InventoryBaseRelationship`

GetManagedObject returns the ManagedObject field if non-nil, zero value otherwise.

### GetManagedObjectOk

`func (o *CondHclStatus) GetManagedObjectOk() (*InventoryBaseRelationship, bool)`

GetManagedObjectOk returns a tuple with the ManagedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedObject

`func (o *CondHclStatus) SetManagedObject(v InventoryBaseRelationship)`

SetManagedObject sets ManagedObject field to given value.

### HasManagedObject

`func (o *CondHclStatus) HasManagedObject() bool`

HasManagedObject returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CondHclStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondHclStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondHclStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondHclStatus) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


