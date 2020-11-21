# ConnectorpackConnectorPackUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connectorpack.ConnectorPackUpdate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connectorpack.ConnectorPackUpdate"]
**CurrentVersion** | Pointer to **string** | Version of connector pack currently running in UCS Director. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the connector pack. | [optional] [readonly] 
**NewVersion** | Pointer to **string** | Version of connector pack to be installed in the next upgrade cycle. | [optional] [readonly] 

## Methods

### NewConnectorpackConnectorPackUpdate

`func NewConnectorpackConnectorPackUpdate(classId string, objectType string, ) *ConnectorpackConnectorPackUpdate`

NewConnectorpackConnectorPackUpdate instantiates a new ConnectorpackConnectorPackUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorpackConnectorPackUpdateWithDefaults

`func NewConnectorpackConnectorPackUpdateWithDefaults() *ConnectorpackConnectorPackUpdate`

NewConnectorpackConnectorPackUpdateWithDefaults instantiates a new ConnectorpackConnectorPackUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorpackConnectorPackUpdate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorpackConnectorPackUpdate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorpackConnectorPackUpdate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorpackConnectorPackUpdate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorpackConnectorPackUpdate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorpackConnectorPackUpdate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentVersion

`func (o *ConnectorpackConnectorPackUpdate) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ConnectorpackConnectorPackUpdate) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ConnectorpackConnectorPackUpdate) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ConnectorpackConnectorPackUpdate) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetName

`func (o *ConnectorpackConnectorPackUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorpackConnectorPackUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorpackConnectorPackUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorpackConnectorPackUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewVersion

`func (o *ConnectorpackConnectorPackUpdate) GetNewVersion() string`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *ConnectorpackConnectorPackUpdate) GetNewVersionOk() (*string, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *ConnectorpackConnectorPackUpdate) SetNewVersion(v string)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *ConnectorpackConnectorPackUpdate) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


