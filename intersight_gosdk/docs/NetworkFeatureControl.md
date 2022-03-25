# NetworkFeatureControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.FeatureControl"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.FeatureControl"]
**AdminState** | Pointer to **string** | The admin state of the feature. | [optional] 
**Instance** | Pointer to **int64** | The number of instances of the feature. | [optional] 
**Name** | Pointer to **string** | The name to identify the feature. | [optional] 
**OperationalState** | Pointer to **string** | The operational state of the feature. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkFeatureControl

`func NewNetworkFeatureControl(classId string, objectType string, ) *NetworkFeatureControl`

NewNetworkFeatureControl instantiates a new NetworkFeatureControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeatureControlWithDefaults

`func NewNetworkFeatureControlWithDefaults() *NetworkFeatureControl`

NewNetworkFeatureControlWithDefaults instantiates a new NetworkFeatureControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkFeatureControl) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkFeatureControl) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkFeatureControl) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkFeatureControl) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkFeatureControl) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkFeatureControl) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *NetworkFeatureControl) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NetworkFeatureControl) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NetworkFeatureControl) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NetworkFeatureControl) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetInstance

`func (o *NetworkFeatureControl) GetInstance() int64`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *NetworkFeatureControl) GetInstanceOk() (*int64, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *NetworkFeatureControl) SetInstance(v int64)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *NetworkFeatureControl) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *NetworkFeatureControl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFeatureControl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFeatureControl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkFeatureControl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationalState

`func (o *NetworkFeatureControl) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *NetworkFeatureControl) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *NetworkFeatureControl) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *NetworkFeatureControl) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkFeatureControl) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkFeatureControl) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkFeatureControl) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkFeatureControl) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkFeatureControl) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkFeatureControl) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkFeatureControl) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkFeatureControl) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


