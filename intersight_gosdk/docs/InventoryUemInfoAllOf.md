# InventoryUemInfoAllOf

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

### NewInventoryUemInfoAllOf

`func NewInventoryUemInfoAllOf(classId string, objectType string, ) *InventoryUemInfoAllOf`

NewInventoryUemInfoAllOf instantiates a new InventoryUemInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryUemInfoAllOfWithDefaults

`func NewInventoryUemInfoAllOfWithDefaults() *InventoryUemInfoAllOf`

NewInventoryUemInfoAllOfWithDefaults instantiates a new InventoryUemInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryUemInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryUemInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryUemInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryUemInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryUemInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryUemInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionStatus

`func (o *InventoryUemInfoAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *InventoryUemInfoAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *InventoryUemInfoAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *InventoryUemInfoAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetLastSequenceId

`func (o *InventoryUemInfoAllOf) GetLastSequenceId() int64`

GetLastSequenceId returns the LastSequenceId field if non-nil, zero value otherwise.

### GetLastSequenceIdOk

`func (o *InventoryUemInfoAllOf) GetLastSequenceIdOk() (*int64, bool)`

GetLastSequenceIdOk returns a tuple with the LastSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSequenceId

`func (o *InventoryUemInfoAllOf) SetLastSequenceId(v int64)`

SetLastSequenceId sets LastSequenceId field to given value.

### HasLastSequenceId

`func (o *InventoryUemInfoAllOf) HasLastSequenceId() bool`

HasLastSequenceId returns a boolean if a field has been set.

### GetPackageVersion

`func (o *InventoryUemInfoAllOf) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *InventoryUemInfoAllOf) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *InventoryUemInfoAllOf) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *InventoryUemInfoAllOf) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *InventoryUemInfoAllOf) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *InventoryUemInfoAllOf) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *InventoryUemInfoAllOf) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *InventoryUemInfoAllOf) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetSwitchId

`func (o *InventoryUemInfoAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *InventoryUemInfoAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *InventoryUemInfoAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *InventoryUemInfoAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


