# AssetConnectionControlMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | The account id to which the device belongs. | [optional] 
**ConnectorVersion** | Pointer to **string** | The version of the device connector currently running on the platform. Deprecated by newer connectors that will report this directly to the device connector gateway in a websocket header, but included to continue to support older versions which report any version change after connect. | [optional] 
**DeviceId** | Pointer to **string** | The Moid of the device under change. Used to route the message to a devices connection. | [optional] 
**DomainGroup** | Pointer to **string** | The domain group id to which the device belongs. | [optional] 
**Evict** | Pointer to **bool** | Flag to force any open connections to be evicted. Used in case device has been deleted or blacklisted. | [optional] 
**Leadership** | Pointer to **string** | The current leadership of a device cluster member. * &#x60;Unknown&#x60; - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight. * &#x60;Primary&#x60; - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary in a given cluster, while the underlying platform may be active-active only one connector will assume the primary role. * &#x60;Secondary&#x60; - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. e.g. file upload will be enabled, but requests to the underlying platform management will be disabled. | [optional] [default to "Unknown"]
**NewIdentity** | Pointer to **string** | The new identity assigned to a device on ownership change (claim/unclaim). | [optional] 
**Partition** | Pointer to **int64** | The partition the device was last connected to, used to address the control message to the device connector gateway instance holding the devices connection. | [optional] 

## Methods

### NewAssetConnectionControlMessageAllOf

`func NewAssetConnectionControlMessageAllOf() *AssetConnectionControlMessageAllOf`

NewAssetConnectionControlMessageAllOf instantiates a new AssetConnectionControlMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetConnectionControlMessageAllOfWithDefaults

`func NewAssetConnectionControlMessageAllOfWithDefaults() *AssetConnectionControlMessageAllOf`

NewAssetConnectionControlMessageAllOfWithDefaults instantiates a new AssetConnectionControlMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AssetConnectionControlMessageAllOf) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetConnectionControlMessageAllOf) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetConnectionControlMessageAllOf) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetConnectionControlMessageAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConnectorVersion

`func (o *AssetConnectionControlMessageAllOf) GetConnectorVersion() string`

GetConnectorVersion returns the ConnectorVersion field if non-nil, zero value otherwise.

### GetConnectorVersionOk

`func (o *AssetConnectionControlMessageAllOf) GetConnectorVersionOk() (*string, bool)`

GetConnectorVersionOk returns a tuple with the ConnectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorVersion

`func (o *AssetConnectionControlMessageAllOf) SetConnectorVersion(v string)`

SetConnectorVersion sets ConnectorVersion field to given value.

### HasConnectorVersion

`func (o *AssetConnectionControlMessageAllOf) HasConnectorVersion() bool`

HasConnectorVersion returns a boolean if a field has been set.

### GetDeviceId

`func (o *AssetConnectionControlMessageAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetConnectionControlMessageAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetConnectionControlMessageAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetConnectionControlMessageAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDomainGroup

`func (o *AssetConnectionControlMessageAllOf) GetDomainGroup() string`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *AssetConnectionControlMessageAllOf) GetDomainGroupOk() (*string, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *AssetConnectionControlMessageAllOf) SetDomainGroup(v string)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *AssetConnectionControlMessageAllOf) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### GetEvict

`func (o *AssetConnectionControlMessageAllOf) GetEvict() bool`

GetEvict returns the Evict field if non-nil, zero value otherwise.

### GetEvictOk

`func (o *AssetConnectionControlMessageAllOf) GetEvictOk() (*bool, bool)`

GetEvictOk returns a tuple with the Evict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvict

`func (o *AssetConnectionControlMessageAllOf) SetEvict(v bool)`

SetEvict sets Evict field to given value.

### HasEvict

`func (o *AssetConnectionControlMessageAllOf) HasEvict() bool`

HasEvict returns a boolean if a field has been set.

### GetLeadership

`func (o *AssetConnectionControlMessageAllOf) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *AssetConnectionControlMessageAllOf) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *AssetConnectionControlMessageAllOf) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *AssetConnectionControlMessageAllOf) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetNewIdentity

`func (o *AssetConnectionControlMessageAllOf) GetNewIdentity() string`

GetNewIdentity returns the NewIdentity field if non-nil, zero value otherwise.

### GetNewIdentityOk

`func (o *AssetConnectionControlMessageAllOf) GetNewIdentityOk() (*string, bool)`

GetNewIdentityOk returns a tuple with the NewIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIdentity

`func (o *AssetConnectionControlMessageAllOf) SetNewIdentity(v string)`

SetNewIdentity sets NewIdentity field to given value.

### HasNewIdentity

`func (o *AssetConnectionControlMessageAllOf) HasNewIdentity() bool`

HasNewIdentity returns a boolean if a field has been set.

### GetPartition

`func (o *AssetConnectionControlMessageAllOf) GetPartition() int64`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *AssetConnectionControlMessageAllOf) GetPartitionOk() (*int64, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *AssetConnectionControlMessageAllOf) SetPartition(v int64)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *AssetConnectionControlMessageAllOf) HasPartition() bool`

HasPartition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


