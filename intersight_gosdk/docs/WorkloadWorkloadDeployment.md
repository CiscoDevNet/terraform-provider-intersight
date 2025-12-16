# WorkloadWorkloadDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.WorkloadDeployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.WorkloadDeployment"]
**Action** | Pointer to **string** | The action is performed on the deployment. * &#x60;None&#x60; - No changes have been made. * &#x60;PrepareToDeploy&#x60; - Marks the deployment as ready once the user has completed all changes and the deployment is in a valid state. Once the deployment is marked as PrepareToDeploy, workload instances will be created, but the actual deployment will occur as part of the deploy action. * &#x60;Deploy&#x60; - Initiates the process of pushing workload configuration to the instances based on the configured schedule. Once deployed, the deployment cannot be reverted to draft status. * &#x60;Suspend&#x60; - Suspends the deployment, preventing any further actions until it is resumed. Making changes to deployment configuration will not be pushed out. The deployment will not take any changes from the attached Workload when it is suspended. * &#x60;Retry&#x60; - Retries the deployment for all instances that previously failed. * &#x60;Resume&#x60; - Resumes a suspended deployment. Any changes made to the deployment when it was suspended or any changes made to the attached Workload will now be pushed out based on configured schedules. * &#x60;Undeploy&#x60; - Undeploy cleans up the policies, templates, and leases generated from the deployment, and marks the deployment as being in the Draft state. The associated definition will be marked as Inactive if there are no other deployments linked to it. * &#x60;ReEvaluateQualifiers&#x60; - Triggers a re-evaluation of resource qualifiers for the deployment to ensure that all associated instances are aligned with the current resource constraints and qualifiers defined in the deployment. | [optional] [default to "None"]
**Blueprints** | Pointer to [**[]WorkloadBlueprintInputReference**](WorkloadBlueprintInputReference.md) |  | [optional] 
**ChangeDetails** | Pointer to [**[]WorkloadDeploymentChangeDetail**](WorkloadDeploymentChangeDetail.md) |  | [optional] 
**Conformance** | Pointer to **string** | The conformance status of the deployment. * &#x60;Ok&#x60; - The deployment conforms to the preferred version of the workload. * &#x60;NonConformant&#x60; - The deployment does not conform to the preferred version of the workload. | [optional] [readonly] [default to "Ok"]
**Description** | Pointer to **string** | A brief description of the deployment. | [optional] 
**DigitCount** | Pointer to **int64** | The minimum digit count to format the instance index with leading zeros,  for example if the digit count is 4 and the start index is 1, then the  first instance will have a suffix 0001. If the number of instances created for the deployment exceeds the 9999, then the suffix will become a 5 digit number. | [optional] [default to 1]
**InstanceConformanceSummary** | Pointer to [**[]WorkloadStateAggregation**](WorkloadStateAggregation.md) |  | [optional] 
**InstanceStatusSummary** | Pointer to [**[]WorkloadStateAggregation**](WorkloadStateAggregation.md) |  | [optional] 
**LastAction** | Pointer to **string** | The last action is performed on the deployment. * &#x60;None&#x60; - No changes have been made. * &#x60;PrepareToDeploy&#x60; - Marks the deployment as ready once the user has completed all changes and the deployment is in a valid state. Once the deployment is marked as PrepareToDeploy, workload instances will be created, but the actual deployment will occur as part of the deploy action. * &#x60;Deploy&#x60; - Initiates the process of pushing workload configuration to the instances based on the configured schedule. Once deployed, the deployment cannot be reverted to draft status. * &#x60;Suspend&#x60; - Suspends the deployment, preventing any further actions until it is resumed. Making changes to deployment configuration will not be pushed out. The deployment will not take any changes from the attached Workload when it is suspended. * &#x60;Retry&#x60; - Retries the deployment for all instances that previously failed. * &#x60;Resume&#x60; - Resumes a suspended deployment. Any changes made to the deployment when it was suspended or any changes made to the attached Workload will now be pushed out based on configured schedules. * &#x60;Undeploy&#x60; - Undeploy cleans up the policies, templates, and leases generated from the deployment, and marks the deployment as being in the Draft state. The associated definition will be marked as Inactive if there are no other deployments linked to it. * &#x60;ReEvaluateQualifiers&#x60; - Triggers a re-evaluation of resource qualifiers for the deployment to ensure that all associated instances are aligned with the current resource constraints and qualifiers defined in the deployment. | [optional] [readonly] [default to "None"]
**LastInstanceIndex** | Pointer to **int64** | Tracks the last numeric index used for workload instances. The next instance index is derived by incrementing this value. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for this Deployment. Name can only contain letters (a-z, A-Z), numbers (0-9), space, hyphen (-), or an underscore (_). The name must be unique within the organization. | [optional] 
**Qualifiers** | Pointer to [**[]ResourceResourceQualifier**](ResourceResourceQualifier.md) |  | [optional] 
**RefName** | Pointer to **string** | A reference name is generated by the system based on the given name by replacing spaces and hyphen in name with underscore. This reference name is used internally and cannot be edited by users. It may only contain letters (a–z, A–Z), numbers (0–9), and underscores (_). | [optional] [readonly] 
**RolloutStrategy** | Pointer to [**NullableWorkloadRolloutStrategy**](WorkloadRolloutStrategy.md) |  | [optional] 
**StartIndexForSuffix** | Pointer to **int64** | The starting index used to generate the suffix for the workload instance name. | [optional] [default to 1]
**Status** | Pointer to **string** | The status of the deployment. * &#x60;Draft&#x60; - The initial state of the deployment, indicating it is still being configured and not ready for deployment. * &#x60;ReadyToDeploy&#x60; - The deployment has passed validation checks and is ready to be deployed to instances. * &#x60;InProgress&#x60; - The initial deployment or changes in the deployment are in progress. * &#x60;Ok&#x60; - All associated instances have been successfully deployed. * &#x60;Failed&#x60; - One or more instances failed during initial deployment or changes in the deployment operations. * &#x60;ChangesScheduled&#x60; - The deployment has changes that need to be pushed to all associated instances. * &#x60;Suspended&#x60; - The deployment is suspended, preventing any further initial deployments or changes to deployment operations. | [optional] [readonly] [default to "Draft"]
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](WorkflowValidationInformation.md) |  | [optional] 
**WorkloadDefinitionReference** | Pointer to [**NullableWorkloadDefinitionMapper**](WorkloadDefinitionMapper.md) |  | [optional] 
**WorkloadInstancePrefix** | Pointer to **string** | The prefix to be used for naming workload instances created by this deployment. Prefix can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), or an underscore (_). This prefix must be unique within the organization. | [optional] 
**DeploymentInput** | Pointer to [**NullableWorkloadDeploymentInputRelationship**](WorkloadDeploymentInputRelationship.md) |  | [optional] 
**DeploymentInputHistory** | Pointer to [**[]WorkloadDeploymentInputRelationship**](WorkloadDeploymentInputRelationship.md) | An array of relationships to workloadDeploymentInput resources. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**QualificationPolicies** | Pointer to [**[]ResourceAbstractResourceQualificationPolicyRelationship**](ResourceAbstractResourceQualificationPolicyRelationship.md) | An array of relationships to resourceAbstractResourceQualificationPolicy resources. | [optional] [readonly] 
**ResourcePool** | Pointer to [**NullableResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**SchedulePolicy** | Pointer to [**NullableSchedulerSchedulePolicyRelationship**](SchedulerSchedulePolicyRelationship.md) |  | [optional] 
**TaskSchedule** | Pointer to [**NullableSchedulerTaskScheduleRelationship**](SchedulerTaskScheduleRelationship.md) |  | [optional] 
**WorkloadDefinition** | Pointer to [**NullableWorkloadWorkloadDefinitionRelationship**](WorkloadWorkloadDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkloadWorkloadDeployment

`func NewWorkloadWorkloadDeployment(classId string, objectType string, ) *WorkloadWorkloadDeployment`

NewWorkloadWorkloadDeployment instantiates a new WorkloadWorkloadDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadDeploymentWithDefaults

`func NewWorkloadWorkloadDeploymentWithDefaults() *WorkloadWorkloadDeployment`

NewWorkloadWorkloadDeploymentWithDefaults instantiates a new WorkloadWorkloadDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadWorkloadDeployment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadWorkloadDeployment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadWorkloadDeployment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadWorkloadDeployment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadWorkloadDeployment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadWorkloadDeployment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *WorkloadWorkloadDeployment) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkloadWorkloadDeployment) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkloadWorkloadDeployment) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkloadWorkloadDeployment) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlueprints

