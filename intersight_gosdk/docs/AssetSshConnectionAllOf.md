# AssetSshConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.SshConnection"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.SshConnection"]
**ManagementAddress** | Pointer to **string** | The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target. | [optional] 
**Port** | Pointer to **int64** | The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection. | [optional] 

## Methods

### NewAssetSshConnectionAllOf

`func NewAssetSshConnectionAllOf(classId string, objectType string, ) *AssetSshConnectionAllOf`

NewAssetSshConnectionAllOf instantiates a new AssetSshConnectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetSshConnectionAllOfWithDefaults

`func NewAssetSshConnectionAllOfWithDefaults() *AssetSshConnectionAllOf`

NewAssetSshConnectionAllOfWithDefaults instantiates a new AssetSshConnectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetSshConnectionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetSshConnectionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetSshConnectionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetSshConnectionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetSshConnectionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetSshConnectionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetManagementAddress

`func (o *AssetSshConnectionAllOf) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *AssetSshConnectionAllOf) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *AssetSshConnectionAllOf) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *AssetSshConnectionAllOf) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPort

`func (o *AssetSshConnectionAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AssetSshConnectionAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AssetSshConnectionAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *AssetSshConnectionAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


