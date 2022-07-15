# ConvergedinfraBaseComplianceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Name** | Pointer to **string** | The name of the component for which the compliance is evaluated. | [optional] [readonly] 
**Reason** | Pointer to **string** | Reason for the status, when the status is Incomplete or NotApproved. Reason should help to identify the component that is not compliant. * &#x60;NotEvaluated&#x60; - The validation for the HCL or Interop status is yet to be performed. * &#x60;Missing-Os-Info&#x60; - This means the Interop status for the sever cannot be computed because we have missing OS information. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Storage-Os&#x60; - The validation failed because the given storage name and version combination is not found in Interop matrix. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS name and version combination is not found in HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor is not found for the given server PID in HCL. * &#x60;Incompatible-Server-Platform&#x60; - The validation failed because the given server platform is not found in the Interop matrix. * &#x60;Incompatible-Server-Model&#x60; - The validation failed because the given server model is not found in the Interop matrix. * &#x60;Incompatible-Adapter-Model&#x60; - The validation failed because the given adapter model is not found in the Interop matrix. * &#x60;Incompatible-Switch-Model&#x60; - The validation failed because the given switch model is not found in the Interop matrix. * &#x60;Incompatible-Server-Firmware&#x60; - The validation failed because the given server firmware version is not found in HCL. * &#x60;Incompatible-Switch-Firmware&#x60; - The validation failed because the given switch firmware version is not found in Interop matrix. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given adapter firmware version is not found in either HCL or Interop matrix. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version is not found in either HCL or Interop matrix. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given adapter firmware and driver version is not found in either HCL or Interop matrix. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed because the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Missing-Os-Or-Driver-Info&#x60; - This means the Interop status for the switch or storage array cannot be computed because we have either missing Host OS information or missing required driver information. Either install ucstools vib or use power shell scripts to tag proper OS information or install required drivers. * &#x60;Incompatible-Components&#x60; - The validation failed for the given server because one or more of its components failed validation. * &#x60;Compatible&#x60; - This means the HCL status and Interop status for the component have passed all validations for all of its related components. | [optional] [readonly] [default to "NotEvaluated"]
**Status** | Pointer to **string** | Compliance status for the component. * &#x60;NotEvaluated&#x60; - The interoperability compliance for the component has not be checked. * &#x60;Approved&#x60; - The component is valid as per the interoperability compliance check. * &#x60;NotApproved&#x60; - The component is not valid as per the interoperability compliance check. * &#x60;Incomplete&#x60; - The interoperability compliance check could not be completed for the component due to incomplete data. | [optional] [readonly] [default to "NotEvaluated"]

## Methods

### NewConvergedinfraBaseComplianceDetails

`func NewConvergedinfraBaseComplianceDetails(classId string, objectType string, ) *ConvergedinfraBaseComplianceDetails`

NewConvergedinfraBaseComplianceDetails instantiates a new ConvergedinfraBaseComplianceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraBaseComplianceDetailsWithDefaults

`func NewConvergedinfraBaseComplianceDetailsWithDefaults() *ConvergedinfraBaseComplianceDetails`

NewConvergedinfraBaseComplianceDetailsWithDefaults instantiates a new ConvergedinfraBaseComplianceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraBaseComplianceDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraBaseComplianceDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraBaseComplianceDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraBaseComplianceDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraBaseComplianceDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraBaseComplianceDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *ConvergedinfraBaseComplianceDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConvergedinfraBaseComplianceDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConvergedinfraBaseComplianceDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConvergedinfraBaseComplianceDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *ConvergedinfraBaseComplianceDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConvergedinfraBaseComplianceDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConvergedinfraBaseComplianceDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ConvergedinfraBaseComplianceDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *ConvergedinfraBaseComplianceDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConvergedinfraBaseComplianceDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConvergedinfraBaseComplianceDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConvergedinfraBaseComplianceDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


