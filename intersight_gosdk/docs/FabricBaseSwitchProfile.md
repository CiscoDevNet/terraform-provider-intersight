# FabricBaseSwitchProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**SwitchId** | Pointer to **string** | Value indicating the switch side on which the switch profile or template has to be deployed. * &#x60;None&#x60; - Switch side not defined for the policy configurations in the switch profile or template. * &#x60;A&#x60; - Policy configurations in the switch profile or template to be deployed on fabric interconnect A. * &#x60;B&#x60; - Policy configurations in the switch profile or template to be deployed on fabric interconnect B. | [optional] [default to "None"]
**ConfigResult** | Pointer to [**NullableFabricConfigResultRelationship**](FabricConfigResultRelationship.md) |  | [optional] 

## Methods

### NewFabricBaseSwitchProfile

`func NewFabricBaseSwitchProfile(classId string, objectType string, ) *FabricBaseSwitchProfile`

NewFabricBaseSwitchProfile instantiates a new FabricBaseSwitchProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricBaseSwitchProfileWithDefaults

`func NewFabricBaseSwitchProfileWithDefaults() *FabricBaseSwitchProfile`

NewFabricBaseSwitchProfileWithDefaults instantiates a new FabricBaseSwitchProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricBaseSwitchProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricBaseSwitchProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricBaseSwitchProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricBaseSwitchProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricBaseSwitchProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricBaseSwitchProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSwitchId

`func (o *FabricBaseSwitchProfile) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FabricBaseSwitchProfile) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FabricBaseSwitchProfile) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FabricBaseSwitchProfile) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetConfigResult

`func (o *FabricBaseSwitchProfile) GetConfigResult() FabricConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *FabricBaseSwitchProfile) GetConfigResultOk() (*FabricConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *FabricBaseSwitchProfile) SetConfigResult(v FabricConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *FabricBaseSwitchProfile) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### SetConfigResultNil

`func (o *FabricBaseSwitchProfile) SetConfigResultNil(b bool)`

 SetConfigResultNil sets the value for ConfigResult to be an explicit nil

### UnsetConfigResult
`func (o *FabricBaseSwitchProfile) UnsetConfigResult()`

UnsetConfigResult ensures that no value is present for ConfigResult, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


