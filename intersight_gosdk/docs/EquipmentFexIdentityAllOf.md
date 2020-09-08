# EquipmentFexIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchId** | Pointer to **int64** | Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B. | [optional] 
**Fex** | Pointer to [**EquipmentFexRelationship**](equipment.Fex.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentFexIdentityAllOf

`func NewEquipmentFexIdentityAllOf() *EquipmentFexIdentityAllOf`

NewEquipmentFexIdentityAllOf instantiates a new EquipmentFexIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFexIdentityAllOfWithDefaults

`func NewEquipmentFexIdentityAllOfWithDefaults() *EquipmentFexIdentityAllOf`

NewEquipmentFexIdentityAllOfWithDefaults instantiates a new EquipmentFexIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchId

`func (o *EquipmentFexIdentityAllOf) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentFexIdentityAllOf) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentFexIdentityAllOf) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentFexIdentityAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetFex

`func (o *EquipmentFexIdentityAllOf) GetFex() EquipmentFexRelationship`

GetFex returns the Fex field if non-nil, zero value otherwise.

### GetFexOk

`func (o *EquipmentFexIdentityAllOf) GetFexOk() (*EquipmentFexRelationship, bool)`

GetFexOk returns a tuple with the Fex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFex

`func (o *EquipmentFexIdentityAllOf) SetFex(v EquipmentFexRelationship)`

SetFex sets Fex field to given value.

### HasFex

`func (o *EquipmentFexIdentityAllOf) HasFex() bool`

HasFex returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentFexIdentityAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentFexIdentityAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentFexIdentityAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentFexIdentityAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


