# FabricClusterAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ClusterAssignment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ClusterAssignment"]
**NetworkElement** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**SourceSwitchProfileOrTemplateName** | Pointer to **string** | Name of the source SwitchProfile or SwitchProfileTemplate whose clone has to be assigned to the network element mentioned in NetworkElement property under ClusterAssignments. | [optional] 

## Methods

### NewFabricClusterAssignment

`func NewFabricClusterAssignment(classId string, objectType string, ) *FabricClusterAssignment`

NewFabricClusterAssignment instantiates a new FabricClusterAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricClusterAssignmentWithDefaults

`func NewFabricClusterAssignmentWithDefaults() *FabricClusterAssignment`

NewFabricClusterAssignmentWithDefaults instantiates a new FabricClusterAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricClusterAssignment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricClusterAssignment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricClusterAssignment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricClusterAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricClusterAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricClusterAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNetworkElement

`func (o *FabricClusterAssignment) GetNetworkElement() MoMoRef`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricClusterAssignment) GetNetworkElementOk() (*MoMoRef, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricClusterAssignment) SetNetworkElement(v MoMoRef)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricClusterAssignment) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetSourceSwitchProfileOrTemplateName

`func (o *FabricClusterAssignment) GetSourceSwitchProfileOrTemplateName() string`

GetSourceSwitchProfileOrTemplateName returns the SourceSwitchProfileOrTemplateName field if non-nil, zero value otherwise.

### GetSourceSwitchProfileOrTemplateNameOk

`func (o *FabricClusterAssignment) GetSourceSwitchProfileOrTemplateNameOk() (*string, bool)`

GetSourceSwitchProfileOrTemplateNameOk returns a tuple with the SourceSwitchProfileOrTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSwitchProfileOrTemplateName

`func (o *FabricClusterAssignment) SetSourceSwitchProfileOrTemplateName(v string)`

SetSourceSwitchProfileOrTemplateName sets SourceSwitchProfileOrTemplateName field to given value.

### HasSourceSwitchProfileOrTemplateName

`func (o *FabricClusterAssignment) HasSourceSwitchProfileOrTemplateName() bool`

HasSourceSwitchProfileOrTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


