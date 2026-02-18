# CondServerBaselineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.ServerBaselineStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.ServerBaselineStatus"]
**AdapterCompliance** | Pointer to **string** | Adapter compliance status of the server as per the referred custom Hcl baseline. * &#x60;NotApplicable&#x60; - Indicates that the custom Hcl baseline is not applicable to the server. * &#x60;Compliant&#x60; - Indicates that the server baseline status is validated and compliant with the defined custom Hcl baseline. * &#x60;NonCompliant&#x60; - Indicates that the server baseline status is not validated or does not meet the defined custom Hcl baseline. | [optional] [readonly] [default to "NotApplicable"]
**AdapterReason** | Pointer to **string** | Reason for the component status as per the referred custom Hcl baseline. * &#x60;NotApplicable&#x60; - Reason is not applicable for the specified baseline status. * &#x60;Compliant&#x60; - Custom Hcl Baseline properties are matched with server&#39;s properties. * &#x60;ServerModelNotMatched&#x60; - Server model of the server does not matched with custom Hcl baseline&#39;s server model. * &#x60;ProcessorNotMatched&#x60; - Processor of the server does not matched with custom Hcl baseline&#39;s processor. * &#x60;FirmwareVersionNotMatched&#x60; - Firmware version of the server does not matched with custom Hcl baseline&#39;s firmware version. * &#x60;ManagementModeNotMatched&#x60; - ManagementMode of the server does not matched with custom Hcl baseline&#39;s management mode. * &#x60;GenerationNotMatched&#x60; - Generation of the server does not matched with custom Hcl baseline&#39;s generation. * &#x60;PersonalityNotMatched&#x60; - Personality of the server does not matched with custom Hcl baseline&#39;s personality. * &#x60;OsVendorNotMatched&#x60; - Operating system vendor of the server does not matched with custom Hcl baseline&#39;s operating system vendor. * &#x60;OsVersionNotMatched&#x60; - Operating system version of the server does not matched with custom Hcl baseline&#39;s operating system version. * &#x60;AdapterModelNotMatched&#x60; - Adapter model of the server does not matched with custom Hcl baseline&#39;s adapter model. * &#x60;AdapterFirmwareNotMatched&#x60; - Adapter firmware version of the server does not matched with custom Hcl baseline&#39;s adapter adapter firmware version. * &#x60;AdapterTypeNotMatched&#x60; - Adapter Type of the server does not matched with custom Hcl baseline&#39;s adapter type. * &#x60;AdapterDriverProtocolNotMatched&#x60; - Adapter driver name of the server does not matched with custom Hcl baseline&#39;s adapter driver name. * &#x60;AdapterDriverVersionNotMatched&#x60; - Adapter driver version of the server does not matched with custom Hcl baseline&#39;s adapter driver version. * &#x60;AdapterNotMatched&#x60; - One or morecustom Hcl  adapter  of the server does not matched with custom Hcl baseline&#39;s adapter details. * &#x60;AdapterVendorNotMatched&#x60; - Adapter vendor of the server does not matched with custom Hcl baseline&#39;s adapter vendor. | [optional] [readonly] [default to "NotApplicable"]
**AdaptersStatus** | Pointer to [**[]CondAdapterBaselineStatus**](CondAdapterBaselineStatus.md) |  | [optional] 
**CustomHclBaseline** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**HardwareCompliance** | Pointer to **string** | Hardware compliance status of the server as per the referred custom Hcl baseline. * &#x60;NotApplicable&#x60; - Indicates that the custom Hcl baseline is not applicable to the server. * &#x60;Compliant&#x60; - Indicates that the server baseline status is validated and compliant with the defined custom Hcl baseline. * &#x60;NonCompliant&#x60; - Indicates that the server baseline status is not validated or does not meet the defined custom Hcl baseline. | [optional] [readonly] [default to "NotApplicable"]
**HardwareReason** | Pointer to **string** | Reason for the hardware status as per the referred custom Hcl baseline. * &#x60;NotApplicable&#x60; - Reason is not applicable for the specified baseline status. * &#x60;Compliant&#x60; - Custom Hcl Baseline properties are matched with server&#39;s properties. * &#x60;ServerModelNotMatched&#x60; - Server model of the server does not matched with custom Hcl baseline&#39;s server model. * &#x60;ProcessorNotMatched&#x60; - Processor of the server does not matched with custom Hcl baseline&#39;s processor. * &#x60;FirmwareVersionNotMatched&#x60; - Firmware version of the server does not matched with custom Hcl baseline&#39;s firmware version. * &#x60;ManagementModeNotMatched&#x60; - ManagementMode of the server does not matched with custom Hcl baseline&#39;s management mode. * &#x60;GenerationNotMatched&#x60; - Generation of the server does not matched with custom Hcl baseline&#39;s generation. * &#x60;PersonalityNotMatched&#x60; - Personality of the server does not matched with custom Hcl baseline&#39;s personality. * &#x60;OsVendorNotMatched&#x60; - Operating system vendor of the server does not matched with custom Hcl baseline&#39;s operating system vendor. * &#x60;OsVersionNotMatched&#x60; - Operating system version of the server does not matched with custom Hcl baseline&#39;s operating system version. * &#x60;AdapterModelNotMatched&#x60; - Adapter model of the server does not matched with custom Hcl baseline&#39;s adapter model. * &#x60;AdapterFirmwareNotMatched&#x60; - Adapter firmware version of the server does not matched with custom Hcl baseline&#39;s adapter adapter firmware version. * &#x60;AdapterTypeNotMatched&#x60; - Adapter Type of the server does not matched with custom Hcl baseline&#39;s adapter type. * &#x60;AdapterDriverProtocolNotMatched&#x60; - Adapter driver name of the server does not matched with custom Hcl baseline&#39;s adapter driver name. * &#x60;AdapterDriverVersionNotMatched&#x60; - Adapter driver version of the server does not matched with custom Hcl baseline&#39;s adapter driver version. * &#x60;AdapterNotMatched&#x60; - One or morecustom Hcl  adapter  of the server does not matched with custom Hcl baseline&#39;s adapter details. * &#x60;AdapterVendorNotMatched&#x60; - Adapter vendor of the server does not matched with custom Hcl baseline&#39;s adapter vendor. | [optional] [readonly] [default to "NotApplicable"]
**Reason** | Pointer to **string** | Reason for the current baseline status. * &#x60;NotApplicable&#x60; - Reason is not applicable for the specified baseline status. * &#x60;Compliant&#x60; - Custom Hcl Baseline properties are matched with server&#39;s properties. * &#x60;ServerModelNotMatched&#x60; - Server model of the server does not matched with custom Hcl baseline&#39;s server model. * &#x60;ProcessorNotMatched&#x60; - Processor of the server does not matched with custom Hcl baseline&#39;s processor. * &#x60;FirmwareVersionNotMatched&#x60; - Firmware version of the server does not matched with custom Hcl baseline&#39;s firmware version. * &#x60;ManagementModeNotMatched&#x60; - ManagementMode of the server does not matched with custom Hcl baseline&#39;s management mode. * &#x60;GenerationNotMatched&#x60; - Generation of the server does not matched with custom Hcl baseline&#39;s generation. * &#x60;PersonalityNotMatched&#x60; - Personality of the server does not matched with custom Hcl baseline&#39;s personality. * &#x60;OsVendorNotMatched&#x60; - Operating system vendor of the server does not matched with custom Hcl baseline&#39;s operating system vendor. * &#x60;OsVersionNotMatched&#x60; - Operating system version of the server does not matched with custom Hcl baseline&#39;s operating system version. * &#x60;AdapterModelNotMatched&#x60; - Adapter model of the server does not matched with custom Hcl baseline&#39;s adapter model. * &#x60;AdapterFirmwareNotMatched&#x60; - Adapter firmware version of the server does not matched with custom Hcl baseline&#39;s adapter adapter firmware version. * &#x60;AdapterTypeNotMatched&#x60; - Adapter Type of the server does not matched with custom Hcl baseline&#39;s adapter type. * &#x60;AdapterDriverProtocolNotMatched&#x60; - Adapter driver name of the server does not matched with custom Hcl baseline&#39;s adapter driver name. * &#x60;AdapterDriverVersionNotMatched&#x60; - Adapter driver version of the server does not matched with custom Hcl baseline&#39;s adapter driver version. * &#x60;AdapterNotMatched&#x60; - One or morecustom Hcl  adapter  of the server does not matched with custom Hcl baseline&#39;s adapter details. * &#x60;AdapterVendorNotMatched&#x60; - Adapter vendor of the server does not matched with custom Hcl baseline&#39;s adapter vendor. | [optional] [readonly] [default to "NotApplicable"]
**SoftwareCompliance** | Pointer to **string** | Software compliance status of the server as per the referred custom Hcl baseline. * &#x60;NotApplicable&#x60; - Indicates that the custom Hcl baseline is not applicable to the server. * &#x60;Compliant&#x60; - Indicates that the server baseline status is validated and compliant with the defined custom Hcl baseline. * &#x60;NonCompliant&#x60; - Indicates that the server baseline status is not validated or does not meet the defined custom Hcl baseline. | [optional] [readonly] [default to "NotApplicable"]
**SoftwareReason** | Pointer to **string** | Reason for the software status as per referred custom Hcl baseline. * &#x60;NotApplicable&#x60; - Reason is not applicable for the specified baseline status. * &#x60;Compliant&#x60; - Custom Hcl Baseline properties are matched with server&#39;s properties. * &#x60;ServerModelNotMatched&#x60; - Server model of the server does not matched with custom Hcl baseline&#39;s server model. * &#x60;ProcessorNotMatched&#x60; - Processor of the server does not matched with custom Hcl baseline&#39;s processor. * &#x60;FirmwareVersionNotMatched&#x60; - Firmware version of the server does not matched with custom Hcl baseline&#39;s firmware version. * &#x60;ManagementModeNotMatched&#x60; - ManagementMode of the server does not matched with custom Hcl baseline&#39;s management mode. * &#x60;GenerationNotMatched&#x60; - Generation of the server does not matched with custom Hcl baseline&#39;s generation. * &#x60;PersonalityNotMatched&#x60; - Personality of the server does not matched with custom Hcl baseline&#39;s personality. * &#x60;OsVendorNotMatched&#x60; - Operating system vendor of the server does not matched with custom Hcl baseline&#39;s operating system vendor. * &#x60;OsVersionNotMatched&#x60; - Operating system version of the server does not matched with custom Hcl baseline&#39;s operating system version. * &#x60;AdapterModelNotMatched&#x60; - Adapter model of the server does not matched with custom Hcl baseline&#39;s adapter model. * &#x60;AdapterFirmwareNotMatched&#x60; - Adapter firmware version of the server does not matched with custom Hcl baseline&#39;s adapter adapter firmware version. * &#x60;AdapterTypeNotMatched&#x60; - Adapter Type of the server does not matched with custom Hcl baseline&#39;s adapter type. * &#x60;AdapterDriverProtocolNotMatched&#x60; - Adapter driver name of the server does not matched with custom Hcl baseline&#39;s adapter driver name. * &#x60;AdapterDriverVersionNotMatched&#x60; - Adapter driver version of the server does not matched with custom Hcl baseline&#39;s adapter driver version. * &#x60;AdapterNotMatched&#x60; - One or morecustom Hcl  adapter  of the server does not matched with custom Hcl baseline&#39;s adapter details. * &#x60;AdapterVendorNotMatched&#x60; - Adapter vendor of the server does not matched with custom Hcl baseline&#39;s adapter vendor. | [optional] [readonly] [default to "NotApplicable"]
**Status** | Pointer to **string** | Status of the server as per the referred custom Hcl baseline. * &#x60;NotApplicable&#x60; - Indicates that the custom Hcl baseline is not applicable to the server. * &#x60;Compliant&#x60; - Indicates that the server baseline status is validated and compliant with the defined custom Hcl baseline. * &#x60;NonCompliant&#x60; - Indicates that the server baseline status is not validated or does not meet the defined custom Hcl baseline. | [optional] [readonly] [default to "NotApplicable"]
**ValidationTime** | Pointer to **time.Time** | It specifies baseline status validation time. | [optional] [readonly] 

