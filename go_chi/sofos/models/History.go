package models

import (
  _ "github.com/lib/pq"
)

type History struct{
  Queries []Query
}
