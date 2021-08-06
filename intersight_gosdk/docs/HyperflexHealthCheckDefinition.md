# HyperflexHealthCheckDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HealthCheckDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HealthCheckDefinition"]
**Category** | Pointer to **string** | Category that the health check belongs to. | [optional] 
**CommonCauses** | Pointer to **string** | Static information detailing the common causes for the health check failure. | [optional] 
**Configuration** | Pointer to **string** | Execution configuration fo the health check script. | [optional] 
**DefaultHealthCheckScriptInfo** | Pointer to [**NullableHyperflexHealthCheckScriptInfo**](hyperflex.HealthCheckScriptInfo.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the health check definition. | [optional] 
**HealthCheckScriptInfos** | Pointer to [**[]HyperflexHealthCheckScriptInfo**](HyperflexHealthCheckScriptInfo.md) |  | [optional] 
**HealthImpact** | Pointer to **string** | Static information detailing the health impact of the health check failure. | [optional] 
**InternalName** | Pointer to **string** | Internal name of the health check definition. | [optional] 
**MinimumHyperFlexVersion** | Pointer to **string** | Minimum HyperFlex version that the check is supported on. Defaults to version 3.0.1. | [optional] [default to "3.0"]
**Name** | Pointer to **string** | Name of the health check definition. | [optional] 
**Reference** | Pointer to **string** | Static information containing additional reference for the health check. | [optional] 
**Resolution** | Pointer to **string** | Static information detailing the possible remediation actions that can be taken to remedy the health check failure. | [optional] 
**ScriptExecutionMode** | Pointer to **string** | Execution mode of the health script on the HyperFlex cluster. * &#x60;ON_DEMAND&#x60; - Execute the health check on-demand. * &#x60;SCHEDULED&#x60; - Execute the health check on a scheduled interval. | [optional] [default to "ON_DEMAND"]
**ScriptExecutionOnComputeNodes** | Pointer to **bool** | Indicates if the script needs to be executed on HyperFlex compute nodes. | Typically, scripts are only executed on the storage Nodes. | [optional] 
**SupportedHypervisorType** | Pointer to **string** | Hypervisor type that the Health Check is supported on (All, if it is hypervisor agnostic). * &#x60;All&#x60; - The Health Check is hypervisor-agnostic. * &#x60;ESXi&#x60; - The Health Check is supported only on Vmware ESXi hypervisor of any version. * &#x60;IWE&#x60; - The Health Check is supported only on Cisco IWE platform. * &#x60;HyperV&#x60; - The Health Check is supported only on Microsoft HyperV hypervisor. | [optional] [readonly] [default to "All"]
**TargetExecutionType** | Pointer to **string** | Indicates whether the health check is executed only on the leader node, or on all nodes in the HyperFlex cluster. * &#x60;EXECUTE_ON_LEADER_NODE&#x60; - Execute the health check script only on the HyperFlex cluster&#39;s leader node. * &#x60;EXECUTE_ON_ALL_NODES&#x60; - Execute health check on all nodes and aggregate the results. * &#x60;EXECUTE_ON_ALL_NODES_AND_AGGREGATE&#x60; - Execute the health check on all Nodes and perform custom aggregation. | [optional] [default to "EXECUTE_ON_LEADER_NODE"]
**Timeout** | Pointer to **int64** | Health check script execution timeout. | [optional] 
**UnsupportedHyperFlexVersions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHyperflexHealthCheckDefinition

`func NewHyperflexHealthCheckDefinition(classId string, objectType string, ) *HyperflexHealthCheckDefinition`

NewHyperflexHealthCheckDefinition instantiates a new HyperflexHealthCheckDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckDefinitionWithDefaults

`func NewHyperflexHealthCheckDefinitionWithDefaults() *HyperflexHealthCheckDefinition`

NewHyperflexHealthCheckDefinitionWithDefaults instantiates a new HyperflexHealthCheckDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HyperflexHealthCheckDefinition) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HyperflexHealthCheckDefinition) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HyperflexHealthCheckDefinition) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HyperflexHealthCheckDefinition) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCommonCauses

`func (o *HyperflexHealthCheckDefinition) GetCommonCauses() string`

