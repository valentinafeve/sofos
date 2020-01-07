package info

type Server struct {
  Address string
  SSL_grade string
  Country string
  Owner string
  Old bool
}

type Info struct {
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

type Query struct {
  Domain string
  Time string
}

type History struct{
  Queries []Query
}
