# EquipmentFexIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.FexIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.FexIdentity"]
**SwitchId** | Pointer to **int64** | Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B. | [optional] [readonly] 
**Fex** | Pointer to [**EquipmentFexRelationship**](EquipmentFexRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewEquipmentFexIdentity

`func NewEquipmentFexIdentity(classId string, objectType string, ) *EquipmentFexIdentity`

NewEquipmentFexIdentity instantiates a new EquipmentFexIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFexIdentityWithDefaults

`func NewEquipmentFexIdentityWithDefaults() *EquipmentFexIdentity`

NewEquipmentFexIdentityWithDefaults instantiates a new EquipmentFexIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentFexIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentFexIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentFexIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentFexIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentFexIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentFexIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSwitchId

`func (o *EquipmentFexIdentity) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentFexIdentity) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentFexIdentity) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentFexIdentity) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetFex

`func (o *EquipmentFexIdentity) GetFex() EquipmentFexRelationship`

GetFex returns the Fex field if non-nil, zero value otherwise.

### GetFexOk

`func (o *EquipmentFexIdentity) GetFexOk() (*EquipmentFexRelationship, bool)`

GetFexOk returns a tuple with the Fex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFex

`func (o *EquipmentFexIdentity) SetFex(v EquipmentFexRelationship)`

SetFex sets Fex field to given value.

### HasFex

`func (o *EquipmentFexIdentity) HasFex() bool`

HasFex returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentFexIdentity) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentFexIdentity) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentFexIdentity) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentFexIdentity) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


