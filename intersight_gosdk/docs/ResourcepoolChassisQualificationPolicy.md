# ResourcepoolChassisQualificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.ChassisQualificationPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.ChassisQualificationPolicy"]
**ExcludeServers** | Pointer to **bool** | When set to true, qualify the chassis alone into pool. When set to false, qualify the servers like Rack server and Blade along with chassis. | [optional] [default to false]
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ResourcePools** | Pointer to [**[]ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) | An array of relationships to resourcepoolPool resources. | [optional] [readonly] 

## Methods

### NewResourcepoolChassisQualificationPolicy

`func NewResourcepoolChassisQualificationPolicy(classId string, objectType string, ) *ResourcepoolChassisQualificationPolicy`

NewResourcepoolChassisQualificationPolicy instantiates a new ResourcepoolChassisQualificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolChassisQualificationPolicyWithDefaults

`func NewResourcepoolChassisQualificationPolicyWithDefaults() *ResourcepoolChassisQualificationPolicy`

NewResourcepoolChassisQualificationPolicyWithDefaults instantiates a new ResourcepoolChassisQualificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolChassisQualificationPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolChassisQualificationPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolChassisQualificationPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolChassisQualificationPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolChassisQualificationPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolChassisQualificationPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeServers

`func (o *ResourcepoolChassisQualificationPolicy) GetExcludeServers() bool`

GetExcludeServers returns the ExcludeServers field if non-nil, zero value otherwise.

### GetExcludeServersOk

`func (o *ResourcepoolChassisQualificationPolicy) GetExcludeServersOk() (*bool, bool)`

GetExcludeServersOk returns a tuple with the ExcludeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeServers

`func (o *ResourcepoolChassisQualificationPolicy) SetExcludeServers(v bool)`

SetExcludeServers sets ExcludeServers field to given value.

### HasExcludeServers

`func (o *ResourcepoolChassisQualificationPolicy) HasExcludeServers() bool`

HasExcludeServers returns a boolean if a field has been set.

### GetOrganization

`func (o *ResourcepoolChassisQualificationPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResourcepoolChassisQualificationPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResourcepoolChassisQualificationPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResourcepoolChassisQualificationPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ResourcepoolChassisQualificationPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ResourcepoolChassisQualificationPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetResourcePools

`func (o *ResourcepoolChassisQualificationPolicy) GetResourcePools() []ResourcepoolPoolRelationship`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *ResourcepoolChassisQualificationPolicy) GetResourcePoolsOk() (*[]ResourcepoolPoolRelationship, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *ResourcepoolChassisQualificationPolicy) SetResourcePools(v []ResourcepoolPoolRelationship)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *ResourcepoolChassisQualificationPolicy) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.

### SetResourcePoolsNil

`func (o *ResourcepoolChassisQualificationPolicy) SetResourcePoolsNil(b bool)`

 SetResourcePoolsNil sets the value for ResourcePools to be an explicit nil

### UnsetResourcePools
`func (o *ResourcepoolChassisQualificationPolicy) UnsetResourcePools()`

UnsetResourcePools ensures that no value is present for ResourcePools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


