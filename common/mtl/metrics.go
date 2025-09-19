package mtl

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

var (
	Registry *prometheus.Registry
)

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())                                       // 注册Go相关指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})) // 注册进程相关指标

	// 将Prometheus服务注册到consul，这样Prometheus能直接知道项目拥有的服务
	r, _ := consul.NewConsulRegister(registryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registryinfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryinfo)

	// 服务关闭前清�?
	server.RegisterShutdownHook(func() {
		r.Deregister(registryinfo)
	})

	// 官网用法 通过访问/metrics来查看Prometheus界面
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)

	// 为了hertz能注�?
	return r, registryinfo
}
