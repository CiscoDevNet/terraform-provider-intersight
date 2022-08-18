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

data "intersight_smtp_policy" "odata_skip_negative_top_positive_zero_result_mo"{
	odata {
	 skip = -600
	 top = 10
	}
}

data "intersight_access_policy" "odata_skip_negative_top_positive_zero_result_mo"{
	odata {
	 skip = -950
	}
}

data "intersight_smtp_policy" "odata_top_skip_zero_results"{
	odata {
	 top = 10
	 skip = 23
	}
}

data "intersight_hyperflex_software_distribution_version" "odata_at1"{
	odata {
	at = "Version eq '1.3(1b)'"      
	}
}

data "intersight_hyperflex_hxdp_version" "odata_at2"{
	odata {
	at = "Version eq '4.0(2e)'"     
	}
}


