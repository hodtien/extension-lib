package main

import (
	"fmt"

	"github.com/hodtien/extension-lib/global"
	"github.com/hodtien/extension-lib/vdd"
	"github.com/hodtien/extension-lib/vdp"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	global.Constructor(
		viper.GetString("nats.url"),
		viper.GetString("domain"),
	)

	var resp map[string]interface{}

	// VDD test
	vddApikey := viper.GetString("api_key.vdd")
	resp = vdd.RetrieveAllDataInBucket(vddApikey, "location_site", "")

	// VDD test
	vdpApikey := viper.GetString("api_key.vdp")
	resp = vdp.NatsGetAllRoleInDomain(vdpApikey)
	fmt.Println(resp)
}
