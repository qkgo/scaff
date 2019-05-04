package cfg

import (
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	ozzo "github.com/go-ozzo/ozzo-config"
	"strconv"
	"sync"
)

var Configuration map[string]interface{}

var Log *logrus.Logger

var LogInfo *logrus.Logger

var LogHttp *logrus.Logger

var LogKey *logrus.Logger

var SqlLog *logrus.Logger

var ManualLog *logrus.Logger

var DataLog *logrus.Logger

var ServerBootStrapId = uuid.Must(uuid.NewV4())

func GetConfig() *ozzo.Config {
	return ConfigParam
}

var ConnectSocketInfo = sync.Map{}

var ConfigParam = ozzo.New()

var AliOssConfig []string
var CanvasUrl string

var DefaultPageSize = strconv.Itoa(10)

var OzConfig OzzConfig

var RabbitServerUrl string
