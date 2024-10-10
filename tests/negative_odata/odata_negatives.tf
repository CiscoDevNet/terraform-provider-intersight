data "intersight_aaa_audit_record" "odata_top_and_skip_negative_values"{
	odata {
	 skip = -90
	 top = -20
	}
}

data "intersight_search_search_item" "odata_skip_negative_value"{
	odata {
	 skip = -200
	}
}

data "intersight_compute_rack_unit" "odata_skip_negative_top_positive_value"{
	odata {
	 skip = -200
	 top = 50
	}
}

