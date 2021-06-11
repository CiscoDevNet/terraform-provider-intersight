# InventoryUemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "inventory.UemInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "inventory.UemInfo"]
**ConnectionStatus** | Pointer to **string** | Connections status of Uem endpoint. * &#x60;Unknown&#x60; - Connection status of UEM is not known. * &#x60;Connected&#x60; - UEM endpoint has an open TCP connection with device connector. * &#x60;Enabled&#x60; - Eventing is enabled on the UEM endpoint. * &#x60;Paused&#x60; - Eventing is paused on the UEM endpoint. * &#x60;Disabled&#x60; - Eventing is disabled on the UEM endpoint. * &#x60;Disconnected&#x60; - There is no TCP connection between the UEM endpoint and device connector. | [optional] [readonly] [default to "Unknown"]
**LastSequenceId** | Pointer to **int64** | Last Sequence ID in the UEM event frame received from the endpoint identified by endpointInfo. | [optional] 
**PackageVersion** | Pointer to **string** | Version of UEM package on the endpoint. | [optional] [readonly] 
**ProtocolVersion** | Pointer to **string** | Version of UEM protocol running on the endpoint. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | The switch ID to identify the path to the endpoint. | [optional] [readonly] 

## Methods

### NewInventoryUemInfo

`func NewInventoryUemInfo(classId string, objectType string, ) *InventoryUemInfo`

NewInventoryUemInfo instantiates a new InventoryUemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryUemInfoWithDefaults

`func NewInventoryUemInfoWithDefaults() *InventoryUemInfo`

NewInventoryUemInfoWithDefaults instantiates a new InventoryUemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryUemInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryUemInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryUemInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryUemInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryUemInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryUemInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionStatus

`func (o *InventoryUemInfo) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *InventoryUemInfo) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *InventoryUemInfo) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *InventoryUemInfo) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetLastSequenceId

`func (o *InventoryUemInfo) GetLastSequenceId() int64`

GetLastSequenceId returns the LastSequenceId field if non-nil, zero value otherwise.

### GetLastSequenceIdOk

`func (o *InventoryUemInfo) GetLastSequenceIdOk() (*int64, bool)`

GetLastSequenceIdOk returns a tuple with the LastSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSequenceId

`func (o *InventoryUemInfo) SetLastSequenceId(v int64)`

SetLastSequenceId sets LastSequenceId field to given value.

### HasLastSequenceId

`func (o *InventoryUemInfo) HasLastSequenceId() bool`

HasLastSequenceId returns a boolean if a field has been set.

### GetPackageVersion

`func (o *InventoryUemInfo) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *InventoryUemInfo) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *InventoryUemInfo) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *InventoryUemInfo) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *InventoryUemInfo) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *InventoryUemInfo) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *InventoryUemInfo) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *InventoryUemInfo) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetSwitchId

`func (o *InventoryUemInfo) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *InventoryUemInfo) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *InventoryUemInfo) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *InventoryUemInfo) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


