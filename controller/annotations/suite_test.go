package annotations

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/api"
)

type AnnotationSuite struct {
	suite.Suite
	client         api.HAProxyClient
	transactionDir string
}

func (suite *AnnotationSuite) SetupSuite() {
	var err error
	suite.transactionDir, err = ioutil.TempDir("/tmp/", "controller-tests")
	if err != nil {
		logger.Panic(err)
	}
	suite.client, err = api.Init(suite.transactionDir, "../../fs/etc/haproxy/haproxy.cfg", "", "")
	suite.Nil(err)
	suite.NoError(suite.client.APIStartTransaction())
}

func (suite *AnnotationSuite) TearDownSuite() {
	os.RemoveAll(suite.transactionDir)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(AnnotationSuite))
}
