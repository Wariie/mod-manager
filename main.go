package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/Wariie/go-woxy/com"
	"github.com/Wariie/go-woxy/modbase"
)

var mods *[]Module

func main() {
	var m modbase.ModuleImpl
	m.Name = "mod-manager"
	m.InstanceName = "mod-manager"

	m.SetHubAddress("127.0.0.1")
	m.SetHubPort("2000")
	//m.SetHubAddress("guilhem-mateo.fr")
	m.SetHubProtocol("http")
	m.SetPort("2001")
	m.SetCommand("pouet", pouet)
	m.Init()
	m.Register("/mod-manager/api", api, "")
	m.Register("/mod-manager", index, "WEB")
	m.Run()
}

func index(w http.ResponseWriter, r *http.Request) {
	modList := refreshModuleList()

	me := modbase.GetModManager().GetMod()

	running := 0
	for k := range modList {
		if modList[k].STATE == Online && modList[k].NAME != "hub" {
			running++
		}
	}

	log.Println("GET / mod.v0", r.RemoteAddr)

	data := IndexPage{
		Title:             me.Name,
		Path:              "/" + me.Name,
		ModNumber:         len(modList) - 1,
		ModNumberRunning:  running,
		Mods:              modList,
		Secret:            modbase.GetModManager().GetSecret(),
		ModuleStateString: ModuleStateString,
	}

	tmpl := template.Must(template.ParseFiles("./views/layouts/master.html", "./views/index.html"))
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Fatalln("Error : ", err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	t, b := com.GetCustomRequestType(r)
	response := ""

	//CHECK SESSION TOKEN

	// CHECK ERROR DURING READING DATA
	if t["error"] == "error" {
		response = "Error reading Request"
	} else if t["Hash"] != "" {
		var mod Module
		modList := *mods

		if t["Type"] == "Command" {

			for m := range modList {
				if modList[m].NAME == t["Hash"] {
					mod = modList[m]
				}
			}

			var cr com.CommandRequest
			cr.Decode(b)

			//CHECK FOR ShutdownOrStart
			if cr.Command == "ShutdownOrStart" {
				if mod.STATE > Stopped && mod.STATE <= Loading {
					cr.Command = "Shutdown"
				} else {
					cr.Command = "Start"
				}
			}

			//Process command
			if cr.Command == "Status" {
				response += string(ModuleStateString[mod.STATE])
			} else {
				body, err := sendCommand(mod, cr.Command, cr.Content)

				if err != nil {
					response += err.Error()
				}
				response += body
			}
			log.Println(cr.Command + " TO " + mod.NAME)

		}
	}
	w.Write([]byte(response))
}

func refreshModuleList() []Module {
	listM := getModules(modbase.GetModManager().GetMod())
	mods = &listM
	return listM
}

func sendCommand(mod Module, command string, content string) (string, error) {
	var cr com.CommandRequest
	cr.Generate(command, mod.PK, mod.NAME, modbase.GetModManager().GetSecret())
	var c interface{} = &cr
	p := (c).(com.Request)
	return com.SendRequest(modbase.GetModManager().GetMod().HubServer, p, false)
}

type IndexPage struct {
	Title             string
	Path              string
	ModNumber         int
	ModNumberRunning  int
	Mods              []Module
	Secret            string
	ModuleStateString [7]string
}

func getModules(m *modbase.ModuleImpl) []Module {
	var cr com.CommandRequest
	cr.Generate("List", "hub", m.Name, modbase.GetModManager().GetSecret())
	srv := m.HubServer
	var mods []Module
	jsonData, err := com.SendRequest(srv, &cr, true)
	if err != nil {
		log.Println("Error retrieving Modules :", err)
	}
	json.NewDecoder(bytes.NewBufferString(jsonData)).Decode(&mods)

	return mods
}

func pouet(r *com.Request, w http.ResponseWriter, re *http.Request, mod *modbase.ModuleImpl) (string, error) {
	return "pouet", nil
}

/*Module - Module configuration */
type Module struct {
	AUTH         ModuleAuthConfig
	BINDING      ServerConfig
	COMMANDS     []string
	EXE          ModuleExecConfig
	NAME         string
	pid          int
	PK           string
	RESOURCEPATH string
	STATE        ModuleState
	TYPES        string
	VERSION      int
}

/*ModuleExecConfig - Module exec file informations */
type ModuleExecConfig struct {
	BIN        string
	MAIN       string
	SRC        string
	SUPERVISED bool
	REMOTE     bool
}

/*ServerConfig - Server configuration*/
type ServerConfig struct {
	ADDRESS  string
	PATH     []Route
	PORT     string
	PROTOCOL string
	ROOT     string
}

/*ModuleAuthConfig - Auth configuration*/
type ModuleAuthConfig struct {
	ENABLED bool
	TYPE    string
}

// Route - Route redirection
type Route struct {
	FROM string
	TO   string
}

//ModuleState - ModuleConfig State
type ModuleState int

//ModuleState list
const (
	Stopped    ModuleState = 0
	Unknown    ModuleState = 1
	Online     ModuleState = 2
	Downloaded ModuleState = 3
	Loading    ModuleState = 4

	Error  ModuleState = 999
	Failed ModuleState = 998
)

var ModuleStateString = [...]string{"STOPPED", "UNKNOWN", "ONLINE", "DOWNLOADED", "LOADING", "ERROR", "FAILED"}
