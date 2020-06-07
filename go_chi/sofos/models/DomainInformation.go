package models

type DomainInformation struct {
  Servers []Server
  Servers_changed bool
  SSL_grade string
  Previous_SSL_grade string
  Logo string
  Title string
  Is_down bool
  Status string
  First bool
}
