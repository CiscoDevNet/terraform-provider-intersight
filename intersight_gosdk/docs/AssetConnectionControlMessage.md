# AssetConnectionControlMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ConnectionControlMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ConnectionControlMessage"]
**Account** | Pointer to **string** | The account id to which the device belongs. | [optional] 
**ConnectorVersion** | Pointer to **string** | The version of the device connector currently running on the platform. Deprecated by newer connectors that will report this directly to the device connector gateway in a websocket header, but included to continue to support older versions which report any version change after connect. | [optional] 
**DeviceId** | Pointer to **string** | The Moid of the device under change. Used to route the message to a device&#39;s connection. | [optional] 
**DomainGroup** | Pointer to **string** | The domain group id to which the device belongs. | [optional] [readonly] 
**Evict** | Pointer to **bool** | Flag to force any open connections to be evicted. Used in case device has been deleted or blacklisted. | [optional] 
**InternalConnectionId** | Pointer to **string** | Uniquely identifies a specific connection this control message is addressed to. Each connection to the gateway is associated with a unique connection id, which is used to identify the connection across any number of connection or connection attempts from the same device endpoint. When an evict message is published from device service to the gateway it may be tagged to a specific connection using this field. e.g. The device re-connects to Intersight before the previous connection has received a close or timeout, in which case we may send an evict specifically to the previous connection, with the new connection ignoring the message. If empty, the control message will be processed by any connections associated with the deviceId. | [optional] [readonly] 
**Leadership** | Pointer to **string** | The current leadership of a device cluster member. * &#x60;Unknown&#x60; - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials, it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight. * &#x60;Primary&#x60; - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary node in a given cluster, while the underlying platform may be active. If it is active, only one connector will assume the primary role. * &#x60;Secondary&#x60; - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. E.g. file upload will be enabled, but requests to the underlying platform management will be disabled. | [optional] [default to "Unknown"]
**NewIdentity** | Pointer to **string** | The new identity assigned to a device on ownership change (claim/unclaim). | [optional] 
**Partition** | Pointer to **int64** | The partition the device was last connected to, used to address the control message to the device connector gateway instance holding the devices connection. | [optional] 

## Methods

### NewAssetConnectionControlMessage

`func NewAssetConnectionControlMessage(classId string, objectType string, ) *AssetConnectionControlMessage`

NewAssetConnectionControlMessage instantiates a new AssetConnectionControlMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetConnectionControlMessageWithDefaults

`func NewAssetConnectionControlMessageWithDefaults() *AssetConnectionControlMessage`

NewAssetConnectionControlMessageWithDefaults instantiates a new AssetConnectionControlMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetConnectionControlMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetConnectionControlMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetConnectionControlMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetConnectionControlMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetConnectionControlMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetConnectionControlMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *AssetConnectionControlMessage) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetConnectionControlMessage) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetConnectionControlMessage) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetConnectionControlMessage) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConnectorVersion

`func (o *AssetConnectionControlMessage) GetConnectorVersion() string`

GetConnectorVersion returns the ConnectorVersion field if non-nil, zero value otherwise.

### GetConnectorVersionOk

`func (o *AssetConnectionControlMessage) GetConnectorVersionOk() (*string, bool)`

GetConnectorVersionOk returns a tuple with the ConnectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorVersion

`func (o *AssetConnectionControlMessage) SetConnectorVersion(v string)`

SetConnectorVersion sets ConnectorVersion field to given value.

### HasConnectorVersion

`func (o *AssetConnectionControlMessage) HasConnectorVersion() bool`

HasConnectorVersion returns a boolean if a field has been set.

### GetDeviceId

`func (o *AssetConnectionControlMessage) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetConnectionControlMessage) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetConnectionControlMessage) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetConnectionControlMessage) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDomainGroup

`func (o *AssetConnectionControlMessage) GetDomainGroup() string`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *AssetConnectionControlMessage) GetDomainGroupOk() (*string, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *AssetConnectionControlMessage) SetDomainGroup(v string)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *AssetConnectionControlMessage) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### GetEvict

`func (o *AssetConnectionControlMessage) GetEvict() bool`

GetEvict returns the Evict field if non-nil, zero value otherwise.

### GetEvictOk

`func (o *AssetConnectionControlMessage) GetEvictOk() (*bool, bool)`

GetEvictOk returns a tuple with the Evict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvict

`func (o *AssetConnectionControlMessage) SetEvict(v bool)`

SetEvict sets Evict field to given value.

### HasEvict

`func (o *AssetConnectionControlMessage) HasEvict() bool`

HasEvict returns a boolean if a field has been set.

### GetInternalConnectionId

`func (o *AssetConnectionControlMessage) GetInternalConnectionId() string`

GetInternalConnectionId returns the InternalConnectionId field if non-nil, zero value otherwise.

### GetInternalConnectionIdOk

`func (o *AssetConnectionControlMessage) GetInternalConnectionIdOk() (*string, bool)`

GetInternalConnectionIdOk returns a tuple with the InternalConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalConnectionId

`func (o *AssetConnectionControlMessage) SetInternalConnectionId(v string)`

SetInternalConnectionId sets InternalConnectionId field to given value.

### HasInternalConnectionId

`func (o *AssetConnectionControlMessage) HasInternalConnectionId() bool`

HasInternalConnectionId returns a boolean if a field has been set.

### GetLeadership

`func (o *AssetConnectionControlMessage) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *AssetConnectionControlMessage) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *AssetConnectionControlMessage) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *AssetConnectionControlMessage) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetNewIdentity

`func (o *AssetConnectionControlMessage) GetNewIdentity() string`

GetNewIdentity returns the NewIdentity field if non-nil, zero value otherwise.

### GetNewIdentityOk

`func (o *AssetConnectionControlMessage) GetNewIdentityOk() (*string, bool)`

GetNewIdentityOk returns a tuple with the NewIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIdentity

`func (o *AssetConnectionControlMessage) SetNewIdentity(v string)`

SetNewIdentity sets NewIdentity field to given value.

### HasNewIdentity

`func (o *AssetConnectionControlMessage) HasNewIdentity() bool`

HasNewIdentity returns a boolean if a field has been set.

### GetPartition

`func (o *AssetConnectionControlMessage) GetPartition() int64`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *AssetConnectionControlMessage) GetPartitionOk() (*int64, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *AssetConnectionControlMessage) SetPartition(v int64)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *AssetConnectionControlMessage) HasPartition() bool`

HasPartition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


