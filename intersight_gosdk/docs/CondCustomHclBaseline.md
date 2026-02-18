# CondCustomHclBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.CustomHclBaseline"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.CustomHclBaseline"]
**Adapters** | Pointer to [**[]CondAdapterDetail**](CondAdapterDetail.md) |  | [optional] 
**CompliantServerCount** | Pointer to **int64** | The number of servers that are compliant with this custom Hcl baseline. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the custom Hcl baseline. | [optional] 
**Firmware** | Pointer to **string** | The firmware version of currently running on the server. | [optional] 
**Generation** | Pointer to **string** | It specifies the generation of the server. like M7, M8 etc. | [optional] 
**HclReason** | Pointer to **string** | The reason of the current Cisco HCL status of the custom Hcl baseline. * &#x60;Missing-Os-Info&#x60; - This means the HclStatus for the server failed HCL validation because we have missing operating system information. Either install ucstools vib or use power shell scripts to tag proper operating system information. * &#x60;Incompatible-Components&#x60; - This means the HclStatus for the server failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails. * &#x60;Compatible&#x60; - This means the HclStatus for the server has passed HCL validation for all of its related components. * &#x60;Not-Evaluated&#x60; - This means the HclStatus for the server has not been evaluated because it is exempted. * &#x60;Not-Applicable&#x60; - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. | [optional] [readonly] [default to "Missing-Os-Info"]
**HclStatus** | Pointer to **string** | The Cisco HCL compatibility status of the custom Hcl baseline. The status can be one of the following \&quot;Incomplete\&quot; - there is not enough information to evaluate against the HCL data \&quot;Validated\&quot; - all components have been evaluated against the HCL and they all have \&quot;Validated\&quot; status \&quot;Not-Listed\&quot; - all components have been evaluated against the HCL and one or more have \&quot;Not-Listed\&quot; status. \&quot;Not-Evaluatecustom Hcl d\&quot; - The provided is insufficient for HCL status evaluation or the server is exempted from HCL validation \&quot;Not-Applicable\&quot; - the custom Hcl baseline is not applicable to the server. * &#x60;Incomplete&#x60; - This means we do not have operating system information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper operating system information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that in component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet or the configurations provided are insufficient for HCL status assessment. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. * &#x60;Not-Applicable&#x60; - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. | [optional] [readonly] [default to "Incomplete"]
**ManagementMode** | Pointer to **string** | The management mode at which server is connected to intersight. | [optional] 
**Name** | Pointer to **string** | Name of the custom Hcl baseline. | [optional] 
**NonCompliantServerCount** | Pointer to **int64** | The number of servers that are non-compliant with this custom Hcl baseline. | [optional] [readonly] 
**OperationState** | Pointer to **string** | Operation state specifies the state of custom Hcl baseline whether the server&#39;s baseline status is evaluated or not. \&quot;InProgress\&quot; - Server&#39;s baseline status assessment is currently in progress \&quot;Pending\&quot; - Server&#39;s baseline status assessment is not yet started \&quot;Active\&quot; - Server&#39;s baseline status assessment is completed. * &#x60;Pending&#x60; - Server&#39;s custom hcl status assessment is not yet started. * &#x60;Active&#x60; - Server&#39;s custom hcl status assessment is completed. * &#x60;InProgress&#x60; - Server&#39;s custom hcl status assessment is currently in progress. * &#x60;Failed&#x60; - Server&#39;s custom hcl status assessment is failed. | [optional] [readonly] [default to "Pending"]
**OsVendor** | Pointer to **string** | The operating system vendor name running on the server. | [optional] 
**OsVersion** | Pointer to **string** | Operating System version running on the server. | [optional] 
**ProcessorFamily** | Pointer to **string** | The processor family associated with the server for example processor model Intel (R) Xeon (R) Platinum 8454H or AMD EPYC 9654. will be validated using its corresponding processor family. Processor model Intel (R) Xeon (R) Platinum 8454H -&gt; 4th Gen Intel Xeon Processor Scalable Family Processor model AMD EPYC 9654 -&gt; 4th Gen AMD EPYC 9004 Series Processors. | [optional] 
**ServerModel** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewCondCustomHclBaseline

`func NewCondCustomHclBaseline(classId string, objectType string, ) *CondCustomHclBaseline`

NewCondCustomHclBaseline instantiates a new CondCustomHclBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondCustomHclBaselineWithDefaults

`func NewCondCustomHclBaselineWithDefaults() *CondCustomHclBaseline`

NewCondCustomHclBaselineWithDefaults instantiates a new CondCustomHclBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondCustomHclBaseline) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondCustomHclBaseline) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondCustomHclBaseline) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondCustomHclBaseline) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondCustomHclBaseline) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondCustomHclBaseline) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapters

`func (o *CondCustomHclBaseline) GetAdapters() []CondAdapterDetail`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *CondCustomHclBaseline) GetAdaptersOk() (*[]CondAdapterDetail, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *CondCustomHclBaseline) SetAdapters(v []CondAdapterDetail)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *CondCustomHclBaseline) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *CondCustomHclBaseline) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *CondCustomHclBaseline) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetCompliantServerCount

`func (o *CondCustomHclBaseline) GetCompliantServerCount() int64`

GetCompliantServerCount returns the CompliantServerCount field if non-nil, zero value otherwise.

### GetCompliantServerCountOk

`func (o *CondCustomHclBaseline) GetCompliantServerCountOk() (*int64, bool)`

GetCompliantServerCountOk returns a tuple with the CompliantServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantServerCount

`func (o *CondCustomHclBaseline) SetCompliantServerCount(v int64)`

SetCompliantServerCount sets CompliantServerCount field to given value.

### HasCompliantServerCount

`func (o *CondCustomHclBaseline) HasCompliantServerCount() bool`

HasCompliantServerCount returns a boolean if a field has been set.

### GetDescription

`func (o *CondCustomHclBaseline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondCustomHclBaseline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondCustomHclBaseline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondCustomHclBaseline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirmware

`func (o *CondCustomHclBaseline) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *CondCustomHclBaseline) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *CondCustomHclBaseline) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *CondCustomHclBaseline) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetGeneration

`func (o *CondCustomHclBaseline) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *CondCustomHclBaseline) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *CondCustomHclBaseline) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *CondCustomHclBaseline) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetHclReason

`func (o *CondCustomHclBaseline) GetHclReason() string`

GetHclReason returns the HclReason field if non-nil, zero value otherwise.

### GetHclReasonOk

`func (o *CondCustomHclBaseline) GetHclReasonOk() (*string, bool)`

GetHclReasonOk returns a tuple with the HclReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclReason

`func (o *CondCustomHclBaseline) SetHclReason(v string)`

SetHclReason sets HclReason field to given value.

### HasHclReason

`func (o *CondCustomHclBaseline) HasHclReason() bool`

HasHclReason returns a boolean if a field has been set.

### GetHclStatus

`func (o *CondCustomHclBaseline) GetHclStatus() string`

GetHclStatus returns the HclStatus field if non-nil, zero value otherwise.

### GetHclStatusOk

`func (o *CondCustomHclBaseline) GetHclStatusOk() (*string, bool)`

GetHclStatusOk returns a tuple with the HclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatus

`func (o *CondCustomHclBaseline) SetHclStatus(v string)`

SetHclStatus sets HclStatus field to given value.

### HasHclStatus

`func (o *CondCustomHclBaseline) HasHclStatus() bool`

HasHclStatus returns a boolean if a field has been set.

### GetManagementMode

`func (o *CondCustomHclBaseline) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *CondCustomHclBaseline) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *CondCustomHclBaseline) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *CondCustomHclBaseline) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetName

`func (o *CondCustomHclBaseline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondCustomHclBaseline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondCustomHclBaseline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondCustomHclBaseline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNonCompliantServerCount

`func (o *CondCustomHclBaseline) GetNonCompliantServerCount() int64`

GetNonCompliantServerCount returns the NonCompliantServerCount field if non-nil, zero value otherwise.

### GetNonCompliantServerCountOk

`func (o *CondCustomHclBaseline) GetNonCompliantServerCountOk() (*int64, bool)`

GetNonCompliantServerCountOk returns a tuple with the NonCompliantServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantServerCount

`func (o *CondCustomHclBaseline) SetNonCompliantServerCount(v int64)`

SetNonCompliantServerCount sets NonCompliantServerCount field to given value.

### HasNonCompliantServerCount

`func (o *CondCustomHclBaseline) HasNonCompliantServerCount() bool`

HasNonCompliantServerCount returns a boolean if a field has been set.

### GetOperationState

`func (o *CondCustomHclBaseline) GetOperationState() string`

GetOperationState returns the OperationState field if non-nil, zero value otherwise.

### GetOperationStateOk

`func (o *CondCustomHclBaseline) GetOperationStateOk() (*string, bool)`

GetOperationStateOk returns a tuple with the OperationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationState

`func (o *CondCustomHclBaseline) SetOperationState(v string)`

SetOperationState sets OperationState field to given value.

### HasOperationState

`func (o *CondCustomHclBaseline) HasOperationState() bool`

HasOperationState returns a boolean if a field has been set.

### GetOsVendor

`func (o *CondCustomHclBaseline) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *CondCustomHclBaseline) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *CondCustomHclBaseline) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *CondCustomHclBaseline) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *CondCustomHclBaseline) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *CondCustomHclBaseline) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *CondCustomHclBaseline) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *CondCustomHclBaseline) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProcessorFamily

`func (o *CondCustomHclBaseline) GetProcessorFamily() string`

GetProcessorFamily returns the ProcessorFamily field if non-nil, zero value otherwise.

### GetProcessorFamilyOk

`func (o *CondCustomHclBaseline) GetProcessorFamilyOk() (*string, bool)`

GetProcessorFamilyOk returns a tuple with the ProcessorFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorFamily

`func (o *CondCustomHclBaseline) SetProcessorFamily(v string)`

SetProcessorFamily sets ProcessorFamily field to given value.

### HasProcessorFamily

`func (o *CondCustomHclBaseline) HasProcessorFamily() bool`

HasProcessorFamily returns a boolean if a field has been set.

### GetServerModel

`func (o *CondCustomHclBaseline) GetServerModel() []string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *CondCustomHclBaseline) GetServerModelOk() (*[]string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *CondCustomHclBaseline) SetServerModel(v []string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *CondCustomHclBaseline) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### SetServerModelNil

`func (o *CondCustomHclBaseline) SetServerModelNil(b bool)`

 SetServerModelNil sets the value for ServerModel to be an explicit nil

### UnsetServerModel
`func (o *CondCustomHclBaseline) UnsetServerModel()`

UnsetServerModel ensures that no value is present for ServerModel, not even an explicit nil
### GetOrganization

`func (o *CondCustomHclBaseline) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CondCustomHclBaseline) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CondCustomHclBaseline) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CondCustomHclBaseline) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *CondCustomHclBaseline) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CondCustomHclBaseline) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


