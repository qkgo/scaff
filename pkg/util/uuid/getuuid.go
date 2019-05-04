package uuid

import (
	"github.com/gofrs/uuid"
)



func GetUuid(){
	uuid.Must(uuid.NewV4())
}
