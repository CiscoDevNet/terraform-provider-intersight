# IamClientMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | Parsed device model from raw User-Agent. | [optional] 
**UserAgent** | Pointer to **string** | The value of the \&quot;User-Agent\&quot; HTTP header, as sent by the HTTP client when it initiate a session to Intersight. This can be used to identify the client operating system, browser type and browser version. Example - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36 It is set when User successfully passed OAuth login flow and receives Access Token. | [optional] 

## Methods

### NewIamClientMetaAllOf

`func NewIamClientMetaAllOf() *IamClientMetaAllOf`

NewIamClientMetaAllOf instantiates a new IamClientMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamClientMetaAllOfWithDefaults

`func NewIamClientMetaAllOfWithDefaults() *IamClientMetaAllOf`

NewIamClientMetaAllOfWithDefaults instantiates a new IamClientMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *IamClientMetaAllOf) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *IamClientMetaAllOf) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *IamClientMetaAllOf) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *IamClientMetaAllOf) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetUserAgent

`func (o *IamClientMetaAllOf) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *IamClientMetaAllOf) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *IamClientMetaAllOf) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *IamClientMetaAllOf) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


