# VirtualizationBaseNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Identity** | Pointer to **string** | The internally generated identity of network. This entity cannot manipulated by users. It aids in uniquely identifying the network object. For VMware, this is MOR (managed object reference). | [optional] 
**Name** | Pointer to **string** | User-provided name to identify the portgroup. | [optional] 

## Methods

### NewVirtualizationBaseNetwork

`func NewVirtualizationBaseNetwork(classId string, objectType string, ) *VirtualizationBaseNetwork`

NewVirtualizationBaseNetwork instantiates a new VirtualizationBaseNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseNetworkWithDefaults

`func NewVirtualizationBaseNetworkWithDefaults() *VirtualizationBaseNetwork`

NewVirtualizationBaseNetworkWithDefaults instantiates a new VirtualizationBaseNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseNetwork) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseNetwork) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseNetwork) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseNetwork) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseNetwork) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseNetwork) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *VirtualizationBaseNetwork) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseNetwork) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseNetwork) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseNetwork) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseNetwork) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


