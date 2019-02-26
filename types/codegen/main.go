package main

import (
	"github.com/jonnyfiveiq/monitor_namespace/types/apis/some.api.group/v1"
	"github.com/rancher/norman/generator"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := generator.DefaultGenerate(v1.Schemas, "github.com/jonnyfiveiq/monitor_namespace/types", false, nil); err != nil {
		logrus.Fatal(err)
	}
}
