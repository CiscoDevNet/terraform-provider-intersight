# SdwanProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**RouterNodes** | Pointer to [**[]SdwanRouterNodeRelationship**](sdwan.RouterNode.Relationship.md) | An array of relationships to sdwanRouterNode resources. | [optional] 
**RouterPolicy** | Pointer to [**SdwanRouterPolicyRelationship**](sdwan.RouterPolicy.Relationship.md) |  | [optional] 
**VmanageAccount** | Pointer to [**SdwanVmanageAccountPolicyRelationship**](sdwan.VmanageAccountPolicy.Relationship.md) |  | [optional] 

## Methods

### NewSdwanProfileAllOf

`func NewSdwanProfileAllOf() *SdwanProfileAllOf`

NewSdwanProfileAllOf instantiates a new SdwanProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanProfileAllOfWithDefaults

`func NewSdwanProfileAllOfWithDefaults() *SdwanProfileAllOf`

NewSdwanProfileAllOfWithDefaults instantiates a new SdwanProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *SdwanProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRouterNodes

`func (o *SdwanProfileAllOf) GetRouterNodes() []SdwanRouterNodeRelationship`

GetRouterNodes returns the RouterNodes field if non-nil, zero value otherwise.

### GetRouterNodesOk

`func (o *SdwanProfileAllOf) GetRouterNodesOk() (*[]SdwanRouterNodeRelationship, bool)`

GetRouterNodesOk returns a tuple with the RouterNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterNodes

`func (o *SdwanProfileAllOf) SetRouterNodes(v []SdwanRouterNodeRelationship)`

SetRouterNodes sets RouterNodes field to given value.

### HasRouterNodes

`func (o *SdwanProfileAllOf) HasRouterNodes() bool`

HasRouterNodes returns a boolean if a field has been set.

### SetRouterNodesNil

`func (o *SdwanProfileAllOf) SetRouterNodesNil(b bool)`

 SetRouterNodesNil sets the value for RouterNodes to be an explicit nil

### UnsetRouterNodes
`func (o *SdwanProfileAllOf) UnsetRouterNodes()`

UnsetRouterNodes ensures that no value is present for RouterNodes, not even an explicit nil
### GetRouterPolicy

`func (o *SdwanProfileAllOf) GetRouterPolicy() SdwanRouterPolicyRelationship`

GetRouterPolicy returns the RouterPolicy field if non-nil, zero value otherwise.

### GetRouterPolicyOk

`func (o *SdwanProfileAllOf) GetRouterPolicyOk() (*SdwanRouterPolicyRelationship, bool)`

GetRouterPolicyOk returns a tuple with the RouterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPolicy

`func (o *SdwanProfileAllOf) SetRouterPolicy(v SdwanRouterPolicyRelationship)`

SetRouterPolicy sets RouterPolicy field to given value.

### HasRouterPolicy

`func (o *SdwanProfileAllOf) HasRouterPolicy() bool`

HasRouterPolicy returns a boolean if a field has been set.

### GetVmanageAccount

`func (o *SdwanProfileAllOf) GetVmanageAccount() SdwanVmanageAccountPolicyRelationship`

GetVmanageAccount returns the VmanageAccount field if non-nil, zero value otherwise.

### GetVmanageAccountOk

`func (o *SdwanProfileAllOf) GetVmanageAccountOk() (*SdwanVmanageAccountPolicyRelationship, bool)`

GetVmanageAccountOk returns a tuple with the VmanageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmanageAccount

`func (o *SdwanProfileAllOf) SetVmanageAccount(v SdwanVmanageAccountPolicyRelationship)`

SetVmanageAccount sets VmanageAccount field to given value.

### HasVmanageAccount

`func (o *SdwanProfileAllOf) HasVmanageAccount() bool`

HasVmanageAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


