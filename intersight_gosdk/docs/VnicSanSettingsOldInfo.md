# VnicSanSettingsOldInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.SanSettingsOldInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.SanSettingsOldInfo"]
**Wwnn** | Pointer to **string** | The WWNN address that is assigned to the server node based on the wwn pool that has been assigned in the SAN Connectivity Policy. | [optional] [readonly] 
**WwnnAddressType** | Pointer to **string** | Type of allocation selected to assign a WWNN address for the server node. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [readonly] [default to "POOL"]
**WwnnLease** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**WwnnPool** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewVnicSanSettingsOldInfo

`func NewVnicSanSettingsOldInfo(classId string, objectType string, ) *VnicSanSettingsOldInfo`

NewVnicSanSettingsOldInfo instantiates a new VnicSanSettingsOldInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicSanSettingsOldInfoWithDefaults

`func NewVnicSanSettingsOldInfoWithDefaults() *VnicSanSettingsOldInfo`

NewVnicSanSettingsOldInfoWithDefaults instantiates a new VnicSanSettingsOldInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicSanSettingsOldInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicSanSettingsOldInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicSanSettingsOldInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicSanSettingsOldInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicSanSettingsOldInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicSanSettingsOldInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetWwnn

`func (o *VnicSanSettingsOldInfo) GetWwnn() string`

GetWwnn returns the Wwnn field if non-nil, zero value otherwise.

### GetWwnnOk

`func (o *VnicSanSettingsOldInfo) GetWwnnOk() (*string, bool)`

GetWwnnOk returns a tuple with the Wwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnn

`func (o *VnicSanSettingsOldInfo) SetWwnn(v string)`

SetWwnn sets Wwnn field to given value.

### HasWwnn

`func (o *VnicSanSettingsOldInfo) HasWwnn() bool`

HasWwnn returns a boolean if a field has been set.

### GetWwnnAddressType

`func (o *VnicSanSettingsOldInfo) GetWwnnAddressType() string`

GetWwnnAddressType returns the WwnnAddressType field if non-nil, zero value otherwise.

### GetWwnnAddressTypeOk

`func (o *VnicSanSettingsOldInfo) GetWwnnAddressTypeOk() (*string, bool)`

GetWwnnAddressTypeOk returns a tuple with the WwnnAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnnAddressType

`func (o *VnicSanSettingsOldInfo) SetWwnnAddressType(v string)`

SetWwnnAddressType sets WwnnAddressType field to given value.

### HasWwnnAddressType

`func (o *VnicSanSettingsOldInfo) HasWwnnAddressType() bool`

HasWwnnAddressType returns a boolean if a field has been set.

### GetWwnnLease

`func (o *VnicSanSettingsOldInfo) GetWwnnLease() MoMoRef`

GetWwnnLease returns the WwnnLease field if non-nil, zero value otherwise.

### GetWwnnLeaseOk

`func (o *VnicSanSettingsOldInfo) GetWwnnLeaseOk() (*MoMoRef, bool)`

GetWwnnLeaseOk returns a tuple with the WwnnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnnLease

`func (o *VnicSanSettingsOldInfo) SetWwnnLease(v MoMoRef)`

SetWwnnLease sets WwnnLease field to given value.

### HasWwnnLease

`func (o *VnicSanSettingsOldInfo) HasWwnnLease() bool`

HasWwnnLease returns a boolean if a field has been set.

### GetWwnnPool

`func (o *VnicSanSettingsOldInfo) GetWwnnPool() MoMoRef`

GetWwnnPool returns the WwnnPool field if non-nil, zero value otherwise.

### GetWwnnPoolOk

`func (o *VnicSanSettingsOldInfo) GetWwnnPoolOk() (*MoMoRef, bool)`

GetWwnnPoolOk returns a tuple with the WwnnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnnPool

`func (o *VnicSanSettingsOldInfo) SetWwnnPool(v MoMoRef)`

SetWwnnPool sets WwnnPool field to given value.

### HasWwnnPool

`func (o *VnicSanSettingsOldInfo) HasWwnnPool() bool`

HasWwnnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


