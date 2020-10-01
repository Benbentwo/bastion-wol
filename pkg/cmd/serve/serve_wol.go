package serve

import (
	"flag"
	"fmt"
	"github.com/Benbentwo/utils/log"
	"github.com/mdlayher/wol"
	"net"
	"net/http"
)

func WolPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "{ \"status\": \"ok\" }")
	log.Logger().Infof("%s: WOL Page Pinged by: %s", now(), r.Host)
}

// Run the wake command.
func wakeCmd(targetFlag string) error {
	flag.Parse()

	target, err := net.ParseMAC(targetFlag)
	if err != nil {
		log.Logger().Fatal(err)
	}

	// Can only do raw or UDP mode, not both.
	if *addrFlag != "" && *ifaceFlag != "" {
		log.Fatalf("must set '-a' or '-i' flag exclusively")
	}

	if *ifaceFlag != "" {
		if err := wakeRaw(*ifaceFlag, target, password); err != nil {
			log.Fatal(err)
		}

		log.Printf("sent raw Wake-on-LAN magic packet using %s to %s", *ifaceFlag, *targetFlag)
		return
	}

	if err := wakeUDP(*addrFlag, target, password); err != nil {
		log.Fatal(err)
	}

	log.Printf("sent UDP Wake-on-LAN magic packet using %s to %s", *addrFlag, *targetFlag)
}

func wakeRaw(iface string, target net.HardwareAddr, password []byte) error {
	ifi, err := net.InterfaceByName(iface)
	if err != nil {
		return err
	}

	c, err := wol.NewRawClient(ifi)
	if err != nil {
		return err
	}
	defer c.Close()

	// Attempt to wake target machine.
	return c.WakePassword(target, password)
}

func wakeUDP(addr string, target net.HardwareAddr, password []byte) error {
	c, err := wol.NewClient()
	if err != nil {
		return err
	}
	defer c.Close()

	// Attempt to wake target machine.
	return c.WakePassword(addr, target, password)
}
wi