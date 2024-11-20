# OsValidRemoteTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.ValidRemoteTarget"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.ValidRemoteTarget"]
**FibreChannelLuns** | Pointer to [**[]OsFibreChannelResponse**](OsFibreChannelResponse.md) |  | [optional] 
**IscsiLuns** | Pointer to [**[]OsIscsiLunResponse**](OsIscsiLunResponse.md) |  | [optional] 
**Src** | Pointer to **string** | Flag to denote the source of the request. If the call is from Orchestration UI, only the flat list of Install targets can be sent as response. | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewOsValidRemoteTarget

`func NewOsValidRemoteTarget(classId string, objectType string, ) *OsValidRemoteTarget`

NewOsValidRemoteTarget instantiates a new OsValidRemoteTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsValidRemoteTargetWithDefaults

`func NewOsValidRemoteTargetWithDefaults() *OsValidRemoteTarget`

NewOsValidRemoteTargetWithDefaults instantiates a new OsValidRemoteTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsValidRemoteTarget) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsValidRemoteTarget) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsValidRemoteTarget) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsValidRemoteTarget) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsValidRemoteTarget) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsValidRemoteTarget) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFibreChannelLuns

`func (o *OsValidRemoteTarget) GetFibreChannelLuns() []OsFibreChannelResponse`

GetFibreChannelLuns returns the FibreChannelLuns field if non-nil, zero value otherwise.

### GetFibreChannelLunsOk

`func (o *OsValidRemoteTarget) GetFibreChannelLunsOk() (*[]OsFibreChannelResponse, bool)`

GetFibreChannelLunsOk returns a tuple with the FibreChannelLuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelLuns

`func (o *OsValidRemoteTarget) SetFibreChannelLuns(v []OsFibreChannelResponse)`

SetFibreChannelLuns sets FibreChannelLuns field to given value.

### HasFibreChannelLuns

`func (o *OsValidRemoteTarget) HasFibreChannelLuns() bool`

HasFibreChannelLuns returns a boolean if a field has been set.

### SetFibreChannelLunsNil

`func (o *OsValidRemoteTarget) SetFibreChannelLunsNil(b bool)`

 SetFibreChannelLunsNil sets the value for FibreChannelLuns to be an explicit nil

### UnsetFibreChannelLuns
`func (o *OsValidRemoteTarget) UnsetFibreChannelLuns()`

UnsetFibreChannelLuns ensures that no value is present for FibreChannelLuns, not even an explicit nil
### GetIscsiLuns

`func (o *OsValidRemoteTarget) GetIscsiLuns() []OsIscsiLunResponse`

GetIscsiLuns returns the IscsiLuns field if non-nil, zero value otherwise.

### GetIscsiLunsOk

`func (o *OsValidRemoteTarget) GetIscsiLunsOk() (*[]OsIscsiLunResponse, bool)`

GetIscsiLunsOk returns a tuple with the IscsiLuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiLuns

`func (o *OsValidRemoteTarget) SetIscsiLuns(v []OsIscsiLunResponse)`

SetIscsiLuns sets IscsiLuns field to given value.

### HasIscsiLuns

`func (o *OsValidRemoteTarget) HasIscsiLuns() bool`

HasIscsiLuns returns a boolean if a field has been set.

### SetIscsiLunsNil

`func (o *OsValidRemoteTarget) SetIscsiLunsNil(b bool)`

 SetIscsiLunsNil sets the value for IscsiLuns to be an explicit nil

### UnsetIscsiLuns
`func (o *OsValidRemoteTarget) UnsetIscsiLuns()`

UnsetIscsiLuns ensures that no value is present for IscsiLuns, not even an explicit nil
### GetSrc

`func (o *OsValidRemoteTarget) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *OsValidRemoteTarget) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *OsValidRemoteTarget) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *OsValidRemoteTarget) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetServer

`func (o *OsValidRemoteTarget) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OsValidRemoteTarget) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OsValidRemoteTarget) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *OsValidRemoteTarget) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *OsValidRemoteTarget) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *OsValidRemoteTarget) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


