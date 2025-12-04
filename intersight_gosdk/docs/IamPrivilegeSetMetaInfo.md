# IamPrivilegeSetMetaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.PrivilegeSetMetaInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.PrivilegeSetMetaInfo"]
**AllDependencyInfo** | Pointer to **interface{}** | All dependency information of the privilege set. | [optional] [readonly] 
**AssociatedPrivilegeSetNames** | Pointer to **[]string** |  | [optional] 
**DownstreamDependencies** | Pointer to **[]string** |  | [optional] 
**MissingDependencyInfo** | Pointer to **interface{}** | Missing dependency information of the privilege set. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the privilege set. | [optional] [readonly] 
**UpstreamDependencies** | Pointer to **[]string** |  | [optional] 
**Uuid** | Pointer to **string** | UUID of the privilege set. | [optional] [readonly] 
**AssociatedPrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] 
**MissingPrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 

## Methods

### NewIamPrivilegeSetMetaInfo

`func NewIamPrivilegeSetMetaInfo(classId string, objectType string, ) *IamPrivilegeSetMetaInfo`

NewIamPrivilegeSetMetaInfo instantiates a new IamPrivilegeSetMetaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivilegeSetMetaInfoWithDefaults

`func NewIamPrivilegeSetMetaInfoWithDefaults() *IamPrivilegeSetMetaInfo`

NewIamPrivilegeSetMetaInfoWithDefaults instantiates a new IamPrivilegeSetMetaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPrivilegeSetMetaInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPrivilegeSetMetaInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPrivilegeSetMetaInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPrivilegeSetMetaInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPrivilegeSetMetaInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPrivilegeSetMetaInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllDependencyInfo

`func (o *IamPrivilegeSetMetaInfo) GetAllDependencyInfo() interface{}`

GetAllDependencyInfo returns the AllDependencyInfo field if non-nil, zero value otherwise.

### GetAllDependencyInfoOk

`func (o *IamPrivilegeSetMetaInfo) GetAllDependencyInfoOk() (*interface{}, bool)`

GetAllDependencyInfoOk returns a tuple with the AllDependencyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDependencyInfo

`func (o *IamPrivilegeSetMetaInfo) SetAllDependencyInfo(v interface{})`

SetAllDependencyInfo sets AllDependencyInfo field to given value.

### HasAllDependencyInfo

`func (o *IamPrivilegeSetMetaInfo) HasAllDependencyInfo() bool`

HasAllDependencyInfo returns a boolean if a field has been set.

### SetAllDependencyInfoNil

`func (o *IamPrivilegeSetMetaInfo) SetAllDependencyInfoNil(b bool)`

 SetAllDependencyInfoNil sets the value for AllDependencyInfo to be an explicit nil

### UnsetAllDependencyInfo
`func (o *IamPrivilegeSetMetaInfo) UnsetAllDependencyInfo()`

UnsetAllDependencyInfo ensures that no value is present for AllDependencyInfo, not even an explicit nil
### GetAssociatedPrivilegeSetNames

`func (o *IamPrivilegeSetMetaInfo) GetAssociatedPrivilegeSetNames() []string`

GetAssociatedPrivilegeSetNames returns the AssociatedPrivilegeSetNames field if non-nil, zero value otherwise.

### GetAssociatedPrivilegeSetNamesOk

`func (o *IamPrivilegeSetMetaInfo) GetAssociatedPrivilegeSetNamesOk() (*[]string, bool)`

GetAssociatedPrivilegeSetNamesOk returns a tuple with the AssociatedPrivilegeSetNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPrivilegeSetNames

`func (o *IamPrivilegeSetMetaInfo) SetAssociatedPrivilegeSetNames(v []string)`

SetAssociatedPrivilegeSetNames sets AssociatedPrivilegeSetNames field to given value.

### HasAssociatedPrivilegeSetNames

`func (o *IamPrivilegeSetMetaInfo) HasAssociatedPrivilegeSetNames() bool`