`func (o *WorkloadWorkloadDeployment) GetBlueprints() []WorkloadBlueprintInputReference`

GetBlueprints returns the Blueprints field if non-nil, zero value otherwise.

### GetBlueprintsOk

`func (o *WorkloadWorkloadDeployment) GetBlueprintsOk() (*[]WorkloadBlueprintInputReference, bool)`

GetBlueprintsOk returns a tuple with the Blueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprints

`func (o *WorkloadWorkloadDeployment) SetBlueprints(v []WorkloadBlueprintInputReference)`

SetBlueprints sets Blueprints field to given value.

### HasBlueprints

`func (o *WorkloadWorkloadDeployment) HasBlueprints() bool`

HasBlueprints returns a boolean if a field has been set.

### SetBlueprintsNil

`func (o *WorkloadWorkloadDeployment) SetBlueprintsNil(b bool)`

 SetBlueprintsNil sets the value for Blueprints to be an explicit nil

### UnsetBlueprints
`func (o *WorkloadWorkloadDeployment) UnsetBlueprints()`

UnsetBlueprints ensures that no value is present for Blueprints, not even an explicit nil
### GetChangeDetails

`func (o *WorkloadWorkloadDeployment) GetChangeDetails() []WorkloadDeploymentChangeDetail`

GetChangeDetails returns the ChangeDetails field if non-nil, zero value otherwise.

