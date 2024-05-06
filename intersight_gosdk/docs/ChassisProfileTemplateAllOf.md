# ChassisProfileTemplateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.ProfileTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.ProfileTemplate"]
**Usage** | Pointer to **int64** | The count of the chassis profiles derived from the template. | [optional] [readonly] [default to 0]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewChassisProfileTemplateAllOf

`func NewChassisProfileTemplateAllOf(classId string, objectType string, ) *ChassisProfileTemplateAllOf`

NewChassisProfileTemplateAllOf instantiates a new ChassisProfileTemplateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisProfileTemplateAllOfWithDefaults

`func NewChassisProfileTemplateAllOfWithDefaults() *ChassisProfileTemplateAllOf`

NewChassisProfileTemplateAllOfWithDefaults instantiates a new ChassisProfileTemplateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisProfileTemplateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisProfileTemplateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisProfileTemplateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisProfileTemplateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisProfileTemplateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisProfileTemplateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUsage

`func (o *ChassisProfileTemplateAllOf) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChassisProfileTemplateAllOf) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChassisProfileTemplateAllOf) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChassisProfileTemplateAllOf) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetOrganization

`func (o *ChassisProfileTemplateAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ChassisProfileTemplateAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ChassisProfileTemplateAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ChassisProfileTemplateAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


