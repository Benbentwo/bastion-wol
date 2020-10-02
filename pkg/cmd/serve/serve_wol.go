package serve

import (
	"fmt"
	"github.com/Benbentwo/utils/log"
	"github.com/mdlayher/wol"
	"net"
	"net/http"
	"os"
)

func WolPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"status\": \"ok\" }")
	log.Logger().Infof("%s: WOL Page Pinged by: %s", now(), r.RemoteAddr)
	macAddr := r.URL.Query().Get("mac")
	if macAddr == "" {
		val, exists := os.LookupEnv("MAC_ADDR_TO_WAKE")
		if exists && val != "" {
			macAddr = val
		} else {
			log.Logger().Errorf("Misconfigured server - no default mac address to send to")
		}
	}

	_ = wakeCmd(macAddr)
}

// Run the wake command.
func wakeCmd(macAddress string) error {
	mac, err := net.ParseMAC(macAddress)
	if err != nil {
		log.Logger().Errorf("convert mac address %s: %s", macAddress, err)
	}
	password := make([]byte, 0)
	if err := wakeUDP("255.255.255.255:9", mac, password); err != nil {
		log.Logger().Fatal(err)
	}

	log.Logger().Printf("sent UDP Wake-on-LAN magic packet using %s to %s", macAddress, "local network")
	return nil
}

//func wakeRaw(iface string, target net.HardwareAddr, password []byte) error {
//	ifi, err := net.InterfaceByName(iface)
//	if err != nil {
//		return err
//	}
//
//	c, err := wol.NewRawClient(ifi)
//	if err != nil {
//		return err
//	}
//	defer c.Close()
//
//	// Attempt to wake target machine.
//	return c.WakePassword(target, password)
//}

func wakeUDP(addr string, target net.HardwareAddr, password []byte) error {
	c, err := wol.NewClient()
	if err != nil {
		return err
	}
	defer c.Close()

	// Attempt to wake target machine.
	return c.WakePassword(addr, target, password)
}
