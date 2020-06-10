package models

type DomainInformation struct {
  Servers []Server
  ServersChanged bool
  SSLGrade string
  PreviousSSLGrade string
  Logo string
  Title string
  Is_down bool
  Status string
  First bool
}
