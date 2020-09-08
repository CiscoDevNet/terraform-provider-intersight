# WorkflowPrimitiveDataPropertyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**WorkflowConstraints**](workflow.Constraints.md) |  | [optional] 
**InventorySelector** | Pointer to [**[]WorkflowMoReferenceProperty**](workflow.MoReferenceProperty.md) |  | [optional] 
**Secure** | Pointer to **bool** | Intersight supports secure properties as task input/output. The values of these properties are encrypted and stored in Intersight. This flag marks the property to be secure when it is set to true. | [optional] 
**Type** | Pointer to **string** | Specify the enum type for primitive data type. * &#x60;string&#x60; - Enum to specify a string data type. * &#x60;integer&#x60; - Enum to specify an integer32 data type. * &#x60;float&#x60; - Enum to specify a float64 data type. * &#x60;boolean&#x60; - Enum to specify a boolean data type. * &#x60;json&#x60; - Enum to specify a json data type. * &#x60;enum&#x60; - Enum to specify a enum data type which is a list of pre-defined strings. | [optional] [default to "string"]

## Methods

### NewWorkflowPrimitiveDataPropertyAllOf

`func NewWorkflowPrimitiveDataPropertyAllOf() *WorkflowPrimitiveDataPropertyAllOf`

NewWorkflowPrimitiveDataPropertyAllOf instantiates a new WorkflowPrimitiveDataPropertyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPrimitiveDataPropertyAllOfWithDefaults

`func NewWorkflowPrimitiveDataPropertyAllOfWithDefaults() *WorkflowPrimitiveDataPropertyAllOf`

NewWorkflowPrimitiveDataPropertyAllOfWithDefaults instantiates a new WorkflowPrimitiveDataPropertyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetConstraints() WorkflowConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetConstraintsOk() (*WorkflowConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WorkflowPrimitiveDataPropertyAllOf) SetConstraints(v WorkflowConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WorkflowPrimitiveDataPropertyAllOf) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetInventorySelector

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetInventorySelector() []WorkflowMoReferenceProperty`

GetInventorySelector returns the InventorySelector field if non-nil, zero value otherwise.

### GetInventorySelectorOk

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetInventorySelectorOk() (*[]WorkflowMoReferenceProperty, bool)`

GetInventorySelectorOk returns a tuple with the InventorySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySelector

`func (o *WorkflowPrimitiveDataPropertyAllOf) SetInventorySelector(v []WorkflowMoReferenceProperty)`

SetInventorySelector sets InventorySelector field to given value.

### HasInventorySelector

`func (o *WorkflowPrimitiveDataPropertyAllOf) HasInventorySelector() bool`

HasInventorySelector returns a boolean if a field has been set.

### GetSecure

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *WorkflowPrimitiveDataPropertyAllOf) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *WorkflowPrimitiveDataPropertyAllOf) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetType

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowPrimitiveDataPropertyAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowPrimitiveDataPropertyAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowPrimitiveDataPropertyAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


