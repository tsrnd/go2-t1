package monitoring

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/utils"
)

const (
	// DefaultIntervalSec default value
	DefaultIntervalSec = 15 // 3sec
)

var (
	logger     *infrastructure.Logger
	instanceID string
	region     string
)

func setLogger(l *infrastructure.Logger) {
	logger = l
}

func initInfo(available bool, doc ec2metadata.EC2InstanceIdentityDocument) error {
	var err error

	if available { // for aws
		logger.Log.Info("init onEC2")
		instanceID = doc.InstanceID
		if instanceID == "" {
			return utils.ErrorsNew("can not find instanceID")

		}
		region = doc.Region
		if region == "" {
			return utils.ErrorsNew("can not find region")
		}
	} else { // for local PC
		logger.Log.Info("init on local environment")
		instanceID, err = os.Hostname()
		if err != nil {
			return utils.ErrorsWrap(err, "can not get hostname")
		}
		region = ""
	}
	return nil
}

// Setup setup monitoring
func Setup(l *infrastructure.Logger) {
	setLogger(l)

	enable := infrastructure.GetConfigBool("monitoring.enable")
	if !enable {
		logger.Log.Info("monitoring log is disabled")
		return
	}

	// Init infomation of instance
	sess := session.Must(session.NewSession())
	svc := ec2metadata.New(sess)
	doc, _ := svc.GetInstanceIdentityDocument()
	if err := initInfo(svc.Available(), doc); err != nil {
		panic(err)
	}

	if infrastructure.GetConfigInt("monitoring.intervalSec") == 0 {
		infrastructure.SetConfig("monitoring.intervalSec", DefaultIntervalSec)
	}

	go logProfile(time.Duration(infrastructure.GetConfigInt64("monitoring.intervalSec")) * time.Second)

	logger.Log.Info("setup monitoring is done")
}

func logProfile(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			logger.Log.Info(
				fmt.Sprintf("monitoring data: instanceID=%s region=%s goroutine=%d heap=%d threadCreate=%d block=%d mutex=%d",
					instanceID, region, countGoroutine(), countHeap(), countThreadCreate(), countBlock(), countMutex()))
			break
		}
	}
}
