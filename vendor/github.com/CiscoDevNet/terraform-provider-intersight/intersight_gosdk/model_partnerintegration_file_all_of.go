/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15711
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// PartnerintegrationFileAllOf Definition of the list of properties defined in 'partnerintegration.File', excluding properties defined in parent classes.
type PartnerintegrationFileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Path of the file being uploaded.
	FilePath *string `json:"FilePath,omitempty"`
	// Type of the file being uploaded. * `None` - Invalid file type for partnerIntegration appliance. * `Model` - Model file of Generic Device. * `Etl` - ETL file of Generic Device. * `Ui` - UI file of Generic Device. * `DeviceConnector` - Generic Device Connector file.
	FileType *string `json:"FileType,omitempty"`
	// The partner integration workspace to use to upload the File.
	WorkspaceName        *string                                `json:"WorkspaceName,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationFileAllOf PartnerintegrationFileAllOf

// NewPartnerintegrationFileAllOf instantiates a new PartnerintegrationFileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationFileAllOf(classId string, objectType string) *PartnerintegrationFileAllOf {
	this := PartnerintegrationFileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fileType string = "None"
	this.FileType = &fileType
	return &this
}

// NewPartnerintegrationFileAllOfWithDefaults instantiates a new PartnerintegrationFileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationFileAllOfWithDefaults() *PartnerintegrationFileAllOf {
	this := PartnerintegrationFileAllOf{}
	var classId string = "partnerintegration.File"
	this.ClassId = classId
	var objectType string = "partnerintegration.File"
	this.ObjectType = objectType
	var fileType string = "None"
	this.FileType = &fileType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationFileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationFileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationFileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationFileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *PartnerintegrationFileAllOf) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFileAllOf) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *PartnerintegrationFileAllOf) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *PartnerintegrationFileAllOf) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *PartnerintegrationFileAllOf) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFileAllOf) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *PartnerintegrationFileAllOf) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *PartnerintegrationFileAllOf) SetFileType(v string) {
	o.FileType = &v
}

// GetWorkspaceName returns the WorkspaceName field value if set, zero value otherwise.
func (o *PartnerintegrationFileAllOf) GetWorkspaceName() string {
	if o == nil || o.WorkspaceName == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceName
}

// GetWorkspaceNameOk returns a tuple with the WorkspaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFileAllOf) GetWorkspaceNameOk() (*string, bool) {
	if o == nil || o.WorkspaceName == nil {
		return nil, false
	}
	return o.WorkspaceName, true
}

// HasWorkspaceName returns a boolean if a field has been set.
func (o *PartnerintegrationFileAllOf) HasWorkspaceName() bool {
	if o != nil && o.WorkspaceName != nil {
		return true
	}

	return false
}

// SetWorkspaceName gets a reference to the given string and assigns it to the WorkspaceName field.
func (o *PartnerintegrationFileAllOf) SetWorkspaceName(v string) {
	o.WorkspaceName = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *PartnerintegrationFileAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFileAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *PartnerintegrationFileAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *PartnerintegrationFileAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o PartnerintegrationFileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FileType != nil {
		toSerialize["FileType"] = o.FileType
	}
	if o.WorkspaceName != nil {
		toSerialize["WorkspaceName"] = o.WorkspaceName
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerintegrationFileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerintegrationFileAllOf := _PartnerintegrationFileAllOf{}

	if err = json.Unmarshal(bytes, &varPartnerintegrationFileAllOf); err == nil {
		*o = PartnerintegrationFileAllOf(varPartnerintegrationFileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FileType")
		delete(additionalProperties, "WorkspaceName")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerintegrationFileAllOf struct {
	value *PartnerintegrationFileAllOf
	isSet bool
}

func (v NullablePartnerintegrationFileAllOf) Get() *PartnerintegrationFileAllOf {
	return v.value
}

func (v *NullablePartnerintegrationFileAllOf) Set(val *PartnerintegrationFileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationFileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationFileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationFileAllOf(val *PartnerintegrationFileAllOf) *NullablePartnerintegrationFileAllOf {
	return &NullablePartnerintegrationFileAllOf{value: val, isSet: true}
}

func (v NullablePartnerintegrationFileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationFileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
