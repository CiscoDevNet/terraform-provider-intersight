# NetworkSecureRouterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.SecureRouterInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.SecureRouterInfo"]
**ConfigMismatch** | Pointer to **bool** | Indicates if there is a configuration mismatch between the detected and configured slots for secure routers. | [optional] [readonly] 
**ConfiguredInSlots** | Pointer to **string** | The slots in which secure router roles are configured. | [optional] [readonly] 
**DetectedInSlots** | Pointer to **string** | The slots in which secure routers are detected. | [optional] [readonly] 
**DiscoveredInSlots** | Pointer to **string** | The slots in which secure routers are discovered. | [optional] [readonly] 

## Methods

### NewNetworkSecureRouterInfo

`func NewNetworkSecureRouterInfo(classId string, objectType string, ) *NetworkSecureRouterInfo`

NewNetworkSecureRouterInfo instantiates a new NetworkSecureRouterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecureRouterInfoWithDefaults

`func NewNetworkSecureRouterInfoWithDefaults() *NetworkSecureRouterInfo`

NewNetworkSecureRouterInfoWithDefaults instantiates a new NetworkSecureRouterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkSecureRouterInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkSecureRouterInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkSecureRouterInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkSecureRouterInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkSecureRouterInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkSecureRouterInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigMismatch

`func (o *NetworkSecureRouterInfo) GetConfigMismatch() bool`

GetConfigMismatch returns the ConfigMismatch field if non-nil, zero value otherwise.

### GetConfigMismatchOk

`func (o *NetworkSecureRouterInfo) GetConfigMismatchOk() (*bool, bool)`

GetConfigMismatchOk returns a tuple with the ConfigMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMismatch

`func (o *NetworkSecureRouterInfo) SetConfigMismatch(v bool)`

SetConfigMismatch sets ConfigMismatch field to given value.

### HasConfigMismatch

`func (o *NetworkSecureRouterInfo) HasConfigMismatch() bool`

HasConfigMismatch returns a boolean if a field has been set.

### GetConfiguredInSlots

`func (o *NetworkSecureRouterInfo) GetConfiguredInSlots() string`

GetConfiguredInSlots returns the ConfiguredInSlots field if non-nil, zero value otherwise.

### GetConfiguredInSlotsOk

`func (o *NetworkSecureRouterInfo) GetConfiguredInSlotsOk() (*string, bool)`

GetConfiguredInSlotsOk returns a tuple with the ConfiguredInSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredInSlots

`func (o *NetworkSecureRouterInfo) SetConfiguredInSlots(v string)`

SetConfiguredInSlots sets ConfiguredInSlots field to given value.

### HasConfiguredInSlots

`func (o *NetworkSecureRouterInfo) HasConfiguredInSlots() bool`

HasConfiguredInSlots returns a boolean if a field has been set.

### GetDetectedInSlots

`func (o *NetworkSecureRouterInfo) GetDetectedInSlots() string`

GetDetectedInSlots returns the DetectedInSlots field if non-nil, zero value otherwise.

### GetDetectedInSlotsOk

`func (o *NetworkSecureRouterInfo) GetDetectedInSlotsOk() (*string, bool)`

GetDetectedInSlotsOk returns a tuple with the DetectedInSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedInSlots

`func (o *NetworkSecureRouterInfo) SetDetectedInSlots(v string)`

SetDetectedInSlots sets DetectedInSlots field to given value.

### HasDetectedInSlots

`func (o *NetworkSecureRouterInfo) HasDetectedInSlots() bool`

HasDetectedInSlots returns a boolean if a field has been set.

### GetDiscoveredInSlots

`func (o *NetworkSecureRouterInfo) GetDiscoveredInSlots() string`

GetDiscoveredInSlots returns the DiscoveredInSlots field if non-nil, zero value otherwise.

### GetDiscoveredInSlotsOk

`func (o *NetworkSecureRouterInfo) GetDiscoveredInSlotsOk() (*string, bool)`

GetDiscoveredInSlotsOk returns a tuple with the DiscoveredInSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredInSlots

`func (o *NetworkSecureRouterInfo) SetDiscoveredInSlots(v string)`

SetDiscoveredInSlots sets DiscoveredInSlots field to given value.

### HasDiscoveredInSlots

`func (o *NetworkSecureRouterInfo) HasDiscoveredInSlots() bool`

HasDiscoveredInSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


