# VirtualizationBasePhysicalNetworkInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwarePhysicalNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwarePhysicalNetworkInterface"]
**Identity** | Pointer to **string** | The internally generated identity of physical network interface. This entity cannot manipulated by users. It aids in uniquely identifying the physical network interface object. For VMware, this is MOR (managed object reference). | [optional] 
**Name** | Pointer to **string** | User-provided name to identify the physical network interface. | [optional] 

## Methods

### NewVirtualizationBasePhysicalNetworkInterfaceAllOf

`func NewVirtualizationBasePhysicalNetworkInterfaceAllOf(classId string, objectType string, ) *VirtualizationBasePhysicalNetworkInterfaceAllOf`

NewVirtualizationBasePhysicalNetworkInterfaceAllOf instantiates a new VirtualizationBasePhysicalNetworkInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBasePhysicalNetworkInterfaceAllOfWithDefaults

`func NewVirtualizationBasePhysicalNetworkInterfaceAllOfWithDefaults() *VirtualizationBasePhysicalNetworkInterfaceAllOf`

NewVirtualizationBasePhysicalNetworkInterfaceAllOfWithDefaults instantiates a new VirtualizationBasePhysicalNetworkInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBasePhysicalNetworkInterfaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


