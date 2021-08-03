# CloudTfcAgentpoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.TfcAgentpool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.TfcAgentpool"]
**Identity** | Pointer to **string** | The ID of the Agent Pool. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the agent pool. | [optional] [readonly] 
**NumActiveAgents** | Pointer to **int64** | The number of active agents used by this pool. The total active agent are sum of idle, busy and unknown agent counts. | [optional] [readonly] 
**NumTokens** | Pointer to **int64** | The number of Tokens in this agent Pool. | [optional] [readonly] 
**Organization** | Pointer to [**CloudTfcOrganizationRelationship**](cloud.TfcOrganization.Relationship.md) |  | [optional] 

## Methods

### NewCloudTfcAgentpoolAllOf

`func NewCloudTfcAgentpoolAllOf(classId string, objectType string, ) *CloudTfcAgentpoolAllOf`

NewCloudTfcAgentpoolAllOf instantiates a new CloudTfcAgentpoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTfcAgentpoolAllOfWithDefaults

`func NewCloudTfcAgentpoolAllOfWithDefaults() *CloudTfcAgentpoolAllOf`

NewCloudTfcAgentpoolAllOfWithDefaults instantiates a new CloudTfcAgentpoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudTfcAgentpoolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudTfcAgentpoolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudTfcAgentpoolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudTfcAgentpoolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudTfcAgentpoolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudTfcAgentpoolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *CloudTfcAgentpoolAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudTfcAgentpoolAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudTfcAgentpoolAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudTfcAgentpoolAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *CloudTfcAgentpoolAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTfcAgentpoolAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTfcAgentpoolAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTfcAgentpoolAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumActiveAgents

`func (o *CloudTfcAgentpoolAllOf) GetNumActiveAgents() int64`

GetNumActiveAgents returns the NumActiveAgents field if non-nil, zero value otherwise.

### GetNumActiveAgentsOk

`func (o *CloudTfcAgentpoolAllOf) GetNumActiveAgentsOk() (*int64, bool)`

GetNumActiveAgentsOk returns a tuple with the NumActiveAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveAgents

`func (o *CloudTfcAgentpoolAllOf) SetNumActiveAgents(v int64)`

SetNumActiveAgents sets NumActiveAgents field to given value.

### HasNumActiveAgents

`func (o *CloudTfcAgentpoolAllOf) HasNumActiveAgents() bool`

HasNumActiveAgents returns a boolean if a field has been set.

### GetNumTokens

`func (o *CloudTfcAgentpoolAllOf) GetNumTokens() int64`

GetNumTokens returns the NumTokens field if non-nil, zero value otherwise.

### GetNumTokensOk

`func (o *CloudTfcAgentpoolAllOf) GetNumTokensOk() (*int64, bool)`

GetNumTokensOk returns a tuple with the NumTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTokens

`func (o *CloudTfcAgentpoolAllOf) SetNumTokens(v int64)`

SetNumTokens sets NumTokens field to given value.

### HasNumTokens

`func (o *CloudTfcAgentpoolAllOf) HasNumTokens() bool`

HasNumTokens returns a boolean if a field has been set.

### GetOrganization

`func (o *CloudTfcAgentpoolAllOf) GetOrganization() CloudTfcOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CloudTfcAgentpoolAllOf) GetOrganizationOk() (*CloudTfcOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CloudTfcAgentpoolAllOf) SetOrganization(v CloudTfcOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CloudTfcAgentpoolAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


