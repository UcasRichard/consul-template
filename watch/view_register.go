package watch
import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	dep "github.com/hashicorp/consul-template/dependency"
	"time"
)

const(
	REGISTER_PREFIX = "listener/"
)

type DependencyRegister struct{
	config *RegisterConfig
	uuid string
}

type RegisterConfig struct {
	Addr string
	Port string
}

func NewDependencyRegister(registerConfig *RegisterConfig) *DependencyRegister{
	return &DependencyRegister{
		config: registerConfig,
		uuid: time.Now().String(),
	}
}

func (register *DependencyRegister) register(v *View) error{
	clients := v.config.Clients
	consul, err := clients.Consul()
	if err != nil {
		return fmt.Errorf("register dependency: error getting client %s", err)
	}
	stores := consul.KV()
	key, err := register.genenateKey(v.Dependency)
	if err != nil{
		log.Printf("%s", err)
		return nil
	}
	value := register.config.value()
	p := &api.KVPair{Key: key, Flags: 0, Value: []byte(value)}
	if _, err = stores.Put(p, nil); err != nil {
		log.Printf("Invalid key not detected: %s", key)
	}
	return err
}

func (register *DependencyRegister) genenateKey(d dep.Dependency)(string, error){
	result := REGISTER_PREFIX
	if storekey, ok := d.(*dep.StoreKey); ok {
		result += storekey.Path + "/" + register.uuid
		return result, nil
	}else{
		return "", fmt.Errorf("now can't gennenate key for %s", d.Display())
	}
}

func (config *RegisterConfig) value() string{
	return config.Addr + ":" + config.Port
}

func (v *View) register() error{
	clients := v.config.Clients
	_, err := clients.Consul()
	if err != nil {
		return fmt.Errorf("register dependency: error getting client %s", err)
	}

	return err
}