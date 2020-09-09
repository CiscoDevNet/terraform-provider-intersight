# AssetManagedDeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPort** | Pointer to **int64** | Port used for the connection to the Cloud by the Device Connector for the Managed Device. | [optional] 
**ConnectionFailureReason** | Pointer to **string** | Maintains the reason for the failure of connection to the Device in case of connection failure. | [optional] 
**ConnectionStatus** | Pointer to **string** | Maintains the status of the connection to the Device. * &#x60;Unknown&#x60; - This represents Connection Status as unknown. * &#x60;Success&#x60; - This represents Connection Status as successful. * &#x60;Failure&#x60; - This represents Connection Status as failure. | [optional] [default to "Unknown"]
**ErrorCode** | Pointer to **int64** | Maintains code related to error from Device Connector, if any. | [optional] 
**ErrorReason** | Pointer to **string** | Maintains the reason for the error from Device Connector, if any. | [optional] 
**ProcessId** | Pointer to **int64** | Maintains the process pid of the Device Connector for the Managed Device. | [optional] 
**ServerPort** | Pointer to **int64** | Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device. | [optional] 
**State** | Pointer to **string** | Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states. * &#x60;New&#x60; - State when the Managed Device is created. * &#x60;StartSent&#x60; - State when Create Managed Device is sent successfully to Intersight Assist to start Managed Device Device Connector. * &#x60;StartSentFailure&#x60; - State when Create Managed Device failed to send start Device Connector to Intersight Assist. * &#x60;StartSuccess&#x60; - State when Device Connector for the Managed Device has been successfully started. * &#x60;StartFailure&#x60; - State when Device Connector for the Managed Device failed to start. * &#x60;UpdateSentFailure&#x60; - State when Update Managed Device request is failed to be sent to Device Connector for the Managed Device. * &#x60;UpdateSent&#x60; - State when Update Managed Device request is successfully sent to Device Connector for the Managed Device. * &#x60;DeleteSentFailure&#x60; - State when Delete Managed Device request is failed to be sent to Intersight Assist. * &#x60;DeleteInProgress&#x60; - State when Intersight Assist received delete Managed Device request successfully. * &#x60;DeleteFailure&#x60; - State when Intersight Assist failed to stop Device Connector for the Managed Device. * &#x60;DeleteSuccess&#x60; - State when Intersight Assist stopped Device Connector for the Managed Device successfully. | [optional] [default to "New"]

## Methods

### NewAssetManagedDeviceStatus

`func NewAssetManagedDeviceStatus() *AssetManagedDeviceStatus`

NewAssetManagedDeviceStatus instantiates a new AssetManagedDeviceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetManagedDeviceStatusWithDefaults

`func NewAssetManagedDeviceStatusWithDefaults() *AssetManagedDeviceStatus`

NewAssetManagedDeviceStatusWithDefaults instantiates a new AssetManagedDeviceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudPort

`func (o *AssetManagedDeviceStatus) GetCloudPort() int64`

GetCloudPort returns the CloudPort field if non-nil, zero value otherwise.

### GetCloudPortOk

`func (o *AssetManagedDeviceStatus) GetCloudPortOk() (*int64, bool)`

GetCloudPortOk returns a tuple with the CloudPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPort

`func (o *AssetManagedDeviceStatus) SetCloudPort(v int64)`

SetCloudPort sets CloudPort field to given value.

### HasCloudPort

`func (o *AssetManagedDeviceStatus) HasCloudPort() bool`

HasCloudPort returns a boolean if a field has been set.

### GetConnectionFailureReason

`func (o *AssetManagedDeviceStatus) GetConnectionFailureReason() string`

GetConnectionFailureReason returns the ConnectionFailureReason field if non-nil, zero value otherwise.

### GetConnectionFailureReasonOk

`func (o *AssetManagedDeviceStatus) GetConnectionFailureReasonOk() (*string, bool)`

GetConnectionFailureReasonOk returns a tuple with the ConnectionFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionFailureReason

`func (o *AssetManagedDeviceStatus) SetConnectionFailureReason(v string)`

SetConnectionFailureReason sets ConnectionFailureReason field to given value.

### HasConnectionFailureReason

`func (o *AssetManagedDeviceStatus) HasConnectionFailureReason() bool`

HasConnectionFailureReason returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AssetManagedDeviceStatus) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AssetManagedDeviceStatus) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AssetManagedDeviceStatus) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AssetManagedDeviceStatus) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *AssetManagedDeviceStatus) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AssetManagedDeviceStatus) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AssetManagedDeviceStatus) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AssetManagedDeviceStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorReason

`func (o *AssetManagedDeviceStatus) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *AssetManagedDeviceStatus) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *AssetManagedDeviceStatus) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *AssetManagedDeviceStatus) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.

### GetProcessId

`func (o *AssetManagedDeviceStatus) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *AssetManagedDeviceStatus) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *AssetManagedDeviceStatus) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *AssetManagedDeviceStatus) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetServerPort

`func (o *AssetManagedDeviceStatus) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AssetManagedDeviceStatus) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AssetManagedDeviceStatus) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AssetManagedDeviceStatus) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetState

`func (o *AssetManagedDeviceStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AssetManagedDeviceStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AssetManagedDeviceStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AssetManagedDeviceStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


