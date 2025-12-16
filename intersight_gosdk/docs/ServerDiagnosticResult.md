# ServerDiagnosticResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.DiagnosticResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.DiagnosticResult"]
**MainComponent** | Pointer to **string** | It represents the main component on which diagnostics were performed. It could be a  physical component such as CPU. When the main component is specified in the request, the response  will include test reports for subcomponents like CPU1 and CPU2, it should be classified under  a main component like CPU. * &#x60;CPU&#x60; - This represents the CPU component. * &#x60;Memory&#x60; - This represents the Memory component. * &#x60;Storage&#x60; - This represents the Storage component. * &#x60;Network&#x60; - This represents the Network component. * &#x60;Graphics&#x60; - This represents the Graphics component. * &#x60;NVME&#x60; - This represents the NVME component. * &#x60;BMC&#x60; - This represents the BMC component. | [optional] [readonly] [default to "CPU"]
**Message** | Pointer to **string** | The message received from the endpoint on completion of the diagnostics. It contains the error message in case the diagnostics failed. | [optional] [readonly] 
**Resolution** | Pointer to **string** | Resolution in case the diagnostics failed. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of the diagnostics test performed on the component. * &#x60;OK&#x60; - The diagnostics severity is successful. * &#x60;Warning&#x60; - The diagnostics severity is warning. * &#x60;Critical&#x60; - The diagnostics severity is critical. | [optional] [readonly] [default to "OK"]
**SubComponentName** | Pointer to **string** | It represents the sub component name on which diagnostics were performed. It could be a  physical component such as CPU1 or CPU2. When the main component, like the CPU, is specified in the  diagnostics request to the endpoint, the response will include test reports for subcomponents  like CPU1 and CPU2. | [optional] [readonly] 
**SubComponentProgress** | Pointer to **string** | Progress of the diagnostics test run on the sub component. * &#x60;New&#x60; - The diagnostic task state representing new state. * &#x60;Running&#x60; - The diagnostic task state representing running state. * &#x60;Exception&#x60; - The diagnostic task state representing exception state. * &#x60;Completed&#x60; - The diagnostic task state representing completed state. | [optional] [readonly] [default to "New"]

## Methods

### NewServerDiagnosticResult

`func NewServerDiagnosticResult(classId string, objectType string, ) *ServerDiagnosticResult`

NewServerDiagnosticResult instantiates a new ServerDiagnosticResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDiagnosticResultWithDefaults

`func NewServerDiagnosticResultWithDefaults() *ServerDiagnosticResult`

NewServerDiagnosticResultWithDefaults instantiates a new ServerDiagnosticResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerDiagnosticResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerDiagnosticResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerDiagnosticResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerDiagnosticResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerDiagnosticResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerDiagnosticResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMainComponent

`func (o *ServerDiagnosticResult) GetMainComponent() string`

GetMainComponent returns the MainComponent field if non-nil, zero value otherwise.

### GetMainComponentOk

`func (o *ServerDiagnosticResult) GetMainComponentOk() (*string, bool)`

GetMainComponentOk returns a tuple with the MainComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainComponent

`func (o *ServerDiagnosticResult) SetMainComponent(v string)`

SetMainComponent sets MainComponent field to given value.

### HasMainComponent

`func (o *ServerDiagnosticResult) HasMainComponent() bool`

HasMainComponent returns a boolean if a field has been set.

### GetMessage

`func (o *ServerDiagnosticResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerDiagnosticResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerDiagnosticResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServerDiagnosticResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResolution

`func (o *ServerDiagnosticResult) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ServerDiagnosticResult) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ServerDiagnosticResult) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ServerDiagnosticResult) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSeverity

`func (o *ServerDiagnosticResult) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ServerDiagnosticResult) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ServerDiagnosticResult) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ServerDiagnosticResult) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSubComponentName

`func (o *ServerDiagnosticResult) GetSubComponentName() string`

GetSubComponentName returns the SubComponentName field if non-nil, zero value otherwise.

### GetSubComponentNameOk

`func (o *ServerDiagnosticResult) GetSubComponentNameOk() (*string, bool)`

GetSubComponentNameOk returns a tuple with the SubComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubComponentName

`func (o *ServerDiagnosticResult) SetSubComponentName(v string)`

SetSubComponentName sets SubComponentName field to given value.

### HasSubComponentName

`func (o *ServerDiagnosticResult) HasSubComponentName() bool`

HasSubComponentName returns a boolean if a field has been set.

### GetSubComponentProgress

`func (o *ServerDiagnosticResult) GetSubComponentProgress() string`

GetSubComponentProgress returns the SubComponentProgress field if non-nil, zero value otherwise.

### GetSubComponentProgressOk

`func (o *ServerDiagnosticResult) GetSubComponentProgressOk() (*string, bool)`

GetSubComponentProgressOk returns a tuple with the SubComponentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubComponentProgress

`func (o *ServerDiagnosticResult) SetSubComponentProgress(v string)`

SetSubComponentProgress sets SubComponentProgress field to given value.

### HasSubComponentProgress

`func (o *ServerDiagnosticResult) HasSubComponentProgress() bool`

HasSubComponentProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


