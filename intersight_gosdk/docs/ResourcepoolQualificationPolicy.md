# ResourcepoolQualificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.QualificationPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.QualificationPolicy"]
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ResourcePools** | Pointer to [**[]ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) | An array of relationships to resourcepoolPool resources. | [optional] 

## Methods

### NewResourcepoolQualificationPolicy

`func NewResourcepoolQualificationPolicy(classId string, objectType string, ) *ResourcepoolQualificationPolicy`

NewResourcepoolQualificationPolicy instantiates a new ResourcepoolQualificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolQualificationPolicyWithDefaults

`func NewResourcepoolQualificationPolicyWithDefaults() *ResourcepoolQualificationPolicy`

NewResourcepoolQualificationPolicyWithDefaults instantiates a new ResourcepoolQualificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolQualificationPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolQualificationPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolQualificationPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolQualificationPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolQualificationPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolQualificationPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOrganization

`func (o *ResourcepoolQualificationPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResourcepoolQualificationPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResourcepoolQualificationPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResourcepoolQualificationPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ResourcepoolQualificationPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ResourcepoolQualificationPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetResourcePools

`func (o *ResourcepoolQualificationPolicy) GetResourcePools() []ResourcepoolPoolRelationship`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *ResourcepoolQualificationPolicy) GetResourcePoolsOk() (*[]ResourcepoolPoolRelationship, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *ResourcepoolQualificationPolicy) SetResourcePools(v []ResourcepoolPoolRelationship)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *ResourcepoolQualificationPolicy) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.

### SetResourcePoolsNil

`func (o *ResourcepoolQualificationPolicy) SetResourcePoolsNil(b bool)`

 SetResourcePoolsNil sets the value for ResourcePools to be an explicit nil

### UnsetResourcePools
`func (o *ResourcepoolQualificationPolicy) UnsetResourcePools()`

UnsetResourcePools ensures that no value is present for ResourcePools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


