# FabricElementIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ElementIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ElementIdentity"]
**Domain** | Pointer to **string** | Name of the Fabric Interconnect domain. | [optional] [readonly] 
**ReplacementType** | Pointer to **string** | Replacement type specifies whether it is single FI or domain replacement. * &#x60;None&#x60; - The default action is none. * &#x60;Individual&#x60; - Replacement of single network element. * &#x60;Domain&#x60; - Domain indicates the replacement of Fabric Interconnect domain. | [optional] [default to "None"]
**SwitchId** | Pointer to **string** | Switch Identifier that uniquely represents the fabric object. * &#x60;A&#x60; - Switch Identifier of Fabric Interconnect A. * &#x60;B&#x60; - Switch Identifier of Fabric Interconnect B. | [optional] [readonly] [default to "A"]
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**ReplacementTarget** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 

## Methods

### NewFabricElementIdentity

`func NewFabricElementIdentity(classId string, objectType string, ) *FabricElementIdentity`

NewFabricElementIdentity instantiates a new FabricElementIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricElementIdentityWithDefaults

`func NewFabricElementIdentityWithDefaults() *FabricElementIdentity`

NewFabricElementIdentityWithDefaults instantiates a new FabricElementIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricElementIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricElementIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricElementIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricElementIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricElementIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricElementIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomain

`func (o *FabricElementIdentity) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *FabricElementIdentity) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *FabricElementIdentity) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *FabricElementIdentity) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetReplacementType

`func (o *FabricElementIdentity) GetReplacementType() string`

GetReplacementType returns the ReplacementType field if non-nil, zero value otherwise.

### GetReplacementTypeOk

`func (o *FabricElementIdentity) GetReplacementTypeOk() (*string, bool)`

GetReplacementTypeOk returns a tuple with the ReplacementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementType

`func (o *FabricElementIdentity) SetReplacementType(v string)`

SetReplacementType sets ReplacementType field to given value.

### HasReplacementType

`func (o *FabricElementIdentity) HasReplacementType() bool`

HasReplacementType returns a boolean if a field has been set.

### GetSwitchId

`func (o *FabricElementIdentity) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FabricElementIdentity) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FabricElementIdentity) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FabricElementIdentity) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricElementIdentity) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricElementIdentity) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricElementIdentity) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricElementIdentity) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetReplacementTarget

`func (o *FabricElementIdentity) GetReplacementTarget() NetworkElementRelationship`

GetReplacementTarget returns the ReplacementTarget field if non-nil, zero value otherwise.

### GetReplacementTargetOk

`func (o *FabricElementIdentity) GetReplacementTargetOk() (*NetworkElementRelationship, bool)`

GetReplacementTargetOk returns a tuple with the ReplacementTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementTarget

`func (o *FabricElementIdentity) SetReplacementTarget(v NetworkElementRelationship)`

SetReplacementTarget sets ReplacementTarget field to given value.

### HasReplacementTarget

`func (o *FabricElementIdentity) HasReplacementTarget() bool`

HasReplacementTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


