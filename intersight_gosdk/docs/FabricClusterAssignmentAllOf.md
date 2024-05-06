# FabricClusterAssignmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ClusterAssignment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ClusterAssignment"]
**NetworkElement** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**SourceSwitchProfileOrTemplateName** | Pointer to **string** | Name of the source SwitchProfile or SwitchProfileTemplate whose clone has to be assigned to the network element mentioned in NetworkElement property under ClusterAssignments. | [optional] 

## Methods

### NewFabricClusterAssignmentAllOf

`func NewFabricClusterAssignmentAllOf(classId string, objectType string, ) *FabricClusterAssignmentAllOf`

NewFabricClusterAssignmentAllOf instantiates a new FabricClusterAssignmentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricClusterAssignmentAllOfWithDefaults

`func NewFabricClusterAssignmentAllOfWithDefaults() *FabricClusterAssignmentAllOf`

NewFabricClusterAssignmentAllOfWithDefaults instantiates a new FabricClusterAssignmentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricClusterAssignmentAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricClusterAssignmentAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricClusterAssignmentAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricClusterAssignmentAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricClusterAssignmentAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricClusterAssignmentAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNetworkElement

`func (o *FabricClusterAssignmentAllOf) GetNetworkElement() MoMoRef`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricClusterAssignmentAllOf) GetNetworkElementOk() (*MoMoRef, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricClusterAssignmentAllOf) SetNetworkElement(v MoMoRef)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricClusterAssignmentAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetSourceSwitchProfileOrTemplateName

`func (o *FabricClusterAssignmentAllOf) GetSourceSwitchProfileOrTemplateName() string`

GetSourceSwitchProfileOrTemplateName returns the SourceSwitchProfileOrTemplateName field if non-nil, zero value otherwise.

### GetSourceSwitchProfileOrTemplateNameOk

`func (o *FabricClusterAssignmentAllOf) GetSourceSwitchProfileOrTemplateNameOk() (*string, bool)`

GetSourceSwitchProfileOrTemplateNameOk returns a tuple with the SourceSwitchProfileOrTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSwitchProfileOrTemplateName

`func (o *FabricClusterAssignmentAllOf) SetSourceSwitchProfileOrTemplateName(v string)`

SetSourceSwitchProfileOrTemplateName sets SourceSwitchProfileOrTemplateName field to given value.

### HasSourceSwitchProfileOrTemplateName

`func (o *FabricClusterAssignmentAllOf) HasSourceSwitchProfileOrTemplateName() bool`

HasSourceSwitchProfileOrTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


