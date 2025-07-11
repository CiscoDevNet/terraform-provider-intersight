# ApicFabricLeafNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.FabricLeafNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.FabricLeafNode"]
**Address** | Pointer to **string** | Object address in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Dn** | Pointer to **string** | Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**FabricLeafNodeDetails** | Pointer to [**NullableApicFabricLeafNodeDetails**](ApicFabricLeafNodeDetails.md) |  | [optional] 
**Name** | Pointer to **string** | Object name in the Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Pod** | Pointer to **string** | Object pod in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 

## Methods

### NewApicFabricLeafNode

`func NewApicFabricLeafNode(classId string, objectType string, ) *ApicFabricLeafNode`

NewApicFabricLeafNode instantiates a new ApicFabricLeafNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicFabricLeafNodeWithDefaults

`func NewApicFabricLeafNodeWithDefaults() *ApicFabricLeafNode`

NewApicFabricLeafNodeWithDefaults instantiates a new ApicFabricLeafNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicFabricLeafNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicFabricLeafNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicFabricLeafNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicFabricLeafNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicFabricLeafNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicFabricLeafNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *ApicFabricLeafNode) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ApicFabricLeafNode) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ApicFabricLeafNode) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ApicFabricLeafNode) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDn

`func (o *ApicFabricLeafNode) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicFabricLeafNode) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicFabricLeafNode) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicFabricLeafNode) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricLeafNodeDetails

`func (o *ApicFabricLeafNode) GetFabricLeafNodeDetails() ApicFabricLeafNodeDetails`

GetFabricLeafNodeDetails returns the FabricLeafNodeDetails field if non-nil, zero value otherwise.

### GetFabricLeafNodeDetailsOk

`func (o *ApicFabricLeafNode) GetFabricLeafNodeDetailsOk() (*ApicFabricLeafNodeDetails, bool)`

GetFabricLeafNodeDetailsOk returns a tuple with the FabricLeafNodeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricLeafNodeDetails

`func (o *ApicFabricLeafNode) SetFabricLeafNodeDetails(v ApicFabricLeafNodeDetails)`

SetFabricLeafNodeDetails sets FabricLeafNodeDetails field to given value.

### HasFabricLeafNodeDetails

`func (o *ApicFabricLeafNode) HasFabricLeafNodeDetails() bool`

HasFabricLeafNodeDetails returns a boolean if a field has been set.

### SetFabricLeafNodeDetailsNil

`func (o *ApicFabricLeafNode) SetFabricLeafNodeDetailsNil(b bool)`

 SetFabricLeafNodeDetailsNil sets the value for FabricLeafNodeDetails to be an explicit nil

### UnsetFabricLeafNodeDetails
`func (o *ApicFabricLeafNode) UnsetFabricLeafNodeDetails()`

UnsetFabricLeafNodeDetails ensures that no value is present for FabricLeafNodeDetails, not even an explicit nil
### GetName

`func (o *ApicFabricLeafNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicFabricLeafNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicFabricLeafNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicFabricLeafNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPod

`func (o *ApicFabricLeafNode) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *ApicFabricLeafNode) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *ApicFabricLeafNode) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *ApicFabricLeafNode) HasPod() bool`

HasPod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


