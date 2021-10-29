# StorageNetAppHighAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppHighAvailability"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppHighAvailability"]
**Enabled** | Pointer to **bool** | Specifies whether or not storage failover is enabled. | [optional] [readonly] 
**GivebackState** | Pointer to **string** | The state of the node that is giving storage back to its HA partner. * &#x60;unknown&#x60; - Default unknown giveback state. * &#x60;nothing_to_giveback&#x60; - The node has nothing to give back to its HA partner. * &#x60;not_attempted&#x60; - The node has not attempted to give back storage to its HA partner. * &#x60;in_progress&#x60; - The node is in progress of giving back storage to its HA partner. * &#x60;failed&#x60; - The node has failed to give back storage to its HA partner. | [optional] [readonly] [default to "unknown"]
**PartnerName** | Pointer to **string** | The partner node name in this node&#39;s High Availability (HA) group. | [optional] [readonly] 
**PartnerUuid** | Pointer to **string** | The partner node uuid in this node&#39;s High Availability (HA) group. | [optional] [readonly] 
**TakeoverState** | Pointer to **string** | The state of the node that is taking over storage from its HA partner. * &#x60;unknown&#x60; - Default unknown takeover state. * &#x60;not_possible&#x60; - It is not possible for the node to take over storage from its HA partner. * &#x60;not_attempted&#x60; - The node has not attempted to take over storage from its HA partner. * &#x60;in_takeover&#x60; - The node has taken over storage from its HA partner. * &#x60;in_progress&#x60; - The node is in progress of taking over storage from its HA partner. * &#x60;failed&#x60; - The node has failed to take over storage from its HA partner. | [optional] [readonly] [default to "unknown"]

## Methods

### NewStorageNetAppHighAvailability

`func NewStorageNetAppHighAvailability(classId string, objectType string, ) *StorageNetAppHighAvailability`

NewStorageNetAppHighAvailability instantiates a new StorageNetAppHighAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppHighAvailabilityWithDefaults

`func NewStorageNetAppHighAvailabilityWithDefaults() *StorageNetAppHighAvailability`

NewStorageNetAppHighAvailabilityWithDefaults instantiates a new StorageNetAppHighAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppHighAvailability) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppHighAvailability) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppHighAvailability) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppHighAvailability) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppHighAvailability) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppHighAvailability) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *StorageNetAppHighAvailability) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppHighAvailability) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppHighAvailability) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppHighAvailability) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGivebackState

`func (o *StorageNetAppHighAvailability) GetGivebackState() string`

GetGivebackState returns the GivebackState field if non-nil, zero value otherwise.

### GetGivebackStateOk

`func (o *StorageNetAppHighAvailability) GetGivebackStateOk() (*string, bool)`

GetGivebackStateOk returns a tuple with the GivebackState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivebackState

`func (o *StorageNetAppHighAvailability) SetGivebackState(v string)`

SetGivebackState sets GivebackState field to given value.

### HasGivebackState

`func (o *StorageNetAppHighAvailability) HasGivebackState() bool`

HasGivebackState returns a boolean if a field has been set.

### GetPartnerName

`func (o *StorageNetAppHighAvailability) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *StorageNetAppHighAvailability) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *StorageNetAppHighAvailability) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *StorageNetAppHighAvailability) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### GetPartnerUuid

`func (o *StorageNetAppHighAvailability) GetPartnerUuid() string`

GetPartnerUuid returns the PartnerUuid field if non-nil, zero value otherwise.

### GetPartnerUuidOk

`func (o *StorageNetAppHighAvailability) GetPartnerUuidOk() (*string, bool)`

GetPartnerUuidOk returns a tuple with the PartnerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerUuid

`func (o *StorageNetAppHighAvailability) SetPartnerUuid(v string)`

SetPartnerUuid sets PartnerUuid field to given value.

### HasPartnerUuid

`func (o *StorageNetAppHighAvailability) HasPartnerUuid() bool`

HasPartnerUuid returns a boolean if a field has been set.

### GetTakeoverState

`func (o *StorageNetAppHighAvailability) GetTakeoverState() string`

GetTakeoverState returns the TakeoverState field if non-nil, zero value otherwise.

### GetTakeoverStateOk

`func (o *StorageNetAppHighAvailability) GetTakeoverStateOk() (*string, bool)`

GetTakeoverStateOk returns a tuple with the TakeoverState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeoverState

`func (o *StorageNetAppHighAvailability) SetTakeoverState(v string)`

SetTakeoverState sets TakeoverState field to given value.

### HasTakeoverState

`func (o *StorageNetAppHighAvailability) HasTakeoverState() bool`

HasTakeoverState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


