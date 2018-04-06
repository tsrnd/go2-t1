package monitoring

import (
	"testing"

	"time"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/stretchr/testify/assert"
	"github.com/tsrnd/trainning/infrastructure"
)

func TestSetupMonitoringNormal(t *testing.T) {
	infrastructure.SetConfig("monitoring.enable", true)
	infrastructure.SetConfig("monitoring.intervalSec", 0)
	Setup(infrastructure.NewLoggerWithType("monitoring"))
}

func TestLogProfileNormal(t *testing.T) {
	infrastructure.SetConfig("monitoring.intervalSec", 1)
	Setup(infrastructure.NewLoggerWithType("monitoring"))
	time.Sleep(1 * time.Second)
}

func TestSetupMonitoringDisable(t *testing.T) {
	infrastructure.SetConfig("monitoring.enable", false)
	Setup(infrastructure.NewLoggerWithType("monitoring"))
}

func TestInitInfoOfInstance(t *testing.T) {
	setLogger(infrastructure.NewLoggerWithType("monitoring"))

	document := ec2metadata.EC2InstanceIdentityDocument{
		InstanceID: "test-instance",
		Region:     "test-region",
	}
	err := initInfo(true, document)
	assert.NoError(t, err)
}

func TestSetupMonitoringFailedInstanceIDNotFound(t *testing.T) {
	setLogger(infrastructure.NewLoggerWithType("monitoring"))

	document := ec2metadata.EC2InstanceIdentityDocument{
		InstanceID: "",
		Region:     "test-region",
	}
	err := initInfo(true, document)
	assert.Error(t, err)
	assert.Regexp(t, "can not find instanceID$", err.Error())
}

func TestSetupMonitoringFailedRegionNotFound(t *testing.T) {
	setLogger(infrastructure.NewLoggerWithType("monitoring"))

	document := ec2metadata.EC2InstanceIdentityDocument{
		InstanceID: "test-instance",
		Region:     "",
	}
	err := initInfo(true, document)
	assert.Error(t, err)
	assert.Regexp(t, "can not find region$", err.Error())
}
