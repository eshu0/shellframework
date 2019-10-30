package shellframework

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/eshu0/shellframework/interfaces"
)

type Environment struct {
	shell      sfinterfaces.IShell
	namevalues map[string]sfinterfaces.IEnvironmentVariable
}

func PointerInvalid(obj interface{}) bool {
	return obj == nil
}

func NewEnvironment(shell sfinterfaces.IShell) sfinterfaces.IEnvironment {

	var env = &Environment{}
	env.namevalues = make(map[string]sfinterfaces.IEnvironmentVariable)
	env.shell = shell
	return env
}

func (env *Environment) SetShell(shell sfinterfaces.IShell) {
	env.shell = shell
}

func (env *Environment) GetShell() sfinterfaces.IShell {
	return env.shell
}

func (env *Environment) SetNameValues(namevals map[string]sfinterfaces.IEnvironmentVariable) {
	env.namevalues = namevals
}

func (env *Environment) GetNameValues() map[string]sfinterfaces.IEnvironmentVariable {
	return env.namevalues
}

func (env *Environment) Print() {
	if PointerInvalid(env) {
		return
	}

	shell := env.GetShell()
	namevalues := env.GetNameValues()

	if namevalues == nil {
		shell.Println("Environment database is empty")
		return
	} else {
		if len(namevalues) == 0 {
			shell.Println("No Environment variables set")
		} else {
			for _, v := range namevalues {
				//v := *value
				shell.Printlnf("Name: %s Value: %s", v.GetName(), strings.Join(v.GetValues(), ","))
			}
		}
	}
}

func (env *Environment) AddStringValue(key string, value string) {
	envvar, exists := env.GetVariable(key)

	if !exists {
		var vals []string
		vals = append(vals, value)
		env.SetVariable(env.MakeMultiVariable(key, vals))
	} else {
		wc := envvar
		lc := wc.GetValues()
		lc = append(lc, value)
		wc.SetValues(lc)
		env.SetVariable(wc)
	}
}

//func (env *SEnvironment) AddStringValue(key string, value string) {
//if PointerInvalid(env) {
//	return
//}
/*
		namevalues := env.GetNameValues()

		if namevalues == nil {
			namevalues = make(map[string]*ishell.IEnvironmentVariable)
		}

	val, exists := env.GetVariable(key)
	if !exists {
		sev := EnvironmentVariable{}
		sev.SetName(key)
		sev.SetValues([]string{})
		sev.SetType(1) // string
		env.SetVariable(&sev)
	}

	//var currentvals []string
	//currentvals = val.GetValues()
	//currentvals = append(currentvals, value)
	//*val.SetValues(currentvals)
	//env.SetVariable(key, *val)
	//env.namevalues[key].Values = append(env.namevalues[key].Values, value)
}
*/

func (env *Environment) MakeMultiVariable(key string, val []string) sfinterfaces.IEnvironmentVariable {
	sev := EnvironmentVariable{}
	sev.SetName(key)
	sev.SetValues(val)
	sev.SetType(1)
	return &sev
}

func (env *Environment) MakeSingleVariable(key string, val string) sfinterfaces.IEnvironmentVariable {
	sev := EnvironmentVariable{}
	sev.SetName(key)
	var vals []string
	vals = append(vals, val)
	sev.SetValues(vals)
	sev.SetType(1)
	return &sev
}

func (env *Environment) SetVariable(value sfinterfaces.IEnvironmentVariable) {
	env.namevalues[value.GetName()] = value
}

func (env *Environment) GetVariable(key string) (sfinterfaces.IEnvironmentVariable, bool) {
	val, exists := env.namevalues[key]
	return val, exists
}

func (env *Environment) GetValues(key string) (bool, []string) {
	if PointerInvalid(env) {
		return false, []string{}
	}

	namevalues := env.GetNameValues()

	_, exists := namevalues[key]
	if !exists {
		return false, []string{}
	} else {
		return true, (namevalues[key]).GetValues()
	}
}

func (env *Environment) SaveToFile(path string) {
	shell := env.GetShell()
	log := *shell.GetLog()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		namevalues := env.GetNameValues()
		bytes, err := json.MarshalIndent(namevalues, "", "\t") //json.Marshal(p)
		if err != nil {
			log.LogError("SaveToFile()", "Marshal json for %s failed with %s ", path, err.Error())
			return
		}

		err = ioutil.WriteFile(path+".json", bytes, 0644)
		if err != nil {
			log.LogError("SaveToFile()", "Saving %s failed with %s ", path, err.Error())
		}
	} else {
		log.LogError("SaveToFile()", "'%s' was not found to save", path)
	}
}

func (env *Environment) LoadFile(path string) {

	shell := env.GetShell()
	log := *shell.GetLog()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		/*
			jsonFile, err := os.Open(filepath)

			if err != nil {
				env.Shell.LogPrintlnf("Loading '%s' failed with %s ", filepath, err.Error())
				return
			}
		*/
		filepath := path + ".json"
		bytes, err := ioutil.ReadFile(filepath) //ReadAll(jsonFile)
		if err != nil {
			log.LogError("LoadFile()", "Reading '%s' failed with %s ", filepath, err.Error())
			return
		}
		var f map[string]EnvironmentVariable

		//var NameValues map[string]ishell.IEnvironmentVariable

		err = json.Unmarshal(bytes, &f)

		if err != nil {
			log.LogError("LoadFile()", " Loading %s failed with %s ", filepath, err.Error())
			return
		}

		for key, ev := range f {
			log.LogDebug("LoadFile()", "SetVariable %s with %s ", key, strings.Join(ev.GetValues(), ","))
			env.SetVariable(env.MakeMultiVariable(key, ev.GetValues()))
		}

		//env.SetNameValues(f)
		//m := f.(map[string]ishell.IEnvironmentVariable)
		//original, ok := f.(map[string]ishell.IEnvironmentVariable)
		//if ok {
		//	println(original.b())
		//}
		/*
			for k, v := range f {
				switch vv := v.(type) {
				case string:
					shell.LogPrintlnf("%s is string", k, vv)
				case []string:
					shell.LogPrintlnf("%s is string array", k, vv)
				case float64:
					shell.LogPrintlnf("%s is float64", k, vv)
				case []interface{}:
					shell.LogPrintlnf("%s is an array:", k)
				case ishell.IEnvironment:
					shell.LogPrintlnf("%s is of a type ishell.IEnvironment", k)
				default:
					shell.LogPrintlnf("%s is of a type I don't know how to handle %s %s", k, vv, v)
				}
			}
		*/
	} else {
		log.LogError("LoadFile()", "'%s' was not found to load", path)
	}
}

// this is here because when in a seperate file
// the golang plus removes the import to interfaces - useful....
type EnvironmentVariable struct {
	name     string
	Values   []string
	Itemtype int
}

func (sevar *EnvironmentVariable) GetName() string {
	return sevar.name
}

func (sevar *EnvironmentVariable) GetValues() []string {
	return sevar.Values
}

func (sevar *EnvironmentVariable) GetType() int {
	return sevar.Itemtype
}

func (sevar *EnvironmentVariable) SetName(name string) {
	sevar.name = name
}

func (sevar *EnvironmentVariable) SetValues(vals []string) {
	sevar.Values = vals
}

func (sevar *EnvironmentVariable) SetType(typ int) {
	sevar.Itemtype = typ
}

func (sevar *EnvironmentVariable) String() string {
	return strings.Join(sevar.GetValues(), ",")
}
