# NiatelemetryNexusDashboardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusDashboardDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusDashboardDetails"]
**ClusterName** | Pointer to **string** | Name of the nexus dashboard cluster. | [optional] 
**DeviceModel** | Pointer to **string** | Model of the nexus dashboard cluster. | [optional] 
**NexusDashboardName** | Pointer to **string** | Name of the NexusDashboard. | [optional] 
**NexusDashboardSerialNumber** | Pointer to **string** | Serial number of NexusDashboard. | [optional] 
**Type** | Pointer to **string** | Node type of the nexus dashboard cluster. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusDashboardDetails

`func NewNiatelemetryNexusDashboardDetails(classId string, objectType string, ) *NiatelemetryNexusDashboardDetails`

NewNiatelemetryNexusDashboardDetails instantiates a new NiatelemetryNexusDashboardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusDashboardDetailsWithDefaults

`func NewNiatelemetryNexusDashboardDetailsWithDefaults() *NiatelemetryNexusDashboardDetails`

NewNiatelemetryNexusDashboardDetailsWithDefaults instantiates a new NiatelemetryNexusDashboardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusDashboardDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusDashboardDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusDashboardDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusDashboardDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusDashboardDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusDashboardDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterName

`func (o *NiatelemetryNexusDashboardDetails) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *NiatelemetryNexusDashboardDetails) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *NiatelemetryNexusDashboardDetails) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *NiatelemetryNexusDashboardDetails) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDeviceModel

`func (o *NiatelemetryNexusDashboardDetails) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *NiatelemetryNexusDashboardDetails) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *NiatelemetryNexusDashboardDetails) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *NiatelemetryNexusDashboardDetails) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetNexusDashboardName

`func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardName() string`

GetNexusDashboardName returns the NexusDashboardName field if non-nil, zero value otherwise.

### GetNexusDashboardNameOk

`func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardNameOk() (*string, bool)`

GetNexusDashboardNameOk returns a tuple with the NexusDashboardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusDashboardName

`func (o *NiatelemetryNexusDashboardDetails) SetNexusDashboardName(v string)`

SetNexusDashboardName sets NexusDashboardName field to given value.

### HasNexusDashboardName

`func (o *NiatelemetryNexusDashboardDetails) HasNexusDashboardName() bool`

HasNexusDashboardName returns a boolean if a field has been set.

### GetNexusDashboardSerialNumber

`func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardSerialNumber() string`

GetNexusDashboardSerialNumber returns the NexusDashboardSerialNumber field if non-nil, zero value otherwise.

### GetNexusDashboardSerialNumberOk

`func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardSerialNumberOk() (*string, bool)`

GetNexusDashboardSerialNumberOk returns a tuple with the NexusDashboardSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusDashboardSerialNumber

`func (o *NiatelemetryNexusDashboardDetails) SetNexusDashboardSerialNumber(v string)`

SetNexusDashboardSerialNumber sets NexusDashboardSerialNumber field to given value.

### HasNexusDashboardSerialNumber

`func (o *NiatelemetryNexusDashboardDetails) HasNexusDashboardSerialNumber() bool`

HasNexusDashboardSerialNumber returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetryNexusDashboardDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetryNexusDashboardDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetryNexusDashboardDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetryNexusDashboardDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNexusDashboardDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNexusDashboardDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNexusDashboardDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNexusDashboardDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


