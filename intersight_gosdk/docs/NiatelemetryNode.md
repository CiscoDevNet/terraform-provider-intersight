# NiatelemetryNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Node"]
**Hostname** | Pointer to **string** | Returns hostname of the node. | [optional] 
**ManagementtIp** | Pointer to **string** | Returns management IP of the node. | [optional] 
**OutofbandIp** | Pointer to **string** | Returns out of band IP of the node. | [optional] 

## Methods

### NewNiatelemetryNode

`func NewNiatelemetryNode(classId string, objectType string, ) *NiatelemetryNode`

NewNiatelemetryNode instantiates a new NiatelemetryNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNodeWithDefaults

`func NewNiatelemetryNodeWithDefaults() *NiatelemetryNode`

NewNiatelemetryNodeWithDefaults instantiates a new NiatelemetryNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *NiatelemetryNode) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NiatelemetryNode) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NiatelemetryNode) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NiatelemetryNode) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetManagementtIp

`func (o *NiatelemetryNode) GetManagementtIp() string`

GetManagementtIp returns the ManagementtIp field if non-nil, zero value otherwise.

### GetManagementtIpOk

`func (o *NiatelemetryNode) GetManagementtIpOk() (*string, bool)`

GetManagementtIpOk returns a tuple with the ManagementtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementtIp

`func (o *NiatelemetryNode) SetManagementtIp(v string)`

SetManagementtIp sets ManagementtIp field to given value.

### HasManagementtIp

`func (o *NiatelemetryNode) HasManagementtIp() bool`

HasManagementtIp returns a boolean if a field has been set.

### GetOutofbandIp

`func (o *NiatelemetryNode) GetOutofbandIp() string`

GetOutofbandIp returns the OutofbandIp field if non-nil, zero value otherwise.

### GetOutofbandIpOk

`func (o *NiatelemetryNode) GetOutofbandIpOk() (*string, bool)`

GetOutofbandIpOk returns a tuple with the OutofbandIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutofbandIp

`func (o *NiatelemetryNode) SetOutofbandIp(v string)`

SetOutofbandIp sets OutofbandIp field to given value.

### HasOutofbandIp

`func (o *NiatelemetryNode) HasOutofbandIp() bool`

HasOutofbandIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


