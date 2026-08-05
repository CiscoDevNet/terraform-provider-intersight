# ServerServerAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.ServerAssignment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.ServerAssignment"]
**Enabled** | Pointer to **bool** | Indicates if this assignment is enabled. | [optional] 
**PoolSelector** | Pointer to **string** | Odata selector that resolves to the server pool to be used for assignment, if applicable. | [optional] 
**ServerSerial** | Pointer to **string** | Serial number of the server. | [optional] 
**ServerType** | Pointer to **string** | The object type of the server - blade or rack. | [optional] 

## Methods

### NewServerServerAssignment

`func NewServerServerAssignment(classId string, objectType string, ) *ServerServerAssignment`

NewServerServerAssignment instantiates a new ServerServerAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerServerAssignmentWithDefaults

`func NewServerServerAssignmentWithDefaults() *ServerServerAssignment`

NewServerServerAssignmentWithDefaults instantiates a new ServerServerAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerServerAssignment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerServerAssignment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerServerAssignment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerServerAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerServerAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerServerAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *ServerServerAssignment) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServerServerAssignment) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServerServerAssignment) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServerServerAssignment) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoolSelector

`func (o *ServerServerAssignment) GetPoolSelector() string`

GetPoolSelector returns the PoolSelector field if non-nil, zero value otherwise.

### GetPoolSelectorOk

`func (o *ServerServerAssignment) GetPoolSelectorOk() (*string, bool)`

GetPoolSelectorOk returns a tuple with the PoolSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSelector

`func (o *ServerServerAssignment) SetPoolSelector(v string)`

SetPoolSelector sets PoolSelector field to given value.

### HasPoolSelector

`func (o *ServerServerAssignment) HasPoolSelector() bool`

HasPoolSelector returns a boolean if a field has been set.

### GetServerSerial

`func (o *ServerServerAssignment) GetServerSerial() string`

GetServerSerial returns the ServerSerial field if non-nil, zero value otherwise.

### GetServerSerialOk

`func (o *ServerServerAssignment) GetServerSerialOk() (*string, bool)`

GetServerSerialOk returns a tuple with the ServerSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSerial

`func (o *ServerServerAssignment) SetServerSerial(v string)`

SetServerSerial sets ServerSerial field to given value.

### HasServerSerial

`func (o *ServerServerAssignment) HasServerSerial() bool`

HasServerSerial returns a boolean if a field has been set.

### GetServerType

`func (o *ServerServerAssignment) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ServerServerAssignment) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ServerServerAssignment) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ServerServerAssignment) HasServerType() bool`

HasServerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


