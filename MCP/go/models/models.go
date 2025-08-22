package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UserCustomFieldValue represents the UserCustomFieldValue schema from the OpenAPI specification
type UserCustomFieldValue struct {
	User_custom_field_attribute_id string `json:"user_custom_field_attribute_id,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Value string `json:"value,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// CompaniesIntegrationFeatureSetting represents the CompaniesIntegrationFeatureSetting schema from the OpenAPI specification
type CompaniesIntegrationFeatureSetting struct {
	Id string `json:"id,omitempty"`
	Integration_feature_setting_id string `json:"integration_feature_setting_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Value string `json:"value,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
}

// ContactCustomFieldAttribute represents the ContactCustomFieldAttribute schema from the OpenAPI specification
type ContactCustomFieldAttribute struct {
	Modified string `json:"modified,omitempty"`
	Access_type string `json:"access_type,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Id string `json:"id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Placement int `json:"placement,omitempty"`
	Created string `json:"created,omitempty"`
	Is_active bool `json:"is_active,omitempty"`
	Name string `json:"name,omitempty"`
	Company_id string `json:"company_id,omitempty"`
}

// ForbiddenError represents the ForbiddenError schema from the OpenAPI specification
type ForbiddenError struct {
	Success bool `json:"success,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// TimeEntryRuleGroup represents the TimeEntryRuleGroup schema from the OpenAPI specification
type TimeEntryRuleGroup struct {
	Company_id string `json:"company_id,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Salary_period_from string `json:"salary_period_from,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Salary_period_days float32 `json:"salary_period_days,omitempty"`
}

// ContactCustomFieldValue represents the ContactCustomFieldValue schema from the OpenAPI specification
type ContactCustomFieldValue struct {
	Modified string `json:"modified,omitempty"`
	Value string `json:"value,omitempty"`
	Contact_custom_field_attribute_id string `json:"contact_custom_field_attribute_id,omitempty"`
	Contact_id string `json:"contact_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
}

// InvoiceLine represents the InvoiceLine schema from the OpenAPI specification
type InvoiceLine struct {
	Discount_text string `json:"discount_text,omitempty"`
	Parent_id string `json:"parent_id,omitempty"`
	Material_id string `json:"material_id,omitempty"`
	Created string `json:"created,omitempty"`
	Discount_percent int `json:"discount_percent,omitempty"`
	Form_id string `json:"form_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	TypeField string `json:"type,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Product_id string `json:"product_id,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	Selling_price float32 `json:"selling_price,omitempty"`
	Invoice_id string `json:"invoice_id,omitempty"`
	Ean_product_id string `json:"ean_product_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Product_bundle_id string `json:"product_bundle_id,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Description string `json:"description,omitempty"`
}

// WallComment represents the WallComment schema from the OpenAPI specification
type WallComment struct {
	Id string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Modified string `json:"modified,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Wall_post_id string `json:"wall_post_id,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
}

// Activity represents the Activity schema from the OpenAPI specification
type Activity struct {
	Name string `json:"name,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Hex_code string `json:"hex_code,omitempty"`
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// ExpenseFile represents the ExpenseFile schema from the OpenAPI specification
type ExpenseFile struct {
	File string `json:"file,omitempty"` // File's name
	File_extension string `json:"file_extension,omitempty"`
	Id string `json:"id,omitempty"` // Read-only
	Created string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Expense_id string `json:"expense_id,omitempty"`
	File_type string `json:"file_type,omitempty"` // File's MIME type
	File_original_name string `json:"file_original_name,omitempty"`
	File_size string `json:"file_size,omitempty"` // File size in bytes
	File_url string `json:"file_url,omitempty"` // Read-only
	Modified string `json:"modified,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
}

// CompaniesFormTemplates represents the CompaniesFormTemplates schema from the OpenAPI specification
type CompaniesFormTemplates struct {
	Modified string `json:"modified,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Form_template_id string `json:"form_template_id,omitempty"`
	Id string `json:"id,omitempty"`
}

// IntegrationSettingsUsers represents the IntegrationSettingsUsers schema from the OpenAPI specification
type IntegrationSettingsUsers struct {
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Integration_setting_id string `json:"integration_setting_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Value string `json:"value,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
}

// UnauthorizedError represents the UnauthorizedError schema from the OpenAPI specification
type UnauthorizedError struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SharedProjectContact represents the SharedProjectContact schema from the OpenAPI specification
type SharedProjectContact struct {
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Sent_email bool `json:"sent_email,omitempty"`
	Contact_id string `json:"contact_id,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Sent_sms bool `json:"sent_sms,omitempty"`
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Id string `json:"id,omitempty"` // Read-only
	Modified_by_id string `json:"modified_by_id,omitempty"` // Read-only
	Project_id string `json:"project_id,omitempty"`
	Start string `json:"start,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Description string `json:"description,omitempty"`
	Company_id string `json:"company_id,omitempty"` // Read-only
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	End string `json:"end,omitempty"`
	Created string `json:"created,omitempty"`
}

// FormTemplate represents the FormTemplate schema from the OpenAPI specification
type FormTemplate struct {
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Pdf_template_identifier string `json:"pdf_template_identifier,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Identifier string `json:"identifier,omitempty"`
	Is_active bool `json:"is_active,omitempty"`
	Description string `json:"description,omitempty"`
	Form_category_id string `json:"form_category_id,omitempty"`
	Form_overview_category_id string `json:"form_overview_category_id,omitempty"`
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
}

// SharedProjectVendor represents the SharedProjectVendor schema from the OpenAPI specification
type SharedProjectVendor struct {
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Sent_email bool `json:"sent_email,omitempty"`
	Sent_sms bool `json:"sent_sms,omitempty"`
	Vendor_id string `json:"vendor_id,omitempty"`
}

// EmptySuccessResponse represents the EmptySuccessResponse schema from the OpenAPI specification
type EmptySuccessResponse struct {
	Data []string `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
}

// InvoiceFile represents the InvoiceFile schema from the OpenAPI specification
type InvoiceFile struct {
	Id string `json:"id,omitempty"`
	Invoice_id string `json:"invoice_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	TypeField string `json:"type,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	File_id string `json:"file_id,omitempty"`
}

// ExpenseLine represents the ExpenseLine schema from the OpenAPI specification
type ExpenseLine struct {
	Text string `json:"text,omitempty"`
	Currency_id string `json:"currency_id,omitempty"`
	Expense_id string `json:"expense_id,omitempty"`
	Is_invoiced string `json:"is_invoiced,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"` // Read-only
	Quantity int `json:"quantity,omitempty"`
	Selling_price float32 `json:"selling_price,omitempty"`
	Buying_price float32 `json:"buying_price,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Modified string `json:"modified,omitempty"`
}

// FormFieldType represents the FormFieldType schema from the OpenAPI specification
type FormFieldType struct {
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"` // Read-only
	Identifier string `json:"identifier,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// PaymentTerm represents the PaymentTerm schema from the OpenAPI specification
type PaymentTerm struct {
	Days_of_credit int `json:"days_of_credit,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Payment_term_type_id string `json:"payment_term_type_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
}

// PlannedProduct represents the PlannedProduct schema from the OpenAPI specification
type PlannedProduct struct {
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Price float64 `json:"price,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Product_number string `json:"product_number,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	Vendor_product_id string `json:"vendor_product_id,omitempty"`
}

// CompanySettings represents the CompanySettings schema from the OpenAPI specification
type CompanySettings struct {
	Company_settings_category_id string `json:"company_settings_category_id,omitempty"`
	Default_value string `json:"default_value,omitempty"`
	TypeField string `json:"type,omitempty"`
	Modified string `json:"modified,omitempty"`
	Id string `json:"id,omitempty"`
	Feature_id string `json:"feature_id,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Options string `json:"options,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Description string `json:"description,omitempty"`
	Placement int `json:"placement,omitempty"`
	Created string `json:"created,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Modified string `json:"modified,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Start_time string `json:"start_time,omitempty"`
	Name string `json:"name"`
	Contact_id string `json:"contact_id,omitempty"`
	Products_total_cost_price float32 `json:"products_total_cost_price,omitempty"`
	Erp_project_id string `json:"erp_project_id,omitempty"`
	Is_offer string `json:"is_offer,omitempty"` // Is the project was offer
	Thumbnail string `json:"thumbnail,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Project_number float64 `json:"project_number,omitempty"`
	Street_name string `json:"street_name,omitempty"`
	Offer_id string `json:"offer_id,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Not_invoiced_amount float32 `json:"not_invoiced_amount,omitempty"`
	Latitude string `json:"latitude,omitempty"`
	Is_fixed_price bool `json:"is_fixed_price,omitempty"`
	Total_sales_price float32 `json:"total_sales_price,omitempty"`
	Created string `json:"created,omitempty"`
	Shared_project_id string `json:"shared_project_id,omitempty"`
	Is_rotten string `json:"is_rotten,omitempty"` // Last form date - read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id"`
	Project_image_url string `json:"project_image_url,omitempty"`
	Description string `json:"description,omitempty"`
	Erp_task_id string `json:"erp_task_id,omitempty"`
	Has_final_invoice bool `json:"has_final_invoice,omitempty"` // Is there at least one final invoice
	Full_name string `json:"full_name,omitempty"` // Project number + name
	Pre_calculation_id string `json:"pre_calculation_id,omitempty"`
	Parent_id string `json:"parent_id,omitempty"`
	Project_status_id string `json:"project_status_id,omitempty"`
	Working_hours_total_cost_price float32 `json:"working_hours_total_cost_price,omitempty"`
	City_id string `json:"city_id,omitempty"`
	End_time string `json:"end_time,omitempty"`
}

// ProjectCustomFieldAttribute represents the ProjectCustomFieldAttribute schema from the OpenAPI specification
type ProjectCustomFieldAttribute struct {
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Modified string `json:"modified,omitempty"`
	Placement int `json:"placement,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Is_active bool `json:"is_active,omitempty"`
	Name string `json:"name,omitempty"`
	Id string `json:"id,omitempty"`
	Access_type string `json:"access_type,omitempty"`
}

// City represents the City schema from the OpenAPI specification
type City struct {
	Name string `json:"name,omitempty"`
	Zip_code int `json:"zip_code,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// ContactPerson represents the ContactPerson schema from the OpenAPI specification
type ContactPerson struct {
	Id string `json:"id,omitempty"`
	Contact_id string `json:"contact_id,omitempty"`
	Created string `json:"created,omitempty"`
	Email string `json:"email,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Title string `json:"title,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
}

// ContactType represents the ContactType schema from the OpenAPI specification
type ContactType struct {
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"` // One of 3 values
	Modified string `json:"modified,omitempty"`
}

// BulkEditIntegrationSettingsUsersRequestBody represents the BulkEditIntegrationSettingsUsersRequestBody schema from the OpenAPI specification
type BulkEditIntegrationSettingsUsersRequestBody struct {
	Integration_settings []BulkEditIntegrationSettingsUsersItemsRequestBody `json:"integration_settings,omitempty"`
	User_id string `json:"user_id,omitempty"`
}

// BulkActionRequestBody represents the BulkActionRequestBody schema from the OpenAPI specification
type BulkActionRequestBody struct {
	Id []string `json:"id,omitempty"`
}

// InvoiceLineText represents the InvoiceLineText schema from the OpenAPI specification
type InvoiceLineText struct {
	Id string `json:"id,omitempty"`
	Invoice_line_id string `json:"invoice_line_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	File_id string `json:"file_id,omitempty"`
	Html string `json:"html,omitempty"`
}

// Role represents the Role schema from the OpenAPI specification
type Role struct {
	Identifier string `json:"identifier,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
}

// UserCustomFieldAttribute represents the UserCustomFieldAttribute schema from the OpenAPI specification
type UserCustomFieldAttribute struct {
	Is_active bool `json:"is_active,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	Placement int `json:"placement,omitempty"`
	Id string `json:"id,omitempty"`
	Access_type string `json:"access_type,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
}

// Material represents the Material schema from the OpenAPI specification
type Material struct {
	Tripletex_id string `json:"tripletex_id,omitempty"`
	Selling_price float32 `json:"selling_price,omitempty"`
	Barcode string `json:"barcode,omitempty"`
	Centiga_id string `json:"centiga_id,omitempty"`
	Is_single_usage bool `json:"is_single_usage,omitempty"`
	Cost_price float32 `json:"cost_price,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Pogo_id string `json:"pogo_id,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Description string `json:"description,omitempty"`
	Modified string `json:"modified,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Billing_cycle string `json:"billing_cycle,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ProjectStatusType represents the ProjectStatusType schema from the OpenAPI specification
type ProjectStatusType struct {
	Name string `json:"name,omitempty"`
	Locale string `json:"_locale,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Modified string `json:"modified,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Id string `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
}

// ProjectStatus represents the ProjectStatus schema from the OpenAPI specification
type ProjectStatus struct {
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"` // One of 3 values
	Is_custom bool `json:"is_custom,omitempty"`
	Name string `json:"name,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// Company represents the Company schema from the OpenAPI specification
type Company struct {
	Receive_form_mails string `json:"receive_form_mails,omitempty"`
	Next_project_number int `json:"next_project_number,omitempty"`
	Country_id string `json:"country_id,omitempty"`
	Expired string `json:"expired,omitempty"`
	Phone string `json:"phone,omitempty"` // Format like eg. `28680133` or `046158971404`
	City_id string `json:"city_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Id string `json:"id,omitempty"`
	Next_offer_number int `json:"next_offer_number,omitempty"`
	File_id string `json:"file_id,omitempty"`
	Street_name string `json:"street_name,omitempty"`
	Cvr string `json:"cvr,omitempty"`
	Invoice_email string `json:"invoice_email,omitempty"`
	Name string `json:"name,omitempty"`
	Website string `json:"website,omitempty"`
	Contact_person_id string `json:"contact_person_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Created string `json:"created,omitempty"`
	Language_id string `json:"language_id,omitempty"`
	Vat_percent int `json:"vat_percent,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Next_invoice_number int `json:"next_invoice_number,omitempty"`
	Phone_countrycode string `json:"phone_countrycode,omitempty"` // Format like eg. `45` or `49`
}

// ProductVariant represents the ProductVariant schema from the OpenAPI specification
type ProductVariant struct {
	Name string `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
	Product_number string `json:"product_number,omitempty"`
	Variant_id string `json:"variant_id,omitempty"`
	Variant_type string `json:"variant_type,omitempty"`
}

// TimeEntryType represents the TimeEntryType schema from the OpenAPI specification
type TimeEntryType struct {
	Modified string `json:"modified,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Time_entry_interval_id string `json:"time_entry_interval_id,omitempty"`
	Created string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Time_entry_value_type_id string `json:"time_entry_value_type_id,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// AddDefaultProjectStatusesSuccess represents the AddDefaultProjectStatusesSuccess schema from the OpenAPI specification
type AddDefaultProjectStatusesSuccess struct {
	Success bool `json:"success,omitempty"`
}

// Vendor represents the Vendor schema from the OpenAPI specification
type Vendor struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Cvr string `json:"cvr,omitempty"`
	Is_custom bool `json:"is_custom,omitempty"`
	Modified string `json:"modified,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Email string `json:"email,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

// ErrorNotFound represents the ErrorNotFound schema from the OpenAPI specification
type ErrorNotFound struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
}

// Expense represents the Expense schema from the OpenAPI specification
type Expense struct {
	Total_selling_price float32 `json:"total_selling_price,omitempty"` // Sum of all `selling_price` from expense lines
	Delivery_date string `json:"delivery_date,omitempty"`
	Short_text string `json:"short_text,omitempty"`
	Contact_id string `json:"contact_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	File_reference string `json:"file_reference,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Supplier_invoice_number string `json:"supplier_invoice_number,omitempty"`
	Order_number string `json:"order_number,omitempty"`
	Reference string `json:"reference,omitempty"`
	Activity_id string `json:"activity_id,omitempty"`
	Created string `json:"created,omitempty"`
	Currency_id string `json:"currency_id,omitempty"`
	Comment string `json:"comment,omitempty"`
	Status string `json:"status,omitempty"`
	Roger_id string `json:"roger_id,omitempty"`
	Is_imported string `json:"is_imported,omitempty"`
	Modified string `json:"modified,omitempty"`
	Total_buying_price float32 `json:"total_buying_price,omitempty"` // Sum of all `buying_price` from expense lines
	Description string `json:"description,omitempty"`
	Due_date string `json:"due_date,omitempty"`
	Id string `json:"id,omitempty"` // Read-only
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Readsoft_id string `json:"readsoft_id,omitempty"`
	Sent_to_email string `json:"sent_to_email,omitempty"`
	Company_id string `json:"company_id,omitempty"` // Read-only
}

// TimeEntryUnitType represents the TimeEntryUnitType schema from the OpenAPI specification
type TimeEntryUnitType struct {
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
}

// AddDefaultProjectStatusesError represents the AddDefaultProjectStatusesError schema from the OpenAPI specification
type AddDefaultProjectStatusesError struct {
	Success bool `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

// IntegrationFeatureSetting represents the IntegrationFeatureSetting schema from the OpenAPI specification
type IntegrationFeatureSetting struct {
	Default_value string `json:"default_value,omitempty"`
	Name string `json:"name,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Is_custom_setting bool `json:"is_custom_setting,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Id string `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Integration_feature_id string `json:"integration_feature_id,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Is_sent bool `json:"is_sent,omitempty"`
	Reply_to string `json:"reply_to,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Recipients string `json:"recipients,omitempty"`
	Api_response string `json:"api_response,omitempty"`
	Carbon_copy string `json:"carbon_copy,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Id string `json:"id,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body string `json:"body,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
}

// SubscriptionSelfServiceRequestBody represents the SubscriptionSelfServiceRequestBody schema from the OpenAPI specification
type SubscriptionSelfServiceRequestBody struct {
	Subscription_self_service_url string `json:"subscription_self_service_url,omitempty"`
}

// TimeEntry represents the TimeEntry schema from the OpenAPI specification
type TimeEntry struct {
	Time_entry_type_id string `json:"time_entry_type_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Id string `json:"id,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Sum int `json:"sum,omitempty"` // Amount of seconds
	Created string `json:"created,omitempty"`
	Form_id string `json:"form_id,omitempty"`
	From_time string `json:"from_time,omitempty"`
	Modified string `json:"modified,omitempty"`
	To_time string `json:"to_time,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Is_all_day bool `json:"is_all_day,omitempty"`
}

// CompaniesIntegrationSetting represents the CompaniesIntegrationSetting schema from the OpenAPI specification
type CompaniesIntegrationSetting struct {
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Integration_setting_id string `json:"integration_setting_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Value string `json:"value,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
}

// ErrorValidation represents the ErrorValidation schema from the OpenAPI specification
type ErrorValidation struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
}

// Product represents the Product schema from the OpenAPI specification
type Product struct {
	Selling_price float64 `json:"selling_price,omitempty"`
	Created string `json:"created,omitempty"`
	Erp_id string `json:"erp_id,omitempty"`
	Name string `json:"name,omitempty"`
	Project_status_type_id string `json:"project_status_type_id,omitempty"`
	Centiga_id string `json:"centiga_id,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Barcode string `json:"barcode,omitempty"`
	Product_number string `json:"product_number,omitempty"`
	Buying_price float64 `json:"buying_price,omitempty"`
	Pogo_id string `json:"pogo_id,omitempty"`
	Id string `json:"id,omitempty"`
	Tripletex_id string `json:"tripletex_id,omitempty"`
	Description string `json:"description,omitempty"`
	Modified string `json:"modified,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Average_cost_price float64 `json:"average_cost_price,omitempty"`
}

// Form represents the Form schema from the OpenAPI specification
type Form struct {
	Is_invoiced string `json:"is_invoiced,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Created string `json:"created,omitempty"`
	Is_draft bool `json:"is_draft,omitempty"`
	Is_shared bool `json:"is_shared,omitempty"`
	Mass_form_id string `json:"mass_form_id,omitempty"`
	Approved_by_id string `json:"approved_by_id,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Form_date string `json:"form_date,omitempty"`
	Form_template_id string `json:"form_template_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"` // Read-only
	Modified string `json:"modified,omitempty"`
}

// StockLocation represents the StockLocation schema from the OpenAPI specification
type StockLocation struct {
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
}

// Countries represents the Countries schema from the OpenAPI specification
type Countries struct {
	Modified string `json:"modified,omitempty"`
	Phone_code string `json:"phone_code,omitempty"`
	Currency_id string `json:"currency_id,omitempty"`
	Language_id string `json:"language_id,omitempty"`
	Name string `json:"name,omitempty"`
	Time_zone string `json:"time_zone,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

// Contact represents the Contact schema from the OpenAPI specification
type Contact struct {
	Pogo_id string `json:"pogo_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Phone string `json:"phone,omitempty"` // Format like eg. `28680133` or `046158971404`
	Company_id string `json:"company_id,omitempty"` // Only filled out if this represents another company within Apacta (used for partners)
	Id string `json:"id,omitempty"`
	Centiga_id string `json:"centiga_id,omitempty"`
	Country_id string `json:"country_id,omitempty"`
	Description string `json:"description,omitempty"`
	Tripletex_id string `json:"tripletex_id,omitempty"`
	Erp_id string `json:"erp_id,omitempty"` // If company has integration to an ERP system, and the contacts are synchronized, this will be the ERP-systems ID of this contact
	Created string `json:"created,omitempty"`
	Cvr string `json:"cvr,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	City_id string `json:"city_id,omitempty"`
	Address string `json:"address,omitempty"` // Street address
	Website string `json:"website,omitempty"`
	Email string `json:"email,omitempty"`
}

// PaymentTermsData represents the PaymentTermsData schema from the OpenAPI specification
type PaymentTermsData struct {
	TypeField string `json:"type,omitempty"`
	Daysofcredit int `json:"daysOfCredit,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Pretty_string string `json:"pretty_string,omitempty"`
}

// TimeEntryInterval represents the TimeEntryInterval schema from the OpenAPI specification
type TimeEntryInterval struct {
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Identifier string `json:"identifier,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
}

// MassMessage represents the MassMessage schema from the OpenAPI specification
type MassMessage struct {
	Company_id string `json:"company_id,omitempty"`
	Content string `json:"content,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// VendorProductPriceFile represents the VendorProductPriceFile schema from the OpenAPI specification
type VendorProductPriceFile struct {
	Vendor_products_count int `json:"vendor_products_count,omitempty"`
	Id string `json:"id,omitempty"`
	File string `json:"file,omitempty"`
	Dir string `json:"dir,omitempty"`
	Progress int `json:"progress,omitempty"`
	Vendor_products_count_total int `json:"vendor_products_count_total,omitempty"`
	Status string `json:"status,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Finished bool `json:"finished,omitempty"`
	Original_file_name string `json:"original_file_name,omitempty"`
	Companies_vendor_id string `json:"companies_vendor_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	TypeField string `json:"type,omitempty"`
	Size int `json:"size,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
}

// WallPost represents the WallPost schema from the OpenAPI specification
type WallPost struct {
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Modified string `json:"modified,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	User_id string `json:"user_id,omitempty"`
}

// Currency represents the Currency schema from the OpenAPI specification
type Currency struct {
	Created string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Tripletex_id string `json:"tripletex_id,omitempty"`
	Currency_sign string `json:"currency_sign,omitempty"`
	Pogo_id string `json:"pogo_id,omitempty"`
	Centiga_id string `json:"centiga_id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

// FormField represents the FormField schema from the OpenAPI specification
type FormField struct {
	Form_template_field_id string `json:"form_template_field_id,omitempty"`
	Placement int `json:"placement,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Modified string `json:"modified,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Comment string `json:"comment,omitempty"`
	Created string `json:"created,omitempty"`
	File_id string `json:"file_id,omitempty"`
	Form_id string `json:"form_id,omitempty"`
	Id string `json:"id,omitempty"` // Read-only
	Content_value string `json:"content_value,omitempty"`
	Form_field_type_id string `json:"form_field_type_id,omitempty"`
}

// OfferStatus represents the OfferStatus schema from the OpenAPI specification
type OfferStatus struct {
	Identifier string `json:"identifier,omitempty"` // One of 6 values
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Is_custom bool `json:"is_custom,omitempty"`
	Modified string `json:"modified,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Name string `json:"name,omitempty"`
	Id string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

// Offer represents the Offer schema from the OpenAPI specification
type Offer struct {
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Description string `json:"description,omitempty"`
	Address string `json:"address,omitempty"` // Street address
	Issue_date string `json:"issue_date,omitempty"`
	Slug string `json:"slug,omitempty"`
	Offer_number int `json:"offer_number,omitempty"`
	Sender_id string `json:"sender_id,omitempty"`
	City_id string `json:"city_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Id string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Modified string `json:"modified,omitempty"`
	Show_products_product_bundle bool `json:"show_products_product_bundle,omitempty"`
	Discount_percent int `json:"discount_percent,omitempty"`
	Vat_percent int `json:"vat_percent,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Payment_term_id string `json:"payment_term_id,omitempty"`
	Show_prices bool `json:"show_prices,omitempty"`
	Title string `json:"title,omitempty"`
	Offer_status_id string `json:"offer_status_id,omitempty"`
	Expiraton_date string `json:"expiraton_date,omitempty"`
	Show_offer_lines bool `json:"show_offer_lines,omitempty"`
	All_lines_one_line bool `json:"all_lines_one_line,omitempty"`
	Show_payment_term bool `json:"show_payment_term,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Show_product_images bool `json:"show_product_images,omitempty"`
	All_working_hours_one_line bool `json:"all_working_hours_one_line,omitempty"`
	Erp_payment_term_id string `json:"erp_payment_term_id,omitempty"`
	Contact_id string `json:"contact_id,omitempty"`
	Rejection_reason string `json:"rejection_reason,omitempty"`
	All_products_one_line bool `json:"all_products_one_line,omitempty"`
	Show_employee_name bool `json:"show_employee_name,omitempty"`
	Created string `json:"created,omitempty"`
}

// CreateSuccessResponse represents the CreateSuccessResponse schema from the OpenAPI specification
type CreateSuccessResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
}

// Invoice represents the Invoice schema from the OpenAPI specification
type Invoice struct {
	Invoice_number int `json:"invoice_number,omitempty"`
	All_products_one_line bool `json:"all_products_one_line,omitempty"`
	Gross_payment float32 `json:"gross_payment,omitempty"`
	Reference string `json:"reference,omitempty"`
	Show_products_product_bundle bool `json:"show_products_product_bundle,omitempty"`
	Group_by_forms bool `json:"group_by_forms,omitempty"`
	All_working_hours_one_line bool `json:"all_working_hours_one_line,omitempty"`
	Created string `json:"created,omitempty"`
	Downloaded string `json:"downloaded,omitempty"`
	Offer_number int `json:"offer_number,omitempty"`
	Eu_customer bool `json:"eu_customer,omitempty"`
	Issued_date string `json:"issued_date,omitempty"`
	Is_final_invoice bool `json:"is_final_invoice,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Show_product_images bool `json:"show_product_images,omitempty"`
	Date_to string `json:"date_to,omitempty"`
	Contact_id string `json:"contact_id,omitempty"`
	Erp_payment_term_id string `json:"erp_payment_term_id,omitempty"`
	Project_overview_attached bool `json:"project_overview_attached,omitempty"`
	Include_invoiced_lines bool `json:"include_invoiced_lines,omitempty"`
	Payment_due_date string `json:"payment_due_date,omitempty"`
	Erp_id string `json:"erp_id,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Total_cost_price float32 `json:"total_cost_price,omitempty"`
	Total_discount_percent float32 `json:"total_discount_percent,omitempty"`
	Show_price_product_bundle bool `json:"show_price_product_bundle,omitempty"`
	Show_prices_products_and_hours bool `json:"show_prices_products_and_hours,omitempty"`
	Vat_percent int `json:"vat_percent,omitempty"`
	Order_line_group_id string `json:"order_line_group_id,omitempty"`
	Show_employee_name bool `json:"show_employee_name,omitempty"`
	Title string `json:"title,omitempty"`
	Modified string `json:"modified,omitempty"`
	Net_payment float32 `json:"net_payment,omitempty"`
	Currency_id string `json:"currency_id,omitempty"`
	Is_offer bool `json:"is_offer,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Integration_id string `json:"integration_id,omitempty"`
	Payment_term_id string `json:"payment_term_id,omitempty"`
	Date_from string `json:"date_from,omitempty"`
	Is_locked bool `json:"is_locked,omitempty"`
	Is_draft bool `json:"is_draft,omitempty"`
	Vendor_id string `json:"vendor_id,omitempty"`
}

// PaymentTermType represents the PaymentTermType schema from the OpenAPI specification
type PaymentTermType struct {
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
}

// MaterialRental represents the MaterialRental schema from the OpenAPI specification
type MaterialRental struct {
	Modified string `json:"modified,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
	To_date string `json:"to_date,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	From_date string `json:"from_date,omitempty"`
	Id string `json:"id,omitempty"`
	Material_id string `json:"material_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Project_id string `json:"project_id,omitempty"`
	Is_invoiced string `json:"is_invoiced,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
}

// PaginationDetails represents the PaginationDetails schema from the OpenAPI specification
type PaginationDetails struct {
	Has_prev_page bool `json:"has_prev_page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Page_count string `json:"page_count,omitempty"`
	Count int `json:"count,omitempty"`
	Current_page string `json:"current_page,omitempty"`
	Has_next_page bool `json:"has_next_page,omitempty"`
}

// Sender represents the Sender schema from the OpenAPI specification
type Sender struct {
	Cvr string `json:"cvr,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Name string `json:"name,omitempty"`
	Language_id string `json:"language_id,omitempty"`
	Id string `json:"id,omitempty"`
	Phone_countrycode string `json:"phone_countrycode,omitempty"` // Format like eg. `45` or `49`
	Website string `json:"website,omitempty"`
	City_id string `json:"city_id,omitempty"`
	Created string `json:"created,omitempty"`
	Email string `json:"email,omitempty"`
	Modified string `json:"modified,omitempty"`
	Phone string `json:"phone,omitempty"` // Format like eg. `28680133` or `046158971404`
	Street_name string `json:"street_name,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
}

// VendorProduct represents the VendorProduct schema from the OpenAPI specification
type VendorProduct struct {
	Name string `json:"name,omitempty"`
	Vendor_id string `json:"vendor_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Product_category_number string `json:"product_category_number,omitempty"`
	Barcode string `json:"barcode,omitempty"`
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Price float64 `json:"price,omitempty"`
	Product_number string `json:"product_number,omitempty"`
	Created string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
}

// ClockingRecord represents the ClockingRecord schema from the OpenAPI specification
type ClockingRecord struct {
	Checkin_latitude string `json:"checkin_latitude,omitempty"`
	Modified string `json:"modified,omitempty"`
	Checkin_longitude string `json:"checkin_longitude,omitempty"`
	Checkout_latitude string `json:"checkout_latitude,omitempty"`
	Checked_out string `json:"checked_out,omitempty"`
	Checkout_longitude string `json:"checkout_longitude,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Checked_in string `json:"checked_in,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Project_id string `json:"project_id,omitempty"`
}

// ProjectCustomFieldValue represents the ProjectCustomFieldValue schema from the OpenAPI specification
type ProjectCustomFieldValue struct {
	Value string `json:"value,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"` // Read-only
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Project_custom_field_attribute_id string `json:"project_custom_field_attribute_id,omitempty"`
	Project_id string `json:"project_id,omitempty"`
}

// CompanyPriceMargins represents the CompanyPriceMargins schema from the OpenAPI specification
type CompanyPriceMargins struct {
	Modified string `json:"modified,omitempty"`
	Percentage_ratio float64 `json:"percentage_ratio,omitempty"`
	Amount_from float64 `json:"amount_from,omitempty"`
	Amount_to float64 `json:"amount_to,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Ratio float32 `json:"ratio,omitempty"`
	TypeField string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Phone string `json:"phone,omitempty"`
	Full_name string `json:"full_name,omitempty"` // READ-ONLY
	Phone_countrycode string `json:"phone_countrycode,omitempty"`
	Cost_price float32 `json:"cost_price,omitempty"` // Cost of salaries
	Api_key string `json:"api_key,omitempty"`
	Email string `json:"email,omitempty"`
	Mobile_countrycode string `json:"mobile_countrycode,omitempty"`
	Password string `json:"password,omitempty"`
	Website string `json:"website,omitempty"`
	Id string `json:"id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Company_id string `json:"company_id,omitempty"`
	Hide_phone bool `json:"hide_phone,omitempty"`
	Language_id string `json:"language_id,omitempty"`
	Is_active bool `json:"is_active,omitempty"`
	Last_name string `json:"last_name,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Sale_price float32 `json:"sale_price,omitempty"` // The price this employee costs per hour when working
	Street_name string `json:"street_name,omitempty"`
	Extra_price float32 `json:"extra_price,omitempty"` // Additional cost on this employee (pension, vacation etc.)
	Hide_address bool `json:"hide_address,omitempty"`
	Receive_form_mails bool `json:"receive_form_mails,omitempty"` // If `true` the employee will receive an email receipt of every form submitted
	First_name string `json:"first_name,omitempty"`
	Initials string `json:"initials,omitempty"`
	Expected_billable_hours float32 `json:"expected_billable_hours,omitempty"`
	City_id string `json:"city_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
}

// TimeEntryRate represents the TimeEntryRate schema from the OpenAPI specification
type TimeEntryRate struct {
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Created_by_id string `json:"created_by_id,omitempty"`
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Selling_amount float32 `json:"selling_amount,omitempty"`
	Time_entry_type_id string `json:"time_entry_type_id,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	Currency_id string `json:"currency_id,omitempty"`
}

// CompaniesVendor represents the CompaniesVendor schema from the OpenAPI specification
type CompaniesVendor struct {
	Delivery_price float32 `json:"delivery_price,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Vendor_department_id string `json:"vendor_department_id,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Company_password string `json:"company_password,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Is_active bool `json:"is_active,omitempty"`
	Vendor_account_reference string `json:"vendor_account_reference,omitempty"`
	Username string `json:"username,omitempty"`
	Company_id string `json:"company_id,omitempty"`
	Use_price_files bool `json:"use_price_files,omitempty"`
	Vendor Vendor `json:"vendor,omitempty"`
	Free_delivery_price float32 `json:"free_delivery_price,omitempty"`
	Receive_invoice_mails bool `json:"receive_invoice_mails,omitempty"`
	Reviewed bool `json:"reviewed,omitempty"`
	Vendor_email string `json:"vendor_email,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Receive_automatic_price_files bool `json:"receive_automatic_price_files,omitempty"`
	Vendor_name string `json:"vendor_name,omitempty"`
	Id string `json:"id,omitempty"`
	Vendor_id string `json:"vendor_id,omitempty"`
}

// IntegrationSettingsUser represents the IntegrationSettingsUser schema from the OpenAPI specification
type IntegrationSettingsUser struct {
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Integration_setting_id string `json:"integration_setting_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Value string `json:"value,omitempty"`
	Created string `json:"created,omitempty"`
}

// MassMessagesUser represents the MassMessagesUser schema from the OpenAPI specification
type MassMessagesUser struct {
	User_id string `json:"user_id,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Is_sent_email bool `json:"is_sent_email,omitempty"`
	Modified string `json:"modified,omitempty"`
	Id string `json:"id,omitempty"`
	Is_read bool `json:"is_read,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Mass_message MassMessage `json:"mass_message,omitempty"`
	Mass_message_id string `json:"mass_message_id,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
}

// InvoiceLineTextTemplate represents the InvoiceLineTextTemplate schema from the OpenAPI specification
type InvoiceLineTextTemplate struct {
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Html string `json:"html,omitempty"`
	Id string `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	Modified string `json:"modified,omitempty"`
	Company_id string `json:"company_id,omitempty"`
}

// BulkEditIntegrationSettingsUsersItemsRequestBody represents the BulkEditIntegrationSettingsUsersItemsRequestBody schema from the OpenAPI specification
type BulkEditIntegrationSettingsUsersItemsRequestBody struct {
	Integration_setting_id string `json:"integration_setting_id,omitempty"`
	Value string `json:"value,omitempty"`
}

// TimeEntryValueType represents the TimeEntryValueType schema from the OpenAPI specification
type TimeEntryValueType struct {
	Description string `json:"description,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Created string `json:"created,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	Time_entry_unit_type_id string `json:"time_entry_unit_type_id,omitempty"`
	Id string `json:"id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
}

// DrivingType represents the DrivingType schema from the OpenAPI specification
type DrivingType struct {
	Created_by_id string `json:"created_by_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Employee_price float64 `json:"employee_price"`
	Erp_id string `json:"erp_id,omitempty"`
	Id string `json:"id"`
	Invoice_price float64 `json:"invoice_price"`
	Name string `json:"name"`
	Company_id string `json:"company_id"`
	Created string `json:"created"`
	Modified string `json:"modified"`
	Salary_id string `json:"salary_id,omitempty"`
}

// OfferLine represents the OfferLine schema from the OpenAPI specification
type OfferLine struct {
	Offer_id string `json:"offer_id,omitempty"`
	Product_bundle_id string `json:"product_bundle_id,omitempty"`
	Deleted string `json:"deleted,omitempty"` // Only present if it's a deleted object
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Sales_price float32 `json:"sales_price,omitempty"`
	Created string `json:"created,omitempty"`
	Created_by_id string `json:"created_by_id,omitempty"`
	Modified string `json:"modified,omitempty"`
	Modified_by_id string `json:"modified_by_id,omitempty"`
	Description string `json:"description,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}
