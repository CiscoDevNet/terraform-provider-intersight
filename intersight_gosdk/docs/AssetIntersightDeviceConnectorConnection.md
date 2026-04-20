# AssetIntersightDeviceConnectorConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.IntersightDeviceConnectorConnection"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.IntersightDeviceConnectorConnection"]
**PreClaim** | Pointer to **bool** | A pre-claim of an embedded target. Does not require a security token to be provided at target create, target will be created in a pre-claim state. | [optional] 
**SecurityToken** | Pointer to **string** | The SecurityToken object holds a time-limited random string used for claiming a device. It is created implicitly for each device connector at the time of registration, providing a secure mechanism for users to assert administrative access during device claim operations. #### Purpose SecurityToken acts as a temporary credential that proves a user&#39;s administrative access to a device, allowing them to claim the device within Intersight. It strengthens the security of claim operations, preventing unauthorized device claims by restricting access to users who possess the token. #### Key Concepts - **Time-Bound Security:** Tokens are generated with expiration times, ensuring they are only valid for a limited duration to reduce the risk of misuse. - **Claim Validation:** Used during claim operations to validate that the user has the necessary privileges to manage the device. - **Access Control:** Integrates with Intersight&#39;s security model, providing controlled access to device claim functionalities. | [optional] 
**SerialNumber** | Pointer to **string** | Obtained from the device connector management UI or API (REST endpoint &#39;/connector/DeviceIdentifiers&#39;). | [optional] 

## Methods

### NewAssetIntersightDeviceConnectorConnection

`func NewAssetIntersightDeviceConnectorConnection(classId string, objectType string, ) *AssetIntersightDeviceConnectorConnection`

NewAssetIntersightDeviceConnectorConnection instantiates a new AssetIntersightDeviceConnectorConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetIntersightDeviceConnectorConnectionWithDefaults

`func NewAssetIntersightDeviceConnectorConnectionWithDefaults() *AssetIntersightDeviceConnectorConnection`

NewAssetIntersightDeviceConnectorConnectionWithDefaults instantiates a new AssetIntersightDeviceConnectorConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetIntersightDeviceConnectorConnection) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetIntersightDeviceConnectorConnection) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetIntersightDeviceConnectorConnection) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetIntersightDeviceConnectorConnection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetIntersightDeviceConnectorConnection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetIntersightDeviceConnectorConnection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPreClaim

`func (o *AssetIntersightDeviceConnectorConnection) GetPreClaim() bool`

GetPreClaim returns the PreClaim field if non-nil, zero value otherwise.

### GetPreClaimOk

`func (o *AssetIntersightDeviceConnectorConnection) GetPreClaimOk() (*bool, bool)`

GetPreClaimOk returns a tuple with the PreClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreClaim

`func (o *AssetIntersightDeviceConnectorConnection) SetPreClaim(v bool)`

SetPreClaim sets PreClaim field to given value.

### HasPreClaim

`func (o *AssetIntersightDeviceConnectorConnection) HasPreClaim() bool`

HasPreClaim returns a boolean if a field has been set.

### GetSecurityToken

`func (o *AssetIntersightDeviceConnectorConnection) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *AssetIntersightDeviceConnectorConnection) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *AssetIntersightDeviceConnectorConnection) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *AssetIntersightDeviceConnectorConnection) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AssetIntersightDeviceConnectorConnection) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AssetIntersightDeviceConnectorConnection) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AssetIntersightDeviceConnectorConnection) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AssetIntersightDeviceConnectorConnection) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


