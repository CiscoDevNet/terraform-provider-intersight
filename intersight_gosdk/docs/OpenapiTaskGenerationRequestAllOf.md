# OpenapiTaskGenerationRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.TaskGenerationRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.TaskGenerationRequest"]
**EndpointType** | Pointer to **string** | Indicates if target endpoint is external or internal. An endpoint is internal if the target is an Intersight resource. For instance, configuring an intersight object using a Task. * &#x60;External&#x60; - Denotes that the target endpoint is an external API endpoint * &#x60;Internal&#x60; - Denotes that the target endpoint is a Intersight API endpoint | [optional] [default to "External"]
**FailedTasks** | Pointer to [**[]OpenapiFailedTask**](OpenapiFailedTask.md) |  | [optional] 
**HeaderParameters** | Pointer to [**[]OpenapiKeyValuePair**](OpenapiKeyValuePair.md) |  | [optional] 
**IsValidateRequest** | Pointer to **bool** | When this value is set to true, the task name validations happen and provides the task validation status against each of the selected API on the property selectedApis When this value is set to false, an internal workflow to generate the tasks is submitted,  conflicting task names are created with an incremented version. | [optional] 
**Protocol** | Pointer to **string** | Specifies the REST API protocol being used, which can be one of HTTP or HTTPS. * &#x60;HTTPS&#x60; - Denotes that the API call uses the HTTPS protocol type * &#x60;HTTP&#x60; - Denotes that the API call uses the HTTP protocol type | [optional] [default to "HTTPS"]
**QueryParameters** | Pointer to [**[]OpenapiKeyValuePair**](OpenapiKeyValuePair.md) |  | [optional] 
**SelectedApis** | Pointer to [**[]OpenapiApiInfo**](OpenapiApiInfo.md) |  | [optional] 
**Status** | Pointer to **string** | Depicts the status of the task creation request. * &#x60;none&#x60; - Indicates the default status. * &#x60;InProgress&#x60; - Request has been picked up for generating tasks from the OpenAPI Specification file. * &#x60;Completed&#x60; - All the tasks from the request have been created. * &#x60;Failed&#x60; - There were failures in generating one or more tasks in the request. | [optional] [readonly] [default to "none"]
**TaskPrefix** | Pointer to **string** | Optional string that can be prefixed to the name of created tasks. | [optional] 
**TaskTags** | Pointer to [**[]OpenapiKeyValuePair**](OpenapiKeyValuePair.md) |  | [optional] 
**Url** | Pointer to **string** | Specifies the URL of the endpoint that the created task communicates with. It is defaulted to intersight.com for internal endpoints. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Source** | Pointer to [**OpenapiProcessFileRelationship**](OpenapiProcessFileRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewOpenapiTaskGenerationRequestAllOf

`func NewOpenapiTaskGenerationRequestAllOf(classId string, objectType string, ) *OpenapiTaskGenerationRequestAllOf`

NewOpenapiTaskGenerationRequestAllOf instantiates a new OpenapiTaskGenerationRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiTaskGenerationRequestAllOfWithDefaults

`func NewOpenapiTaskGenerationRequestAllOfWithDefaults() *OpenapiTaskGenerationRequestAllOf`

NewOpenapiTaskGenerationRequestAllOfWithDefaults instantiates a new OpenapiTaskGenerationRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiTaskGenerationRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiTaskGenerationRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiTaskGenerationRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiTaskGenerationRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndpointType

`func (o *OpenapiTaskGenerationRequestAllOf) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *OpenapiTaskGenerationRequestAllOf) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *OpenapiTaskGenerationRequestAllOf) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetFailedTasks

`func (o *OpenapiTaskGenerationRequestAllOf) GetFailedTasks() []OpenapiFailedTask`

GetFailedTasks returns the FailedTasks field if non-nil, zero value otherwise.

### GetFailedTasksOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetFailedTasksOk() (*[]OpenapiFailedTask, bool)`

GetFailedTasksOk returns a tuple with the FailedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTasks

`func (o *OpenapiTaskGenerationRequestAllOf) SetFailedTasks(v []OpenapiFailedTask)`

SetFailedTasks sets FailedTasks field to given value.

### HasFailedTasks

`func (o *OpenapiTaskGenerationRequestAllOf) HasFailedTasks() bool`

HasFailedTasks returns a boolean if a field has been set.

### SetFailedTasksNil

`func (o *OpenapiTaskGenerationRequestAllOf) SetFailedTasksNil(b bool)`

 SetFailedTasksNil sets the value for FailedTasks to be an explicit nil

### UnsetFailedTasks
`func (o *OpenapiTaskGenerationRequestAllOf) UnsetFailedTasks()`

UnsetFailedTasks ensures that no value is present for FailedTasks, not even an explicit nil
### GetHeaderParameters

`func (o *OpenapiTaskGenerationRequestAllOf) GetHeaderParameters() []OpenapiKeyValuePair`

GetHeaderParameters returns the HeaderParameters field if non-nil, zero value otherwise.

### GetHeaderParametersOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetHeaderParametersOk() (*[]OpenapiKeyValuePair, bool)`

GetHeaderParametersOk returns a tuple with the HeaderParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderParameters

`func (o *OpenapiTaskGenerationRequestAllOf) SetHeaderParameters(v []OpenapiKeyValuePair)`

SetHeaderParameters sets HeaderParameters field to given value.

### HasHeaderParameters

`func (o *OpenapiTaskGenerationRequestAllOf) HasHeaderParameters() bool`

HasHeaderParameters returns a boolean if a field has been set.

### SetHeaderParametersNil

`func (o *OpenapiTaskGenerationRequestAllOf) SetHeaderParametersNil(b bool)`

 SetHeaderParametersNil sets the value for HeaderParameters to be an explicit nil

### UnsetHeaderParameters
`func (o *OpenapiTaskGenerationRequestAllOf) UnsetHeaderParameters()`

UnsetHeaderParameters ensures that no value is present for HeaderParameters, not even an explicit nil
### GetIsValidateRequest

`func (o *OpenapiTaskGenerationRequestAllOf) GetIsValidateRequest() bool`

GetIsValidateRequest returns the IsValidateRequest field if non-nil, zero value otherwise.

### GetIsValidateRequestOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetIsValidateRequestOk() (*bool, bool)`

GetIsValidateRequestOk returns a tuple with the IsValidateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidateRequest

`func (o *OpenapiTaskGenerationRequestAllOf) SetIsValidateRequest(v bool)`

SetIsValidateRequest sets IsValidateRequest field to given value.

### HasIsValidateRequest

`func (o *OpenapiTaskGenerationRequestAllOf) HasIsValidateRequest() bool`

HasIsValidateRequest returns a boolean if a field has been set.

### GetProtocol

`func (o *OpenapiTaskGenerationRequestAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OpenapiTaskGenerationRequestAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *OpenapiTaskGenerationRequestAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetQueryParameters

`func (o *OpenapiTaskGenerationRequestAllOf) GetQueryParameters() []OpenapiKeyValuePair`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetQueryParametersOk() (*[]OpenapiKeyValuePair, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *OpenapiTaskGenerationRequestAllOf) SetQueryParameters(v []OpenapiKeyValuePair)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *OpenapiTaskGenerationRequestAllOf) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### SetQueryParametersNil

`func (o *OpenapiTaskGenerationRequestAllOf) SetQueryParametersNil(b bool)`

 SetQueryParametersNil sets the value for QueryParameters to be an explicit nil

### UnsetQueryParameters
`func (o *OpenapiTaskGenerationRequestAllOf) UnsetQueryParameters()`

UnsetQueryParameters ensures that no value is present for QueryParameters, not even an explicit nil
### GetSelectedApis

`func (o *OpenapiTaskGenerationRequestAllOf) GetSelectedApis() []OpenapiApiInfo`

GetSelectedApis returns the SelectedApis field if non-nil, zero value otherwise.

### GetSelectedApisOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetSelectedApisOk() (*[]OpenapiApiInfo, bool)`

GetSelectedApisOk returns a tuple with the SelectedApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedApis

`func (o *OpenapiTaskGenerationRequestAllOf) SetSelectedApis(v []OpenapiApiInfo)`

SetSelectedApis sets SelectedApis field to given value.

### HasSelectedApis

`func (o *OpenapiTaskGenerationRequestAllOf) HasSelectedApis() bool`

HasSelectedApis returns a boolean if a field has been set.

### SetSelectedApisNil

`func (o *OpenapiTaskGenerationRequestAllOf) SetSelectedApisNil(b bool)`

 SetSelectedApisNil sets the value for SelectedApis to be an explicit nil

### UnsetSelectedApis
`func (o *OpenapiTaskGenerationRequestAllOf) UnsetSelectedApis()`

UnsetSelectedApis ensures that no value is present for SelectedApis, not even an explicit nil
### GetStatus

`func (o *OpenapiTaskGenerationRequestAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenapiTaskGenerationRequestAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenapiTaskGenerationRequestAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskPrefix

`func (o *OpenapiTaskGenerationRequestAllOf) GetTaskPrefix() string`

GetTaskPrefix returns the TaskPrefix field if non-nil, zero value otherwise.

### GetTaskPrefixOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetTaskPrefixOk() (*string, bool)`

GetTaskPrefixOk returns a tuple with the TaskPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPrefix

`func (o *OpenapiTaskGenerationRequestAllOf) SetTaskPrefix(v string)`

SetTaskPrefix sets TaskPrefix field to given value.

### HasTaskPrefix

`func (o *OpenapiTaskGenerationRequestAllOf) HasTaskPrefix() bool`

HasTaskPrefix returns a boolean if a field has been set.

### GetTaskTags

`func (o *OpenapiTaskGenerationRequestAllOf) GetTaskTags() []OpenapiKeyValuePair`

GetTaskTags returns the TaskTags field if non-nil, zero value otherwise.

### GetTaskTagsOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetTaskTagsOk() (*[]OpenapiKeyValuePair, bool)`

GetTaskTagsOk returns a tuple with the TaskTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTags

`func (o *OpenapiTaskGenerationRequestAllOf) SetTaskTags(v []OpenapiKeyValuePair)`

SetTaskTags sets TaskTags field to given value.

### HasTaskTags

`func (o *OpenapiTaskGenerationRequestAllOf) HasTaskTags() bool`

HasTaskTags returns a boolean if a field has been set.

### SetTaskTagsNil

`func (o *OpenapiTaskGenerationRequestAllOf) SetTaskTagsNil(b bool)`

 SetTaskTagsNil sets the value for TaskTags to be an explicit nil

### UnsetTaskTags
`func (o *OpenapiTaskGenerationRequestAllOf) UnsetTaskTags()`

UnsetTaskTags ensures that no value is present for TaskTags, not even an explicit nil
### GetUrl

`func (o *OpenapiTaskGenerationRequestAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OpenapiTaskGenerationRequestAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OpenapiTaskGenerationRequestAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOrganization

`func (o *OpenapiTaskGenerationRequestAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OpenapiTaskGenerationRequestAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OpenapiTaskGenerationRequestAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSource

`func (o *OpenapiTaskGenerationRequestAllOf) GetSource() OpenapiProcessFileRelationship`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetSourceOk() (*OpenapiProcessFileRelationship, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OpenapiTaskGenerationRequestAllOf) SetSource(v OpenapiProcessFileRelationship)`

SetSource sets Source field to given value.

### HasSource

`func (o *OpenapiTaskGenerationRequestAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetWorkflow

`func (o *OpenapiTaskGenerationRequestAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *OpenapiTaskGenerationRequestAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *OpenapiTaskGenerationRequestAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *OpenapiTaskGenerationRequestAllOf) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