HasAssociatedPrivilegeSetNames returns a boolean if a field has been set.

### SetAssociatedPrivilegeSetNamesNil

`func (o *IamPrivilegeSetMetaInfo) SetAssociatedPrivilegeSetNamesNil(b bool)`

 SetAssociatedPrivilegeSetNamesNil sets the value for AssociatedPrivilegeSetNames to be an explicit nil

### UnsetAssociatedPrivilegeSetNames
`func (o *IamPrivilegeSetMetaInfo) UnsetAssociatedPrivilegeSetNames()`

UnsetAssociatedPrivilegeSetNames ensures that no value is present for AssociatedPrivilegeSetNames, not even an explicit nil
### GetDownstreamDependencies

`func (o *IamPrivilegeSetMetaInfo) GetDownstreamDependencies() []string`

GetDownstreamDependencies returns the DownstreamDependencies field if non-nil, zero value otherwise.

### GetDownstreamDependenciesOk

`func (o *IamPrivilegeSetMetaInfo) GetDownstreamDependenciesOk() (*[]string, bool)`

GetDownstreamDependenciesOk returns a tuple with the DownstreamDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamDependencies

`func (o *IamPrivilegeSetMetaInfo) SetDownstreamDependencies(v []string)`

SetDownstreamDependencies sets DownstreamDependencies field to given value.

### HasDownstreamDependencies

`func (o *IamPrivilegeSetMetaInfo) HasDownstreamDependencies() bool`

HasDownstreamDependencies returns a boolean if a field has been set.

### SetDownstreamDependenciesNil

`func (o *IamPrivilegeSetMetaInfo) SetDownstreamDependenciesNil(b bool)`

 SetDownstreamDependenciesNil sets the value for DownstreamDependencies to be an explicit nil

### UnsetDownstreamDependencies
`func (o *IamPrivilegeSetMetaInfo) UnsetDownstreamDependencies()`

UnsetDownstreamDependencies ensures that no value is present for DownstreamDependencies, not even an explicit nil
### GetMissingDependencyInfo

`func (o *IamPrivilegeSetMetaInfo) GetMissingDependencyInfo() interface{}`

GetMissingDependencyInfo returns the MissingDependencyInfo field if non-nil, zero value otherwise.

### GetMissingDependencyInfoOk

`func (o *IamPrivilegeSetMetaInfo) GetMissingDependencyInfoOk() (*interface{}, bool)`

GetMissingDependencyInfoOk returns a tuple with the MissingDependencyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingDependencyInfo

`func (o *IamPrivilegeSetMetaInfo) SetMissingDependencyInfo(v interface{})`

SetMissingDependencyInfo sets MissingDependencyInfo field to given value.

### HasMissingDependencyInfo

`func (o *IamPrivilegeSetMetaInfo) HasMissingDependencyInfo() bool`

HasMissingDependencyInfo returns a boolean if a field has been set.

### SetMissingDependencyInfoNil

`func (o *IamPrivilegeSetMetaInfo) SetMissingDependencyInfoNil(b bool)`

 SetMissingDependencyInfoNil sets the value for MissingDependencyInfo to be an explicit nil

### UnsetMissingDependencyInfo
`func (o *IamPrivilegeSetMetaInfo) UnsetMissingDependencyInfo()`

UnsetMissingDependencyInfo ensures that no value is present for MissingDependencyInfo, not even an explicit nil
### GetName

`func (o *IamPrivilegeSetMetaInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamPrivilegeSetMetaInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamPrivilegeSetMetaInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamPrivilegeSetMetaInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpstreamDependencies

`func (o *IamPrivilegeSetMetaInfo) GetUpstreamDependencies() []string`

GetUpstreamDependencies returns the UpstreamDependencies field if non-nil, zero value otherwise.

### GetUpstreamDependenciesOk

`func (o *IamPrivilegeSetMetaInfo) GetUpstreamDependenciesOk() (*[]string, bool)`

GetUpstreamDependenciesOk returns a tuple with the UpstreamDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamDependencies

`func (o *IamPrivilegeSetMetaInfo) SetUpstreamDependencies(v []string)`

SetUpstreamDependencies sets UpstreamDependencies field to given value.

### HasUpstreamDependencies

`func (o *IamPrivilegeSetMetaInfo) HasUpstreamDependencies() bool`

HasUpstreamDependencies returns a boolean if a field has been set.

### SetUpstreamDependenciesNil

`func (o *IamPrivilegeSetMetaInfo) SetUpstreamDependenciesNil(b bool)`

 SetUpstreamDependenciesNil sets the value for UpstreamDependencies to be an explicit nil

### UnsetUpstreamDependencies
`func (o *IamPrivilegeSetMetaInfo) UnsetUpstreamDependencies()`

UnsetUpstreamDependencies ensures that no value is present for UpstreamDependencies, not even an explicit nil
### GetUuid

`func (o *IamPrivilegeSetMetaInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IamPrivilegeSetMetaInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IamPrivilegeSetMetaInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IamPrivilegeSetMetaInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAssociatedPrivilegeSets

`func (o *IamPrivilegeSetMetaInfo) GetAssociatedPrivilegeSets() []IamPrivilegeSetRelationship`

GetAssociatedPrivilegeSets returns the AssociatedPrivilegeSets field if non-nil, zero value otherwise.

### GetAssociatedPrivilegeSetsOk

`func (o *IamPrivilegeSetMetaInfo) GetAssociatedPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetAssociatedPrivilegeSetsOk returns a tuple with the AssociatedPrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPrivilegeSets

`func (o *IamPrivilegeSetMetaInfo) SetAssociatedPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetAssociatedPrivilegeSets sets AssociatedPrivilegeSets field to given value.

### HasAssociatedPrivilegeSets

`func (o *IamPrivilegeSetMetaInfo) HasAssociatedPrivilegeSets() bool`

HasAssociatedPrivilegeSets returns a boolean if a field has been set.

### SetAssociatedPrivilegeSetsNil

`func (o *IamPrivilegeSetMetaInfo) SetAssociatedPrivilegeSetsNil(b bool)`

 SetAssociatedPrivilegeSetsNil sets the value for AssociatedPrivilegeSets to be an explicit nil

### UnsetAssociatedPrivilegeSets
`func (o *IamPrivilegeSetMetaInfo) UnsetAssociatedPrivilegeSets()`

UnsetAssociatedPrivilegeSets ensures that no value is present for AssociatedPrivilegeSets, not even an explicit nil
### GetMissingPrivilegeSets

`func (o *IamPrivilegeSetMetaInfo) GetMissingPrivilegeSets() []IamPrivilegeSetRelationship`

GetMissingPrivilegeSets returns the MissingPrivilegeSets field if non-nil, zero value otherwise.

### GetMissingPrivilegeSetsOk

`func (o *IamPrivilegeSetMetaInfo) GetMissingPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetMissingPrivilegeSetsOk returns a tuple with the MissingPrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPrivilegeSets

`func (o *IamPrivilegeSetMetaInfo) SetMissingPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetMissingPrivilegeSets sets MissingPrivilegeSets field to given value.

### HasMissingPrivilegeSets

`func (o *IamPrivilegeSetMetaInfo) HasMissingPrivilegeSets() bool`

HasMissingPrivilegeSets returns a boolean if a field has been set.

### SetMissingPrivilegeSetsNil

`func (o *IamPrivilegeSetMetaInfo) SetMissingPrivilegeSetsNil(b bool)`

 SetMissingPrivilegeSetsNil sets the value for MissingPrivilegeSets to be an explicit nil

### UnsetMissingPrivilegeSets
`func (o *IamPrivilegeSetMetaInfo) UnsetMissingPrivilegeSets()`

UnsetMissingPrivilegeSets ensures that no value is present for MissingPrivilegeSets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


