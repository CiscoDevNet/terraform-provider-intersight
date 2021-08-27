# SdwanProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdwan.Profile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdwan.Profile"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RouterNodes** | Pointer to [**[]SdwanRouterNodeRelationship**](SdwanRouterNodeRelationship.md) | An array of relationships to sdwanRouterNode resources. | [optional] 
**RouterPolicy** | Pointer to [**SdwanRouterPolicyRelationship**](SdwanRouterPolicyRelationship.md) |  | [optional] 
**VmanageAccount** | Pointer to [**SdwanVmanageAccountPolicyRelationship**](SdwanVmanageAccountPolicyRelationship.md) |  | [optional] 

## Methods

### NewSdwanProfile

`func NewSdwanProfile(classId string, objectType string, ) *SdwanProfile`

NewSdwanProfile instantiates a new SdwanProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanProfileWithDefaults

`func NewSdwanProfileWithDefaults() *SdwanProfile`

NewSdwanProfileWithDefaults instantiates a new SdwanProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdwanProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdwanProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