## Methods

### NewCondServerBaselineStatus

`func NewCondServerBaselineStatus(classId string, objectType string, ) *CondServerBaselineStatus`

NewCondServerBaselineStatus instantiates a new CondServerBaselineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondServerBaselineStatusWithDefaults

`func NewCondServerBaselineStatusWithDefaults() *CondServerBaselineStatus`

NewCondServerBaselineStatusWithDefaults instantiates a new CondServerBaselineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondServerBaselineStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondServerBaselineStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondServerBaselineStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondServerBaselineStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondServerBaselineStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondServerBaselineStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterCompliance

`func (o *CondServerBaselineStatus) GetAdapterCompliance() string`

GetAdapterCompliance returns the AdapterCompliance field if non-nil, zero value otherwise.

### GetAdapterComplianceOk

`func (o *CondServerBaselineStatus) GetAdapterComplianceOk() (*string, bool)`

GetAdapterComplianceOk returns a tuple with the AdapterCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterCompliance

`func (o *CondServerBaselineStatus) SetAdapterCompliance(v string)`

SetAdapterCompliance sets AdapterCompliance field to given value.

### HasAdapterCompliance

`func (o *CondServerBaselineStatus) HasAdapterCompliance() bool`

HasAdapterCompliance returns a boolean if a field has been set.

