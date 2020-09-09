# ServerConfigImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the imported profile. | [optional] 
**PolicyPrefix** | Pointer to **string** | Policy prefix for the policies of the imported server profile. | [optional] 
**PolicyTypes** | Pointer to **[]string** |  | [optional] 
**ProfileName** | Pointer to **string** | Profile name for the imported server profile. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**ServerProfile** | Pointer to [**ServerProfileRelationship**](server.Profile.Relationship.md) |  | [optional] 

## Methods

### NewServerConfigImport

`func NewServerConfigImport() *ServerConfigImport`

NewServerConfigImport instantiates a new ServerConfigImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigImportWithDefaults

`func NewServerConfigImportWithDefaults() *ServerConfigImport`

NewServerConfigImportWithDefaults instantiates a new ServerConfigImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServerConfigImport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerConfigImport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerConfigImport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerConfigImport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyPrefix

`func (o *ServerConfigImport) GetPolicyPrefix() string`

GetPolicyPrefix returns the PolicyPrefix field if non-nil, zero value otherwise.

### GetPolicyPrefixOk

`func (o *ServerConfigImport) GetPolicyPrefixOk() (*string, bool)`

GetPolicyPrefixOk returns a tuple with the PolicyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPrefix

`func (o *ServerConfigImport) SetPolicyPrefix(v string)`

SetPolicyPrefix sets PolicyPrefix field to given value.

### HasPolicyPrefix

`func (o *ServerConfigImport) HasPolicyPrefix() bool`

HasPolicyPrefix returns a boolean if a field has been set.

### GetPolicyTypes

`func (o *ServerConfigImport) GetPolicyTypes() []string`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *ServerConfigImport) GetPolicyTypesOk() (*[]string, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *ServerConfigImport) SetPolicyTypes(v []string)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *ServerConfigImport) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.

### GetProfileName

`func (o *ServerConfigImport) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ServerConfigImport) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ServerConfigImport) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ServerConfigImport) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetOrganization

`func (o *ServerConfigImport) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServerConfigImport) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServerConfigImport) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServerConfigImport) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetServer

`func (o *ServerConfigImport) GetServer() ComputeRackUnitRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerConfigImport) GetServerOk() (*ComputeRackUnitRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerConfigImport) SetServer(v ComputeRackUnitRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ServerConfigImport) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServerProfile

`func (o *ServerConfigImport) GetServerProfile() ServerProfileRelationship`

GetServerProfile returns the ServerProfile field if non-nil, zero value otherwise.

### GetServerProfileOk

`func (o *ServerConfigImport) GetServerProfileOk() (*ServerProfileRelationship, bool)`

GetServerProfileOk returns a tuple with the ServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProfile

`func (o *ServerConfigImport) SetServerProfile(v ServerProfileRelationship)`

SetServerProfile sets ServerProfile field to given value.

### HasServerProfile

`func (o *ServerConfigImport) HasServerProfile() bool`

HasServerProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


