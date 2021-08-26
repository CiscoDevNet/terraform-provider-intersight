# ServerConfigImportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.ConfigImport"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.ConfigImport"]
**Description** | Pointer to **string** | Description of the imported profile. | [optional] 
**PolicyPrefix** | Pointer to **string** | Policy prefix for the policies of the imported server profile. | [optional] 
**PolicyTypes** | Pointer to **[]string** |  | [optional] 
**ProfileName** | Pointer to **string** | Profile name for the imported server profile. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Server** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**ServerProfile** | Pointer to [**ServerProfileRelationship**](ServerProfileRelationship.md) |  | [optional] 

## Methods

### NewServerConfigImportAllOf

`func NewServerConfigImportAllOf(classId string, objectType string, ) *ServerConfigImportAllOf`

NewServerConfigImportAllOf instantiates a new ServerConfigImportAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigImportAllOfWithDefaults

`func NewServerConfigImportAllOfWithDefaults() *ServerConfigImportAllOf`

NewServerConfigImportAllOfWithDefaults instantiates a new ServerConfigImportAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerConfigImportAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerConfigImportAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerConfigImportAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerConfigImportAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerConfigImportAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerConfigImportAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ServerConfigImportAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerConfigImportAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerConfigImportAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerConfigImportAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyPrefix

`func (o *ServerConfigImportAllOf) GetPolicyPrefix() string`

GetPolicyPrefix returns the PolicyPrefix field if non-nil, zero value otherwise.

### GetPolicyPrefixOk

`func (o *ServerConfigImportAllOf) GetPolicyPrefixOk() (*string, bool)`

GetPolicyPrefixOk returns a tuple with the PolicyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPrefix

`func (o *ServerConfigImportAllOf) SetPolicyPrefix(v string)`

SetPolicyPrefix sets PolicyPrefix field to given value.

### HasPolicyPrefix

`func (o *ServerConfigImportAllOf) HasPolicyPrefix() bool`

HasPolicyPrefix returns a boolean if a field has been set.

### GetPolicyTypes

`func (o *ServerConfigImportAllOf) GetPolicyTypes() []string`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *ServerConfigImportAllOf) GetPolicyTypesOk() (*[]string, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *ServerConfigImportAllOf) SetPolicyTypes(v []string)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *ServerConfigImportAllOf) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.

### SetPolicyTypesNil

`func (o *ServerConfigImportAllOf) SetPolicyTypesNil(b bool)`

 SetPolicyTypesNil sets the value for PolicyTypes to be an explicit nil

### UnsetPolicyTypes
`func (o *ServerConfigImportAllOf) UnsetPolicyTypes()`

UnsetPolicyTypes ensures that no value is present for PolicyTypes, not even an explicit nil
### GetProfileName

`func (o *ServerConfigImportAllOf) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ServerConfigImportAllOf) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ServerConfigImportAllOf) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ServerConfigImportAllOf) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetOrganization

`func (o *ServerConfigImportAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServerConfigImportAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServerConfigImportAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServerConfigImportAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetServer

`func (o *ServerConfigImportAllOf) GetServer() ComputeRackUnitRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerConfigImportAllOf) GetServerOk() (*ComputeRackUnitRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerConfigImportAllOf) SetServer(v ComputeRackUnitRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ServerConfigImportAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServerProfile

`func (o *ServerConfigImportAllOf) GetServerProfile() ServerProfileRelationship`

GetServerProfile returns the ServerProfile field if non-nil, zero value otherwise.

### GetServerProfileOk

`func (o *ServerConfigImportAllOf) GetServerProfileOk() (*ServerProfileRelationship, bool)`

GetServerProfileOk returns a tuple with the ServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProfile

`func (o *ServerConfigImportAllOf) SetServerProfile(v ServerProfileRelationship)`

SetServerProfile sets ServerProfile field to given value.

### HasServerProfile

`func (o *ServerConfigImportAllOf) HasServerProfile() bool`

HasServerProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


