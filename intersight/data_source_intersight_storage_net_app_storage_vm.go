package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageNetAppStorageVm() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageNetAppStorageVmRead,
		Schema: map[string]*schema.Schema{
			"cifs_enabled": {
				Description: "Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fcp_enabled": {
				Description: "Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"iscsi_enabled": {
				Description: "Status for iSCSI protocol allowed to run on Vservers.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the tenant in storage array.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nfs_enabled": {
				Description: "Status for Network File System Protocol ( NFS ) allowed to run on  Vservers.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"nvme_enabled": {
				Description: "Status for NVME protocol allowed to run on Vservers.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Description: "The state of this tenant.\n* `Unknown` - Component state is not available.\n* `Starting` - Component is being started.\n* `Running` - Component is currently running.\n* `Stopping` - Component is being stopped.\n* `Stopped` - Component has been stopped.\n* `Deleting` - Component deletion is in progress.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "The uuid of this tenant in storage array.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"array": {
						Description: "A reference to a storageNetAppCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"cifs_enabled": {
						Description: "Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"disk_pool": {
						Description: "An array of relationships to storageNetAppAggregate resources.",
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
					"fcp_enabled": {
						Description: "Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"iscsi_enabled": {
						Description: "Status for iSCSI protocol allowed to run on Vservers.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Name of the tenant in storage array.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"nfs_enabled": {
						Description: "Status for Network File System Protocol ( NFS ) allowed to run on  Vservers.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"nvme_enabled": {
						Description: "Status for NVME protocol allowed to run on Vservers.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"state": {
						Description: "The state of this tenant.\n* `Unknown` - Component state is not available.\n* `Starting` - Component is being started.\n* `Running` - Component is currently running.\n* `Stopping` - Component is being stopped.\n* `Stopped` - Component has been stopped.\n* `Deleting` - Component deletion is in progress.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
					"uuid": {
						Description: "The uuid of this tenant in storage array.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageNetAppStorageVmRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageNetAppStorageVm{}
	if v, ok := d.GetOk("cifs_enabled"); ok {
		x := (v.(bool))
		o.SetCifsEnabled(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("fcp_enabled"); ok {
		x := (v.(bool))
		o.SetFcpEnabled(x)
	}
	if v, ok := d.GetOk("iscsi_enabled"); ok {
		x := (v.(bool))
		o.SetIscsiEnabled(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("nfs_enabled"); ok {
		x := (v.(bool))
		o.SetNfsEnabled(x)
	}
	if v, ok := d.GetOk("nvme_enabled"); ok {
		x := (v.(bool))
		o.SetNvmeEnabled(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageNetAppStorageVm object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppStorageVmList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageNetAppStorageVm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageNetAppStorageVmList.GetCount()
	var i int32
	var storageNetAppStorageVmResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppStorageVmList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageNetAppStorageVm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageNetAppStorageVmList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageNetAppStorageVm data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["array"] = flattenMapStorageNetAppClusterRelationship(s.GetArray(), d)
				temp["cifs_enabled"] = (s.GetCifsEnabled())
				temp["class_id"] = (s.GetClassId())

				temp["disk_pool"] = flattenListStorageNetAppAggregateRelationship(s.GetDiskPool(), d)
				temp["fcp_enabled"] = (s.GetFcpEnabled())
				temp["iscsi_enabled"] = (s.GetIscsiEnabled())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["nfs_enabled"] = (s.GetNfsEnabled())
				temp["nvme_enabled"] = (s.GetNvmeEnabled())
				temp["object_type"] = (s.GetObjectType())
				temp["state"] = (s.GetState())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["uuid"] = (s.GetUuid())
				storageNetAppStorageVmResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageNetAppStorageVmResults))
	if err := d.Set("results", storageNetAppStorageVmResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageNetAppStorageVmResults[0]["moid"].(string))
	return de
}
