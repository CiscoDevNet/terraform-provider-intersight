# ConnectorWinrmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.WinrmRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.WinrmRequest"]
**AssetTargetMoid** | Pointer to **string** | The Target endpoint Moid which is used to fetch the previously persisted Target information in Intersight. | [optional] 
**Body** | Pointer to **string** | Contains the file content to be copied or the script to be executed. | [optional] 
**Filename** | Pointer to **string** | The absolute filename to which body is to be written to. This field is ignored if the opType is ExecuteScript. | [optional] 
**OpType** | Pointer to **string** | The type of operation to be performed on the endpoint. File copy and script execution are the possible options. Only script execution is supported for now. * &#x60;ExecuteScript&#x60; - Executes the given script on the target windows machine. * &#x60;CopyToFile&#x60; - The plugin copies the body of the incoming message to the given location. | [optional] [default to "ExecuteScript"]
**SessionId** | Pointer to **string** | A unique id to track long running script executions. Every execution request is identified by a unique session id. Scripts that have longer execution times can be tracked to completion by using their unique session id. | [optional] 
**Timeout** | Pointer to **int64** | The time within which the script execution must be completed. If the script execution exceeds the given time limit an appropriate response is sent back to the calling service. | [optional] 

## Methods

### NewConnectorWinrmRequest

`func NewConnectorWinrmRequest(classId string, objectType string, ) *ConnectorWinrmRequest`

NewConnectorWinrmRequest instantiates a new ConnectorWinrmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorWinrmRequestWithDefaults

`func NewConnectorWinrmRequestWithDefaults() *ConnectorWinrmRequest`

NewConnectorWinrmRequestWithDefaults instantiates a new ConnectorWinrmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorWinrmRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorWinrmRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorWinrmRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorWinrmRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorWinrmRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorWinrmRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetTargetMoid

`func (o *ConnectorWinrmRequest) GetAssetTargetMoid() string`

GetAssetTargetMoid returns the AssetTargetMoid field if non-nil, zero value otherwise.

### GetAssetTargetMoidOk

`func (o *ConnectorWinrmRequest) GetAssetTargetMoidOk() (*string, bool)`

GetAssetTargetMoidOk returns a tuple with the AssetTargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTargetMoid

`func (o *ConnectorWinrmRequest) SetAssetTargetMoid(v string)`

SetAssetTargetMoid sets AssetTargetMoid field to given value.

### HasAssetTargetMoid

`func (o *ConnectorWinrmRequest) HasAssetTargetMoid() bool`

HasAssetTargetMoid returns a boolean if a field has been set.

### GetBody

`func (o *ConnectorWinrmRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ConnectorWinrmRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ConnectorWinrmRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ConnectorWinrmRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFilename

`func (o *ConnectorWinrmRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ConnectorWinrmRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ConnectorWinrmRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ConnectorWinrmRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetOpType

`func (o *ConnectorWinrmRequest) GetOpType() string`

GetOpType returns the OpType field if non-nil, zero value otherwise.

### GetOpTypeOk

`func (o *ConnectorWinrmRequest) GetOpTypeOk() (*string, bool)`

GetOpTypeOk returns a tuple with the OpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpType

`func (o *ConnectorWinrmRequest) SetOpType(v string)`

SetOpType sets OpType field to given value.

### HasOpType

`func (o *ConnectorWinrmRequest) HasOpType() bool`

HasOpType returns a boolean if a field has been set.

### GetSessionId

`func (o *ConnectorWinrmRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ConnectorWinrmRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ConnectorWinrmRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ConnectorWinrmRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimeout

`func (o *ConnectorWinrmRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectorWinrmRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectorWinrmRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectorWinrmRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