GetCommonCauses returns the CommonCauses field if non-nil, zero value otherwise.

### GetCommonCausesOk

`func (o *HyperflexHealthCheckDefinition) GetCommonCausesOk() (*string, bool)`

GetCommonCausesOk returns a tuple with the CommonCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonCauses

`func (o *HyperflexHealthCheckDefinition) SetCommonCauses(v string)`

SetCommonCauses sets CommonCauses field to given value.

### HasCommonCauses

`func (o *HyperflexHealthCheckDefinition) HasCommonCauses() bool`

HasCommonCauses returns a boolean if a field has been set.

### GetConfiguration

`func (o *HyperflexHealthCheckDefinition) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *HyperflexHealthCheckDefinition) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *HyperflexHealthCheckDefinition) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *HyperflexHealthCheckDefinition) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDefaultHealthCheckScriptInfo

`func (o *HyperflexHealthCheckDefinition) GetDefaultHealthCheckScriptInfo() HyperflexHealthCheckScriptInfo`

GetDefaultHealthCheckScriptInfo returns the DefaultHealthCheckScriptInfo field if non-nil, zero value otherwise.

### GetDefaultHealthCheckScriptInfoOk

`func (o *HyperflexHealthCheckDefinition) GetDefaultHealthCheckScriptInfoOk() (*HyperflexHealthCheckScriptInfo, bool)`

GetDefaultHealthCheckScriptInfoOk returns a tuple with the DefaultHealthCheckScriptInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHealthCheckScriptInfo

`func (o *HyperflexHealthCheckDefinition) SetDefaultHealthCheckScriptInfo(v HyperflexHealthCheckScriptInfo)`

SetDefaultHealthCheckScriptInfo sets DefaultHealthCheckScriptInfo field to given value.

### HasDefaultHealthCheckScriptInfo

`func (o *HyperflexHealthCheckDefinition) HasDefaultHealthCheckScriptInfo() bool`

HasDefaultHealthCheckScriptInfo returns a boolean if a field has been set.

### SetDefaultHealthCheckScriptInfoNil

`func (o *HyperflexHealthCheckDefinition) SetDefaultHealthCheckScriptInfoNil(b bool)`

 SetDefaultHealthCheckScriptInfoNil sets the value for DefaultHealthCheckScriptInfo to be an explicit nil

### UnsetDefaultHealthCheckScriptInfo
`func (o *HyperflexHealthCheckDefinition) UnsetDefaultHealthCheckScriptInfo()`

UnsetDefaultHealthCheckScriptInfo ensures that no value is present for DefaultHealthCheckScriptInfo, not even an explicit nil
### GetDescription

`func (o *HyperflexHealthCheckDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexHealthCheckDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexHealthCheckDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexHealthCheckDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHealthCheckScriptInfos

`func (o *HyperflexHealthCheckDefinition) GetHealthCheckScriptInfos() []HyperflexHealthCheckScriptInfo`

GetHealthCheckScriptInfos returns the HealthCheckScriptInfos field if non-nil, zero value otherwise.

### GetHealthCheckScriptInfosOk

`func (o *HyperflexHealthCheckDefinition) GetHealthCheckScriptInfosOk() (*[]HyperflexHealthCheckScriptInfo, bool)`

GetHealthCheckScriptInfosOk returns a tuple with the HealthCheckScriptInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckScriptInfos

`func (o *HyperflexHealthCheckDefinition) SetHealthCheckScriptInfos(v []HyperflexHealthCheckScriptInfo)`

SetHealthCheckScriptInfos sets HealthCheckScriptInfos field to given value.

### HasHealthCheckScriptInfos

`func (o *HyperflexHealthCheckDefinition) HasHealthCheckScriptInfos() bool`

HasHealthCheckScriptInfos returns a boolean if a field has been set.

### SetHealthCheckScriptInfosNil

`func (o *HyperflexHealthCheckDefinition) SetHealthCheckScriptInfosNil(b bool)`

 SetHealthCheckScriptInfosNil sets the value for HealthCheckScriptInfos to be an explicit nil

### UnsetHealthCheckScriptInfos
`func (o *HyperflexHealthCheckDefinition) UnsetHealthCheckScriptInfos()`

UnsetHealthCheckScriptInfos ensures that no value is present for HealthCheckScriptInfos, not even an explicit nil
### GetHealthImpact

`func (o *HyperflexHealthCheckDefinition) GetHealthImpact() string`

GetHealthImpact returns the HealthImpact field if non-nil, zero value otherwise.

### GetHealthImpactOk

`func (o *HyperflexHealthCheckDefinition) GetHealthImpactOk() (*string, bool)`

GetHealthImpactOk returns a tuple with the HealthImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthImpact

`func (o *HyperflexHealthCheckDefinition) SetHealthImpact(v string)`

SetHealthImpact sets HealthImpact field to given value.

### HasHealthImpact

`func (o *HyperflexHealthCheckDefinition) HasHealthImpact() bool`

HasHealthImpact returns a boolean if a field has been set.

### GetInternalName

`func (o *HyperflexHealthCheckDefinition) GetInternalName() string`

GetInternalName returns the InternalName field if non-nil, zero value otherwise.

### GetInternalNameOk

`func (o *HyperflexHealthCheckDefinition) GetInternalNameOk() (*string, bool)`

GetInternalNameOk returns a tuple with the InternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalName

`func (o *HyperflexHealthCheckDefinition) SetInternalName(v string)`

SetInternalName sets InternalName field to given value.

### HasInternalName

`func (o *HyperflexHealthCheckDefinition) HasInternalName() bool`

HasInternalName returns a boolean if a field has been set.

### GetMinimumHyperFlexVersion

`func (o *HyperflexHealthCheckDefinition) GetMinimumHyperFlexVersion() string`

GetMinimumHyperFlexVersion returns the MinimumHyperFlexVersion field if non-nil, zero value otherwise.

### GetMinimumHyperFlexVersionOk

`func (o *HyperflexHealthCheckDefinition) GetMinimumHyperFlexVersionOk() (*string, bool)`

GetMinimumHyperFlexVersionOk returns a tuple with the MinimumHyperFlexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumHyperFlexVersion

`func (o *HyperflexHealthCheckDefinition) SetMinimumHyperFlexVersion(v string)`

SetMinimumHyperFlexVersion sets MinimumHyperFlexVersion field to given value.

### HasMinimumHyperFlexVersion

`func (o *HyperflexHealthCheckDefinition) HasMinimumHyperFlexVersion() bool`

HasMinimumHyperFlexVersion returns a boolean if a field has been set.

### GetName

`func (o *HyperflexHealthCheckDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHealthCheckDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHealthCheckDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHealthCheckDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *HyperflexHealthCheckDefinition) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *HyperflexHealthCheckDefinition) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *HyperflexHealthCheckDefinition) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *HyperflexHealthCheckDefinition) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetResolution

`func (o *HyperflexHealthCheckDefinition) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *HyperflexHealthCheckDefinition) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *HyperflexHealthCheckDefinition) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *HyperflexHealthCheckDefinition) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetScriptExecutionMode

