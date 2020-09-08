# ConnectorFileMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgType** | Pointer to **string** | Message type carrying the file operation to perform. * &#x60;OpenFile&#x60; - Open a file for write to on the platforms file system. Must be the first message sent to the plugin when writing files, attempting to write to or close a file that has not been opened will return error. If file does not exist file will be created, if it does exist file will be truncated. If the files parent directory does not exist an error will be returned. * &#x60;FileContent&#x60; - Carries content to write to an open file. The complete file content can be buffered across many FileContent messages, with each subsequent message received appending to the end of the file. If the total size of FileContent messages exceeds the file size limit of 5MB the plugin will return error and remove the opened file. * &#x60;CloseFile&#x60; - Close the file. Close message must be sent at the end of file transfer. If no close is sent and no input is received after a timeout of 10 minutes the file will be removed. After a CloseFile is received no further FileContent messages will be accepted and will return error. * &#x60;DeleteFile&#x60; - Delete a file from the filesystem. If the file does not exist operation is a no-op and no error is returned. File must be within one of the platforms allowed writable directories. * &#x60;ListDir&#x60; - Open a directory and return a list of files in the directory. Does not recursively walk subdirectories, returns only the top level files in the directory. Output will be a list of json encoded FileInfos. * &#x60;&#x60; - Open a file and read the contents into a response message. The response message body is the entire contents of the file with no encoding or compression. If the file does not exist an error will be returned. If the size of the file exceeds 500KB an error will be returned without reading any file contents. * &#x60;GetFreeSpace&#x60; - Check the disk space availability of folder and returns the free disk space in percentage. Output will be a list of encoded FileInfos. | [optional] [default to "OpenFile"]
**Path** | Pointer to **string** | The absolute path of the file to open on the platforms file system. Must be a sub-directory of a directory defined within the platform configurations WriteableDirectories. The file system device to write to must also have sufficient free space to write to (&lt;75% full). Must be set for each message that is sent. | [optional] 
**Stream** | Pointer to **string** | The stream of bytes to write to file. Only applicable for FileContent message. Ignored for OpenFile and CloseFile messages. | [optional] 

## Methods

### NewConnectorFileMessage

`func NewConnectorFileMessage() *ConnectorFileMessage`

NewConnectorFileMessage instantiates a new ConnectorFileMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorFileMessageWithDefaults

`func NewConnectorFileMessageWithDefaults() *ConnectorFileMessage`

NewConnectorFileMessageWithDefaults instantiates a new ConnectorFileMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgType

`func (o *ConnectorFileMessage) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *ConnectorFileMessage) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *ConnectorFileMessage) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *ConnectorFileMessage) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetPath

`func (o *ConnectorFileMessage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConnectorFileMessage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConnectorFileMessage) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ConnectorFileMessage) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStream

`func (o *ConnectorFileMessage) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConnectorFileMessage) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConnectorFileMessage) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConnectorFileMessage) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


