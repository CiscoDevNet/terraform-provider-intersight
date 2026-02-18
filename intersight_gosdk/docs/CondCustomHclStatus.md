# CondCustomHclStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.CustomHclStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.CustomHclStatus"]
**Adapters** | Pointer to [**[]CondAdapterInstance**](CondAdapterInstance.md) |  | [optional] 
**BaselineStatusDetails** | Pointer to [**[]CondServerBaselineStatus**](CondServerBaselineStatus.md) |  | [optional] 
**CiscoHclStatus** | Pointer to **string** | The HCL compatibility status of the server. The status can be one of the following \&quot;Incomplete\&quot; - there is no enough information to evaluate against the HCL data \&quot;Validated\&quot; - all components have been evaluated against the HCL and they all have \&quot;Validated\&quot; status \&quot;Not-Listed\&quot; - all components have been evaluated against the HCL and one or more have \&quot;Not-Listed\&quot; status. \&quot;Not-Evaluated\&quot; - The server was not evaluated against the HCL because it is exempcustom Hcl t or the provided is insufficient for HCL status assessment. * &#x60;Incomplete&#x60; - This means we do not have operating system information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper operating system information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that in component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet or the configurations provided are insufficient for HCL status assessment. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. * &#x60;Not-Applicable&#x60; - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. | [optional] [readonly] [default to "Incomplete"]
**CustomHclStatus** | Pointer to **string** | The custom HCL overall status against all the defined custom Hcl baseline. The status can be one of the following \&quot;CompliantToBaseline\&quot; - The server is compliant to one or all the defined custom Hcl baseline. \&quot;NonCompliantToBaseline\&quot; - The server is non compliant to any custom Hcl baseline. * &#x60;NonCompliantToBaseline&#x60; - The server is non compliant to any custom Hcl baseline. * &#x60;CompliantToBaseline&#x60; - The server is compliant to one or all the defined custom Hcl baseline. | [optional] [readonly] [default to "NonCompliantToBaseline"]
**Firmware** | Pointer to **string** | The firmware version currently running on the server. | [optional] [readonly] 
**Generation** | Pointer to **string** | It specifies the generation of the server like M6, M7, M8 etc. | [optional] [readonly] 
**ManagementMode** | Pointer to **string** | The management mode through which the server is connected to Intersight, such as standalone or managed mode. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [readonly] [default to "IntersightStandalone"]
**Name** | Pointer to **string** | It specifies the server name or model. | [optional] [readonly] 
**OsVendor** | Pointer to **string** | The operating System vendor of the server. | [optional] [readonly] 
**OsVersion** | Pointer to **string** | The operating System version of the server. | [optional] [readonly] 
**Personality** | Pointer to **string** | Unique identity of added software personality. | [optional] [readonly] 
**ProcessorFamily** | Pointer to **string** | The processor family of the specified processor model associated with the server. | [optional] [readonly] 
**ProcessorModel** | Pointer to **string** | The processor model associated with the server. | [optional] [readonly] 
**ServerModel** | Pointer to **string** | It specifies the server model or Product ID (PID). | [optional] [readonly] 
**CompliantBaseline** | Pointer to [**[]CondCustomHclBaselineRelationship**](CondCustomHclBaselineRelationship.md) | An array of relationships to condCustomHclBaseline resources. | [optional] [readonly] 
**ManagedObject** | Pointer to [**NullableInventoryBaseRelationship**](InventoryBaseRelationship.md) |  | [optional] 
**NonCompliantBaseline** | Pointer to [**[]CondCustomHclBaselineRelationship**](CondCustomHclBaselineRelationship.md) | An array of relationships to condCustomHclBaseline resources. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewCondCustomHclStatus

`func NewCondCustomHclStatus(classId string, objectType string, ) *CondCustomHclStatus`

NewCondCustomHclStatus instantiates a new CondCustomHclStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondCustomHclStatusWithDefaults

`func NewCondCustomHclStatusWithDefaults() *CondCustomHclStatus`

NewCondCustomHclStatusWithDefaults instantiates a new CondCustomHclStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondCustomHclStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondCustomHclStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondCustomHclStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondCustomHclStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondCustomHclStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondCustomHclStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapters

`func (o *CondCustomHclStatus) GetAdapters() []CondAdapterInstance`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *CondCustomHclStatus) GetAdaptersOk() (*[]CondAdapterInstance, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *CondCustomHclStatus) SetAdapters(v []CondAdapterInstance)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *CondCustomHclStatus) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *CondCustomHclStatus) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *CondCustomHclStatus) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetBaselineStatusDetails

`func (o *CondCustomHclStatus) GetBaselineStatusDetails() []CondServerBaselineStatus`

GetBaselineStatusDetails returns the BaselineStatusDetails field if non-nil, zero value otherwise.

