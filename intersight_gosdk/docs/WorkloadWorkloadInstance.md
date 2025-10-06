# WorkloadWorkloadInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.WorkloadInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.WorkloadInstance"]
**Action** | Pointer to **string** | The action to be performed on the workload instance. * &#x60;None&#x60; - Absence of any action on the workload instance. * &#x60;Suspend&#x60; - Pauses the execution of the workload instance, temporarily stopping its operations without permanently removing it. * &#x60;Resume&#x60; - Restarts a suspended workload instance, allowing it to continue operations from where it left off. * &#x60;Deploy&#x60; - Initiates the deployment of the workload instance, provisioning the necessary resources and starting its execution. * &#x60;Retry&#x60; - Attempts to re-deploy the workload instance, either due to a previous failure or to apply changes made to the instance. * &#x60;RetryAll&#x60; - Attempts to re-deploy all workload instances associated with the same deployment, either due to a previous failure or to apply changes made to the instances. * &#x60;Attach&#x60; - Associates the workload instance with its assigned resources, allowing it to utilize the resources for its operations. * &#x60;Detach&#x60; - Disassociates the workload instance from its assigned resources, preventing it from using the resources for its operations. * &#x60;UnAssign&#x60; - Detaches assigned resources from the workload instance while keeping the instance active. | [optional] [default to "None"]
**AssignedResources** | Pointer to [**[]WorkloadAssignedResource**](WorkloadAssignedResource.md) |  | [optional] 
**ChangeTypes** | Pointer to **[]string** |  | [optional] 
**Conformance** | Pointer to **string** | The conformance status of the deployment. * &#x60;Ok&#x60; - The deployment conforms to the preferred version of the workload. * &#x60;NonConformant&#x60; - The deployment does not conform to the preferred version of the workload. | [optional] [readonly] [default to "Ok"]
**LastAction** | Pointer to **string** | The last action performed on the workload instance. * &#x60;None&#x60; - Absence of any action on the workload instance. * &#x60;Suspend&#x60; - Pauses the execution of the workload instance, temporarily stopping its operations without permanently removing it. * &#x60;Resume&#x60; - Restarts a suspended workload instance, allowing it to continue operations from where it left off. * &#x60;Deploy&#x60; - Initiates the deployment of the workload instance, provisioning the necessary resources and starting its execution. * &#x60;Retry&#x60; - Attempts to re-deploy the workload instance, either due to a previous failure or to apply changes made to the instance. * &#x60;RetryAll&#x60; - Attempts to re-deploy all workload instances associated with the same deployment, either due to a previous failure or to apply changes made to the instances. * &#x60;Attach&#x60; - Associates the workload instance with its assigned resources, allowing it to utilize the resources for its operations. * &#x60;Detach&#x60; - Disassociates the workload instance from its assigned resources, preventing it from using the resources for its operations. * &#x60;UnAssign&#x60; - Detaches assigned resources from the workload instance while keeping the instance active. | [optional] [readonly] [default to "None"]
**LocationDetails** | Pointer to [**NullableCommGeoLocationDetails**](CommGeoLocationDetails.md) |  | [optional] 
**Name** | Pointer to **string** | The name for this Workload instance. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of the workload instance. * &#x60;Staging&#x60; - The instance is in the staging phase, awaiting further actions. * &#x60;ReadyToDeploy&#x60; - The instance is fully configured and ready for deployment. * &#x60;InProgress&#x60; - Deployment or modification of the instance is currently in progress. * &#x60;Ok&#x60; - The instance is running successfully without issues. * &#x60;Failed&#x60; - The instance has encountered an error or failure preventing normal operation. * &#x60;Suspended&#x60; - The instance has been temporarily paused and is inactive. * &#x60;ChangesScheduled&#x60; - There is a change in the configuration that needs to be pushed to the instance. * &#x60;InSufficientResource&#x60; - The instance lacks the necessary resources to operate. * &#x60;OutOfService&#x60; - The instance is no longer available or operational. * &#x60;UnAssigning&#x60; - The instance is being unassigned or removed from service. | [optional] [readonly] [default to "Staging"]
**StatusChangeReason** | Pointer to **string** | The context or justification for the status transition. * &#x60;None&#x60; - No changes have been made. * &#x60;ResourceDisqualified&#x60; - The change in resource status triggered due to the resource being disqualified. | [optional] [readonly] [default to "None"]
**ChassisPoolMember** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**DeploymentInput** | Pointer to [**NullableWorkloadDeploymentInputRelationship**](WorkloadDeploymentInputRelationship.md) |  | [optional] 
**LastDeploymentInput** | Pointer to [**NullableWorkloadDeploymentInputRelationship**](WorkloadDeploymentInputRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ServiceItemInstances** | Pointer to [**[]WorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) | An array of relationships to workflowServiceItemInstance resources. | [optional] [readonly] 
**WorkloadDefinition** | Pointer to [**NullableWorkloadWorkloadDefinitionRelationship**](WorkloadWorkloadDefinitionRelationship.md) |  | [optional] 
**WorkloadDeployment** | Pointer to [**NullableWorkloadWorkloadDeploymentRelationship**](WorkloadWorkloadDeploymentRelationship.md) |  | [optional] 

## Methods

### NewWorkloadWorkloadInstance

`func NewWorkloadWorkloadInstance(classId string, objectType string, ) *WorkloadWorkloadInstance`

NewWorkloadWorkloadInstance instantiates a new WorkloadWorkloadInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadInstanceWithDefaults

`func NewWorkloadWorkloadInstanceWithDefaults() *WorkloadWorkloadInstance`

NewWorkloadWorkloadInstanceWithDefaults instantiates a new WorkloadWorkloadInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadWorkloadInstance) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadWorkloadInstance) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadWorkloadInstance) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadWorkloadInstance) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadWorkloadInstance) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadWorkloadInstance) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *WorkloadWorkloadInstance) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkloadWorkloadInstance) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkloadWorkloadInstance) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkloadWorkloadInstance) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAssignedResources

