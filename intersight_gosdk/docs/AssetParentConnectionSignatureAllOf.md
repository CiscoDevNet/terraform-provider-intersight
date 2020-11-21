# AssetParentConnectionSignatureAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ParentConnectionSignature"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ParentConnectionSignature"]
**DeviceId** | Pointer to **string** | The moid of the parent device registration. | [optional] 
**NodeId** | Pointer to **string** | The node identity of the parent device, corresponds to the parents ClusterMember.memberIdentity. Used on connect to establish through which device in a cluster the child is connected through. | [optional] 
**Signature** | Pointer to **string** | The result of signing the deviceId appended with the timeStamp fields with the devices private key. | [optional] 
**TimeStamp** | Pointer to [**time.Time**](time.Time.md) | The time at which the signature was generated. Date is accurate to Intersights clock. Used to expire the signature. | [optional] 

## Methods

### NewAssetParentConnectionSignatureAllOf

`func NewAssetParentConnectionSignatureAllOf(classId string, objectType string, ) *AssetParentConnectionSignatureAllOf`

NewAssetParentConnectionSignatureAllOf instantiates a new AssetParentConnectionSignatureAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetParentConnectionSignatureAllOfWithDefaults

`func NewAssetParentConnectionSignatureAllOfWithDefaults() *AssetParentConnectionSignatureAllOf`

NewAssetParentConnectionSignatureAllOfWithDefaults instantiates a new AssetParentConnectionSignatureAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetParentConnectionSignatureAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetParentConnectionSignatureAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetParentConnectionSignatureAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetParentConnectionSignatureAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetParentConnectionSignatureAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetParentConnectionSignatureAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *AssetParentConnectionSignatureAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetParentConnectionSignatureAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetParentConnectionSignatureAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetParentConnectionSignatureAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetNodeId

`func (o *AssetParentConnectionSignatureAllOf) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *AssetParentConnectionSignatureAllOf) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *AssetParentConnectionSignatureAllOf) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *AssetParentConnectionSignatureAllOf) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetSignature

`func (o *AssetParentConnectionSignatureAllOf) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AssetParentConnectionSignatureAllOf) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AssetParentConnectionSignatureAllOf) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AssetParentConnectionSignatureAllOf) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTimeStamp

`func (o *AssetParentConnectionSignatureAllOf) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AssetParentConnectionSignatureAllOf) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AssetParentConnectionSignatureAllOf) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *AssetParentConnectionSignatureAllOf) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


