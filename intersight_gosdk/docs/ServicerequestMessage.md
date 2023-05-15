# ServicerequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "servicerequest.Message"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "servicerequest.Message"]
**ServiceItemName** | Pointer to **string** | The service item in which operation is perfomed. | [optional] [readonly] 

## Methods

### NewServicerequestMessage

`func NewServicerequestMessage(classId string, objectType string, ) *ServicerequestMessage`

NewServicerequestMessage instantiates a new ServicerequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicerequestMessageWithDefaults

`func NewServicerequestMessageWithDefaults() *ServicerequestMessage`

NewServicerequestMessageWithDefaults instantiates a new ServicerequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServicerequestMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServicerequestMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServicerequestMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServicerequestMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServicerequestMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServicerequestMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetServiceItemName

`func (o *ServicerequestMessage) GetServiceItemName() string`

GetServiceItemName returns the ServiceItemName field if non-nil, zero value otherwise.

### GetServiceItemNameOk

`func (o *ServicerequestMessage) GetServiceItemNameOk() (*string, bool)`

GetServiceItemNameOk returns a tuple with the ServiceItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemName

`func (o *ServicerequestMessage) SetServiceItemName(v string)`

SetServiceItemName sets ServiceItemName field to given value.

### HasServiceItemName

`func (o *ServicerequestMessage) HasServiceItemName() bool`

HasServiceItemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