### GetAdapterReason

`func (o *CondServerBaselineStatus) GetAdapterReason() string`

GetAdapterReason returns the AdapterReason field if non-nil, zero value otherwise.

### GetAdapterReasonOk

`func (o *CondServerBaselineStatus) GetAdapterReasonOk() (*string, bool)`

GetAdapterReasonOk returns a tuple with the AdapterReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterReason

`func (o *CondServerBaselineStatus) SetAdapterReason(v string)`

SetAdapterReason sets AdapterReason field to given value.

### HasAdapterReason

`func (o *CondServerBaselineStatus) HasAdapterReason() bool`

HasAdapterReason returns a boolean if a field has been set.

### GetAdaptersStatus

`func (o *CondServerBaselineStatus) GetAdaptersStatus() []CondAdapterBaselineStatus`

GetAdaptersStatus returns the AdaptersStatus field if non-nil, zero value otherwise.

### GetAdaptersStatusOk

`func (o *CondServerBaselineStatus) GetAdaptersStatusOk() (*[]CondAdapterBaselineStatus, bool)`

GetAdaptersStatusOk returns a tuple with the AdaptersStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptersStatus

`func (o *CondServerBaselineStatus) SetAdaptersStatus(v []CondAdapterBaselineStatus)`

SetAdaptersStatus sets AdaptersStatus field to given value.

