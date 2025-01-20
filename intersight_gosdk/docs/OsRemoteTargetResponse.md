# OsRemoteTargetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**InterfaceName** | Pointer to **string** | Interface name of the iSCSI/FC used as initiator. | [optional] 

## Methods

### NewOsRemoteTargetResponse

`func NewOsRemoteTargetResponse(classId string, objectType string, ) *OsRemoteTargetResponse`

NewOsRemoteTargetResponse instantiates a new OsRemoteTargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsRemoteTargetResponseWithDefaults

`func NewOsRemoteTargetResponseWithDefaults() *OsRemoteTargetResponse`

NewOsRemoteTargetResponseWithDefaults instantiates a new OsRemoteTargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsRemoteTargetResponse) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsRemoteTargetResponse) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsRemoteTargetResponse) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsRemoteTargetResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsRemoteTargetResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsRemoteTargetResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaceName

`func (o *OsRemoteTargetResponse) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *OsRemoteTargetResponse) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *OsRemoteTargetResponse) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *OsRemoteTargetResponse) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


