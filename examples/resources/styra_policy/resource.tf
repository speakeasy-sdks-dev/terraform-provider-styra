resource "styra_policy" "my_policy" {
  if_none_match = "...my_if_none_match..."
  modules = {
    "see" : "documentation",
  }
  policy = "...my_policy..."
  signature = {
    excluded   = "{ \"see\": \"documentation\" }"
    signatures = "{ \"see\": \"documentation\" }"
  }
}