### HasAdaptersStatus

`func (o *CondServerBaselineStatus) HasAdaptersStatus() bool`

HasAdaptersStatus returns a boolean if a field has been set.

### SetAdaptersStatusNil

`func (o *CondServerBaselineStatus) SetAdaptersStatusNil(b bool)`

 SetAdaptersStatusNil sets the value for AdaptersStatus to be an explicit nil

### UnsetAdaptersStatus
`func (o *CondServerBaselineStatus) UnsetAdaptersStatus()`

UnsetAdaptersStatus ensures that no value is present for AdaptersStatus, not even an explicit nil
### GetCustomHclBaseline

`func (o *CondServerBaselineStatus) GetCustomHclBaseline() MoMoRef`

GetCustomHclBaseline returns the CustomHclBaseline field if non-nil, zero value otherwise.

### GetCustomHclBaselineOk

`func (o *CondServerBaselineStatus) GetCustomHclBaselineOk() (*MoMoRef, bool)`

GetCustomHclBaselineOk returns a tuple with the CustomHclBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHclBaseline

`func (o *CondServerBaselineStatus) SetCustomHclBaseline(v MoMoRef)`

SetCustomHclBaseline sets CustomHclBaseline field to given value.

### HasCustomHclBaseline

`func (o *CondServerBaselineStatus) HasCustomHclBaseline() bool`

