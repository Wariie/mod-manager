package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	com "github.com/Wariie/go-woxy/com"
	modbase "github.com/Wariie/go-woxy/modbase"
)

var mod *modbase.ModuleImpl

func main() {
	var m modbase.ModuleImpl
	mod = &m
	mod.Name = "mod-manager"
	mod.InstanceName = "module-manager"
	modbase.HubAddress = "localhost"
	modbase.ModulePort = "2001"
	mod.Init()
	mod.Register("GET", "/mod-manager", index, "WEB")
	mod.Run()
}

func index(ctx *gin.Context) {

	var mods map[string]Module
	me := modbase.GetModManager().GetMod()
	mods = getModules(me)
	log.Println(mods)

	running := 0
	for k := range mods {
		if mods[k].STATE == Online {
			running++
		}
	}

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":              me.Name,
		"path":               "/" + me.Name,
		"mod_number":         len(mods),
		"mod_number_running": running,
		"mods":               mods,
	})
	log.Println("GET / mod.v0", ctx.Request.RemoteAddr)
}

func getModules(m *modbase.ModuleImpl) map[string]Module {
	var cr com.CommandRequest
	cr.Generate(m.Name, "hub", "Get:List:Module")
	srv := com.Server{IP: modbase.HubAddress, Port: modbase.HubPort, Path: "", Protocol: "http"}
	var mods map[string]Module
	jsonData, err := com.SendRequest(srv, &cr, true)
	if err != nil {
		log.Println("Error retrieving Modules :", err)
	}
	json.NewDecoder(bytes.NewBufferString(jsonData)).Decode(&mods)

	return mods
}

/*Module - Module configuration */
type Module struct {
	NAME    string
	VERSION int
	TYPES   string
	EXE     ModuleExecConfig
	BINDING ServerConfig
	STATE   ModuleState
	PK      string
}

/*ModuleExecConfig - Module exec file informations */
type ModuleExecConfig struct {
	SRC  string
	MAIN string
	BIN  string
}

/*ServerConfig - Server configuration*/
type ServerConfig struct {
	ADDRESS  string
	PORT     string
	PATH     []Route
	PROTOCOL string
	ROOT     string
}

// Route - Route redirection
type Route struct {
	FROM string
	TO   string
}

//ModuleState - State of ModuleConfig
type ModuleState string

const (
	Unknown    ModuleState = "UNKNOWN"
	Loading    ModuleState = "LOADING"
	Online     ModuleState = "ONLINE"
	Stopped    ModuleState = "STOPPED"
	Downloaded ModuleState = "DOWNLOADED"
	Error      ModuleState = "ERROR"
)
