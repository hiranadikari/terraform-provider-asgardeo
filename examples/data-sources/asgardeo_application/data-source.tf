# Look up an existing application by name.
data "asgardeo_application" "myapp" {
  name = "My Existing App"
}

output "existing_app_client_id" {
  value = data.asgardeo_application.myapp.client_id
}
