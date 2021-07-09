# FabricElementIdentityAllOf

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

### NewFabricElementIdentityAllOf

`func NewFabricElementIdentityAllOf(classId string, objectType string, ) *FabricElementIdentityAllOf`

NewFabricElementIdentityAllOf instantiates a new FabricElementIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricElementIdentityAllOfWithDefaults

`func NewFabricElementIdentityAllOfWithDefaults() *FabricElementIdentityAllOf`

NewFabricElementIdentityAllOfWithDefaults instantiates a new FabricElementIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricElementIdentityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricElementIdentityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricElementIdentityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricElementIdentityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricElementIdentityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricElementIdentityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomain

`func (o *FabricElementIdentityAllOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *FabricElementIdentityAllOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *FabricElementIdentityAllOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *FabricElementIdentityAllOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetReplacementType

`func (o *FabricElementIdentityAllOf) GetReplacementType() string`

GetReplacementType returns the ReplacementType field if non-nil, zero value otherwise.

### GetReplacementTypeOk

`func (o *FabricElementIdentityAllOf) GetReplacementTypeOk() (*string, bool)`

GetReplacementTypeOk returns a tuple with the ReplacementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementType

`func (o *FabricElementIdentityAllOf) SetReplacementType(v string)`

SetReplacementType sets ReplacementType field to given value.

### HasReplacementType

`func (o *FabricElementIdentityAllOf) HasReplacementType() bool`

HasReplacementType returns a boolean if a field has been set.

### GetSwitchId

`func (o *FabricElementIdentityAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FabricElementIdentityAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FabricElementIdentityAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FabricElementIdentityAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricElementIdentityAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricElementIdentityAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricElementIdentityAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricElementIdentityAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetReplacementTarget

`func (o *FabricElementIdentityAllOf) GetReplacementTarget() NetworkElementRelationship`

GetReplacementTarget returns the ReplacementTarget field if non-nil, zero value otherwise.

### GetReplacementTargetOk

`func (o *FabricElementIdentityAllOf) GetReplacementTargetOk() (*NetworkElementRelationship, bool)`

GetReplacementTargetOk returns a tuple with the ReplacementTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementTarget

`func (o *FabricElementIdentityAllOf) SetReplacementTarget(v NetworkElementRelationship)`

SetReplacementTarget sets ReplacementTarget field to given value.

### HasReplacementTarget

`func (o *FabricElementIdentityAllOf) HasReplacementTarget() bool`

HasReplacementTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


