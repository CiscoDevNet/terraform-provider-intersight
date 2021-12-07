# VirtualizationBaseNetworkPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.NetworkPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.NetworkPort"]
**Name** | Pointer to **string** | The name of the network port. | [optional] 

## Methods

### NewVirtualizationBaseNetworkPort

`func NewVirtualizationBaseNetworkPort(classId string, objectType string, ) *VirtualizationBaseNetworkPort`

NewVirtualizationBaseNetworkPort instantiates a new VirtualizationBaseNetworkPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseNetworkPortWithDefaults

`func NewVirtualizationBaseNetworkPortWithDefaults() *VirtualizationBaseNetworkPort`

NewVirtualizationBaseNetworkPortWithDefaults instantiates a new VirtualizationBaseNetworkPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseNetworkPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseNetworkPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseNetworkPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseNetworkPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseNetworkPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseNetworkPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *VirtualizationBaseNetworkPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseNetworkPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseNetworkPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseNetworkPort) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


