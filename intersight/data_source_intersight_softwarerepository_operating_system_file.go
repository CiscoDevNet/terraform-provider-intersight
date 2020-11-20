package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSoftwarerepositoryOperatingSystemFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSoftwarerepositoryOperatingSystemFileRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"catalog": {
				Description: "A reference to a softwarerepositoryCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_count": {
				Description: "The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"import_action": {
				Description: "The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.\n* `None` - No action should be taken on the imported file.\n* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.\n* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.\n* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.\n* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.\n* `PreCache` - Cache the file into the Intersight Appliance.\n* `Cancel` - The cancel import process for the file into the repository.\n* `Extract` - The action to extract the file in the external repository.\n* `Evict` - Evict the cached file from the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"import_state": {
				Description: "The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.\n* `ReadyForImport` - The image is ready to be imported into the repository.\n* `Importing` - The image is being imported into the repository.\n* `Imported` - The image has been extracted and imported into the repository.\n* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.\n* `Extracting` - Indicates that the image is being extracted into the repository.\n* `Extracted` - Indicates that the image has been extracted into the repository.\n* `Failed` - The image import from an external source to the repository has failed.\n* `MetaOnly` - The image is present in an external repository.\n* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.\n* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.\n* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.\n* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.\n* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.\n* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"md5e_tag": {
				Description: "The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"md5sum": {
				Description: "The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The name of the file. It is populated as part of the image import operation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sha512sum": {
				Description: "The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"size": {
				Description: "The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"software_advisory_url": {
				Description: "The software advisory, if any, provided by the vendor for this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_source": {
				Description: "Location of the file in an external repository.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
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
			"vendor": {
				Description: "The vendor or publisher of this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "Vendor provided version for the file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSoftwarerepositoryOperatingSystemFileRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewSoftwarerepositoryOperatingSystemFileWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("download_count"); ok {
		x := int64(v.(int))
		o.SetDownloadCount(x)
	}
	if v, ok := d.GetOk("import_action"); ok {
		x := (v.(string))
		o.SetImportAction(x)
	}
	if v, ok := d.GetOk("import_state"); ok {
		x := (v.(string))
		o.SetImportState(x)
	}
	if v, ok := d.GetOk("md5e_tag"); ok {
		x := (v.(string))
		o.SetMd5eTag(x)
	}
	if v, ok := d.GetOk("md5sum"); ok {
		x := (v.(string))
		o.SetMd5sum(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("sha512sum"); ok {
		x := (v.(string))
		o.SetSha512sum(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("software_advisory_url"); ok {
		x := (v.(string))
		o.SetSoftwareAdvisoryUrl(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryOperatingSystemFileList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.SoftwarerepositoryOperatingSystemFileList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to SoftwarerepositoryOperatingSystemFile: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewSoftwarerepositoryOperatingSystemFileWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}

			if err := d.Set("catalog", flattenMapSoftwarerepositoryCatalogRelationship(s.GetCatalog(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Catalog: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return fmt.Errorf("error occurred while setting property Description: %+v", err)
			}
			if err := d.Set("download_count", (s.GetDownloadCount())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadCount: %+v", err)
			}
			if err := d.Set("import_action", (s.GetImportAction())); err != nil {
				return fmt.Errorf("error occurred while setting property ImportAction: %+v", err)
			}
			if err := d.Set("import_state", (s.GetImportState())); err != nil {
				return fmt.Errorf("error occurred while setting property ImportState: %+v", err)
			}
			if err := d.Set("md5e_tag", (s.GetMd5eTag())); err != nil {
				return fmt.Errorf("error occurred while setting property Md5eTag: %+v", err)
			}
			if err := d.Set("md5sum", (s.GetMd5sum())); err != nil {
				return fmt.Errorf("error occurred while setting property Md5sum: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("sha512sum", (s.GetSha512sum())); err != nil {
				return fmt.Errorf("error occurred while setting property Sha512sum: %+v", err)
			}
			if err := d.Set("size", (s.GetSize())); err != nil {
				return fmt.Errorf("error occurred while setting property Size: %+v", err)
			}
			if err := d.Set("software_advisory_url", (s.GetSoftwareAdvisoryUrl())); err != nil {
				return fmt.Errorf("error occurred while setting property SoftwareAdvisoryUrl: %+v", err)
			}

			if err := d.Set("nr_source", flattenMapSoftwarerepositoryFileServer(s.GetSource(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Source: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return fmt.Errorf("error occurred while setting property Vendor: %+v", err)
			}
			if err := d.Set("nr_version", (s.GetVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property Version: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
