# ConvergedinfraAdapterComplianceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.AdapterComplianceDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.AdapterComplianceDetails"]
**DriverName** | Pointer to **string** | The driver name (e.g. nenic, nfnic) of the adapter. | [optional] [readonly] 
**DriverVersion** | Pointer to **string** | The driver version of the adapter. | [optional] [readonly] 
**Firmware** | Pointer to **string** | The firmware version of the adapter. | [optional] [readonly] 
**HclStatus** | Pointer to **string** | The HCL compatibility status for the adapter. * &#x60;NotEvaluated&#x60; - The interoperability compliance for the component has not be checked. * &#x60;Approved&#x60; - The component is valid as per the interoperability compliance check. * &#x60;NotApproved&#x60; - The component is not valid as per the interoperability compliance check. * &#x60;Incomplete&#x60; - The interoperability compliance check could not be completed for the component due to incomplete data. | [optional] [readonly] [default to "NotEvaluated"]
**HclStatusReason** | Pointer to **string** | The reason for the HCL status when it is not Approved. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version. | [optional] [readonly] [default to "Missing-Os-Driver-Info"]
**Model** | Pointer to **string** | The model information of the adapter. | [optional] [readonly] 
**Adapter** | Pointer to [**AdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**ServerComplianceDetails** | Pointer to [**ConvergedinfraServerComplianceDetailsRelationship**](ConvergedinfraServerComplianceDetailsRelationship.md) |  | [optional] 
**StorageCompliances** | Pointer to [**[]ConvergedinfraStorageComplianceDetailsRelationship**](ConvergedinfraStorageComplianceDetailsRelationship.md) | An array of relationships to convergedinfraStorageComplianceDetails resources. | [optional] [readonly] 

## Methods

### NewConvergedinfraAdapterComplianceDetails

`func NewConvergedinfraAdapterComplianceDetails(classId string, objectType string, ) *ConvergedinfraAdapterComplianceDetails`

NewConvergedinfraAdapterComplianceDetails instantiates a new ConvergedinfraAdapterComplianceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraAdapterComplianceDetailsWithDefaults

`func NewConvergedinfraAdapterComplianceDetailsWithDefaults() *ConvergedinfraAdapterComplianceDetails`

NewConvergedinfraAdapterComplianceDetailsWithDefaults instantiates a new ConvergedinfraAdapterComplianceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraAdapterComplianceDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraAdapterComplianceDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraAdapterComplianceDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraAdapterComplianceDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriverName

`func (o *ConvergedinfraAdapterComplianceDetails) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *ConvergedinfraAdapterComplianceDetails) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *ConvergedinfraAdapterComplianceDetails) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetDriverVersion

`func (o *ConvergedinfraAdapterComplianceDetails) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *ConvergedinfraAdapterComplianceDetails) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *ConvergedinfraAdapterComplianceDetails) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetFirmware

`func (o *ConvergedinfraAdapterComplianceDetails) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *ConvergedinfraAdapterComplianceDetails) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *ConvergedinfraAdapterComplianceDetails) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetHclStatus

`func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatus() string`

GetHclStatus returns the HclStatus field if non-nil, zero value otherwise.

### GetHclStatusOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatusOk() (*string, bool)`

GetHclStatusOk returns a tuple with the HclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatus

`func (o *ConvergedinfraAdapterComplianceDetails) SetHclStatus(v string)`

SetHclStatus sets HclStatus field to given value.

### HasHclStatus

`func (o *ConvergedinfraAdapterComplianceDetails) HasHclStatus() bool`

HasHclStatus returns a boolean if a field has been set.

### GetHclStatusReason

`func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatusReason() string`

GetHclStatusReason returns the HclStatusReason field if non-nil, zero value otherwise.

### GetHclStatusReasonOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatusReasonOk() (*string, bool)`

GetHclStatusReasonOk returns a tuple with the HclStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatusReason

`func (o *ConvergedinfraAdapterComplianceDetails) SetHclStatusReason(v string)`

SetHclStatusReason sets HclStatusReason field to given value.

### HasHclStatusReason

`func (o *ConvergedinfraAdapterComplianceDetails) HasHclStatusReason() bool`

HasHclStatusReason returns a boolean if a field has been set.

### GetModel

`func (o *ConvergedinfraAdapterComplianceDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConvergedinfraAdapterComplianceDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConvergedinfraAdapterComplianceDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAdapter

`func (o *ConvergedinfraAdapterComplianceDetails) GetAdapter() AdapterUnitRelationship`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetAdapterOk() (*AdapterUnitRelationship, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *ConvergedinfraAdapterComplianceDetails) SetAdapter(v AdapterUnitRelationship)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *ConvergedinfraAdapterComplianceDetails) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetServerComplianceDetails

`func (o *ConvergedinfraAdapterComplianceDetails) GetServerComplianceDetails() ConvergedinfraServerComplianceDetailsRelationship`

GetServerComplianceDetails returns the ServerComplianceDetails field if non-nil, zero value otherwise.

### GetServerComplianceDetailsOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetServerComplianceDetailsOk() (*ConvergedinfraServerComplianceDetailsRelationship, bool)`

GetServerComplianceDetailsOk returns a tuple with the ServerComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComplianceDetails

`func (o *ConvergedinfraAdapterComplianceDetails) SetServerComplianceDetails(v ConvergedinfraServerComplianceDetailsRelationship)`

SetServerComplianceDetails sets ServerComplianceDetails field to given value.

### HasServerComplianceDetails

`func (o *ConvergedinfraAdapterComplianceDetails) HasServerComplianceDetails() bool`

HasServerComplianceDetails returns a boolean if a field has been set.

### GetStorageCompliances

`func (o *ConvergedinfraAdapterComplianceDetails) GetStorageCompliances() []ConvergedinfraStorageComplianceDetailsRelationship`

GetStorageCompliances returns the StorageCompliances field if non-nil, zero value otherwise.

### GetStorageCompliancesOk

`func (o *ConvergedinfraAdapterComplianceDetails) GetStorageCompliancesOk() (*[]ConvergedinfraStorageComplianceDetailsRelationship, bool)`

GetStorageCompliancesOk returns a tuple with the StorageCompliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCompliances

`func (o *ConvergedinfraAdapterComplianceDetails) SetStorageCompliances(v []ConvergedinfraStorageComplianceDetailsRelationship)`

SetStorageCompliances sets StorageCompliances field to given value.

### HasStorageCompliances

`func (o *ConvergedinfraAdapterComplianceDetails) HasStorageCompliances() bool`

HasStorageCompliances returns a boolean if a field has been set.

### SetStorageCompliancesNil

`func (o *ConvergedinfraAdapterComplianceDetails) SetStorageCompliancesNil(b bool)`

 SetStorageCompliancesNil sets the value for StorageCompliances to be an explicit nil

### UnsetStorageCompliances
`func (o *ConvergedinfraAdapterComplianceDetails) UnsetStorageCompliances()`

UnsetStorageCompliances ensures that no value is present for StorageCompliances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