HasCustomHclBaseline returns a boolean if a field has been set.

### GetHardwareCompliance

`func (o *CondServerBaselineStatus) GetHardwareCompliance() string`

GetHardwareCompliance returns the HardwareCompliance field if non-nil, zero value otherwise.

### GetHardwareComplianceOk

`func (o *CondServerBaselineStatus) GetHardwareComplianceOk() (*string, bool)`

GetHardwareComplianceOk returns a tuple with the HardwareCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareCompliance

`func (o *CondServerBaselineStatus) SetHardwareCompliance(v string)`

SetHardwareCompliance sets HardwareCompliance field to given value.

### HasHardwareCompliance

`func (o *CondServerBaselineStatus) HasHardwareCompliance() bool`

HasHardwareCompliance returns a boolean if a field has been set.

### GetHardwareReason

`func (o *CondServerBaselineStatus) GetHardwareReason() string`

GetHardwareReason returns the HardwareReason field if non-nil, zero value otherwise.

### GetHardwareReasonOk

`func (o *CondServerBaselineStatus) GetHardwareReasonOk() (*string, bool)`

GetHardwareReasonOk returns a tuple with the HardwareReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReason

`func (o *CondServerBaselineStatus) SetHardwareReason(v string)`

SetHardwareReason sets HardwareReason field to given value.

### HasHardwareReason

`func (o *CondServerBaselineStatus) HasHardwareReason() bool`

HasHardwareReason returns a boolean if a field has been set.

### GetReason

`func (o *CondServerBaselineStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CondServerBaselineStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CondServerBaselineStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CondServerBaselineStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSoftwareCompliance

`func (o *CondServerBaselineStatus) GetSoftwareCompliance() string`

GetSoftwareCompliance returns the SoftwareCompliance field if non-nil, zero value otherwise.

### GetSoftwareComplianceOk

`func (o *CondServerBaselineStatus) GetSoftwareComplianceOk() (*string, bool)`

GetSoftwareComplianceOk returns a tuple with the SoftwareCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareCompliance

`func (o *CondServerBaselineStatus) SetSoftwareCompliance(v string)`

SetSoftwareCompliance sets SoftwareCompliance field to given value.

### HasSoftwareCompliance

`func (o *CondServerBaselineStatus) HasSoftwareCompliance() bool`

HasSoftwareCompliance returns a boolean if a field has been set.

### GetSoftwareReason

`func (o *CondServerBaselineStatus) GetSoftwareReason() string`

GetSoftwareReason returns the SoftwareReason field if non-nil, zero value otherwise.

### GetSoftwareReasonOk

`func (o *CondServerBaselineStatus) GetSoftwareReasonOk() (*string, bool)`

GetSoftwareReasonOk returns a tuple with the SoftwareReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareReason

`func (o *CondServerBaselineStatus) SetSoftwareReason(v string)`

SetSoftwareReason sets SoftwareReason field to given value.

### HasSoftwareReason

`func (o *CondServerBaselineStatus) HasSoftwareReason() bool`

HasSoftwareReason returns a boolean if a field has been set.

### GetStatus

`func (o *CondServerBaselineStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondServerBaselineStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondServerBaselineStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondServerBaselineStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidationTime

`func (o *CondServerBaselineStatus) GetValidationTime() time.Time`

GetValidationTime returns the ValidationTime field if non-nil, zero value otherwise.

### GetValidationTimeOk

`func (o *CondServerBaselineStatus) GetValidationTimeOk() (*time.Time, bool)`

GetValidationTimeOk returns a tuple with the ValidationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationTime

`func (o *CondServerBaselineStatus) SetValidationTime(v time.Time)`

SetValidationTime sets ValidationTime field to given value.

### HasValidationTime

`func (o *CondServerBaselineStatus) HasValidationTime() bool`

HasValidationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