`func (o *WorkloadWorkloadInstance) GetAssignedResources() []WorkloadAssignedResource`

GetAssignedResources returns the AssignedResources field if non-nil, zero value otherwise.

### GetAssignedResourcesOk

`func (o *WorkloadWorkloadInstance) GetAssignedResourcesOk() (*[]WorkloadAssignedResource, bool)`

GetAssignedResourcesOk returns a tuple with the AssignedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedResources

`func (o *WorkloadWorkloadInstance) SetAssignedResources(v []WorkloadAssignedResource)`

SetAssignedResources sets AssignedResources field to given value.

### HasAssignedResources

`func (o *WorkloadWorkloadInstance) HasAssignedResources() bool`

HasAssignedResources returns a boolean if a field has been set.

### SetAssignedResourcesNil

`func (o *WorkloadWorkloadInstance) SetAssignedResourcesNil(b bool)`

 SetAssignedResourcesNil sets the value for AssignedResources to be an explicit nil

### UnsetAssignedResources
`func (o *WorkloadWorkloadInstance) UnsetAssignedResources()`

UnsetAssignedResources ensures that no value is present for AssignedResources, not even an explicit nil
### GetChangeTypes

`func (o *WorkloadWorkloadInstance) GetChangeTypes() []string`

GetChangeTypes returns the ChangeTypes field if non-nil, zero value otherwise.

### GetChangeTypesOk

`func (o *WorkloadWorkloadInstance) GetChangeTypesOk() (*[]string, bool)`

GetChangeTypesOk returns a tuple with the ChangeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTypes

`func (o *WorkloadWorkloadInstance) SetChangeTypes(v []string)`

SetChangeTypes sets ChangeTypes field to given value.

### HasChangeTypes

`func (o *WorkloadWorkloadInstance) HasChangeTypes() bool`

HasChangeTypes returns a boolean if a field has been set.

### SetChangeTypesNil

`func (o *WorkloadWorkloadInstance) SetChangeTypesNil(b bool)`

 SetChangeTypesNil sets the value for ChangeTypes to be an explicit nil

### UnsetChangeTypes
`func (o *WorkloadWorkloadInstance) UnsetChangeTypes()`

UnsetChangeTypes ensures that no value is present for ChangeTypes, not even an explicit nil
### GetConformance

`func (o *WorkloadWorkloadInstance) GetConformance() string`

GetConformance returns the Conformance field if non-nil, zero value otherwise.

### GetConformanceOk

`func (o *WorkloadWorkloadInstance) GetConformanceOk() (*string, bool)`

GetConformanceOk returns a tuple with the Conformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConformance

`func (o *WorkloadWorkloadInstance) SetConformance(v string)`

SetConformance sets Conformance field to given value.

### HasConformance

`func (o *WorkloadWorkloadInstance) HasConformance() bool`

HasConformance returns a boolean if a field has been set.

### GetLastAction

