# FabricSecureRouterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SecureRouterRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SecureRouterRole"]
**EthNetworkGroupPolicy** | Pointer to [**[]FabricEthNetworkGroupPolicyRelationship**](FabricEthNetworkGroupPolicyRelationship.md) | An array of relationships to fabricEthNetworkGroupPolicy resources. | [optional] 

## Methods

### NewFabricSecureRouterRole

`func NewFabricSecureRouterRole(classId string, objectType string, ) *FabricSecureRouterRole`

NewFabricSecureRouterRole instantiates a new FabricSecureRouterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSecureRouterRoleWithDefaults

`func NewFabricSecureRouterRoleWithDefaults() *FabricSecureRouterRole`

NewFabricSecureRouterRoleWithDefaults instantiates a new FabricSecureRouterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSecureRouterRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSecureRouterRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSecureRouterRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSecureRouterRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSecureRouterRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSecureRouterRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEthNetworkGroupPolicy

`func (o *FabricSecureRouterRole) GetEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship`

GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetEthNetworkGroupPolicyOk

`func (o *FabricSecureRouterRole) GetEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool)`

GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkGroupPolicy

`func (o *FabricSecureRouterRole) SetEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship)`

SetEthNetworkGroupPolicy sets EthNetworkGroupPolicy field to given value.

### HasEthNetworkGroupPolicy

`func (o *FabricSecureRouterRole) HasEthNetworkGroupPolicy() bool`

HasEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetEthNetworkGroupPolicyNil

`func (o *FabricSecureRouterRole) SetEthNetworkGroupPolicyNil(b bool)`

 SetEthNetworkGroupPolicyNil sets the value for EthNetworkGroupPolicy to be an explicit nil

### UnsetEthNetworkGroupPolicy
`func (o *FabricSecureRouterRole) UnsetEthNetworkGroupPolicy()`

UnsetEthNetworkGroupPolicy ensures that no value is present for EthNetworkGroupPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


