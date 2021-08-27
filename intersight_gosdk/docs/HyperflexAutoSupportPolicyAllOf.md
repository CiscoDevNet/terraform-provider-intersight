# HyperflexAutoSupportPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.AutoSupportPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.AutoSupportPolicy"]
**AdminState** | Pointer to **bool** | Enable or disable Auto-Support. | [optional] [default to true]
**ServiceTicketReceipient** | Pointer to **string** | The recipient email address for support tickets. | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexAutoSupportPolicyAllOf

`func NewHyperflexAutoSupportPolicyAllOf(classId string, objectType string, ) *HyperflexAutoSupportPolicyAllOf`

NewHyperflexAutoSupportPolicyAllOf instantiates a new HyperflexAutoSupportPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAutoSupportPolicyAllOfWithDefaults

`func NewHyperflexAutoSupportPolicyAllOfWithDefaults() *HyperflexAutoSupportPolicyAllOf`

NewHyperflexAutoSupportPolicyAllOfWithDefaults instantiates a new HyperflexAutoSupportPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexAutoSupportPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexAutoSupportPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexAutoSupportPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexAutoSupportPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexAutoSupportPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexAutoSupportPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *HyperflexAutoSupportPolicyAllOf) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *HyperflexAutoSupportPolicyAllOf) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *HyperflexAutoSupportPolicyAllOf) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *HyperflexAutoSupportPolicyAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetServiceTicketReceipient

`func (o *HyperflexAutoSupportPolicyAllOf) GetServiceTicketReceipient() string`

GetServiceTicketReceipient returns the ServiceTicketReceipient field if non-nil, zero value otherwise.

### GetServiceTicketReceipientOk

`func (o *HyperflexAutoSupportPolicyAllOf) GetServiceTicketReceipientOk() (*string, bool)`

GetServiceTicketReceipientOk returns a tuple with the ServiceTicketReceipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTicketReceipient

`func (o *HyperflexAutoSupportPolicyAllOf) SetServiceTicketReceipient(v string)`

SetServiceTicketReceipient sets ServiceTicketReceipient field to given value.

### HasServiceTicketReceipient

`func (o *HyperflexAutoSupportPolicyAllOf) HasServiceTicketReceipient() bool`

HasServiceTicketReceipient returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexAutoSupportPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexAutoSupportPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexAutoSupportPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexAutoSupportPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexAutoSupportPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexAutoSupportPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexAutoSupportPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexAutoSupportPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexAutoSupportPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexAutoSupportPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


