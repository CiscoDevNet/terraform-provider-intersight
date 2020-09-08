# CondHclStatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareStatus** | Pointer to **string** | The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \&quot;Incompatible-Server-With-Component\&quot; - the server model and component combination is not listed in HCL \&quot;Incompatible-Firmware\&quot; - The server&#39;s firmware is not listed for this component&#39;s hardware profile \&quot;Incompatible-Component\&quot; - the component&#39;s model is not listed in the HCL \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \&quot;Not-Evaluated\&quot; - the hardware profile was not evaulated for the component because the server&#39;s hw/sw status is not listed or server is exempted. \&quot;Compatible\&quot; - this component&#39;s hardware profile is listed in the HCL. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version. | [optional] [default to "Missing-Os-Driver-Info"]
**HclCimcVersion** | Pointer to **string** | The current CIMC version for the server normalized for querying HCL data. | [optional] 
**HclDriverName** | Pointer to **string** | The current driver name of the component we are validating normalized for querying HCL data. | [optional] 
**HclDriverVersion** | Pointer to **string** | The current driver version of the component we are validating normalized for querying HCL data. | [optional] 
**HclFirmwareVersion** | Pointer to **string** | The current firmware version of the component model normalized for querying HCL data. | [optional] 
**HclModel** | Pointer to **string** | The component model we are trying to validate normalized for querying HCL data. | [optional] 
**InvCimcVersion** | Pointer to **string** | The current CIMC version for the server as received from inventory. | [optional] 
**InvDriverName** | Pointer to **string** | The current driver name of the component we are validating as received from inventory. | [optional] 
**InvDriverVersion** | Pointer to **string** | The current driver version of the component we are validating as received from inventory. | [optional] 
**InvFirmwareVersion** | Pointer to **string** | The current firmware version of the component model as received from inventory. | [optional] 
**InvModel** | Pointer to **string** | The component model we are trying to validate as received from inventory. | [optional] 
**Reason** | Pointer to **string** | The reason for the status. The reason can be one of \&quot;Incompatible-Server-With-Component\&quot; - HCL validation has failed because the server model is not validated with this component \&quot;Incompatible-Processor\&quot; - HCL validation has failed because the processor is not validated with this server \&quot;Incompatible-Os-Info\&quot; - HCL validation has failed because the os vendor and version is not validated with this server \&quot;Incompatible-Component-Model\&quot; - HCL validation has failed because the component model is not validated \&quot;Incompatible-Firmware\&quot; - HCL validation has failed because the component or server firmware version is not validated \&quot;Incompatible-Driver\&quot; - HCL validation has failed because the driver version is not validated \&quot;Incompatible-Firmware-Driver\&quot; - HCL validation has failed because the firmware version and driver version is not validated \&quot;Missing-Os-Driver-Info\&quot; - HCL validation was not performed because we are missing os driver information form the inventory \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \&quot;Service-Error\&quot; - HCL data service is available but an error occured when making the request or parsing the response \&quot;Unrecognized-Protocol\&quot; - This service does not recognize the reason code in the response from the HCL data service \&quot;Compatible\&quot; - this component&#39;s inventory data has \&quot;Validated\&quot; status with the HCL. \&quot;Not-Evaluated\&quot; - The component is not evaluated against the HCL because the server is exempted. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version. | [optional] [default to "Missing-Os-Driver-Info"]
**SoftwareStatus** | Pointer to **string** | The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \&quot;Incompatible-Firmware\&quot; - the component&#39;s firmware is not listed under the server&#39;s hardware and software profile and the component&#39;s hardware profile \&quot;Incompatible-Driver\&quot; - the component&#39;s driver is not listed under the server&#39;s hardware and software profile and the component&#39;s hardware profile \&quot;Incompatible-Firmware-Driver\&quot; - the component&#39;s firmware and driver are not listed under the server&#39;s hardware and software profile and the component&#39;s hardware profile \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \&quot;Not-Evaluated\&quot; - the component&#39;s hardware status was not evaluated because the server&#39;s hardware or software profile is not listed or server is exempted. \&quot;Compatible\&quot; - this component&#39;s hardware profile is listed in the HCL. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version. | [optional] [default to "Missing-Os-Driver-Info"]
**Status** | Pointer to **string** | The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \&quot;Unknown\&quot; - we do not have enough information to evaluate against the HCL data \&quot;Validated\&quot; - we have validated this component against the HCL and it has \&quot;Validated\&quot; status \&quot;Not-Validated\&quot; - we have validated this component against the HCL and it has \&quot;Not-Validated\&quot; status. \&quot;Not-Evaluated\&quot; - The component is not evaluated against the HCL because the server is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**Component** | Pointer to [**InventoryBaseRelationship**](inventory.Base.Relationship.md) |  | [optional] 
**HclStatus** | Pointer to [**CondHclStatusRelationship**](cond.HclStatus.Relationship.md) |  | [optional] 

## Methods