### GetBaselineStatusDetailsOk

`func (o *CondCustomHclStatus) GetBaselineStatusDetailsOk() (*[]CondServerBaselineStatus, bool)`

GetBaselineStatusDetailsOk returns a tuple with the BaselineStatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineStatusDetails

`func (o *CondCustomHclStatus) SetBaselineStatusDetails(v []CondServerBaselineStatus)`

SetBaselineStatusDetails sets BaselineStatusDetails field to given value.

### HasBaselineStatusDetails

`func (o *CondCustomHclStatus) HasBaselineStatusDetails() bool`

HasBaselineStatusDetails returns a boolean if a field has been set.

### SetBaselineStatusDetailsNil

`func (o *CondCustomHclStatus) SetBaselineStatusDetailsNil(b bool)`

 SetBaselineStatusDetailsNil sets the value for BaselineStatusDetails to be an explicit nil

### UnsetBaselineStatusDetails
`func (o *CondCustomHclStatus) UnsetBaselineStatusDetails()`

UnsetBaselineStatusDetails ensures that no value is present for BaselineStatusDetails, not even an explicit nil
### GetCiscoHclStatus

`func (o *CondCustomHclStatus) GetCiscoHclStatus() string`

GetCiscoHclStatus returns the CiscoHclStatus field if non-nil, zero value otherwise.

### GetCiscoHclStatusOk

`func (o *CondCustomHclStatus) GetCiscoHclStatusOk() (*string, bool)`

GetCiscoHclStatusOk returns a tuple with the CiscoHclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoHclStatus

`func (o *CondCustomHclStatus) SetCiscoHclStatus(v string)`

SetCiscoHclStatus sets CiscoHclStatus field to given value.

### HasCiscoHclStatus

`func (o *CondCustomHclStatus) HasCiscoHclStatus() bool`

HasCiscoHclStatus returns a boolean if a field has been set.

### GetCustomHclStatus

`func (o *CondCustomHclStatus) GetCustomHclStatus() string`

GetCustomHclStatus returns the CustomHclStatus field if non-nil, zero value otherwise.

### GetCustomHclStatusOk

`func (o *CondCustomHclStatus) GetCustomHclStatusOk() (*string, bool)`

GetCustomHclStatusOk returns a tuple with the CustomHclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHclStatus

`func (o *CondCustomHclStatus) SetCustomHclStatus(v string)`

SetCustomHclStatus sets CustomHclStatus field to given value.

### HasCustomHclStatus

`func (o *CondCustomHclStatus) HasCustomHclStatus() bool`

HasCustomHclStatus returns a boolean if a field has been set.

### GetFirmware

`func (o *CondCustomHclStatus) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *CondCustomHclStatus) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *CondCustomHclStatus) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *CondCustomHclStatus) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetGeneration

`func (o *CondCustomHclStatus) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *CondCustomHclStatus) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *CondCustomHclStatus) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *CondCustomHclStatus) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetManagementMode

`func (o *CondCustomHclStatus) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *CondCustomHclStatus) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *CondCustomHclStatus) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *CondCustomHclStatus) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetName

`func (o *CondCustomHclStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondCustomHclStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondCustomHclStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondCustomHclStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsVendor

`func (o *CondCustomHclStatus) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *CondCustomHclStatus) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *CondCustomHclStatus) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *CondCustomHclStatus) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *CondCustomHclStatus) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *CondCustomHclStatus) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *CondCustomHclStatus) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *CondCustomHclStatus) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPersonality

`func (o *CondCustomHclStatus) GetPersonality() string`

GetPersonality returns the Personality field if non-nil, zero value otherwise.

### GetPersonalityOk

`func (o *CondCustomHclStatus) GetPersonalityOk() (*string, bool)`

GetPersonalityOk returns a tuple with the Personality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonality

`func (o *CondCustomHclStatus) SetPersonality(v string)`

SetPersonality sets Personality field to given value.

### HasPersonality

`func (o *CondCustomHclStatus) HasPersonality() bool`

HasPersonality returns a boolean if a field has been set.

### GetProcessorFamily

`func (o *CondCustomHclStatus) GetProcessorFamily() string`

GetProcessorFamily returns the ProcessorFamily field if non-nil, zero value otherwise.

### GetProcessorFamilyOk

`func (o *CondCustomHclStatus) GetProcessorFamilyOk() (*string, bool)`

GetProcessorFamilyOk returns a tuple with the ProcessorFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorFamily

`func (o *CondCustomHclStatus) SetProcessorFamily(v string)`

SetProcessorFamily sets ProcessorFamily field to given value.

### HasProcessorFamily

`func (o *CondCustomHclStatus) HasProcessorFamily() bool`

HasProcessorFamily returns a boolean if a field has been set.

### GetProcessorModel

`func (o *CondCustomHclStatus) GetProcessorModel() string`

GetProcessorModel returns the ProcessorModel field if non-nil, zero value otherwise.

### GetProcessorModelOk

`func (o *CondCustomHclStatus) GetProcessorModelOk() (*string, bool)`

GetProcessorModelOk returns a tuple with the ProcessorModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorModel

`func (o *CondCustomHclStatus) SetProcessorModel(v string)`

SetProcessorModel sets ProcessorModel field to given value.

### HasProcessorModel

`func (o *CondCustomHclStatus) HasProcessorModel() bool`

HasProcessorModel returns a boolean if a field has been set.

### GetServerModel

`func (o *CondCustomHclStatus) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *CondCustomHclStatus) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *CondCustomHclStatus) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *CondCustomHclStatus) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### GetCompliantBaseline

