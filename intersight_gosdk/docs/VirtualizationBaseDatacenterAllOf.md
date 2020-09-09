# VirtualizationBaseDatacenterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **string** | Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference). | [optional] [readonly] 
**Name** | Pointer to **string** | User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter. | [optional] 

## Methods

### NewVirtualizationBaseDatacenterAllOf

`func NewVirtualizationBaseDatacenterAllOf() *VirtualizationBaseDatacenterAllOf`

NewVirtualizationBaseDatacenterAllOf instantiates a new VirtualizationBaseDatacenterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseDatacenterAllOfWithDefaults

`func NewVirtualizationBaseDatacenterAllOfWithDefaults() *VirtualizationBaseDatacenterAllOf`

NewVirtualizationBaseDatacenterAllOfWithDefaults instantiates a new VirtualizationBaseDatacenterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *VirtualizationBaseDatacenterAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseDatacenterAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseDatacenterAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseDatacenterAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseDatacenterAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseDatacenterAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseDatacenterAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseDatacenterAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