### NewCondHclStatusDetail

`func NewCondHclStatusDetail() *CondHclStatusDetail`

NewCondHclStatusDetail instantiates a new CondHclStatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondHclStatusDetailWithDefaults

`func NewCondHclStatusDetailWithDefaults() *CondHclStatusDetail`

NewCondHclStatusDetailWithDefaults instantiates a new CondHclStatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareStatus

`func (o *CondHclStatusDetail) GetHardwareStatus() string`

GetHardwareStatus returns the HardwareStatus field if non-nil, zero value otherwise.

### GetHardwareStatusOk

`func (o *CondHclStatusDetail) GetHardwareStatusOk() (*string, bool)`

GetHardwareStatusOk returns a tuple with the HardwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareStatus

`func (o *CondHclStatusDetail) SetHardwareStatus(v string)`

SetHardwareStatus sets HardwareStatus field to given value.

### HasHardwareStatus

`func (o *CondHclStatusDetail) HasHardwareStatus() bool`

HasHardwareStatus returns a boolean if a field has been set.

### GetHclCimcVersion

`func (o *CondHclStatusDetail) GetHclCimcVersion() string`

GetHclCimcVersion returns the HclCimcVersion field if non-nil, zero value otherwise.

### GetHclCimcVersionOk

`func (o *CondHclStatusDetail) GetHclCimcVersionOk() (*string, bool)`

GetHclCimcVersionOk returns a tuple with the HclCimcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclCimcVersion

`func (o *CondHclStatusDetail) SetHclCimcVersion(v string)`

SetHclCimcVersion sets HclCimcVersion field to given value.

### HasHclCimcVersion

`func (o *CondHclStatusDetail) HasHclCimcVersion() bool`

HasHclCimcVersion returns a boolean if a field has been set.

### GetHclDriverName

`func (o *CondHclStatusDetail) GetHclDriverName() string`

GetHclDriverName returns the HclDriverName field if non-nil, zero value otherwise.

### GetHclDriverNameOk

`func (o *CondHclStatusDetail) GetHclDriverNameOk() (*string, bool)`

GetHclDriverNameOk returns a tuple with the HclDriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclDriverName

`func (o *CondHclStatusDetail) SetHclDriverName(v string)`

SetHclDriverName sets HclDriverName field to given value.

### HasHclDriverName

`func (o *CondHclStatusDetail) HasHclDriverName() bool`

HasHclDriverName returns a boolean if a field has been set.

### GetHclDriverVersion

`func (o *CondHclStatusDetail) GetHclDriverVersion() string`

GetHclDriverVersion returns the HclDriverVersion field if non-nil, zero value otherwise.

### GetHclDriverVersionOk

`func (o *CondHclStatusDetail) GetHclDriverVersionOk() (*string, bool)`

GetHclDriverVersionOk returns a tuple with the HclDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclDriverVersion

`func (o *CondHclStatusDetail) SetHclDriverVersion(v string)`

SetHclDriverVersion sets HclDriverVersion field to given value.

### HasHclDriverVersion

`func (o *CondHclStatusDetail) HasHclDriverVersion() bool`

HasHclDriverVersion returns a boolean if a field has been set.

### GetHclFirmwareVersion

`func (o *CondHclStatusDetail) GetHclFirmwareVersion() string`

GetHclFirmwareVersion returns the HclFirmwareVersion field if non-nil, zero value otherwise.

### GetHclFirmwareVersionOk

`func (o *CondHclStatusDetail) GetHclFirmwareVersionOk() (*string, bool)`

GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclFirmwareVersion

`func (o *CondHclStatusDetail) SetHclFirmwareVersion(v string)`

SetHclFirmwareVersion sets HclFirmwareVersion field to given value.

### HasHclFirmwareVersion

`func (o *CondHclStatusDetail) HasHclFirmwareVersion() bool`

HasHclFirmwareVersion returns a boolean if a field has been set.

### GetHclModel

`func (o *CondHclStatusDetail) GetHclModel() string`

GetHclModel returns the HclModel field if non-nil, zero value otherwise.

### GetHclModelOk

`func (o *CondHclStatusDetail) GetHclModelOk() (*string, bool)`

GetHclModelOk returns a tuple with the HclModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclModel

`func (o *CondHclStatusDetail) SetHclModel(v string)`

SetHclModel sets HclModel field to given value.

### HasHclModel

`func (o *CondHclStatusDetail) HasHclModel() bool`

HasHclModel returns a boolean if a field has been set.

### GetInvCimcVersion

`func (o *CondHclStatusDetail) GetInvCimcVersion() string`

GetInvCimcVersion returns the InvCimcVersion field if non-nil, zero value otherwise.

### GetInvCimcVersionOk

`func (o *CondHclStatusDetail) GetInvCimcVersionOk() (*string, bool)`

GetInvCimcVersionOk returns a tuple with the InvCimcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvCimcVersion

`func (o *CondHclStatusDetail) SetInvCimcVersion(v string)`

SetInvCimcVersion sets InvCimcVersion field to given value.

### HasInvCimcVersion

`func (o *CondHclStatusDetail) HasInvCimcVersion() bool`

HasInvCimcVersion returns a boolean if a field has been set.

### GetInvDriverName

`func (o *CondHclStatusDetail) GetInvDriverName() string`

GetInvDriverName returns the InvDriverName field if non-nil, zero value otherwise.

### GetInvDriverNameOk

`func (o *CondHclStatusDetail) GetInvDriverNameOk() (*string, bool)`

GetInvDriverNameOk returns a tuple with the InvDriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvDriverName

`func (o *CondHclStatusDetail) SetInvDriverName(v string)`

SetInvDriverName sets InvDriverName field to given value.

### HasInvDriverName

`func (o *CondHclStatusDetail) HasInvDriverName() bool`

HasInvDriverName returns a boolean if a field has been set.

### GetInvDriverVersion

`func (o *CondHclStatusDetail) GetInvDriverVersion() string`

GetInvDriverVersion returns the InvDriverVersion field if non-nil, zero value otherwise.

### GetInvDriverVersionOk

`func (o *CondHclStatusDetail) GetInvDriverVersionOk() (*string, bool)`

GetInvDriverVersionOk returns a tuple with the InvDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvDriverVersion

`func (o *CondHclStatusDetail) SetInvDriverVersion(v string)`

SetInvDriverVersion sets InvDriverVersion field to given value.

### HasInvDriverVersion

`func (o *CondHclStatusDetail) HasInvDriverVersion() bool`

HasInvDriverVersion returns a boolean if a field has been set.

### GetInvFirmwareVersion

`func (o *CondHclStatusDetail) GetInvFirmwareVersion() string`

GetInvFirmwareVersion returns the InvFirmwareVersion field if non-nil, zero value otherwise.

### GetInvFirmwareVersionOk

`func (o *CondHclStatusDetail) GetInvFirmwareVersionOk() (*string, bool)`

GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvFirmwareVersion

`func (o *CondHclStatusDetail) SetInvFirmwareVersion(v string)`

SetInvFirmwareVersion sets InvFirmwareVersion field to given value.

### HasInvFirmwareVersion

`func (o *CondHclStatusDetail) HasInvFirmwareVersion() bool`

HasInvFirmwareVersion returns a boolean if a field has been set.

### GetInvModel

`func (o *CondHclStatusDetail) GetInvModel() string`

GetInvModel returns the InvModel field if non-nil, zero value otherwise.

### GetInvModelOk

`func (o *CondHclStatusDetail) GetInvModelOk() (*string, bool)`

GetInvModelOk returns a tuple with the InvModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvModel

`func (o *CondHclStatusDetail) SetInvModel(v string)`

SetInvModel sets InvModel field to given value.

### HasInvModel

`func (o *CondHclStatusDetail) HasInvModel() bool`

HasInvModel returns a boolean if a field has been set.

### GetReason

`func (o *CondHclStatusDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CondHclStatusDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CondHclStatusDetail) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CondHclStatusDetail) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSoftwareStatus

`func (o *CondHclStatusDetail) GetSoftwareStatus() string`

GetSoftwareStatus returns the SoftwareStatus field if non-nil, zero value otherwise.

### GetSoftwareStatusOk

`func (o *CondHclStatusDetail) GetSoftwareStatusOk() (*string, bool)`

GetSoftwareStatusOk returns a tuple with the SoftwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareStatus

`func (o *CondHclStatusDetail) SetSoftwareStatus(v string)`

SetSoftwareStatus sets SoftwareStatus field to given value.

### HasSoftwareStatus

`func (o *CondHclStatusDetail) HasSoftwareStatus() bool`

HasSoftwareStatus returns a boolean if a field has been set.

### GetStatus

`func (o *CondHclStatusDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondHclStatusDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondHclStatusDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondHclStatusDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComponent

`func (o *CondHclStatusDetail) GetComponent() InventoryBaseRelationship`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CondHclStatusDetail) GetComponentOk() (*InventoryBaseRelationship, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CondHclStatusDetail) SetComponent(v InventoryBaseRelationship)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *CondHclStatusDetail) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetHclStatus

`func (o *CondHclStatusDetail) GetHclStatus() CondHclStatusRelationship`

GetHclStatus returns the HclStatus field if non-nil, zero value otherwise.

### GetHclStatusOk

`func (o *CondHclStatusDetail) GetHclStatusOk() (*CondHclStatusRelationship, bool)`

GetHclStatusOk returns a tuple with the HclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatus

`func (o *CondHclStatusDetail) SetHclStatus(v CondHclStatusRelationship)`

SetHclStatus sets HclStatus field to given value.

### HasHclStatus

`func (o *CondHclStatusDetail) HasHclStatus() bool`

HasHclStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