`func (o *CondCustomHclStatus) GetCompliantBaseline() []CondCustomHclBaselineRelationship`

GetCompliantBaseline returns the CompliantBaseline field if non-nil, zero value otherwise.

### GetCompliantBaselineOk

`func (o *CondCustomHclStatus) GetCompliantBaselineOk() (*[]CondCustomHclBaselineRelationship, bool)`

GetCompliantBaselineOk returns a tuple with the CompliantBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantBaseline

`func (o *CondCustomHclStatus) SetCompliantBaseline(v []CondCustomHclBaselineRelationship)`

SetCompliantBaseline sets CompliantBaseline field to given value.

### HasCompliantBaseline

`func (o *CondCustomHclStatus) HasCompliantBaseline() bool`

HasCompliantBaseline returns a boolean if a field has been set.

### SetCompliantBaselineNil

`func (o *CondCustomHclStatus) SetCompliantBaselineNil(b bool)`

 SetCompliantBaselineNil sets the value for CompliantBaseline to be an explicit nil

### UnsetCompliantBaseline
`func (o *CondCustomHclStatus) UnsetCompliantBaseline()`

UnsetCompliantBaseline ensures that no value is present for CompliantBaseline, not even an explicit nil
### GetManagedObject

`func (o *CondCustomHclStatus) GetManagedObject() InventoryBaseRelationship`

GetManagedObject returns the ManagedObject field if non-nil, zero value otherwise.

### GetManagedObjectOk

`func (o *CondCustomHclStatus) GetManagedObjectOk() (*InventoryBaseRelationship, bool)`

GetManagedObjectOk returns a tuple with the ManagedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedObject

`func (o *CondCustomHclStatus) SetManagedObject(v InventoryBaseRelationship)`

SetManagedObject sets ManagedObject field to given value.

### HasManagedObject

`func (o *CondCustomHclStatus) HasManagedObject() bool`

HasManagedObject returns a boolean if a field has been set.

### SetManagedObjectNil

`func (o *CondCustomHclStatus) SetManagedObjectNil(b bool)`

 SetManagedObjectNil sets the value for ManagedObject to be an explicit nil

### UnsetManagedObject
`func (o *CondCustomHclStatus) UnsetManagedObject()`

UnsetManagedObject ensures that no value is present for ManagedObject, not even an explicit nil
### GetNonCompliantBaseline

`func (o *CondCustomHclStatus) GetNonCompliantBaseline() []CondCustomHclBaselineRelationship`

GetNonCompliantBaseline returns the NonCompliantBaseline field if non-nil, zero value otherwise.

### GetNonCompliantBaselineOk

`func (o *CondCustomHclStatus) GetNonCompliantBaselineOk() (*[]CondCustomHclBaselineRelationship, bool)`

GetNonCompliantBaselineOk returns a tuple with the NonCompliantBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantBaseline

`func (o *CondCustomHclStatus) SetNonCompliantBaseline(v []CondCustomHclBaselineRelationship)`

SetNonCompliantBaseline sets NonCompliantBaseline field to given value.

### HasNonCompliantBaseline

`func (o *CondCustomHclStatus) HasNonCompliantBaseline() bool`

HasNonCompliantBaseline returns a boolean if a field has been set.

### SetNonCompliantBaselineNil

`func (o *CondCustomHclStatus) SetNonCompliantBaselineNil(b bool)`

 SetNonCompliantBaselineNil sets the value for NonCompliantBaseline to be an explicit nil

### UnsetNonCompliantBaseline
`func (o *CondCustomHclStatus) UnsetNonCompliantBaseline()`

UnsetNonCompliantBaseline ensures that no value is present for NonCompliantBaseline, not even an explicit nil
### GetOrganization

`func (o *CondCustomHclStatus) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CondCustomHclStatus) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CondCustomHclStatus) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CondCustomHclStatus) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *CondCustomHclStatus) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CondCustomHclStatus) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