`func (o *HyperflexHealthCheckDefinition) GetScriptExecutionMode() string`

GetScriptExecutionMode returns the ScriptExecutionMode field if non-nil, zero value otherwise.

### GetScriptExecutionModeOk

`func (o *HyperflexHealthCheckDefinition) GetScriptExecutionModeOk() (*string, bool)`

GetScriptExecutionModeOk returns a tuple with the ScriptExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptExecutionMode

`func (o *HyperflexHealthCheckDefinition) SetScriptExecutionMode(v string)`

SetScriptExecutionMode sets ScriptExecutionMode field to given value.

### HasScriptExecutionMode

`func (o *HyperflexHealthCheckDefinition) HasScriptExecutionMode() bool`

HasScriptExecutionMode returns a boolean if a field has been set.

### GetScriptExecutionOnComputeNodes

`func (o *HyperflexHealthCheckDefinition) GetScriptExecutionOnComputeNodes() bool`

GetScriptExecutionOnComputeNodes returns the ScriptExecutionOnComputeNodes field if non-nil, zero value otherwise.

### GetScriptExecutionOnComputeNodesOk

`func (o *HyperflexHealthCheckDefinition) GetScriptExecutionOnComputeNodesOk() (*bool, bool)`

GetScriptExecutionOnComputeNodesOk returns a tuple with the ScriptExecutionOnComputeNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptExecutionOnComputeNodes

