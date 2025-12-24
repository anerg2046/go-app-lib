package genv

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		Load()
	})
}

// EnvVariable contains information about the environment variable, such as key,
// value, and default value.
type EnvVariable struct {
	Key          string
	Val          string
	DefaultValue interface{}
	IsDefined    bool
}

// EnvVariables is where environment variables are stored.
var EnvVariables = make(map[string]*EnvVariable)

// Key is used to determine the path of the environment variable to be accessed.
//
//	genv.Key("env-key").String()
func Key(key string) *EnvVariable {

	envVar, ok := EnvVariables[key]
	if !ok {

		val, ok := os.LookupEnv(key)
		EnvVariables[key] = &EnvVariable{Key: key, Val: val, IsDefined: ok}

		return EnvVariables[key]
	}

	return envVar
}

// Default is used to specify the default value for the environment
// variable to be accessed.
//
//	genv.Key("env-key").Default("defaultValue").String()
func (e *EnvVariable) Default(defaultValue interface{}) *EnvVariable {

	e.DefaultValue = defaultValue

	return e
}

// Update is used to update the value of the corresponding environment variable.
//
//	genv.Key("env-key").Update("updatedValue")
func (e *EnvVariable) Update(value interface{}) {

	switch value := value.(type) {
	case bool:
		e.Val = strconv.FormatBool(value)
	case float64:
		e.Val = strconv.FormatFloat(value, 'f', -1, 64)
	case int:
		e.Val = strconv.FormatInt(int64(value), 10)
	case string:
		e.Val = value
	}

	e.IsDefined = true
	os.Setenv(e.Key, e.Val)
}

// Bool method is used for environment variables of type bool.
//
//	genv.Key("env-key").Bool()
func (e *EnvVariable) Bool() bool {

	var dv bool
	if !e.IsDefined {
		if e.DefaultValue != nil {
			dv = e.DefaultValue.(bool)
		}

		return dv
	}

	val, _ := strconv.ParseBool(e.Val)

	return val
}

// Float method is used for environment variables of type float.
//
//	genv.Key("env-key").Float()
func (e *EnvVariable) Float() float64 {

	var dv float64
	if !e.IsDefined {
		if e.DefaultValue != nil {
			dv = e.DefaultValue.(float64)
		}

		return dv
	}

	val, _ := strconv.ParseFloat(e.Val, 64)

	return val
}

// Int method is used for environment variables of type int.
//
//	genv.Key("env-key").Int()
func (e *EnvVariable) Int() int {

	var dv int
	if !e.IsDefined {
		if e.DefaultValue != nil {
			dv = e.DefaultValue.(int)
		}

		return dv
	}

	val, _ := strconv.ParseInt(e.Val, 10, 32)

	return int(val)
}

// String method is used for environment variables of type string.
//
//	genv.Key("env-key").String()
func (e *EnvVariable) String() string {

	var dv string
	if !e.IsDefined {
		if e.DefaultValue != nil {
			dv = e.DefaultValue.(string)
		}

		return dv
	}

	return e.Val
}

const keyValRegex = `^\s*([\w.-]+)\s*=\s*(.*)?\s*$`

// Load method allows environment variables to be loaded from the desired file.
//
//	dotenv.Load(".envfile")
func Load(path ...string) (err error) {

	var dotenvPath string
	if len(path) > 0 {
		dotenvPath = path[0]
	} else {
		dotenvPath = ".env"
	}

	source, err := os.Open(dotenvPath)
	if err != nil {
		log.Print(err)
		return
	}
	defer source.Close()

	variables, err := parse(source)
	if err != nil {
		log.Print(err)
		return
	}

	for key, val := range variables {
		if _, ok := os.LookupEnv(key); !ok {
			os.Setenv(key, val)
		}
	}

	return
}

func parse(source *os.File) (variables map[string]string, err error) {

	variables = make(map[string]string)

	r, err := regexp.Compile(keyValRegex)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		key, val := parseLine(scanner.Text(), r)
		if key != "" && val != "" {
			variables[key] = val
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func parseLine(line string, r *regexp.Regexp) (key string, val string) {

	matches := r.FindStringSubmatch(line)

	if len(matches) == 3 {
		key, val = matches[1], matches[2]

		if val != "" {
			end := len(val) - 1
			isDoubleQuoted := val[0] == '"' && val[end] == '"'
			isSingleQuoted := val[0] == '\'' && val[end] == '\''

			if isSingleQuoted || isDoubleQuoted {
				val = val[1:end]
			} else {
				val = strings.Trim(val, " ")
			}
		}
	}

	return
}
