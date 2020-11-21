# PolicyConfigResultContextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ConfigResultContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ConfigResultContext"]
**EntityData** | Pointer to **interface{}** | The data of the object present in config result context. | [optional] 
**EntityMoid** | Pointer to **string** | The Moid of the object present in config result context. | [optional] 
**EntityName** | Pointer to **string** | The name of the object present in config result context. | [optional] 
**EntityType** | Pointer to **string** | The type of the object present in config result context. | [optional] 

## Methods

### NewPolicyConfigResultContextAllOf

`func NewPolicyConfigResultContextAllOf(classId string, objectType string, ) *PolicyConfigResultContextAllOf`

NewPolicyConfigResultContextAllOf instantiates a new PolicyConfigResultContextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigResultContextAllOfWithDefaults

`func NewPolicyConfigResultContextAllOfWithDefaults() *PolicyConfigResultContextAllOf`

NewPolicyConfigResultContextAllOfWithDefaults instantiates a new PolicyConfigResultContextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyConfigResultContextAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyConfigResultContextAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyConfigResultContextAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyConfigResultContextAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyConfigResultContextAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyConfigResultContextAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEntityData

`func (o *PolicyConfigResultContextAllOf) GetEntityData() interface{}`

GetEntityData returns the EntityData field if non-nil, zero value otherwise.

### GetEntityDataOk

`func (o *PolicyConfigResultContextAllOf) GetEntityDataOk() (*interface{}, bool)`

GetEntityDataOk returns a tuple with the EntityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityData

`func (o *PolicyConfigResultContextAllOf) SetEntityData(v interface{})`

SetEntityData sets EntityData field to given value.

### HasEntityData

`func (o *PolicyConfigResultContextAllOf) HasEntityData() bool`

HasEntityData returns a boolean if a field has been set.

### SetEntityDataNil

`func (o *PolicyConfigResultContextAllOf) SetEntityDataNil(b bool)`

 SetEntityDataNil sets the value for EntityData to be an explicit nil

### UnsetEntityData
`func (o *PolicyConfigResultContextAllOf) UnsetEntityData()`

UnsetEntityData ensures that no value is present for EntityData, not even an explicit nil
### GetEntityMoid

`func (o *PolicyConfigResultContextAllOf) GetEntityMoid() string`

GetEntityMoid returns the EntityMoid field if non-nil, zero value otherwise.

### GetEntityMoidOk

`func (o *PolicyConfigResultContextAllOf) GetEntityMoidOk() (*string, bool)`

GetEntityMoidOk returns a tuple with the EntityMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMoid

`func (o *PolicyConfigResultContextAllOf) SetEntityMoid(v string)`

SetEntityMoid sets EntityMoid field to given value.

### HasEntityMoid

`func (o *PolicyConfigResultContextAllOf) HasEntityMoid() bool`

HasEntityMoid returns a boolean if a field has been set.

### GetEntityName

`func (o *PolicyConfigResultContextAllOf) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *PolicyConfigResultContextAllOf) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *PolicyConfigResultContextAllOf) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *PolicyConfigResultContextAllOf) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityType

`func (o *PolicyConfigResultContextAllOf) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PolicyConfigResultContextAllOf) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PolicyConfigResultContextAllOf) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *PolicyConfigResultContextAllOf) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


