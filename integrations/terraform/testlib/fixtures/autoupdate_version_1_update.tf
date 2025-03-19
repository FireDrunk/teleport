resource "teleport_autoupdate_version" "test" {
  version = "v1"
  spec = {
    tools = {
      target_version = "1.2.4"
    }
    agents = {
      start_version  = "1.2.3"
      target_version = "1.2.5"
      schedule       = "regular"
      mode           = "enabled"
    }
  }
}
