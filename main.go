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
	m.Name = "mod-manager"
	m.InstanceName = "mod-manager"

	m.SetServer("", "", "2001", "")
	m.SetCommand("pouet", pouet)
	m.Init()
	m.Register("GET", "/mod-manager", index, "WEB")
	m.Run()
}

func index(ctx *gin.Context) {

	var mods map[string]Module
	me := modbase.GetModManager().GetMod()
	mods = getModules(me)
	log.Println(mods)

	running := 0
	for k := range mods {
		if mods[k].STATE == Online && mods[k].NAME != "hub" {
			running++
		}
	}

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":              me.Name,
		"path":               "/" + me.Name,
		"mod_number":         len(mods) - 1,
		"mod_number_running": running,
		"mods":               mods,
		"secret":             modbase.GetModManager().GetSecret(),
	})
	log.Println("GET / mod.v0", ctx.Request.RemoteAddr)
}

func getModules(m *modbase.ModuleImpl) map[string]Module {
	var cr com.CommandRequest
	cr.Generate("List", "hub", m.Name, modbase.GetModManager().GetSecret())
	srv := m.HubServer
	var mods map[string]Module
	jsonData, err := com.SendRequest(srv, &cr, true)
	if err != nil {
		log.Println("Error retrieving Modules :", err)
	}
	json.NewDecoder(bytes.NewBufferString(jsonData)).Decode(&mods)

	return mods
}

func pouet(r com.Request, c *gin.Context, mod *modbase.ModuleImpl) (string, error) {
	return "pouet", nil
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
