# KvmKvmSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.KvmSession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.KvmSession"]
**KvmMuxUrl** | Pointer to **string** | The URL of the KVM MUX to connect to. | [optional] [readonly] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 

## Methods

### NewKvmKvmSession

`func NewKvmKvmSession(classId string, objectType string, ) *KvmKvmSession`

NewKvmKvmSession instantiates a new KvmKvmSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmKvmSessionWithDefaults

`func NewKvmKvmSessionWithDefaults() *KvmKvmSession`

NewKvmKvmSessionWithDefaults instantiates a new KvmKvmSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmKvmSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmKvmSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmKvmSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmKvmSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmKvmSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmKvmSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKvmMuxUrl

`func (o *KvmKvmSession) GetKvmMuxUrl() string`

GetKvmMuxUrl returns the KvmMuxUrl field if non-nil, zero value otherwise.

### GetKvmMuxUrlOk

`func (o *KvmKvmSession) GetKvmMuxUrlOk() (*string, bool)`

GetKvmMuxUrlOk returns a tuple with the KvmMuxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmMuxUrl

`func (o *KvmKvmSession) SetKvmMuxUrl(v string)`

SetKvmMuxUrl sets KvmMuxUrl field to given value.

### HasKvmMuxUrl

`func (o *KvmKvmSession) HasKvmMuxUrl() bool`

HasKvmMuxUrl returns a boolean if a field has been set.

### GetDevice

`func (o *KvmKvmSession) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *KvmKvmSession) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *KvmKvmSession) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *KvmKvmSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *KvmKvmSession) GetServer() ComputeRackUnitRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KvmKvmSession) GetServerOk() (*ComputeRackUnitRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KvmKvmSession) SetServer(v ComputeRackUnitRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *KvmKvmSession) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


