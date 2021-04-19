# VirtualizationBaseNetworkPortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "hyperflex.NetworkPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "hyperflex.NetworkPort"]
**Name** | Pointer to **string** | The name of the network port. | [optional] 

## Methods

### NewVirtualizationBaseNetworkPortAllOf

`func NewVirtualizationBaseNetworkPortAllOf(classId string, objectType string, ) *VirtualizationBaseNetworkPortAllOf`

NewVirtualizationBaseNetworkPortAllOf instantiates a new VirtualizationBaseNetworkPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseNetworkPortAllOfWithDefaults

`func NewVirtualizationBaseNetworkPortAllOfWithDefaults() *VirtualizationBaseNetworkPortAllOf`

NewVirtualizationBaseNetworkPortAllOfWithDefaults instantiates a new VirtualizationBaseNetworkPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseNetworkPortAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseNetworkPortAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseNetworkPortAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseNetworkPortAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseNetworkPortAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseNetworkPortAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *VirtualizationBaseNetworkPortAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseNetworkPortAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseNetworkPortAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseNetworkPortAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


