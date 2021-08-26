# CondHclStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.HclStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.HclStatus"]
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
**Details** | Pointer to [**[]CondHclStatusDetailRelationship**](CondHclStatusDetailRelationship.md) | An array of relationships to condHclStatusDetail resources. | [optional] [readonly] 
**ManagedObject** | Pointer to [**InventoryBaseRelationship**](InventoryBaseRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewCondHclStatusAllOf

`func NewCondHclStatusAllOf(classId string, objectType string, ) *CondHclStatusAllOf`

NewCondHclStatusAllOf instantiates a new CondHclStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondHclStatusAllOfWithDefaults

`func NewCondHclStatusAllOfWithDefaults() *CondHclStatusAllOf`

NewCondHclStatusAllOfWithDefaults instantiates a new CondHclStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondHclStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondHclStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondHclStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondHclStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondHclStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondHclStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComponentStatus

`func (o *CondHclStatusAllOf) GetComponentStatus() string`

GetComponentStatus returns the ComponentStatus field if non-nil, zero value otherwise.

### GetComponentStatusOk

`func (o *CondHclStatusAllOf) GetComponentStatusOk() (*string, bool)`

GetComponentStatusOk returns a tuple with the ComponentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentStatus

`func (o *CondHclStatusAllOf) SetComponentStatus(v string)`

SetComponentStatus sets ComponentStatus field to given value.

### HasComponentStatus

`func (o *CondHclStatusAllOf) HasComponentStatus() bool`

HasComponentStatus returns a boolean if a field has been set.

### GetHardwareStatus

`func (o *CondHclStatusAllOf) GetHardwareStatus() string`

GetHardwareStatus returns the HardwareStatus field if non-nil, zero value otherwise.

### GetHardwareStatusOk

`func (o *CondHclStatusAllOf) GetHardwareStatusOk() (*string, bool)`

GetHardwareStatusOk returns a tuple with the HardwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareStatus

`func (o *CondHclStatusAllOf) SetHardwareStatus(v string)`

SetHardwareStatus sets HardwareStatus field to given value.

### HasHardwareStatus

`func (o *CondHclStatusAllOf) HasHardwareStatus() bool`

HasHardwareStatus returns a boolean if a field has been set.

### GetHclFirmwareVersion

`func (o *CondHclStatusAllOf) GetHclFirmwareVersion() string`

GetHclFirmwareVersion returns the HclFirmwareVersion field if non-nil, zero value otherwise.

### GetHclFirmwareVersionOk

`func (o *CondHclStatusAllOf) GetHclFirmwareVersionOk() (*string, bool)`

GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclFirmwareVersion

`func (o *CondHclStatusAllOf) SetHclFirmwareVersion(v string)`

SetHclFirmwareVersion sets HclFirmwareVersion field to given value.

### HasHclFirmwareVersion

`func (o *CondHclStatusAllOf) HasHclFirmwareVersion() bool`

HasHclFirmwareVersion returns a boolean if a field has been set.

### GetHclModel

`func (o *CondHclStatusAllOf) GetHclModel() string`

GetHclModel returns the HclModel field if non-nil, zero value otherwise.

### GetHclModelOk

`func (o *CondHclStatusAllOf) GetHclModelOk() (*string, bool)`

GetHclModelOk returns a tuple with the HclModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclModel

`func (o *CondHclStatusAllOf) SetHclModel(v string)`

SetHclModel sets HclModel field to given value.

### HasHclModel

`func (o *CondHclStatusAllOf) HasHclModel() bool`

HasHclModel returns a boolean if a field has been set.

### GetHclOsVendor

`func (o *CondHclStatusAllOf) GetHclOsVendor() string`

GetHclOsVendor returns the HclOsVendor field if non-nil, zero value otherwise.

### GetHclOsVendorOk

`func (o *CondHclStatusAllOf) GetHclOsVendorOk() (*string, bool)`

GetHclOsVendorOk returns a tuple with the HclOsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclOsVendor

`func (o *CondHclStatusAllOf) SetHclOsVendor(v string)`

SetHclOsVendor sets HclOsVendor field to given value.

### HasHclOsVendor

`func (o *CondHclStatusAllOf) HasHclOsVendor() bool`

HasHclOsVendor returns a boolean if a field has been set.

### GetHclOsVersion

`func (o *CondHclStatusAllOf) GetHclOsVersion() string`

GetHclOsVersion returns the HclOsVersion field if non-nil, zero value otherwise.

### GetHclOsVersionOk

`func (o *CondHclStatusAllOf) GetHclOsVersionOk() (*string, bool)`

GetHclOsVersionOk returns a tuple with the HclOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclOsVersion

`func (o *CondHclStatusAllOf) SetHclOsVersion(v string)`

SetHclOsVersion sets HclOsVersion field to given value.

### HasHclOsVersion

`func (o *CondHclStatusAllOf) HasHclOsVersion() bool`

HasHclOsVersion returns a boolean if a field has been set.

### GetHclProcessor

`func (o *CondHclStatusAllOf) GetHclProcessor() string`

GetHclProcessor returns the HclProcessor field if non-nil, zero value otherwise.

### GetHclProcessorOk

`func (o *CondHclStatusAllOf) GetHclProcessorOk() (*string, bool)`

GetHclProcessorOk returns a tuple with the HclProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclProcessor

`func (o *CondHclStatusAllOf) SetHclProcessor(v string)`

SetHclProcessor sets HclProcessor field to given value.

### HasHclProcessor

`func (o *CondHclStatusAllOf) HasHclProcessor() bool`

HasHclProcessor returns a boolean if a field has been set.

### GetInvFirmwareVersion

`func (o *CondHclStatusAllOf) GetInvFirmwareVersion() string`

GetInvFirmwareVersion returns the InvFirmwareVersion field if non-nil, zero value otherwise.

### GetInvFirmwareVersionOk

`func (o *CondHclStatusAllOf) GetInvFirmwareVersionOk() (*string, bool)`

GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvFirmwareVersion

`func (o *CondHclStatusAllOf) SetInvFirmwareVersion(v string)`

SetInvFirmwareVersion sets InvFirmwareVersion field to given value.

### HasInvFirmwareVersion

`func (o *CondHclStatusAllOf) HasInvFirmwareVersion() bool`

HasInvFirmwareVersion returns a boolean if a field has been set.

### GetInvModel

`func (o *CondHclStatusAllOf) GetInvModel() string`

GetInvModel returns the InvModel field if non-nil, zero value otherwise.

### GetInvModelOk

`func (o *CondHclStatusAllOf) GetInvModelOk() (*string, bool)`

GetInvModelOk returns a tuple with the InvModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvModel

`func (o *CondHclStatusAllOf) SetInvModel(v string)`

SetInvModel sets InvModel field to given value.

### HasInvModel

`func (o *CondHclStatusAllOf) HasInvModel() bool`

HasInvModel returns a boolean if a field has been set.

### GetInvOsVendor

`func (o *CondHclStatusAllOf) GetInvOsVendor() string`

GetInvOsVendor returns the InvOsVendor field if non-nil, zero value otherwise.

### GetInvOsVendorOk

`func (o *CondHclStatusAllOf) GetInvOsVendorOk() (*string, bool)`

GetInvOsVendorOk returns a tuple with the InvOsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvOsVendor

`func (o *CondHclStatusAllOf) SetInvOsVendor(v string)`

SetInvOsVendor sets InvOsVendor field to given value.

### HasInvOsVendor

`func (o *CondHclStatusAllOf) HasInvOsVendor() bool`

HasInvOsVendor returns a boolean if a field has been set.

### GetInvOsVersion

`func (o *CondHclStatusAllOf) GetInvOsVersion() string`

GetInvOsVersion returns the InvOsVersion field if non-nil, zero value otherwise.

### GetInvOsVersionOk

`func (o *CondHclStatusAllOf) GetInvOsVersionOk() (*string, bool)`

GetInvOsVersionOk returns a tuple with the InvOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvOsVersion

`func (o *CondHclStatusAllOf) SetInvOsVersion(v string)`

SetInvOsVersion sets InvOsVersion field to given value.

### HasInvOsVersion

`func (o *CondHclStatusAllOf) HasInvOsVersion() bool`

HasInvOsVersion returns a boolean if a field has been set.

### GetInvProcessor

`func (o *CondHclStatusAllOf) GetInvProcessor() string`

GetInvProcessor returns the InvProcessor field if non-nil, zero value otherwise.

### GetInvProcessorOk

`func (o *CondHclStatusAllOf) GetInvProcessorOk() (*string, bool)`

GetInvProcessorOk returns a tuple with the InvProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvProcessor

`func (o *CondHclStatusAllOf) SetInvProcessor(v string)`

SetInvProcessor sets InvProcessor field to given value.

### HasInvProcessor

`func (o *CondHclStatusAllOf) HasInvProcessor() bool`

HasInvProcessor returns a boolean if a field has been set.

### GetReason

`func (o *CondHclStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CondHclStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CondHclStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CondHclStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetServerReason

`func (o *CondHclStatusAllOf) GetServerReason() string`

GetServerReason returns the ServerReason field if non-nil, zero value otherwise.

### GetServerReasonOk

`func (o *CondHclStatusAllOf) GetServerReasonOk() (*string, bool)`

GetServerReasonOk returns a tuple with the ServerReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerReason

`func (o *CondHclStatusAllOf) SetServerReason(v string)`

SetServerReason sets ServerReason field to given value.

### HasServerReason

`func (o *CondHclStatusAllOf) HasServerReason() bool`

HasServerReason returns a boolean if a field has been set.

### GetSoftwareStatus

`func (o *CondHclStatusAllOf) GetSoftwareStatus() string`

GetSoftwareStatus returns the SoftwareStatus field if non-nil, zero value otherwise.

### GetSoftwareStatusOk

`func (o *CondHclStatusAllOf) GetSoftwareStatusOk() (*string, bool)`

GetSoftwareStatusOk returns a tuple with the SoftwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareStatus

`func (o *CondHclStatusAllOf) SetSoftwareStatus(v string)`

SetSoftwareStatus sets SoftwareStatus field to given value.

### HasSoftwareStatus

`func (o *CondHclStatusAllOf) HasSoftwareStatus() bool`

HasSoftwareStatus returns a boolean if a field has been set.

### GetStatus

`func (o *CondHclStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondHclStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondHclStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondHclStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetails

`func (o *CondHclStatusAllOf) GetDetails() []CondHclStatusDetailRelationship`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CondHclStatusAllOf) GetDetailsOk() (*[]CondHclStatusDetailRelationship, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CondHclStatusAllOf) SetDetails(v []CondHclStatusDetailRelationship)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CondHclStatusAllOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CondHclStatusAllOf) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CondHclStatusAllOf) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetManagedObject

`func (o *CondHclStatusAllOf) GetManagedObject() InventoryBaseRelationship`

GetManagedObject returns the ManagedObject field if non-nil, zero value otherwise.

### GetManagedObjectOk

`func (o *CondHclStatusAllOf) GetManagedObjectOk() (*InventoryBaseRelationship, bool)`

GetManagedObjectOk returns a tuple with the ManagedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedObject

`func (o *CondHclStatusAllOf) SetManagedObject(v InventoryBaseRelationship)`

SetManagedObject sets ManagedObject field to given value.

### HasManagedObject

`func (o *CondHclStatusAllOf) HasManagedObject() bool`

HasManagedObject returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CondHclStatusAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondHclStatusAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondHclStatusAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondHclStatusAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


