package main

import (
	"github.com/apacta/mcp-server/config"
	"github.com/apacta/mcp-server/models"
	tools_events "github.com/apacta/mcp-server/tools/events"
	tools_contactcustomfieldattributes "github.com/apacta/mcp-server/tools/contactcustomfieldattributes"
	tools_paymenttermtypes "github.com/apacta/mcp-server/tools/paymenttermtypes"
	tools_companies "github.com/apacta/mcp-server/tools/companies"
	tools_activities "github.com/apacta/mcp-server/tools/activities"
	tools_massmessagesusers "github.com/apacta/mcp-server/tools/massmessagesusers"
	tools_companysettings "github.com/apacta/mcp-server/tools/companysettings"
	tools_contacts "github.com/apacta/mcp-server/tools/contacts"
	tools_timeentries "github.com/apacta/mcp-server/tools/timeentries"
	tools_users "github.com/apacta/mcp-server/tools/users"
	tools_usercustomfieldvalue "github.com/apacta/mcp-server/tools/usercustomfieldvalue"
	tools_timeentrytypes "github.com/apacta/mcp-server/tools/timeentrytypes"
	tools_projects "github.com/apacta/mcp-server/tools/projects"
	tools_financial_statistics "github.com/apacta/mcp-server/tools/financial_statistics"
	tools_formfields "github.com/apacta/mcp-server/tools/formfields"
	tools_projectstatustypes "github.com/apacta/mcp-server/tools/projectstatustypes"
	tools_invoices "github.com/apacta/mcp-server/tools/invoices"
	tools_offerstatuses "github.com/apacta/mcp-server/tools/offerstatuses"
	tools_cities "github.com/apacta/mcp-server/tools/cities"
	tools_countries "github.com/apacta/mcp-server/tools/countries"
	tools_formtemplates "github.com/apacta/mcp-server/tools/formtemplates"
	tools_usercustomfieldattributes "github.com/apacta/mcp-server/tools/usercustomfieldattributes"
	tools_clockingrecords "github.com/apacta/mcp-server/tools/clockingrecords"
	tools_contactpersons "github.com/apacta/mcp-server/tools/contactpersons"
	tools_paymentterms "github.com/apacta/mcp-server/tools/paymentterms"
	tools_formfieldtypes "github.com/apacta/mcp-server/tools/formfieldtypes"
	tools_integrations "github.com/apacta/mcp-server/tools/integrations"
	tools_forms "github.com/apacta/mcp-server/tools/forms"
	tools_expenselines "github.com/apacta/mcp-server/tools/expenselines"
	tools_offers "github.com/apacta/mcp-server/tools/offers"
	tools_currencies "github.com/apacta/mcp-server/tools/currencies"
	tools_invoicelines "github.com/apacta/mcp-server/tools/invoicelines"
	tools_products "github.com/apacta/mcp-server/tools/products"
	tools_materialrentals "github.com/apacta/mcp-server/tools/materialrentals"
	tools_wallcomments "github.com/apacta/mcp-server/tools/wallcomments"
	tools_timeentryunittypes "github.com/apacta/mcp-server/tools/timeentryunittypes"
	tools_expenses "github.com/apacta/mcp-server/tools/expenses"
	tools_invoicelinetexttemplates "github.com/apacta/mcp-server/tools/invoicelinetexttemplates"
	tools_rejectionreasons "github.com/apacta/mcp-server/tools/rejectionreasons"
	tools_wages "github.com/apacta/mcp-server/tools/wages"
	tools_productvariants "github.com/apacta/mcp-server/tools/productvariants"
	tools_vendorproducts "github.com/apacta/mcp-server/tools/vendorproducts"
	tools_vendors "github.com/apacta/mcp-server/tools/vendors"
	tools_materials "github.com/apacta/mcp-server/tools/materials"
	tools_wallposts "github.com/apacta/mcp-server/tools/wallposts"
	tools_roles "github.com/apacta/mcp-server/tools/roles"
	tools_drivingtypes "github.com/apacta/mcp-server/tools/drivingtypes"
	tools_timeentryvaluetypes "github.com/apacta/mcp-server/tools/timeentryvaluetypes"
	tools_companiesvendors "github.com/apacta/mcp-server/tools/companiesvendors"
	tools_projectstatuses "github.com/apacta/mcp-server/tools/projectstatuses"
	tools_financialstatistics "github.com/apacta/mcp-server/tools/financialstatistics"
	tools_contacttypes "github.com/apacta/mcp-server/tools/contacttypes"
	tools_expense_oioubl_files "github.com/apacta/mcp-server/tools/expense_oioubl_files"
	tools_projectcustomfieldattributes "github.com/apacta/mcp-server/tools/projectcustomfieldattributes"
	tools_reports "github.com/apacta/mcp-server/tools/reports"
	tools_vendorproductpricefiles "github.com/apacta/mcp-server/tools/vendorproductpricefiles"
	tools_changelog "github.com/apacta/mcp-server/tools/changelog"
	tools_timeentryrates "github.com/apacta/mcp-server/tools/timeentryrates"
	tools_expensefiles "github.com/apacta/mcp-server/tools/expensefiles"
	tools_timeentryintervals "github.com/apacta/mcp-server/tools/timeentryintervals"
	tools_stocklocations "github.com/apacta/mcp-server/tools/stocklocations"
	tools_invoiceemails "github.com/apacta/mcp-server/tools/invoiceemails"
	tools_contactcustomfieldvalue "github.com/apacta/mcp-server/tools/contactcustomfieldvalue"
	tools_timeentryrate "github.com/apacta/mcp-server/tools/timeentryrate"
	tools_invoicefiles "github.com/apacta/mcp-server/tools/invoicefiles"
	tools_defaultprojectstatuses "github.com/apacta/mcp-server/tools/defaultprojectstatuses"
	tools_ping "github.com/apacta/mcp-server/tools/ping"
	tools_employeehours "github.com/apacta/mcp-server/tools/employeehours"
	tools_timeentryrulegroups "github.com/apacta/mcp-server/tools/timeentryrulegroups"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_events.CreateGet_eventsTool(cfg),
		tools_events.CreatePost_eventsTool(cfg),
		tools_contactcustomfieldattributes.CreateGet_contact_custom_field_attributesTool(cfg),
		tools_paymenttermtypes.CreateGet_payment_term_typesTool(cfg),
		tools_events.CreateGet_events_is_user_freeTool(cfg),
		tools_companies.CreateGet_companies_subscription_self_serviceTool(cfg),
		tools_activities.CreateDelete_activities_bulkdeleteTool(cfg),
		tools_massmessagesusers.CreateGet_mass_messages_users_mass_messages_user_idTool(cfg),
		tools_massmessagesusers.CreatePut_mass_messages_users_mass_messages_user_idTool(cfg),
		tools_companysettings.CreateGetcompaysettingslistTool(cfg),
		tools_massmessagesusers.CreateGet_mass_messages_usersTool(cfg),
		tools_contacts.CreateGet_contactsTool(cfg),
		tools_contacts.CreatePost_contactsTool(cfg),
		tools_timeentries.CreateGet_time_entriesTool(cfg),
		tools_timeentries.CreatePost_time_entriesTool(cfg),
		tools_users.CreateGet_users_user_id_integration_settingsTool(cfg),
		tools_users.CreatePost_users_user_id_integration_settingsTool(cfg),
		tools_usercustomfieldvalue.CreateGet_users_user_id_user_custom_field_value_user_custom_field_value_idTool(cfg),
		tools_usercustomfieldvalue.CreatePut_users_user_id_user_custom_field_value_user_custom_field_value_idTool(cfg),
		tools_timeentrytypes.CreateBulkdeactivatetimeentrytypesTool(cfg),
		tools_companies.CreateGet_companies_company_id_integration_feature_settingsTool(cfg),
		tools_projects.CreateGet_projects_project_id_files_file_idTool(cfg),
		tools_projects.CreatePut_projects_project_id_files_file_idTool(cfg),
		tools_projects.CreateDelete_projects_project_id_files_file_idTool(cfg),
		tools_financial_statistics.CreateGetmarginTool(cfg),
		tools_formfields.CreatePost_form_fieldsTool(cfg),
		tools_projectstatustypes.CreateGet_project_status_typesTool(cfg),
		tools_invoices.CreatePost_invoices_invoice_id_copyTool(cfg),
		tools_offerstatuses.CreatePut_offer_statuses_offer_status_idTool(cfg),
		tools_offerstatuses.CreateDelete_offer_statuses_offer_status_idTool(cfg),
		tools_offerstatuses.CreateGet_offer_statuses_offer_status_idTool(cfg),
		tools_timeentrytypes.CreateGet_time_entry_typesTool(cfg),
		tools_timeentrytypes.CreatePost_time_entry_typesTool(cfg),
		tools_timeentrytypes.CreateBulkactivatetimeentrytypesTool(cfg),
		tools_cities.CreateGet_cities_city_idTool(cfg),
		tools_users.CreateGet_users_user_id_integration_settings_integration_settings_user_idTool(cfg),
		tools_users.CreatePut_users_user_id_integration_settings_integration_settings_user_idTool(cfg),
		tools_users.CreateDelete_users_user_id_integration_settings_integration_settings_user_idTool(cfg),
		tools_formfields.CreateGet_form_fields_form_field_idTool(cfg),
		tools_countries.CreateGet_countries_country_idTool(cfg),
		tools_formtemplates.CreateGet_form_templatesTool(cfg),
		tools_companies.CreateGet_companies_company_id_integration_settingsTool(cfg),
		tools_companies.CreatePost_companies_company_id_integration_settingsTool(cfg),
		tools_usercustomfieldattributes.CreateGet_user_custom_field_attributes_user_custom_field_attribute_idTool(cfg),
		tools_contacts.CreateGet_contacts_contact_idTool(cfg),
		tools_contacts.CreatePut_contacts_contact_idTool(cfg),
		tools_contacts.CreateDelete_contacts_contact_idTool(cfg),
		tools_clockingrecords.CreatePost_clocking_records_checkoutTool(cfg),
		tools_contactpersons.CreateDelete_contacts_contact_id_contact_persons_contact_person_idTool(cfg),
		tools_contactpersons.CreateGetcontactpersonTool(cfg),
		tools_contactpersons.CreateEditcontactpersonTool(cfg),
		tools_offerstatuses.CreateGet_offer_statusesTool(cfg),
		tools_offerstatuses.CreatePost_offer_statusesTool(cfg),
		tools_paymentterms.CreateGet_payment_terms_payment_term_idTool(cfg),
		tools_formfieldtypes.CreateGet_form_field_types_form_field_type_idTool(cfg),
		tools_integrations.CreateGet_integrations_contactssyncTool(cfg),
		tools_clockingrecords.CreatePost_clocking_recordsTool(cfg),
		tools_clockingrecords.CreateGet_clocking_recordsTool(cfg),
		tools_activities.CreateDelete_activities_activity_idTool(cfg),
		tools_activities.CreatePut_activities_activity_idTool(cfg),
		tools_companies.CreateGet_companies_company_idTool(cfg),
		tools_projects.CreateDelete_projects_project_id_users_user_idTool(cfg),
		tools_projects.CreateGet_projects_project_id_users_user_idTool(cfg),
		tools_forms.CreateGet_forms_undelete_form_idTool(cfg),
		tools_forms.CreateDelete_forms_form_idTool(cfg),
		tools_forms.CreateGet_forms_form_idTool(cfg),
		tools_forms.CreatePut_forms_form_idTool(cfg),
		tools_expenselines.CreateGet_expense_linesTool(cfg),
		tools_expenselines.CreatePost_expense_linesTool(cfg),
		tools_offers.CreateDelete_offers_offer_idTool(cfg),
		tools_offers.CreateGet_offers_offer_idTool(cfg),
		tools_paymentterms.CreateGet_payment_terms_erpTool(cfg),
		tools_currencies.CreateGet_currenciesTool(cfg),
		tools_invoicelines.CreateGet_invoice_linesTool(cfg),
		tools_invoicelines.CreatePost_invoice_linesTool(cfg),
		tools_offerstatuses.CreateDelete_offer_statuses_bulkdeleteTool(cfg),
		tools_products.CreateGet_productsTool(cfg),
		tools_products.CreatePost_productsTool(cfg),
		tools_usercustomfieldattributes.CreateGet_user_custom_field_attributesTool(cfg),
		tools_materialrentals.CreateDelete_materials_material_id_rentals_material_rental_idTool(cfg),
		tools_materialrentals.CreateGet_materials_material_id_rentals_material_rental_idTool(cfg),
		tools_materialrentals.CreatePut_materials_material_id_rentals_material_rental_idTool(cfg),
		tools_financial_statistics.CreateGetfinancialstatisticsTool(cfg),
		tools_wallcomments.CreateGet_wall_comments_wall_comment_idTool(cfg),
		tools_timeentryunittypes.CreateGet_time_entry_unit_types_time_entry_unit_type_idTool(cfg),
		tools_financial_statistics.CreateGetproductscostpriceTool(cfg),
		tools_expenses.CreateGet_expenses_highest_amountTool(cfg),
		tools_invoicelinetexttemplates.CreateDelete_invoice_line_text_template_invoice_line_text_template_idTool(cfg),
		tools_invoicelinetexttemplates.CreateGet_invoice_line_text_template_invoice_line_text_template_idTool(cfg),
		tools_rejectionreasons.CreateGet_overview_rejection_reasonsTool(cfg),
		tools_wages.CreateGet_wages_downloadsalaryfileTool(cfg),
		tools_companies.CreateDelete_companies_company_id_price_margins_price_margins_idTool(cfg),
		tools_companies.CreateGet_companies_company_id_price_margins_price_margins_idTool(cfg),
		tools_companies.CreatePost_companies_company_id_price_margins_price_margins_idTool(cfg),
		tools_productvariants.CreateGet_products_product_id_variantsTool(cfg),
		tools_users.CreateGet_users_resendwelcomesmsTool(cfg),
		tools_vendorproducts.CreateGet_vendor_products_vendor_product_idTool(cfg),
		tools_formtemplates.CreateGet_form_templates_form_template_idTool(cfg),
		tools_vendors.CreateGetvendorslistTool(cfg),
		tools_vendors.CreateAddvendorTool(cfg),
		tools_financial_statistics.CreateGetexpensessalespriceTool(cfg),
		tools_materials.CreateGet_materialsTool(cfg),
		tools_materials.CreatePost_materialsTool(cfg),
		tools_wallposts.CreateGet_wall_posts_wall_post_idTool(cfg),
		tools_expenses.CreateBulkdeleteexpensesTool(cfg),
		tools_projects.CreateGet_projects_project_id_usersTool(cfg),
		tools_projects.CreatePost_projects_project_id_usersTool(cfg),
		tools_users.CreateDelete_users_user_idTool(cfg),
		tools_users.CreateGet_users_user_idTool(cfg),
		tools_users.CreatePut_users_user_idTool(cfg),
		tools_roles.CreateGet_rolesTool(cfg),
		tools_forms.CreateGet_formsTool(cfg),
		tools_forms.CreatePost_formsTool(cfg),
		tools_drivingtypes.CreateDelete_driving_types_driving_type_idTool(cfg),
		tools_drivingtypes.CreateGet_driving_types_driving_type_idTool(cfg),
		tools_drivingtypes.CreatePut_driving_types_driving_type_idTool(cfg),
		tools_materialrentals.CreatePost_materials_material_id_rentals_checkoutTool(cfg),
		tools_users.CreateGet_usersTool(cfg),
		tools_users.CreatePost_usersTool(cfg),
		tools_productvariants.CreateDelete_products_product_id_variants_variant_type_variant_idTool(cfg),
		tools_timeentryvaluetypes.CreateGet_time_entry_value_typesTool(cfg),
		tools_projects.CreateGet_projectsTool(cfg),
		tools_projects.CreatePost_projectsTool(cfg),
		tools_forms.CreateGet_forms_view_time_form_pdf_form_idTool(cfg),
		tools_financial_statistics.CreateGetfinancialstatisticsoverviewTool(cfg),
		tools_companiesvendors.CreateDelete_companies_vendors_companies_vendor_idTool(cfg),
		tools_companiesvendors.CreateGetcompaniesvendorTool(cfg),
		tools_companiesvendors.CreateEditcompaniesvendorTool(cfg),
		tools_integrations.CreatePost_integrations_billysauthenticateTool(cfg),
		tools_projectstatuses.CreateGet_project_statusesTool(cfg),
		tools_projectstatuses.CreatePost_project_statusesTool(cfg),
		tools_financialstatistics.CreateGet_financial_statistics_workinghoursTool(cfg),
		tools_contacttypes.CreateGet_contact_typesTool(cfg),
		tools_expense_oioubl_files.CreateGet_expenses_expense_id_original_files_file_idTool(cfg),
		tools_projects.CreateGet_projects_project_id_project_filesTool(cfg),
		tools_projectcustomfieldattributes.CreateGet_project_custom_field_attributesTool(cfg),
		tools_companiesvendors.CreateGetcompaiesvendorslistTool(cfg),
		tools_companiesvendors.CreateAddcompaniesvendorTool(cfg),
		tools_materialrentals.CreateGet_materials_material_id_rentalsTool(cfg),
		tools_materialrentals.CreatePost_materials_material_id_rentalsTool(cfg),
		tools_reports.CreateGet_reportsTool(cfg),
		tools_vendorproductpricefiles.CreateGet_vendor_product_price_files_vendor_product_price_file_idTool(cfg),
		tools_financial_statistics.CreateGetinvoicedamountTool(cfg),
		tools_paymenttermtypes.CreateGet_payment_term_types_payment_term_type_idTool(cfg),
		tools_contacttypes.CreateGet_contact_types_contact_type_idTool(cfg),
		tools_changelog.CreateGet_offers_offer_id_changelogTool(cfg),
		tools_timeentryrates.CreateGet_time_entry_ratesTool(cfg),
		tools_timeentryrates.CreatePost_time_entry_ratesTool(cfg),
		tools_timeentryvaluetypes.CreateGet_time_entry_value_types_time_entry_value_type_idTool(cfg),
		tools_contactcustomfieldattributes.CreateGet_contact_custom_field_attributes_contact_custom_field_attribute_idTool(cfg),
		tools_paymentterms.CreateGet_payment_termsTool(cfg),
		tools_invoices.CreatePost_invoices_invoice_id_unlinkprojectpdfTool(cfg),
		tools_expensefiles.CreateGet_expense_filesTool(cfg),
		tools_timeentryintervals.CreateGet_time_entry_intervals_time_entry_interval_idTool(cfg),
		tools_timeentryunittypes.CreateGet_time_entry_unit_typesTool(cfg),
		tools_projectstatuses.CreateDelete_project_statuses_project_status_idTool(cfg),
		tools_projectstatuses.CreateGet_project_statuses_project_status_idTool(cfg),
		tools_projectstatuses.CreatePut_project_statuses_project_status_idTool(cfg),
		tools_expenses.CreateSendemailsexpensesTool(cfg),
		tools_companiesvendors.CreateBulkcompaniesvendorsTool(cfg),
		tools_drivingtypes.CreateGet_driving_typesTool(cfg),
		tools_drivingtypes.CreatePost_driving_typesTool(cfg),
		tools_expense_oioubl_files.CreateGet_expenses_expense_id_original_filesTool(cfg),
		tools_projects.CreatePost_projects_project_id_send_bulk_pdfTool(cfg),
		tools_stocklocations.CreatePut_stock_locations_location_idTool(cfg),
		tools_stocklocations.CreateDelete_stock_locations_location_idTool(cfg),
		tools_stocklocations.CreateGet_stock_locations_location_idTool(cfg),
		tools_clockingrecords.CreatePut_clocking_records_clocking_record_idTool(cfg),
		tools_clockingrecords.CreateDelete_clocking_records_clocking_record_idTool(cfg),
		tools_clockingrecords.CreateGet_clocking_records_clocking_record_idTool(cfg),
		tools_vendorproducts.CreateGet_vendor_productsTool(cfg),
		tools_invoiceemails.CreateGetoneinvoiceemailsTool(cfg),
		tools_companies.CreateDelete_companies_company_id_integration_settings_companies_integration_setting_idTool(cfg),
		tools_companies.CreateGet_companies_company_id_integration_settings_companies_integration_setting_idTool(cfg),
		tools_companies.CreatePut_companies_company_id_integration_settings_companies_integration_setting_idTool(cfg),
		tools_users.CreateUsersbulkdeactivateTool(cfg),
		tools_products.CreatePut_products_product_idTool(cfg),
		tools_products.CreateDelete_products_product_idTool(cfg),
		tools_products.CreateGet_products_product_idTool(cfg),
		tools_invoicelines.CreateDelete_invoice_lines_invoice_line_idTool(cfg),
		tools_invoicelines.CreateGet_invoice_lines_invoice_line_idTool(cfg),
		tools_invoicelines.CreatePut_invoice_lines_invoice_line_idTool(cfg),
		tools_companiesvendors.CreateGetcompaniesvendorsexpensestatisticsTool(cfg),
		tools_contactcustomfieldvalue.CreateGet_contacts_contact_id_contact_custom_field_valuesTool(cfg),
		tools_wallposts.CreateGet_wall_posts_wall_post_id_wall_commentsTool(cfg),
		tools_timeentryrate.CreateDelete_time_entry_rates_time_entry_rate_idTool(cfg),
		tools_timeentryrates.CreateGet_time_entry_rates_time_entry_rate_idTool(cfg),
		tools_timeentryrates.CreatePut_time_entry_rates_time_entry_rate_idTool(cfg),
		tools_timeentrytypes.CreateBulkdeletetimeentrytypesTool(cfg),
		tools_invoicefiles.CreateDelete_invoices_invoice_id_files_file_idTool(cfg),
		tools_invoicefiles.CreateGetoneinvoicefilesTool(cfg),
		tools_companies.CreateDelete_companies_company_id_form_templates_form_template_idTool(cfg),
		tools_companies.CreateGet_companies_company_id_form_templates_form_template_idTool(cfg),
		tools_activities.CreateGet_activitiesTool(cfg),
		tools_activities.CreatePost_activitiesTool(cfg),
		tools_defaultprojectstatuses.CreatePost_project_statuses_add_defaultTool(cfg),
		tools_projects.CreateDelete_projects_project_id_project_files_project_file_idTool(cfg),
		tools_projects.CreateGet_projects_project_id_project_files_project_file_idTool(cfg),
		tools_projects.CreatePut_projects_project_id_project_files_project_file_idTool(cfg),
		tools_contactpersons.CreateGetcontactpersonslistTool(cfg),
		tools_contactpersons.CreateAddcontactpersonTool(cfg),
		tools_currencies.CreateGet_currencies_currency_idTool(cfg),
		tools_ping.CreateGet_pingTool(cfg),
		tools_companies.CreateGet_companiesTool(cfg),
		tools_timeentryintervals.CreateGet_time_entry_intervalsTool(cfg),
		tools_projectstatuses.CreateDelete_project_statuses_bulkdeleteTool(cfg),
		tools_formfieldtypes.CreateGet_form_field_typesTool(cfg),
		tools_products.CreateBulkdeleteproductsTool(cfg),
		tools_projectcustomfieldattributes.CreateGet_project_custom_field_attributes_project_custom_field_attribute_idTool(cfg),
		tools_expenselines.CreateDelete_expense_lines_expense_line_idTool(cfg),
		tools_expenselines.CreateGet_expense_lines_expense_line_idTool(cfg),
		tools_expenselines.CreatePut_expense_lines_expense_line_idTool(cfg),
		tools_companies.CreateGet_companies_company_id_companies_integration_feature_settings_c_integration_feature_setting_idTool(cfg),
		tools_companies.CreatePut_companies_company_id_companies_integration_feature_settings_c_integration_feature_setting_idTool(cfg),
		tools_expenses.CreateDelete_expenses_expense_idTool(cfg),
		tools_expenses.CreateGet_expenses_expense_idTool(cfg),
		tools_expenses.CreatePut_expenses_expense_idTool(cfg),
		tools_projects.CreateGet_projects_project_id_all_filesTool(cfg),
		tools_companies.CreateGet_companies_company_id_companies_integration_feature_settingsTool(cfg),
		tools_companies.CreatePost_companies_company_id_companies_integration_feature_settingsTool(cfg),
		tools_employeehours.CreateGet_employee_hoursTool(cfg),
		tools_drivingtypes.CreateBulkdeletedrivingtypesTool(cfg),
		tools_vendorproductpricefiles.CreateGet_vendor_product_price_filesTool(cfg),
		tools_invoices.CreateDelete_invoices_invoice_idTool(cfg),
		tools_invoices.CreateGet_invoices_invoice_idTool(cfg),
		tools_invoices.CreatePut_invoices_invoice_idTool(cfg),
		tools_integrations.CreateGet_integrations_listTool(cfg),
		tools_invoicelinetexttemplates.CreateGet_invoice_line_text_templateTool(cfg),
		tools_offers.CreateGet_offersTool(cfg),
		tools_offers.CreatePost_offersTool(cfg),
		tools_projectstatuses.CreateGet_projects_has_projects_with_custom_statusesTool(cfg),
		tools_companies.CreateGet_companies_company_id_integration_feature_settings_integration_feature_setting_idTool(cfg),
		tools_expensefiles.CreateGet_expense_files_expense_file_idTool(cfg),
		tools_expensefiles.CreatePut_expense_files_expense_file_idTool(cfg),
		tools_expensefiles.CreateDelete_expense_files_expense_file_idTool(cfg),
		tools_timeentries.CreateDelete_time_entries_time_entry_idTool(cfg),
		tools_timeentries.CreateGet_time_entries_time_entry_idTool(cfg),
		tools_timeentries.CreatePut_time_entries_time_entry_idTool(cfg),
		tools_invoices.CreatePost_invoices_invoice_id_linkprojectpdfTool(cfg),
		tools_invoices.CreateGet_invoices_vatoptionsTool(cfg),
		tools_events.CreateDelete_events_event_idTool(cfg),
		tools_events.CreateGet_events_event_idTool(cfg),
		tools_events.CreatePut_events_event_idTool(cfg),
		tools_integrations.CreateGet_integrations_viewTool(cfg),
		tools_contacts.CreateBulkdeletecontactsTool(cfg),
		tools_integrations.CreateGet_integrations_productssyncTool(cfg),
		tools_vendors.CreateDelete_vendors_vendor_idTool(cfg),
		tools_vendors.CreateGetvendorTool(cfg),
		tools_vendors.CreateEditvendorTool(cfg),
		tools_invoices.CreateBulkdeleteinvoicesTool(cfg),
		tools_countries.CreateGet_countriesTool(cfg),
		tools_companies.CreateGet_companies_company_id_form_templatesTool(cfg),
		tools_timeentryrulegroups.CreateGet_time_entry_rule_groupsTool(cfg),
		tools_expenses.CreateGet_expensesTool(cfg),
		tools_expenses.CreatePost_expensesTool(cfg),
		tools_financial_statistics.CreateGetmaterialrentalscostpriceTool(cfg),
		tools_wallcomments.CreatePost_wall_commentsTool(cfg),
		tools_usercustomfieldvalue.CreateGet_users_user_id_user_custom_field_valueTool(cfg),
		tools_projects.CreateDelete_projects_project_idTool(cfg),
		tools_projects.CreateGet_projects_project_idTool(cfg),
		tools_projects.CreatePut_projects_project_idTool(cfg),
		tools_invoices.CreatePost_invoicesTool(cfg),
		tools_invoices.CreateGet_invoicesTool(cfg),
		tools_stocklocations.CreateGet_stock_locationsTool(cfg),
		tools_stocklocations.CreatePost_stock_locationsTool(cfg),
		tools_timeentrytypes.CreateDelete_time_entry_types_time_entry_type_idTool(cfg),
		tools_timeentrytypes.CreateGet_time_entry_types_time_entry_type_idTool(cfg),
		tools_timeentrytypes.CreatePut_time_entry_types_time_entry_type_idTool(cfg),
		tools_wallposts.CreateGet_wall_postsTool(cfg),
		tools_wallposts.CreatePost_wall_postsTool(cfg),
		tools_cities.CreateGet_citiesTool(cfg),
		tools_projects.CreateGet_projects_project_id_filesTool(cfg),
		tools_products.CreatePost_products_undelete_product_idTool(cfg),
		tools_invoicefiles.CreateGetinvoicefilesTool(cfg),
		tools_materials.CreateDelete_materials_material_idTool(cfg),
		tools_materials.CreateGet_materials_material_idTool(cfg),
		tools_materials.CreatePut_materials_material_idTool(cfg),
		tools_users.CreateUsersbulkactivateTool(cfg),
	}
}
