# ChassisIomProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.IomProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.IomProfile"]
**ConfigChanges** | Pointer to [**NullablePolicyConfigChange**](PolicyConfigChange.md) |  | [optional] 
**IomEntity** | Pointer to **string** | IOM in chassis for which IOM profile is applicable. or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;IOMA&#x60; - IOM on left side of chassis. * &#x60;IOMB&#x60; - IOM on right side of chassis. | [optional] [default to "IOMA"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profile** | Pointer to [**ChassisProfileRelationship**](ChassisProfileRelationship.md) |  | [optional] 

## Methods

### NewChassisIomProfileAllOf

`func NewChassisIomProfileAllOf(classId string, objectType string, ) *ChassisIomProfileAllOf`

NewChassisIomProfileAllOf instantiates a new ChassisIomProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisIomProfileAllOfWithDefaults

`func NewChassisIomProfileAllOfWithDefaults() *ChassisIomProfileAllOf`

NewChassisIomProfileAllOfWithDefaults instantiates a new ChassisIomProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisIomProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisIomProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisIomProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisIomProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisIomProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisIomProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigChanges

`func (o *ChassisIomProfileAllOf) GetConfigChanges() PolicyConfigChange`

GetConfigChanges returns the ConfigChanges field if non-nil, zero value otherwise.

### GetConfigChangesOk

`func (o *ChassisIomProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool)`

GetConfigChangesOk returns a tuple with the ConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChanges

`func (o *ChassisIomProfileAllOf) SetConfigChanges(v PolicyConfigChange)`

SetConfigChanges sets ConfigChanges field to given value.

### HasConfigChanges

`func (o *ChassisIomProfileAllOf) HasConfigChanges() bool`

HasConfigChanges returns a boolean if a field has been set.

### SetConfigChangesNil

`func (o *ChassisIomProfileAllOf) SetConfigChangesNil(b bool)`

 SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil

### UnsetConfigChanges
`func (o *ChassisIomProfileAllOf) UnsetConfigChanges()`

UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
### GetIomEntity

`func (o *ChassisIomProfileAllOf) GetIomEntity() string`

GetIomEntity returns the IomEntity field if non-nil, zero value otherwise.

### GetIomEntityOk

`func (o *ChassisIomProfileAllOf) GetIomEntityOk() (*string, bool)`

GetIomEntityOk returns a tuple with the IomEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomEntity

`func (o *ChassisIomProfileAllOf) SetIomEntity(v string)`

SetIomEntity sets IomEntity field to given value.

### HasIomEntity

`func (o *ChassisIomProfileAllOf) HasIomEntity() bool`

HasIomEntity returns a boolean if a field has been set.

### GetOrganization

`func (o *ChassisIomProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ChassisIomProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ChassisIomProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ChassisIomProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfile

`func (o *ChassisIomProfileAllOf) GetProfile() ChassisProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ChassisIomProfileAllOf) GetProfileOk() (*ChassisProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ChassisIomProfileAllOf) SetProfile(v ChassisProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ChassisIomProfileAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


