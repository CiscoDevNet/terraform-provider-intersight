# FabricSwitchProfileTemplateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchProfileTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchProfileTemplate"]
**SwitchClusterProfileTemplate** | Pointer to [**FabricSwitchClusterProfileTemplateRelationship**](FabricSwitchClusterProfileTemplateRelationship.md) |  | [optional] 

## Methods

### NewFabricSwitchProfileTemplateAllOf

`func NewFabricSwitchProfileTemplateAllOf(classId string, objectType string, ) *FabricSwitchProfileTemplateAllOf`

NewFabricSwitchProfileTemplateAllOf instantiates a new FabricSwitchProfileTemplateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchProfileTemplateAllOfWithDefaults

`func NewFabricSwitchProfileTemplateAllOfWithDefaults() *FabricSwitchProfileTemplateAllOf`

NewFabricSwitchProfileTemplateAllOfWithDefaults instantiates a new FabricSwitchProfileTemplateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchProfileTemplateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchProfileTemplateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchProfileTemplateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchProfileTemplateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchProfileTemplateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchProfileTemplateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSwitchClusterProfileTemplate

`func (o *FabricSwitchProfileTemplateAllOf) GetSwitchClusterProfileTemplate() FabricSwitchClusterProfileTemplateRelationship`

GetSwitchClusterProfileTemplate returns the SwitchClusterProfileTemplate field if non-nil, zero value otherwise.

### GetSwitchClusterProfileTemplateOk

`func (o *FabricSwitchProfileTemplateAllOf) GetSwitchClusterProfileTemplateOk() (*FabricSwitchClusterProfileTemplateRelationship, bool)`

GetSwitchClusterProfileTemplateOk returns a tuple with the SwitchClusterProfileTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchClusterProfileTemplate

`func (o *FabricSwitchProfileTemplateAllOf) SetSwitchClusterProfileTemplate(v FabricSwitchClusterProfileTemplateRelationship)`

SetSwitchClusterProfileTemplate sets SwitchClusterProfileTemplate field to given value.

### HasSwitchClusterProfileTemplate

`func (o *FabricSwitchProfileTemplateAllOf) HasSwitchClusterProfileTemplate() bool`

HasSwitchClusterProfileTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