### GetChangeDetailsOk

`func (o *WorkloadWorkloadDeployment) GetChangeDetailsOk() (*[]WorkloadDeploymentChangeDetail, bool)`

GetChangeDetailsOk returns a tuple with the ChangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDetails

`func (o *WorkloadWorkloadDeployment) SetChangeDetails(v []WorkloadDeploymentChangeDetail)`

SetChangeDetails sets ChangeDetails field to given value.

### HasChangeDetails

`func (o *WorkloadWorkloadDeployment) HasChangeDetails() bool`

HasChangeDetails returns a boolean if a field has been set.

### SetChangeDetailsNil

`func (o *WorkloadWorkloadDeployment) SetChangeDetailsNil(b bool)`

 SetChangeDetailsNil sets the value for ChangeDetails to be an explicit nil

### UnsetChangeDetails
`func (o *WorkloadWorkloadDeployment) UnsetChangeDetails()`

UnsetChangeDetails ensures that no value is present for ChangeDetails, not even an explicit nil
### GetConformance

`func (o *WorkloadWorkloadDeployment) GetConformance() string`

GetConformance returns the Conformance field if non-nil, zero value otherwise.

### GetConformanceOk

`func (o *WorkloadWorkloadDeployment) GetConformanceOk() (*string, bool)`

GetConformanceOk returns a tuple with the Conformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConformance

`func (o *WorkloadWorkloadDeployment) SetConformance(v string)`

SetConformance sets Conformance field to given value.

### HasConformance

`func (o *WorkloadWorkloadDeployment) HasConformance() bool`

HasConformance returns a boolean if a field has been set.

### GetDescription

`func (o *WorkloadWorkloadDeployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkloadWorkloadDeployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkloadWorkloadDeployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkloadWorkloadDeployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigitCount

`func (o *WorkloadWorkloadDeployment) GetDigitCount() int64`

GetDigitCount returns the DigitCount field if non-nil, zero value otherwise.

### GetDigitCountOk

`func (o *WorkloadWorkloadDeployment) GetDigitCountOk() (*int64, bool)`

GetDigitCountOk returns a tuple with the DigitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitCount

`func (o *WorkloadWorkloadDeployment) SetDigitCount(v int64)`

SetDigitCount sets DigitCount field to given value.

### HasDigitCount

`func (o *WorkloadWorkloadDeployment) HasDigitCount() bool`

HasDigitCount returns a boolean if a field has been set.

### GetInstanceConformanceSummary

`func (o *WorkloadWorkloadDeployment) GetInstanceConformanceSummary() []WorkloadStateAggregation`

GetInstanceConformanceSummary returns the InstanceConformanceSummary field if non-nil, zero value otherwise.

### GetInstanceConformanceSummaryOk

`func (o *WorkloadWorkloadDeployment) GetInstanceConformanceSummaryOk() (*[]WorkloadStateAggregation, bool)`

GetInstanceConformanceSummaryOk returns a tuple with the InstanceConformanceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceConformanceSummary

`func (o *WorkloadWorkloadDeployment) SetInstanceConformanceSummary(v []WorkloadStateAggregation)`

SetInstanceConformanceSummary sets InstanceConformanceSummary field to given value.

### HasInstanceConformanceSummary

`func (o *WorkloadWorkloadDeployment) HasInstanceConformanceSummary() bool`

HasInstanceConformanceSummary returns a boolean if a field has been set.

### SetInstanceConformanceSummaryNil

`func (o *WorkloadWorkloadDeployment) SetInstanceConformanceSummaryNil(b bool)`

 SetInstanceConformanceSummaryNil sets the value for InstanceConformanceSummary to be an explicit nil

### UnsetInstanceConformanceSummary
`func (o *WorkloadWorkloadDeployment) UnsetInstanceConformanceSummary()`

UnsetInstanceConformanceSummary ensures that no value is present for InstanceConformanceSummary, not even an explicit nil
### GetInstanceStatusSummary

`func (o *WorkloadWorkloadDeployment) GetInstanceStatusSummary() []WorkloadStateAggregation`

GetInstanceStatusSummary returns the InstanceStatusSummary field if non-nil, zero value otherwise.

### GetInstanceStatusSummaryOk

`func (o *WorkloadWorkloadDeployment) GetInstanceStatusSummaryOk() (*[]WorkloadStateAggregation, bool)`

GetInstanceStatusSummaryOk returns a tuple with the InstanceStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStatusSummary

`func (o *WorkloadWorkloadDeployment) SetInstanceStatusSummary(v []WorkloadStateAggregation)`

SetInstanceStatusSummary sets InstanceStatusSummary field to given value.

### HasInstanceStatusSummary

`func (o *WorkloadWorkloadDeployment) HasInstanceStatusSummary() bool`

HasInstanceStatusSummary returns a boolean if a field has been set.

### SetInstanceStatusSummaryNil

`func (o *WorkloadWorkloadDeployment) SetInstanceStatusSummaryNil(b bool)`

 SetInstanceStatusSummaryNil sets the value for InstanceStatusSummary to be an explicit nil

### UnsetInstanceStatusSummary
`func (o *WorkloadWorkloadDeployment) UnsetInstanceStatusSummary()`

UnsetInstanceStatusSummary ensures that no value is present for InstanceStatusSummary, not even an explicit nil
### GetLastAction

`func (o *WorkloadWorkloadDeployment) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkloadWorkloadDeployment) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkloadWorkloadDeployment) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkloadWorkloadDeployment) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetLastInstanceIndex

`func (o *WorkloadWorkloadDeployment) GetLastInstanceIndex() int64`

GetLastInstanceIndex returns the LastInstanceIndex field if non-nil, zero value otherwise.

### GetLastInstanceIndexOk

`func (o *WorkloadWorkloadDeployment) GetLastInstanceIndexOk() (*int64, bool)`

