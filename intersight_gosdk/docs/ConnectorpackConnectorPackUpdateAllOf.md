# ConnectorpackConnectorPackUpdateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connectorpack.ConnectorPackUpdate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connectorpack.ConnectorPackUpdate"]
**CurrentVersion** | Pointer to **string** | Version of connector pack currently running in UCS Director. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the connector pack. | [optional] [readonly] 
**NewVersion** | Pointer to **string** | Version of connector pack to be installed in the next upgrade cycle. | [optional] [readonly] 

## Methods

### NewConnectorpackConnectorPackUpdateAllOf

`func NewConnectorpackConnectorPackUpdateAllOf(classId string, objectType string, ) *ConnectorpackConnectorPackUpdateAllOf`

NewConnectorpackConnectorPackUpdateAllOf instantiates a new ConnectorpackConnectorPackUpdateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorpackConnectorPackUpdateAllOfWithDefaults

`func NewConnectorpackConnectorPackUpdateAllOfWithDefaults() *ConnectorpackConnectorPackUpdateAllOf`

NewConnectorpackConnectorPackUpdateAllOfWithDefaults instantiates a new ConnectorpackConnectorPackUpdateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorpackConnectorPackUpdateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorpackConnectorPackUpdateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentVersion

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ConnectorpackConnectorPackUpdateAllOf) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ConnectorpackConnectorPackUpdateAllOf) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetName

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorpackConnectorPackUpdateAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorpackConnectorPackUpdateAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewVersion

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetNewVersion() string`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *ConnectorpackConnectorPackUpdateAllOf) GetNewVersionOk() (*string, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *ConnectorpackConnectorPackUpdateAllOf) SetNewVersion(v string)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *ConnectorpackConnectorPackUpdateAllOf) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


