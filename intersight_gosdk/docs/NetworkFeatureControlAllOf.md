# NetworkFeatureControlAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.FeatureControl"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.FeatureControl"]
**AdminState** | Pointer to **string** | The admin state of the feature. | [optional] [readonly] 
**Instance** | Pointer to **int64** | The number of instances of the feature. | [optional] [readonly] 
**Name** | Pointer to **string** | The name to identify the feature. | [optional] [readonly] 
**OperationalState** | Pointer to **string** | The operational state of the feature. | [optional] [readonly] 
**StatusMsg** | Pointer to **string** | The status message to capture admin state detailed information. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkFeatureControlAllOf

`func NewNetworkFeatureControlAllOf(classId string, objectType string, ) *NetworkFeatureControlAllOf`

NewNetworkFeatureControlAllOf instantiates a new NetworkFeatureControlAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeatureControlAllOfWithDefaults

`func NewNetworkFeatureControlAllOfWithDefaults() *NetworkFeatureControlAllOf`

NewNetworkFeatureControlAllOfWithDefaults instantiates a new NetworkFeatureControlAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkFeatureControlAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkFeatureControlAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkFeatureControlAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkFeatureControlAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkFeatureControlAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkFeatureControlAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *NetworkFeatureControlAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NetworkFeatureControlAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NetworkFeatureControlAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NetworkFeatureControlAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetInstance

`func (o *NetworkFeatureControlAllOf) GetInstance() int64`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *NetworkFeatureControlAllOf) GetInstanceOk() (*int64, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *NetworkFeatureControlAllOf) SetInstance(v int64)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *NetworkFeatureControlAllOf) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *NetworkFeatureControlAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFeatureControlAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFeatureControlAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkFeatureControlAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationalState

`func (o *NetworkFeatureControlAllOf) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *NetworkFeatureControlAllOf) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *NetworkFeatureControlAllOf) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *NetworkFeatureControlAllOf) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetStatusMsg

`func (o *NetworkFeatureControlAllOf) GetStatusMsg() string`

GetStatusMsg returns the StatusMsg field if non-nil, zero value otherwise.

### GetStatusMsgOk

`func (o *NetworkFeatureControlAllOf) GetStatusMsgOk() (*string, bool)`

GetStatusMsgOk returns a tuple with the StatusMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMsg

`func (o *NetworkFeatureControlAllOf) SetStatusMsg(v string)`

SetStatusMsg sets StatusMsg field to given value.

### HasStatusMsg

`func (o *NetworkFeatureControlAllOf) HasStatusMsg() bool`

HasStatusMsg returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkFeatureControlAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkFeatureControlAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkFeatureControlAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkFeatureControlAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkFeatureControlAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkFeatureControlAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkFeatureControlAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkFeatureControlAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


