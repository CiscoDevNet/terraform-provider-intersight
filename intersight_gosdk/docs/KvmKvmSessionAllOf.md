# KvmKvmSessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.KvmSession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.KvmSession"]
**KvmMuxUrl** | Pointer to **string** | The URL of the KVM MUX to connect to. | [optional] [readonly] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 

## Methods

### NewKvmKvmSessionAllOf

`func NewKvmKvmSessionAllOf(classId string, objectType string, ) *KvmKvmSessionAllOf`

NewKvmKvmSessionAllOf instantiates a new KvmKvmSessionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmKvmSessionAllOfWithDefaults

`func NewKvmKvmSessionAllOfWithDefaults() *KvmKvmSessionAllOf`

NewKvmKvmSessionAllOfWithDefaults instantiates a new KvmKvmSessionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmKvmSessionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmKvmSessionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmKvmSessionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmKvmSessionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmKvmSessionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmKvmSessionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKvmMuxUrl

`func (o *KvmKvmSessionAllOf) GetKvmMuxUrl() string`

GetKvmMuxUrl returns the KvmMuxUrl field if non-nil, zero value otherwise.

### GetKvmMuxUrlOk

`func (o *KvmKvmSessionAllOf) GetKvmMuxUrlOk() (*string, bool)`

GetKvmMuxUrlOk returns a tuple with the KvmMuxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmMuxUrl

`func (o *KvmKvmSessionAllOf) SetKvmMuxUrl(v string)`

SetKvmMuxUrl sets KvmMuxUrl field to given value.

### HasKvmMuxUrl

`func (o *KvmKvmSessionAllOf) HasKvmMuxUrl() bool`

HasKvmMuxUrl returns a boolean if a field has been set.

### GetDevice

`func (o *KvmKvmSessionAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *KvmKvmSessionAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *KvmKvmSessionAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *KvmKvmSessionAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *KvmKvmSessionAllOf) GetServer() ComputeRackUnitRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KvmKvmSessionAllOf) GetServerOk() (*ComputeRackUnitRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KvmKvmSessionAllOf) SetServer(v ComputeRackUnitRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *KvmKvmSessionAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