GetLastInstanceIndexOk returns a tuple with the LastInstanceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInstanceIndex

`func (o *WorkloadWorkloadDeployment) SetLastInstanceIndex(v int64)`

SetLastInstanceIndex sets LastInstanceIndex field to given value.

### HasLastInstanceIndex

`func (o *WorkloadWorkloadDeployment) HasLastInstanceIndex() bool`

HasLastInstanceIndex returns a boolean if a field has been set.

### GetName

`func (o *WorkloadWorkloadDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadWorkloadDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadWorkloadDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadWorkloadDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQualifiers

`func (o *WorkloadWorkloadDeployment) GetQualifiers() []ResourceResourceQualifier`

GetQualifiers returns the Qualifiers field if non-nil, zero value otherwise.

### GetQualifiersOk

`func (o *WorkloadWorkloadDeployment) GetQualifiersOk() (*[]ResourceResourceQualifier, bool)`

GetQualifiersOk returns a tuple with the Qualifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiers

`func (o *WorkloadWorkloadDeployment) SetQualifiers(v []ResourceResourceQualifier)`

SetQualifiers sets Qualifiers field to given value.

### HasQualifiers

`func (o *WorkloadWorkloadDeployment) HasQualifiers() bool`

HasQualifiers returns a boolean if a field has been set.

### SetQualifiersNil

`func (o *WorkloadWorkloadDeployment) SetQualifiersNil(b bool)`

 SetQualifiersNil sets the value for Qualifiers to be an explicit nil

### UnsetQualifiers
`func (o *WorkloadWorkloadDeployment) UnsetQualifiers()`

UnsetQualifiers ensures that no value is present for Qualifiers, not even an explicit nil
### GetRefName

`func (o *WorkloadWorkloadDeployment) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *WorkloadWorkloadDeployment) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *WorkloadWorkloadDeployment) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *WorkloadWorkloadDeployment) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetRolloutStrategy

`func (o *WorkloadWorkloadDeployment) GetRolloutStrategy() WorkloadRolloutStrategy`

GetRolloutStrategy returns the RolloutStrategy field if non-nil, zero value otherwise.

### GetRolloutStrategyOk

`func (o *WorkloadWorkloadDeployment) GetRolloutStrategyOk() (*WorkloadRolloutStrategy, bool)`

GetRolloutStrategyOk returns a tuple with the RolloutStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutStrategy

`func (o *WorkloadWorkloadDeployment) SetRolloutStrategy(v WorkloadRolloutStrategy)`

SetRolloutStrategy sets RolloutStrategy field to given value.

### HasRolloutStrategy

`func (o *WorkloadWorkloadDeployment) HasRolloutStrategy() bool`

HasRolloutStrategy returns a boolean if a field has been set.

### SetRolloutStrategyNil

`func (o *WorkloadWorkloadDeployment) SetRolloutStrategyNil(b bool)`

 SetRolloutStrategyNil sets the value for RolloutStrategy to be an explicit nil

### UnsetRolloutStrategy
`func (o *WorkloadWorkloadDeployment) UnsetRolloutStrategy()`

UnsetRolloutStrategy ensures that no value is present for RolloutStrategy, not even an explicit nil
### GetStartIndexForSuffix

`func (o *WorkloadWorkloadDeployment) GetStartIndexForSuffix() int64`

GetStartIndexForSuffix returns the StartIndexForSuffix field if non-nil, zero value otherwise.

### GetStartIndexForSuffixOk

`func (o *WorkloadWorkloadDeployment) GetStartIndexForSuffixOk() (*int64, bool)`

GetStartIndexForSuffixOk returns a tuple with the StartIndexForSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndexForSuffix

`func (o *WorkloadWorkloadDeployment) SetStartIndexForSuffix(v int64)`

SetStartIndexForSuffix sets StartIndexForSuffix field to given value.

### HasStartIndexForSuffix

`func (o *WorkloadWorkloadDeployment) HasStartIndexForSuffix() bool`

