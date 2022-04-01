package cfg

import (
	ozzo "github.com/go-ozzo/ozzo-config"
	"github.com/gofrs/uuid"
	"github.com/qkgo/scaff/pkg/log"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

var Configuration map[string]interface{}

var Log *logrus.Logger

var LogDebug *logrus.Logger

var LogInfo *logrus.Logger

var LogHttp *logrus.Logger

var LogRpc *logrus.Logger

var LogKey *logrus.Logger

var SqlLog *logrus.Logger

var XOrmLog *log.SqlLogger

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
