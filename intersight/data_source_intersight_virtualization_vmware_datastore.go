package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualizationVmwareDatastore() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualizationVmwareDatastoreRead,
		Schema: map[string]*schema.Schema{
			"accessible": {
				Description: "Shows if this datastore is accessible.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"host_count": {
				Description: "Number of hosts attached to or supported-by this datastore.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"identity": {
				Description: "The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"maintenance_mode": {
				Description: "Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"multiple_host_access": {
				Description: "Indicates if this datastore is connected to multiple hosts.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"name": {
				Description: "Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Datastore health status, as reported by the hypervisor platform.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"thin_provisioning_supported": {
				Description: "Indicates if this datastore supports thin provisioning for files.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"type": {
				Description: "A string indicating the type of the datastore (VMFS, NFS, etc).\n* `Unknown` - The nature of the file system is unknown.\n* `VMFS` - It is a Virtual Machine Filesystem.\n* `NFS` - It is a Network File System.\n* `vSAN` - It is a virtual Storage Area Network file system.\n* `VirtualVolume` - A Virtual Volume datastore represents a storage container in a hypervisor server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"un_committed": {
				Description: "Space uncommitted in this datastore in bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"url": {
				Description: "The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/').",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_count": {
				Description: "Number of virtual machines relying on (using) this datastore.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"accessible": {
					Description: "Shows if this datastore is accessible.",
					Type:        schema.TypeBool,
					Optional:    true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"capacity": {
						Description: "Storage utilization for this datastore.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"capacity": {
									Description: "The total capacity of the entity (bytes).",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"free": {
									Description: "Free storage space remaining in the entity (bytes) as a point-in-time snapshot. The available space is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used storage capacity is also reported.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"used": {
									Description: "Spaced already used by this entity (bytes), as a point-in-time snapshot.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster": {
						Description: "A reference to a virtualizationVmwareCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"datacenter": {
						Description: "A reference to a virtualizationVmwareDatacenter resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"host_count": {
						Description: "Number of hosts attached to or supported-by this datastore.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"hosts": {
						Description: "An array of relationships to virtualizationVmwareHost resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"identity": {
						Description: "The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"maintenance_mode": {
						Description: "Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"multiple_host_access": {
						Description: "Indicates if this datastore is connected to multiple hosts.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"name": {
						Description: "Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"status": {
						Description: "Datastore health status, as reported by the hypervisor platform.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"thin_provisioning_supported": {
						Description: "Indicates if this datastore supports thin provisioning for files.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"type": {
						Description: "A string indicating the type of the datastore (VMFS, NFS, etc).\n* `Unknown` - The nature of the file system is unknown.\n* `VMFS` - It is a Virtual Machine Filesystem.\n* `NFS` - It is a Network File System.\n* `vSAN` - It is a virtual Storage Area Network file system.\n* `VirtualVolume` - A Virtual Volume datastore represents a storage container in a hypervisor server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"un_committed": {
						Description: "Space uncommitted in this datastore in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"url": {
						Description: "The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vm_count": {
						Description: "Number of virtual machines relying on (using) this datastore.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceVirtualizationVmwareDatastoreRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VirtualizationVmwareDatastore{}
	if v, ok := d.GetOk("accessible"); ok {
		x := (v.(bool))
		o.SetAccessible(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("host_count"); ok {
		x := int64(v.(int))
		o.SetHostCount(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("maintenance_mode"); ok {
		x := (v.(bool))
		o.SetMaintenanceMode(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("multiple_host_access"); ok {
		x := (v.(bool))
		o.SetMultipleHostAccess(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("thin_provisioning_supported"); ok {
		x := (v.(bool))
		o.SetThinProvisioningSupported(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("un_committed"); ok {
		x := int64(v.(int))
		o.SetUnCommitted(x)
	}
	if v, ok := d.GetOk("url"); ok {
		x := (v.(string))
		o.SetUrl(x)
	}
	if v, ok := d.GetOk("vm_count"); ok {
		x := int64(v.(int))
		o.SetVmCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VirtualizationVmwareDatastore object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVmwareDatastoreList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VirtualizationVmwareDatastore: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VirtualizationVmwareDatastoreList.GetCount()
	var i int32
	var virtualizationVmwareDatastoreResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVmwareDatastoreList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VirtualizationVmwareDatastore: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VirtualizationVmwareDatastoreList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VirtualizationVmwareDatastore data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["accessible"] = (s.GetAccessible())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["capacity"] = flattenMapVirtualizationStorageCapacity(s.GetCapacity(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster"] = flattenMapVirtualizationVmwareClusterRelationship(s.GetCluster(), d)

				temp["datacenter"] = flattenMapVirtualizationVmwareDatacenterRelationship(s.GetDatacenter(), d)
				temp["host_count"] = (s.GetHostCount())

				temp["hosts"] = flattenListVirtualizationVmwareHostRelationship(s.GetHosts(), d)
				temp["identity"] = (s.GetIdentity())
				temp["maintenance_mode"] = (s.GetMaintenanceMode())
				temp["moid"] = (s.GetMoid())
				temp["multiple_host_access"] = (s.GetMultipleHostAccess())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["thin_provisioning_supported"] = (s.GetThinProvisioningSupported())
				temp["type"] = (s.GetType())
				temp["un_committed"] = (s.GetUnCommitted())
				temp["url"] = (s.GetUrl())
				temp["vm_count"] = (s.GetVmCount())
				virtualizationVmwareDatastoreResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(virtualizationVmwareDatastoreResults))
	if err := d.Set("results", virtualizationVmwareDatastoreResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(virtualizationVmwareDatastoreResults[0]["moid"].(string))
	return de
}
