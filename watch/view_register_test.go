package watch
import (
	"testing"
	"fmt"
)

func TestNewRegister(t *testing.T){
	config := &RegisterConfig{
		Addr: "localhost",
		Port: "9090",
	}
	register := NewDependencyRegister(config)
	fmt.Printf(register.uuid)
	if register.config.value() != "localhost:9090" {
		t.Errorf("expected localhost:9090, but the true value is %s", register.config.value())
	}
}