HasStartIndexForSuffix returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadWorkloadDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadWorkloadDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadWorkloadDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadWorkloadDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidationInformation

`func (o *WorkloadWorkloadDeployment) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkloadWorkloadDeployment) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkloadWorkloadDeployment) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkloadWorkloadDeployment) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkloadWorkloadDeployment) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkloadWorkloadDeployment) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetWorkloadDefinitionReference

`func (o *WorkloadWorkloadDeployment) GetWorkloadDefinitionReference() WorkloadDefinitionMapper`

GetWorkloadDefinitionReference returns the WorkloadDefinitionReference field if non-nil, zero value otherwise.

### GetWorkloadDefinitionReferenceOk

`func (o *WorkloadWorkloadDeployment) GetWorkloadDefinitionReferenceOk() (*WorkloadDefinitionMapper, bool)`

GetWorkloadDefinitionReferenceOk returns a tuple with the WorkloadDefinitionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDefinitionReference

`func (o *WorkloadWorkloadDeployment) SetWorkloadDefinitionReference(v WorkloadDefinitionMapper)`

SetWorkloadDefinitionReference sets WorkloadDefinitionReference field to given value.

### HasWorkloadDefinitionReference

`func (o *WorkloadWorkloadDeployment) HasWorkloadDefinitionReference() bool`

HasWorkloadDefinitionReference returns a boolean if a field has been set.

### SetWorkloadDefinitionReferenceNil

`func (o *WorkloadWorkloadDeployment) SetWorkloadDefinitionReferenceNil(b bool)`

 SetWorkloadDefinitionReferenceNil sets the value for WorkloadDefinitionReference to be an explicit nil

### UnsetWorkloadDefinitionReference
`func (o *WorkloadWorkloadDeployment) UnsetWorkloadDefinitionReference()`

UnsetWorkloadDefinitionReference ensures that no value is present for WorkloadDefinitionReference, not even an explicit nil
### GetWorkloadInstancePrefix

`func (o *WorkloadWorkloadDeployment) GetWorkloadInstancePrefix() string`

GetWorkloadInstancePrefix returns the WorkloadInstancePrefix field if non-nil, zero value otherwise.

### GetWorkloadInstancePrefixOk

`func (o *WorkloadWorkloadDeployment) GetWorkloadInstancePrefixOk() (*string, bool)`

GetWorkloadInstancePrefixOk returns a tuple with the WorkloadInstancePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadInstancePrefix

`func (o *WorkloadWorkloadDeployment) SetWorkloadInstancePrefix(v string)`

SetWorkloadInstancePrefix sets WorkloadInstancePrefix field to given value.

### HasWorkloadInstancePrefix

`func (o *WorkloadWorkloadDeployment) HasWorkloadInstancePrefix() bool`

HasWorkloadInstancePrefix returns a boolean if a field has been set.

### GetDeploymentInput

`func (o *WorkloadWorkloadDeployment) GetDeploymentInput() WorkloadDeploymentInputRelationship`

GetDeploymentInput returns the DeploymentInput field if non-nil, zero value otherwise.

### GetDeploymentInputOk

`func (o *WorkloadWorkloadDeployment) GetDeploymentInputOk() (*WorkloadDeploymentInputRelationship, bool)`

GetDeploymentInputOk returns a tuple with the DeploymentInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentInput

`func (o *WorkloadWorkloadDeployment) SetDeploymentInput(v WorkloadDeploymentInputRelationship)`

SetDeploymentInput sets DeploymentInput field to given value.

### HasDeploymentInput

`func (o *WorkloadWorkloadDeployment) HasDeploymentInput() bool`

HasDeploymentInput returns a boolean if a field has been set.

### SetDeploymentInputNil

`func (o *WorkloadWorkloadDeployment) SetDeploymentInputNil(b bool)`

 SetDeploymentInputNil sets the value for DeploymentInput to be an explicit nil

### UnsetDeploymentInput
`func (o *WorkloadWorkloadDeployment) UnsetDeploymentInput()`

UnsetDeploymentInput ensures that no value is present for DeploymentInput, not even an explicit nil
### GetDeploymentInputHistory

`func (o *WorkloadWorkloadDeployment) GetDeploymentInputHistory() []WorkloadDeploymentInputRelationship`

GetDeploymentInputHistory returns the DeploymentInputHistory field if non-nil, zero value otherwise.

### GetDeploymentInputHistoryOk

`func (o *WorkloadWorkloadDeployment) GetDeploymentInputHistoryOk() (*[]WorkloadDeploymentInputRelationship, bool)`

GetDeploymentInputHistoryOk returns a tuple with the DeploymentInputHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentInputHistory

`func (o *WorkloadWorkloadDeployment) SetDeploymentInputHistory(v []WorkloadDeploymentInputRelationship)`

SetDeploymentInputHistory sets DeploymentInputHistory field to given value.

### HasDeploymentInputHistory

`func (o *WorkloadWorkloadDeployment) HasDeploymentInputHistory() bool`

HasDeploymentInputHistory returns a boolean if a field has been set.

### SetDeploymentInputHistoryNil

`func (o *WorkloadWorkloadDeployment) SetDeploymentInputHistoryNil(b bool)`

 SetDeploymentInputHistoryNil sets the value for DeploymentInputHistory to be an explicit nil

### UnsetDeploymentInputHistory
`func (o *WorkloadWorkloadDeployment) UnsetDeploymentInputHistory()`

UnsetDeploymentInputHistory ensures that no value is present for DeploymentInputHistory, not even an explicit nil
### GetOrganization

`func (o *WorkloadWorkloadDeployment) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkloadWorkloadDeployment) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkloadWorkloadDeployment) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkloadWorkloadDeployment) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *WorkloadWorkloadDeployment) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *WorkloadWorkloadDeployment) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetQualificationPolicies

