# WorkflowDisplayMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DisplayMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DisplayMeta"]
**InventorySelector** | Pointer to **bool** | Inventory selector specified for primitive data property should be used in Intersight User Interface. | [optional] [default to true]
**WidgetType** | Pointer to **string** | Specify the widget type for data display. * &#x60;None&#x60; - Display none of the widget types. * &#x60;Radio&#x60; - Display the widget as a radio button. * &#x60;Dropdown&#x60; - Display the widget as a dropdown. * &#x60;GridSelector&#x60; - Display the widget as a selector. * &#x60;DrawerSelector&#x60; - Display the widget as a selector. * &#x60;MultiSelect&#x60; - Display the widget as a multi-select. | [optional] [default to "None"]

## Methods

### NewWorkflowDisplayMeta

`func NewWorkflowDisplayMeta(classId string, objectType string, ) *WorkflowDisplayMeta`

NewWorkflowDisplayMeta instantiates a new WorkflowDisplayMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDisplayMetaWithDefaults

`func NewWorkflowDisplayMetaWithDefaults() *WorkflowDisplayMeta`

NewWorkflowDisplayMetaWithDefaults instantiates a new WorkflowDisplayMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDisplayMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDisplayMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDisplayMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDisplayMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDisplayMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDisplayMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInventorySelector

`func (o *WorkflowDisplayMeta) GetInventorySelector() bool`

GetInventorySelector returns the InventorySelector field if non-nil, zero value otherwise.

### GetInventorySelectorOk

`func (o *WorkflowDisplayMeta) GetInventorySelectorOk() (*bool, bool)`

GetInventorySelectorOk returns a tuple with the InventorySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySelector

`func (o *WorkflowDisplayMeta) SetInventorySelector(v bool)`

SetInventorySelector sets InventorySelector field to given value.

### HasInventorySelector

`func (o *WorkflowDisplayMeta) HasInventorySelector() bool`

HasInventorySelector returns a boolean if a field has been set.

### GetWidgetType

`func (o *WorkflowDisplayMeta) GetWidgetType() string`

GetWidgetType returns the WidgetType field if non-nil, zero value otherwise.

### GetWidgetTypeOk

`func (o *WorkflowDisplayMeta) GetWidgetTypeOk() (*string, bool)`

GetWidgetTypeOk returns a tuple with the WidgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetType

`func (o *WorkflowDisplayMeta) SetWidgetType(v string)`

SetWidgetType sets WidgetType field to given value.

### HasWidgetType

`func (o *WorkflowDisplayMeta) HasWidgetType() bool`

HasWidgetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


