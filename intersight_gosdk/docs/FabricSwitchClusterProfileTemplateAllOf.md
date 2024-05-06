# FabricSwitchClusterProfileTemplateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchClusterProfileTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchClusterProfileTemplate"]
**Usage** | Pointer to **int64** | The count of switch cluster profiles derived from the template. | [optional] [readonly] [default to 0]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SwitchProfileTemplates** | Pointer to [**[]FabricSwitchProfileTemplateRelationship**](FabricSwitchProfileTemplateRelationship.md) | An array of relationships to fabricSwitchProfileTemplate resources. | [optional] 

## Methods

### NewFabricSwitchClusterProfileTemplateAllOf

`func NewFabricSwitchClusterProfileTemplateAllOf(classId string, objectType string, ) *FabricSwitchClusterProfileTemplateAllOf`

NewFabricSwitchClusterProfileTemplateAllOf instantiates a new FabricSwitchClusterProfileTemplateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchClusterProfileTemplateAllOfWithDefaults

`func NewFabricSwitchClusterProfileTemplateAllOfWithDefaults() *FabricSwitchClusterProfileTemplateAllOf`

NewFabricSwitchClusterProfileTemplateAllOfWithDefaults instantiates a new FabricSwitchClusterProfileTemplateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchClusterProfileTemplateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchClusterProfileTemplateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUsage

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FabricSwitchClusterProfileTemplateAllOf) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FabricSwitchClusterProfileTemplateAllOf) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchClusterProfileTemplateAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchClusterProfileTemplateAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSwitchProfileTemplates

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetSwitchProfileTemplates() []FabricSwitchProfileTemplateRelationship`

GetSwitchProfileTemplates returns the SwitchProfileTemplates field if non-nil, zero value otherwise.

### GetSwitchProfileTemplatesOk

`func (o *FabricSwitchClusterProfileTemplateAllOf) GetSwitchProfileTemplatesOk() (*[]FabricSwitchProfileTemplateRelationship, bool)`

GetSwitchProfileTemplatesOk returns a tuple with the SwitchProfileTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfileTemplates

`func (o *FabricSwitchClusterProfileTemplateAllOf) SetSwitchProfileTemplates(v []FabricSwitchProfileTemplateRelationship)`

SetSwitchProfileTemplates sets SwitchProfileTemplates field to given value.

### HasSwitchProfileTemplates

`func (o *FabricSwitchClusterProfileTemplateAllOf) HasSwitchProfileTemplates() bool`

HasSwitchProfileTemplates returns a boolean if a field has been set.

### SetSwitchProfileTemplatesNil

`func (o *FabricSwitchClusterProfileTemplateAllOf) SetSwitchProfileTemplatesNil(b bool)`

 SetSwitchProfileTemplatesNil sets the value for SwitchProfileTemplates to be an explicit nil

### UnsetSwitchProfileTemplates
`func (o *FabricSwitchClusterProfileTemplateAllOf) UnsetSwitchProfileTemplates()`

UnsetSwitchProfileTemplates ensures that no value is present for SwitchProfileTemplates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