`func (o *WorkloadWorkloadDeployment) GetQualificationPolicies() []ResourceAbstractResourceQualificationPolicyRelationship`

GetQualificationPolicies returns the QualificationPolicies field if non-nil, zero value otherwise.

### GetQualificationPoliciesOk

`func (o *WorkloadWorkloadDeployment) GetQualificationPoliciesOk() (*[]ResourceAbstractResourceQualificationPolicyRelationship, bool)`

GetQualificationPoliciesOk returns a tuple with the QualificationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualificationPolicies

`func (o *WorkloadWorkloadDeployment) SetQualificationPolicies(v []ResourceAbstractResourceQualificationPolicyRelationship)`

SetQualificationPolicies sets QualificationPolicies field to given value.

### HasQualificationPolicies

`func (o *WorkloadWorkloadDeployment) HasQualificationPolicies() bool`

HasQualificationPolicies returns a boolean if a field has been set.

### SetQualificationPoliciesNil

`func (o *WorkloadWorkloadDeployment) SetQualificationPoliciesNil(b bool)`

 SetQualificationPoliciesNil sets the value for QualificationPolicies to be an explicit nil

### UnsetQualificationPolicies
`func (o *WorkloadWorkloadDeployment) UnsetQualificationPolicies()`

UnsetQualificationPolicies ensures that no value is present for QualificationPolicies, not even an explicit nil
### GetResourcePool

`func (o *WorkloadWorkloadDeployment) GetResourcePool() ResourcepoolPoolRelationship`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *WorkloadWorkloadDeployment) GetResourcePoolOk() (*ResourcepoolPoolRelationship, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *WorkloadWorkloadDeployment) SetResourcePool(v ResourcepoolPoolRelationship)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *WorkloadWorkloadDeployment) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### SetResourcePoolNil

`func (o *WorkloadWorkloadDeployment) SetResourcePoolNil(b bool)`

 SetResourcePoolNil sets the value for ResourcePool to be an explicit nil

### UnsetResourcePool
`func (o *WorkloadWorkloadDeployment) UnsetResourcePool()`

UnsetResourcePool ensures that no value is present for ResourcePool, not even an explicit nil
### GetSchedulePolicy