`func (o *HyperflexHealthCheckDefinition) SetScriptExecutionOnComputeNodes(v bool)`

SetScriptExecutionOnComputeNodes sets ScriptExecutionOnComputeNodes field to given value.

### HasScriptExecutionOnComputeNodes

`func (o *HyperflexHealthCheckDefinition) HasScriptExecutionOnComputeNodes() bool`

HasScriptExecutionOnComputeNodes returns a boolean if a field has been set.

### GetSupportedHypervisorType

`func (o *HyperflexHealthCheckDefinition) GetSupportedHypervisorType() string`

GetSupportedHypervisorType returns the SupportedHypervisorType field if non-nil, zero value otherwise.

### GetSupportedHypervisorTypeOk

`func (o *HyperflexHealthCheckDefinition) GetSupportedHypervisorTypeOk() (*string, bool)`

GetSupportedHypervisorTypeOk returns a tuple with the SupportedHypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedHypervisorType

`func (o *HyperflexHealthCheckDefinition) SetSupportedHypervisorType(v string)`

SetSupportedHypervisorType sets SupportedHypervisorType field to given value.

### HasSupportedHypervisorType

`func (o *HyperflexHealthCheckDefinition) HasSupportedHypervisorType() bool`

HasSupportedHypervisorType returns a boolean if a field has been set.

### GetTargetExecutionType

`func (o *HyperflexHealthCheckDefinition) GetTargetExecutionType() string`

GetTargetExecutionType returns the TargetExecutionType field if non-nil, zero value otherwise.

### GetTargetExecutionTypeOk

`func (o *HyperflexHealthCheckDefinition) GetTargetExecutionTypeOk() (*string, bool)`

GetTargetExecutionTypeOk returns a tuple with the TargetExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetExecutionType

`func (o *HyperflexHealthCheckDefinition) SetTargetExecutionType(v string)`

SetTargetExecutionType sets TargetExecutionType field to given value.

### HasTargetExecutionType

`func (o *HyperflexHealthCheckDefinition) HasTargetExecutionType() bool`

HasTargetExecutionType returns a boolean if a field has been set.

### GetTimeout

`func (o *HyperflexHealthCheckDefinition) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HyperflexHealthCheckDefinition) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HyperflexHealthCheckDefinition) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HyperflexHealthCheckDefinition) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUnsupportedHyperFlexVersions

`func (o *HyperflexHealthCheckDefinition) GetUnsupportedHyperFlexVersions() []string`

GetUnsupportedHyperFlexVersions returns the UnsupportedHyperFlexVersions field if non-nil, zero value otherwise.

### GetUnsupportedHyperFlexVersionsOk

`func (o *HyperflexHealthCheckDefinition) GetUnsupportedHyperFlexVersionsOk() (*[]string, bool)`

GetUnsupportedHyperFlexVersionsOk returns a tuple with the UnsupportedHyperFlexVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedHyperFlexVersions

`func (o *HyperflexHealthCheckDefinition) SetUnsupportedHyperFlexVersions(v []string)`

SetUnsupportedHyperFlexVersions sets UnsupportedHyperFlexVersions field to given value.

### HasUnsupportedHyperFlexVersions

`func (o *HyperflexHealthCheckDefinition) HasUnsupportedHyperFlexVersions() bool`

HasUnsupportedHyperFlexVersions returns a boolean if a field has been set.

### SetUnsupportedHyperFlexVersionsNil

`func (o *HyperflexHealthCheckDefinition) SetUnsupportedHyperFlexVersionsNil(b bool)`

 SetUnsupportedHyperFlexVersionsNil sets the value for UnsupportedHyperFlexVersions to be an explicit nil

### UnsetUnsupportedHyperFlexVersions
`func (o *HyperflexHealthCheckDefinition) UnsetUnsupportedHyperFlexVersions()`

UnsetUnsupportedHyperFlexVersions ensures that no value is present for UnsupportedHyperFlexVersions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