`func (o *WorkloadWorkloadInstance) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkloadWorkloadInstance) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkloadWorkloadInstance) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkloadWorkloadInstance) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetLocationDetails

`func (o *WorkloadWorkloadInstance) GetLocationDetails() CommGeoLocationDetails`

GetLocationDetails returns the LocationDetails field if non-nil, zero value otherwise.

### GetLocationDetailsOk

`func (o *WorkloadWorkloadInstance) GetLocationDetailsOk() (*CommGeoLocationDetails, bool)`

GetLocationDetailsOk returns a tuple with the LocationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDetails

`func (o *WorkloadWorkloadInstance) SetLocationDetails(v CommGeoLocationDetails)`

SetLocationDetails sets LocationDetails field to given value.

### HasLocationDetails

`func (o *WorkloadWorkloadInstance) HasLocationDetails() bool`

HasLocationDetails returns a boolean if a field has been set.

### SetLocationDetailsNil

`func (o *WorkloadWorkloadInstance) SetLocationDetailsNil(b bool)`

 SetLocationDetailsNil sets the value for LocationDetails to be an explicit nil

### UnsetLocationDetails
`func (o *WorkloadWorkloadInstance) UnsetLocationDetails()`

UnsetLocationDetails ensures that no value is present for LocationDetails, not even an explicit nil
### GetName

`func (o *WorkloadWorkloadInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadWorkloadInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadWorkloadInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadWorkloadInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadWorkloadInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadWorkloadInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadWorkloadInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadWorkloadInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusChangeReason

`func (o *WorkloadWorkloadInstance) GetStatusChangeReason() string`

GetStatusChangeReason returns the StatusChangeReason field if non-nil, zero value otherwise.

### GetStatusChangeReasonOk

`func (o *WorkloadWorkloadInstance) GetStatusChangeReasonOk() (*string, bool)`

GetStatusChangeReasonOk returns a tuple with the StatusChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangeReason

`func (o *WorkloadWorkloadInstance) SetStatusChangeReason(v string)`

SetStatusChangeReason sets StatusChangeReason field to given value.

### HasStatusChangeReason

`func (o *WorkloadWorkloadInstance) HasStatusChangeReason() bool`

HasStatusChangeReason returns a boolean if a field has been set.

### GetChassisPoolMember

`func (o *WorkloadWorkloadInstance) GetChassisPoolMember() MoBaseMoRelationship`

GetChassisPoolMember returns the ChassisPoolMember field if non-nil, zero value otherwise.

### GetChassisPoolMemberOk

`func (o *WorkloadWorkloadInstance) GetChassisPoolMemberOk() (*MoBaseMoRelationship, bool)`

GetChassisPoolMemberOk returns a tuple with the ChassisPoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisPoolMember

`func (o *WorkloadWorkloadInstance) SetChassisPoolMember(v MoBaseMoRelationship)`

SetChassisPoolMember sets ChassisPoolMember field to given value.

### HasChassisPoolMember

`func (o *WorkloadWorkloadInstance) HasChassisPoolMember() bool`

HasChassisPoolMember returns a boolean if a field has been set.

### SetChassisPoolMemberNil

`func (o *WorkloadWorkloadInstance) SetChassisPoolMemberNil(b bool)`

 SetChassisPoolMemberNil sets the value for ChassisPoolMember to be an explicit nil

### UnsetChassisPoolMember
`func (o *WorkloadWorkloadInstance) UnsetChassisPoolMember()`

UnsetChassisPoolMember ensures that no value is present for ChassisPoolMember, not even an explicit nil
### GetDeploymentInput

`func (o *WorkloadWorkloadInstance) GetDeploymentInput() WorkloadDeploymentInputRelationship`

GetDeploymentInput returns the DeploymentInput field if non-nil, zero value otherwise.

### GetDeploymentInputOk

`func (o *WorkloadWorkloadInstance) GetDeploymentInputOk() (*WorkloadDeploymentInputRelationship, bool)`

GetDeploymentInputOk returns a tuple with the DeploymentInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentInput

`func (o *WorkloadWorkloadInstance) SetDeploymentInput(v WorkloadDeploymentInputRelationship)`

SetDeploymentInput sets DeploymentInput field to given value.

### HasDeploymentInput

`func (o *WorkloadWorkloadInstance) HasDeploymentInput() bool`

HasDeploymentInput returns a boolean if a field has been set.

### SetDeploymentInputNil

`func (o *WorkloadWorkloadInstance) SetDeploymentInputNil(b bool)`

 SetDeploymentInputNil sets the value for DeploymentInput to be an explicit nil

### UnsetDeploymentInput
`func (o *WorkloadWorkloadInstance) UnsetDeploymentInput()`

UnsetDeploymentInput ensures that no value is present for DeploymentInput, not even an explicit nil
### GetLastDeploymentInput

`func (o *WorkloadWorkloadInstance) GetLastDeploymentInput() WorkloadDeploymentInputRelationship`

GetLastDeploymentInput returns the LastDeploymentInput field if non-nil, zero value otherwise.

### GetLastDeploymentInputOk

`func (o *WorkloadWorkloadInstance) GetLastDeploymentInputOk() (*WorkloadDeploymentInputRelationship, bool)`

GetLastDeploymentInputOk returns a tuple with the LastDeploymentInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentInput

`func (o *WorkloadWorkloadInstance) SetLastDeploymentInput(v WorkloadDeploymentInputRelationship)`

SetLastDeploymentInput sets LastDeploymentInput field to given value.

### HasLastDeploymentInput

`func (o *WorkloadWorkloadInstance) HasLastDeploymentInput() bool`

HasLastDeploymentInput returns a boolean if a field has been set.

### SetLastDeploymentInputNil

`func (o *WorkloadWorkloadInstance) SetLastDeploymentInputNil(b bool)`

 SetLastDeploymentInputNil sets the value for LastDeploymentInput to be an explicit nil

### UnsetLastDeploymentInput
`func (o *WorkloadWorkloadInstance) UnsetLastDeploymentInput()`

UnsetLastDeploymentInput ensures that no value is present for LastDeploymentInput, not even an explicit nil
### GetOrganization

`func (o *WorkloadWorkloadInstance) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkloadWorkloadInstance) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkloadWorkloadInstance) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkloadWorkloadInstance) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *WorkloadWorkloadInstance) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *WorkloadWorkloadInstance) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetServiceItemInstances

`func (o *WorkloadWorkloadInstance) GetServiceItemInstances() []WorkflowServiceItemInstanceRelationship`

GetServiceItemInstances returns the ServiceItemInstances field if non-nil, zero value otherwise.

### GetServiceItemInstancesOk

`func (o *WorkloadWorkloadInstance) GetServiceItemInstancesOk() (*[]WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstancesOk returns a tuple with the ServiceItemInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstances

`func (o *WorkloadWorkloadInstance) SetServiceItemInstances(v []WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstances sets ServiceItemInstances field to given value.

### HasServiceItemInstances

`func (o *WorkloadWorkloadInstance) HasServiceItemInstances() bool`

HasServiceItemInstances returns a boolean if a field has been set.

### SetServiceItemInstancesNil

`func (o *WorkloadWorkloadInstance) SetServiceItemInstancesNil(b bool)`

 SetServiceItemInstancesNil sets the value for ServiceItemInstances to be an explicit nil

### UnsetServiceItemInstances
`func (o *WorkloadWorkloadInstance) UnsetServiceItemInstances()`

UnsetServiceItemInstances ensures that no value is present for ServiceItemInstances, not even an explicit nil
### GetWorkloadDefinition

`func (o *WorkloadWorkloadInstance) GetWorkloadDefinition() WorkloadWorkloadDefinitionRelationship`

GetWorkloadDefinition returns the WorkloadDefinition field if non-nil, zero value otherwise.

### GetWorkloadDefinitionOk

`func (o *WorkloadWorkloadInstance) GetWorkloadDefinitionOk() (*WorkloadWorkloadDefinitionRelationship, bool)`

GetWorkloadDefinitionOk returns a tuple with the WorkloadDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDefinition

`func (o *WorkloadWorkloadInstance) SetWorkloadDefinition(v WorkloadWorkloadDefinitionRelationship)`

SetWorkloadDefinition sets WorkloadDefinition field to given value.

### HasWorkloadDefinition

`func (o *WorkloadWorkloadInstance) HasWorkloadDefinition() bool`

HasWorkloadDefinition returns a boolean if a field has been set.

### SetWorkloadDefinitionNil

`func (o *WorkloadWorkloadInstance) SetWorkloadDefinitionNil(b bool)`

 SetWorkloadDefinitionNil sets the value for WorkloadDefinition to be an explicit nil

### UnsetWorkloadDefinition
`func (o *WorkloadWorkloadInstance) UnsetWorkloadDefinition()`

UnsetWorkloadDefinition ensures that no value is present for WorkloadDefinition, not even an explicit nil
### GetWorkloadDeployment

`func (o *WorkloadWorkloadInstance) GetWorkloadDeployment() WorkloadWorkloadDeploymentRelationship`

GetWorkloadDeployment returns the WorkloadDeployment field if non-nil, zero value otherwise.

### GetWorkloadDeploymentOk

`func (o *WorkloadWorkloadInstance) GetWorkloadDeploymentOk() (*WorkloadWorkloadDeploymentRelationship, bool)`

GetWorkloadDeploymentOk returns a tuple with the WorkloadDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDeployment

`func (o *WorkloadWorkloadInstance) SetWorkloadDeployment(v WorkloadWorkloadDeploymentRelationship)`

SetWorkloadDeployment sets WorkloadDeployment field to given value.

### HasWorkloadDeployment

`func (o *WorkloadWorkloadInstance) HasWorkloadDeployment() bool`

HasWorkloadDeployment returns a boolean if a field has been set.

### SetWorkloadDeploymentNil

`func (o *WorkloadWorkloadInstance) SetWorkloadDeploymentNil(b bool)`

 SetWorkloadDeploymentNil sets the value for WorkloadDeployment to be an explicit nil

### UnsetWorkloadDeployment
`func (o *WorkloadWorkloadInstance) UnsetWorkloadDeployment()`

UnsetWorkloadDeployment ensures that no value is present for WorkloadDeployment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


