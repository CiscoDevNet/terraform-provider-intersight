# MotemplateSyncError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "motemplate.SyncError"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "motemplate.SyncError"]
**Message** | Pointer to **string** | The localized message based on the locale setting of the user&#39;s context providing the error description. | [optional] [readonly] 
**Type** | Pointer to **string** | The error type that indicates the point of failure. * &#x60;Transient&#x60; - Any error which is a runtime error due to some other action in progress on the derived object that is blocking the sync action. This error type is retriable.For example, when vNIC Template is updated, but the derived vNIC or vNICs are part of a LAN Connectivity policy associated with a profile being deployed to endpoint. In this scenario, the derived vNIC update displays this error. * &#x60;Validation&#x60; - When the template sync on the derived object fails due to an incorrect configuration, it displays a validation error. This error type is considered fatal and not retried.For example, when a new policy is attached to a server profile template, the sync to a derived server profile fails due to the policy attach errors. * &#x60;User&#x60; - Any configuration error due to incorrect or invalid input and that requires user intervention for correction, is displayed under this category. This error type is considered fatal and not retried.For example, when a new policy is attached to a server profile template, the sync to a derived server profile fails. This happens when the policyis not applicable to the server assigned to the server profile, such as the Power policy that is not applicable for UCS Rack servers. * &#x60;Internal&#x60; - Any application internal errors are displayed under this category. This error type is considered fatal and not retried. | [optional] [readonly] [default to "Transient"]

## Methods

### NewMotemplateSyncError

`func NewMotemplateSyncError(classId string, objectType string, ) *MotemplateSyncError`

NewMotemplateSyncError instantiates a new MotemplateSyncError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotemplateSyncErrorWithDefaults

`func NewMotemplateSyncErrorWithDefaults() *MotemplateSyncError`

NewMotemplateSyncErrorWithDefaults instantiates a new MotemplateSyncError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MotemplateSyncError) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MotemplateSyncError) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MotemplateSyncError) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MotemplateSyncError) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MotemplateSyncError) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MotemplateSyncError) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *MotemplateSyncError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MotemplateSyncError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MotemplateSyncError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MotemplateSyncError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *MotemplateSyncError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MotemplateSyncError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MotemplateSyncError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MotemplateSyncError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


