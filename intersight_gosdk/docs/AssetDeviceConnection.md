# AssetDeviceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **int64** | The version of the connector API, describes the capability of the connector&#39;s framework. If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded. | [optional] [readonly] 
**AppPartitionNumber** | Pointer to **int64** | The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter. | [optional] [readonly] 
**ConnectionReason** | Pointer to **string** | If &#39;connectionStatus&#39; is not equal to Connected, connectionReason provides further details about why the device is not connected with Intersight. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | The status of the persistent connection between the device connector and Intersight. * &#x60;&#x60; - The device registered with Intersight but subsequently did not establish a persistent websocket connection. * &#x60;Connected&#x60; - The device&#39;s connection to Intersight has been established and is active. * &#x60;NotConnected&#x60; - The device&#39;s connection to Intersight has been disconnected. * &#x60;ClaimInProgress&#x60; - Claim of the device is in progress. * &#x60;Unclaimed&#x60; - The device was un-claimed from the users account by an Administrator of the device. | [optional] [readonly] [default to ""]
**ConnectionStatusLastChangeTime** | Pointer to [**time.Time**](time.Time.md) | The last time at which the &#39;connectionStatus&#39; property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between Intersight and Device Connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with Intersight. | [optional] [readonly] 
**ConnectorVersion** | Pointer to **string** | The version of the device connector running on the managed device. | [optional] [readonly] 
**DeviceExternalIpAddress** | Pointer to **string** | The IP Address of the managed device as seen from Intersight at the time of registration. This could be the IP address of the managed device&#39;s interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network. | [optional] [readonly] 
**ProxyApp** | Pointer to **string** | The name of the app which will proxy the messages to the device connector. | [optional] [readonly] 

## Methods

### NewAssetDeviceConnection

`func NewAssetDeviceConnection() *AssetDeviceConnection`

NewAssetDeviceConnection instantiates a new AssetDeviceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceConnectionWithDefaults

`func NewAssetDeviceConnectionWithDefaults() *AssetDeviceConnection`

NewAssetDeviceConnectionWithDefaults instantiates a new AssetDeviceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AssetDeviceConnection) GetApiVersion() int64`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AssetDeviceConnection) GetApiVersionOk() (*int64, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AssetDeviceConnection) SetApiVersion(v int64)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AssetDeviceConnection) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAppPartitionNumber

`func (o *AssetDeviceConnection) GetAppPartitionNumber() int64`

GetAppPartitionNumber returns the AppPartitionNumber field if non-nil, zero value otherwise.

### GetAppPartitionNumberOk

`func (o *AssetDeviceConnection) GetAppPartitionNumberOk() (*int64, bool)`

GetAppPartitionNumberOk returns a tuple with the AppPartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPartitionNumber

`func (o *AssetDeviceConnection) SetAppPartitionNumber(v int64)`

SetAppPartitionNumber sets AppPartitionNumber field to given value.

### HasAppPartitionNumber

`func (o *AssetDeviceConnection) HasAppPartitionNumber() bool`

HasAppPartitionNumber returns a boolean if a field has been set.

### GetConnectionId

`func (o *AssetDeviceConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AssetDeviceConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AssetDeviceConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AssetDeviceConnection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionReason

`func (o *AssetDeviceConnection) GetConnectionReason() string`

GetConnectionReason returns the ConnectionReason field if non-nil, zero value otherwise.

### GetConnectionReasonOk

`func (o *AssetDeviceConnection) GetConnectionReasonOk() (*string, bool)`

GetConnectionReasonOk returns a tuple with the ConnectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionReason

`func (o *AssetDeviceConnection) SetConnectionReason(v string)`

SetConnectionReason sets ConnectionReason field to given value.

### HasConnectionReason

`func (o *AssetDeviceConnection) HasConnectionReason() bool`

HasConnectionReason returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AssetDeviceConnection) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AssetDeviceConnection) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AssetDeviceConnection) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AssetDeviceConnection) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetConnectionStatusLastChangeTime

`func (o *AssetDeviceConnection) GetConnectionStatusLastChangeTime() time.Time`

GetConnectionStatusLastChangeTime returns the ConnectionStatusLastChangeTime field if non-nil, zero value otherwise.

### GetConnectionStatusLastChangeTimeOk

`func (o *AssetDeviceConnection) GetConnectionStatusLastChangeTimeOk() (*time.Time, bool)`

GetConnectionStatusLastChangeTimeOk returns a tuple with the ConnectionStatusLastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatusLastChangeTime

`func (o *AssetDeviceConnection) SetConnectionStatusLastChangeTime(v time.Time)`

SetConnectionStatusLastChangeTime sets ConnectionStatusLastChangeTime field to given value.

### HasConnectionStatusLastChangeTime

`func (o *AssetDeviceConnection) HasConnectionStatusLastChangeTime() bool`

HasConnectionStatusLastChangeTime returns a boolean if a field has been set.

### GetConnectorVersion

`func (o *AssetDeviceConnection) GetConnectorVersion() string`

GetConnectorVersion returns the ConnectorVersion field if non-nil, zero value otherwise.

### GetConnectorVersionOk

`func (o *AssetDeviceConnection) GetConnectorVersionOk() (*string, bool)`

GetConnectorVersionOk returns a tuple with the ConnectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorVersion

`func (o *AssetDeviceConnection) SetConnectorVersion(v string)`

SetConnectorVersion sets ConnectorVersion field to given value.

### HasConnectorVersion

`func (o *AssetDeviceConnection) HasConnectorVersion() bool`

HasConnectorVersion returns a boolean if a field has been set.

### GetDeviceExternalIpAddress

`func (o *AssetDeviceConnection) GetDeviceExternalIpAddress() string`

GetDeviceExternalIpAddress returns the DeviceExternalIpAddress field if non-nil, zero value otherwise.

### GetDeviceExternalIpAddressOk

`func (o *AssetDeviceConnection) GetDeviceExternalIpAddressOk() (*string, bool)`

GetDeviceExternalIpAddressOk returns a tuple with the DeviceExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceExternalIpAddress

`func (o *AssetDeviceConnection) SetDeviceExternalIpAddress(v string)`

SetDeviceExternalIpAddress sets DeviceExternalIpAddress field to given value.

### HasDeviceExternalIpAddress

`func (o *AssetDeviceConnection) HasDeviceExternalIpAddress() bool`

HasDeviceExternalIpAddress returns a boolean if a field has been set.

### GetProxyApp

`func (o *AssetDeviceConnection) GetProxyApp() string`

GetProxyApp returns the ProxyApp field if non-nil, zero value otherwise.

### GetProxyAppOk

`func (o *AssetDeviceConnection) GetProxyAppOk() (*string, bool)`

GetProxyAppOk returns a tuple with the ProxyApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyApp

`func (o *AssetDeviceConnection) SetProxyApp(v string)`

SetProxyApp sets ProxyApp field to given value.

### HasProxyApp

`func (o *AssetDeviceConnection) HasProxyApp() bool`

HasProxyApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


