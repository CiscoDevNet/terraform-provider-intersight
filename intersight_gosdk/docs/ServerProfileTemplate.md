# ServerProfileTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.ProfileTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.ProfileTemplate"]
**Usage** | Pointer to **int64** | The count of the server profiles derived from the template. | [optional] [readonly] [default to 0]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewServerProfileTemplate

`func NewServerProfileTemplate(classId string, objectType string, ) *ServerProfileTemplate`

NewServerProfileTemplate instantiates a new ServerProfileTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerProfileTemplateWithDefaults

`func NewServerProfileTemplateWithDefaults() *ServerProfileTemplate`

NewServerProfileTemplateWithDefaults instantiates a new ServerProfileTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerProfileTemplate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerProfileTemplate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerProfileTemplate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerProfileTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerProfileTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerProfileTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUsage

`func (o *ServerProfileTemplate) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ServerProfileTemplate) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ServerProfileTemplate) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ServerProfileTemplate) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetOrganization

`func (o *ServerProfileTemplate) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServerProfileTemplate) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServerProfileTemplate) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServerProfileTemplate) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


