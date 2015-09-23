package ycode

import (
	"vals"
	"github.com/spf13/viper"
	"bytes"
	"io/ioutil"
	"strings"
	"fmt"
	"reflect"
	"log"
)

// Source structs provide access to the api for rendering a mapping source
// to a custom structure.
type Source struct {
	val *vals.Value
}

// Pairs up a human name with the sql code.  The name can be space or '_' separated.
type Sql struct {
	Name string
	Sql string
}

// A lookup of Sql instances by their name.
type CodeLookup map[string]*Sql

func (code CodeLookup) LoadSqlFields(scriptStruct interface{}) {
	v := reflect.ValueOf(scriptStruct).Elem()
	n := v.NumField()
	for i := 0; i < n; i++ {
		name := v.Type().Field(i).Name
		if val, ok := code[name]; ok {
			v.FieldByName(name).Set(reflect.ValueOf(val))
		}
	}
}

// Helper function for logging errors as fatal should they exist.
func maybePanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Brute force capitalizing words where the first letter is capitalized
// and the rest of the word is lower-cased.
func capitalized(s string) string {
	a := strings.ToUpper(s[0:1])
	b := strings.ToLower(s[1:])
	return fmt.Sprintf("%s%s", a, b)
}

// Breaks apart name by either _ or space.
func split(name string) []string {
	parts := strings.FieldsFunc(name, func(c rune) bool {
		switch c {
		case ' ', '_':
			return true
		default:
			return false
		}
	})
	return parts
}

// Splits the word over ' ' and '_'; capitalizes each word then joins them.
func capitalizeWords(word string) string {
	parts := split(word)
	for i,e := range parts {
		parts[i] = capitalized(e)
	}
	return strings.Join(parts, "")
}

// Provides the name of the Sql instance in Pascal format.
func (s *Sql) PascalName() string {
	return capitalizeWords(s.Name)
}

// Turns the given Value as a CodeLookup map to ease mapping it to a struct.
func (val *Source) YamlToSqlScripts() CodeLookup {
	n := val.Len()
	scripts := make(map[string]*Sql, n)
	for i := 0; i < n; i++ {
		v := val.In(i)
		s := &Sql{
			Name: v.At("name").AsString(),
			Sql: v.At("sql").AsString(),
		}
		scripts[s.PascalName()] = s
	}
	return scripts
}

// Fills the destination (should be a pointer) with the Sql code found in the
// yaml file mapping the name properties to the Pascal cased struct fields.
func (v *Source) FromYaml(dest interface{}) error {
	v.YamlToSqlScripts().LoadSqlFields(dest)
	return v.Error()
}

// Loads the given yaml file as a Value.
func LoadYaml(file string) *Source {
	y, err := ioutil.ReadFile(file)
	maybePanic(err)
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(y))
	conf := viper.AllSettings()

	return &Source{
		val: vals.New(conf),
	}
}

// Delegates to underlying value but return the Source.
func (s *Source) At(key string) *Source {
	s.val = s.val.At(key)
	return s
}

// Delegates to underlying value but return the Source.
func (s *Source) AsString() string {
	return s.val.AsString()
}

// Delegates to underlying value but return the Source.
func (s *Source) In(n int) *Source {
	s.val = s.val.In(n)
	return s
}

// Delegates to underlying value but return the Source.
func (s *Source) Error() error {
	return s.val.Error()
}

// Delegates to underlying value but return the Source.
func (s *Source) Len() int {
	return s.val.Len()
}





