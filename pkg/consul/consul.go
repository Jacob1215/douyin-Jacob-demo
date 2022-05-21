package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {//接口里面这个东西已经被实现了，所以后面的NewRegisterClient可以返回struct格式的RegistryClient
	Register(address string,port int,name string,tags []string, id string,check *api.AgentServiceCheck) error
	DeRegister(serviceId string)error
}

func NewRegistryClient(host string,port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

//服务注册
func (r *Registry)Register(address string,port int,name string,tags []string, id string,check *api.AgentServiceCheck) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",r.Host,r.Port)

	client,err := api.NewClient(cfg)
	if err!=nil{
		panic(err)
	}
	//生成注册对象
	registration:= new(api.AgentServiceRegistration)
	registration.Name=name
	registration.ID= id
	registration.Port=port
	registration.Tags=tags
	registration.Address=address
	//生成对应的检查对象。
	registration.Check= check//检查
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *Registry) DeRegister(serviceId string)error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",r.Host,r.Port)

	client,err := api.NewClient(cfg)
	if err!=nil{
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err

}