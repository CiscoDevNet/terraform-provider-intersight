# SdwanProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**RouterNodes** | Pointer to [**[]SdwanRouterNodeRelationship**](sdwan.RouterNode.Relationship.md) | An array of relationships to sdwanRouterNode resources. | [optional] 
**RouterPolicy** | Pointer to [**SdwanRouterPolicyRelationship**](sdwan.RouterPolicy.Relationship.md) |  | [optional] 
**VmanageAccount** | Pointer to [**SdwanVmanageAccountPolicyRelationship**](sdwan.VmanageAccountPolicy.Relationship.md) |  | [optional] 

## Methods

### NewSdwanProfile

`func NewSdwanProfile() *SdwanProfile`

NewSdwanProfile instantiates a new SdwanProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanProfileWithDefaults

`func NewSdwanProfileWithDefaults() *SdwanProfile`

NewSdwanProfileWithDefaults instantiates a new SdwanProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *SdwanProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRouterNodes

`func (o *SdwanProfile) GetRouterNodes() []SdwanRouterNodeRelationship`

GetRouterNodes returns the RouterNodes field if non-nil, zero value otherwise.

### GetRouterNodesOk

`func (o *SdwanProfile) GetRouterNodesOk() (*[]SdwanRouterNodeRelationship, bool)`

GetRouterNodesOk returns a tuple with the RouterNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterNodes

`func (o *SdwanProfile) SetRouterNodes(v []SdwanRouterNodeRelationship)`

SetRouterNodes sets RouterNodes field to given value.

### HasRouterNodes

`func (o *SdwanProfile) HasRouterNodes() bool`

HasRouterNodes returns a boolean if a field has been set.

### SetRouterNodesNil

`func (o *SdwanProfile) SetRouterNodesNil(b bool)`

 SetRouterNodesNil sets the value for RouterNodes to be an explicit nil

### UnsetRouterNodes
`func (o *SdwanProfile) UnsetRouterNodes()`

UnsetRouterNodes ensures that no value is present for RouterNodes, not even an explicit nil
### GetRouterPolicy

`func (o *SdwanProfile) GetRouterPolicy() SdwanRouterPolicyRelationship`

GetRouterPolicy returns the RouterPolicy field if non-nil, zero value otherwise.

### GetRouterPolicyOk

`func (o *SdwanProfile) GetRouterPolicyOk() (*SdwanRouterPolicyRelationship, bool)`

GetRouterPolicyOk returns a tuple with the RouterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPolicy

`func (o *SdwanProfile) SetRouterPolicy(v SdwanRouterPolicyRelationship)`

SetRouterPolicy sets RouterPolicy field to given value.

### HasRouterPolicy

`func (o *SdwanProfile) HasRouterPolicy() bool`

HasRouterPolicy returns a boolean if a field has been set.

### GetVmanageAccount

`func (o *SdwanProfile) GetVmanageAccount() SdwanVmanageAccountPolicyRelationship`

GetVmanageAccount returns the VmanageAccount field if non-nil, zero value otherwise.

### GetVmanageAccountOk

`func (o *SdwanProfile) GetVmanageAccountOk() (*SdwanVmanageAccountPolicyRelationship, bool)`

GetVmanageAccountOk returns a tuple with the VmanageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmanageAccount

`func (o *SdwanProfile) SetVmanageAccount(v SdwanVmanageAccountPolicyRelationship)`

SetVmanageAccount sets VmanageAccount field to given value.

### HasVmanageAccount

`func (o *SdwanProfile) HasVmanageAccount() bool`

HasVmanageAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


