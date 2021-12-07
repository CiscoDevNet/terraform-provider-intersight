# ConnectorXmlApiMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.XmlApiMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.XmlApiMessage"]
**WithAuth** | Pointer to **bool** | Flag to disable authentication bypassing. If set to true it is expected a valid cookie/login is provided within the XML API request body. | [optional] 
**XmlRequest** | Pointer to **string** | The XML request body to proxy to the platform. | [optional] 

## Methods

### NewConnectorXmlApiMessage

`func NewConnectorXmlApiMessage(classId string, objectType string, ) *ConnectorXmlApiMessage`

NewConnectorXmlApiMessage instantiates a new ConnectorXmlApiMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorXmlApiMessageWithDefaults

`func NewConnectorXmlApiMessageWithDefaults() *ConnectorXmlApiMessage`

NewConnectorXmlApiMessageWithDefaults instantiates a new ConnectorXmlApiMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorXmlApiMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorXmlApiMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorXmlApiMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorXmlApiMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorXmlApiMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorXmlApiMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetWithAuth

`func (o *ConnectorXmlApiMessage) GetWithAuth() bool`

GetWithAuth returns the WithAuth field if non-nil, zero value otherwise.

### GetWithAuthOk

`func (o *ConnectorXmlApiMessage) GetWithAuthOk() (*bool, bool)`

GetWithAuthOk returns a tuple with the WithAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAuth

`func (o *ConnectorXmlApiMessage) SetWithAuth(v bool)`

SetWithAuth sets WithAuth field to given value.

### HasWithAuth

`func (o *ConnectorXmlApiMessage) HasWithAuth() bool`

HasWithAuth returns a boolean if a field has been set.

### GetXmlRequest

`func (o *ConnectorXmlApiMessage) GetXmlRequest() string`

GetXmlRequest returns the XmlRequest field if non-nil, zero value otherwise.

### GetXmlRequestOk

`func (o *ConnectorXmlApiMessage) GetXmlRequestOk() (*string, bool)`

GetXmlRequestOk returns a tuple with the XmlRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlRequest

`func (o *ConnectorXmlApiMessage) SetXmlRequest(v string)`

SetXmlRequest sets XmlRequest field to given value.

### HasXmlRequest

`func (o *ConnectorXmlApiMessage) HasXmlRequest() bool`

HasXmlRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


