# NetworkVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Vrf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Vrf"]
**AdminState** | Pointer to **string** | This field identifies the adminState of the given component. | [optional] 
**Name** | Pointer to **string** | This field identifies the name of the given component. | [optional] 
**VrfId** | Pointer to **string** | This field identifies the id of the given component. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVrf

`func NewNetworkVrf(classId string, objectType string, ) *NetworkVrf`

NewNetworkVrf instantiates a new NetworkVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVrfWithDefaults

`func NewNetworkVrfWithDefaults() *NetworkVrf`

NewNetworkVrfWithDefaults instantiates a new NetworkVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVrf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVrf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVrf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVrf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVrf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVrf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *NetworkVrf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NetworkVrf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NetworkVrf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NetworkVrf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetName

`func (o *NetworkVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkVrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkVrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVrfId

`func (o *NetworkVrf) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *NetworkVrf) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *NetworkVrf) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *NetworkVrf) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVrf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVrf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVrf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVrf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVrf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVrf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVrf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVrf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