`func (o *WorkloadWorkloadDeployment) GetSchedulePolicy() SchedulerSchedulePolicyRelationship`

GetSchedulePolicy returns the SchedulePolicy field if non-nil, zero value otherwise.

### GetSchedulePolicyOk

`func (o *WorkloadWorkloadDeployment) GetSchedulePolicyOk() (*SchedulerSchedulePolicyRelationship, bool)`

GetSchedulePolicyOk returns a tuple with the SchedulePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulePolicy

`func (o *WorkloadWorkloadDeployment) SetSchedulePolicy(v SchedulerSchedulePolicyRelationship)`

SetSchedulePolicy sets SchedulePolicy field to given value.

### HasSchedulePolicy

`func (o *WorkloadWorkloadDeployment) HasSchedulePolicy() bool`

HasSchedulePolicy returns a boolean if a field has been set.

### SetSchedulePolicyNil

`func (o *WorkloadWorkloadDeployment) SetSchedulePolicyNil(b bool)`

 SetSchedulePolicyNil sets the value for SchedulePolicy to be an explicit nil

### UnsetSchedulePolicy
`func (o *WorkloadWorkloadDeployment) UnsetSchedulePolicy()`

UnsetSchedulePolicy ensures that no value is present for SchedulePolicy, not even an explicit nil
### GetTaskSchedule

`func (o *WorkloadWorkloadDeployment) GetTaskSchedule() SchedulerTaskScheduleRelationship`

GetTaskSchedule returns the TaskSchedule field if non-nil, zero value otherwise.

### GetTaskScheduleOk

`func (o *WorkloadWorkloadDeployment) GetTaskScheduleOk() (*SchedulerTaskScheduleRelationship, bool)`

GetTaskScheduleOk returns a tuple with the TaskSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSchedule

`func (o *WorkloadWorkloadDeployment) SetTaskSchedule(v SchedulerTaskScheduleRelationship)`

SetTaskSchedule sets TaskSchedule field to given value.

### HasTaskSchedule

`func (o *WorkloadWorkloadDeployment) HasTaskSchedule() bool`

HasTaskSchedule returns a boolean if a field has been set.

### SetTaskScheduleNil

`func (o *WorkloadWorkloadDeployment) SetTaskScheduleNil(b bool)`

 SetTaskScheduleNil sets the value for TaskSchedule to be an explicit nil

### UnsetTaskSchedule
`func (o *WorkloadWorkloadDeployment) UnsetTaskSchedule()`

UnsetTaskSchedule ensures that no value is present for TaskSchedule, not even an explicit nil
### GetWorkloadDefinition

`func (o *WorkloadWorkloadDeployment) GetWorkloadDefinition() WorkloadWorkloadDefinitionRelationship`

GetWorkloadDefinition returns the WorkloadDefinition field if non-nil, zero value otherwise.

### GetWorkloadDefinitionOk

`func (o *WorkloadWorkloadDeployment) GetWorkloadDefinitionOk() (*WorkloadWorkloadDefinitionRelationship, bool)`

GetWorkloadDefinitionOk returns a tuple with the WorkloadDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDefinition

`func (o *WorkloadWorkloadDeployment) SetWorkloadDefinition(v WorkloadWorkloadDefinitionRelationship)`

SetWorkloadDefinition sets WorkloadDefinition field to given value.

### HasWorkloadDefinition

`func (o *WorkloadWorkloadDeployment) HasWorkloadDefinition() bool`

HasWorkloadDefinition returns a boolean if a field has been set.

### SetWorkloadDefinitionNil

`func (o *WorkloadWorkloadDeployment) SetWorkloadDefinitionNil(b bool)`

 SetWorkloadDefinitionNil sets the value for WorkloadDefinition to be an explicit nil

### UnsetWorkloadDefinition
`func (o *WorkloadWorkloadDeployment) UnsetWorkloadDefinition()`

UnsetWorkloadDefinition ensures that no value is present for WorkloadDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


