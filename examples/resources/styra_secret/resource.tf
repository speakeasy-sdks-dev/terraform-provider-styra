resource "styra_secret" "my_secret" {
  description   = "...my_description..."
  if_none_match = "...my_if_none_match..."
  name          = "...my_name..."
  secret        = "...my_secret..."
  secret_id     = "...my_secret_id..."
}