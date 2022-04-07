package main

import (
	"context"
	"k8s.io/klog/v2"
	"log"

	ofctx "github.com/lixf311/fc-framework-go/context"
	"github.com/lixf311/fc-framework-go/framework"
	"github.com/lixf311/fc-framework-go/plugin"
	"github.com/lixf311/fc-framework-go/plugin/skywalking"
)

func topicFunction(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	if in != nil {
		log.Printf("pubsub - Data: %s", in)
	} else {
		log.Print("pubsub - Data: Received")
	}
	return ctx.ReturnOnSuccess().WithData([]byte("hello there")), nil
}

func main() {
	ctx := context.Background()
	fwk, err := framework.NewFramework()
	if err != nil {
		klog.Fatal(err)
	}
	fwk.RegisterPlugins(map[string]plugin.Plugin{
		"skywalking": &skywalking.PluginSkywalking{},
	})

	err = fwk.Register(ctx, topicFunction)
	if err != nil {
		klog.Fatal(err)
	}

	err = fwk.Start(ctx)
	if err != nil {
		klog.Fatal(err)
	}
}
