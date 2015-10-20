package ycode
   
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
	"vals"
)

func TestEmbeddedYamlCode(t *testing.T) {

	Convey("Traversing source should create new instance of source", t, func() {
		m := map[string]string{
			"a": "x",
			"b": "y",
			"c": "z",
		}
		v := vals.New(m)
		src := NewSource(v)

		fmt.Printf("a: '%s' '%s'\n", src.At("a").AsString(), v.At("a").AsString())

		So(src.At("a").AsString(), ShouldEqual, "x")
		So(src.At("b").AsString(), ShouldEqual, "y")
		So(src.At("c").AsString(), ShouldEqual, "z")
		So(src.At("c"), ShouldNotEqual, src)
	})

	Convey("Should pascal case the values as shown", t, func() {
		m := map[string]string{
			"build_a_table": "BuildATable",
			"build a table": "BuildATable",
			"create": "Create",
		}
		for k,v := range m {
			So(capitalizeWords(k), ShouldEqual, v)
		}
	})

	Convey("Should split string around spaces and _", t, func() {
		m := map[string][]string{
			"build_a_table": []string{"build", "a", "table"},
			"build a table": []string{"build", "a", "table"},
			"create": []string{"create"},
		}
		for k,v := range m {
			for i,e := range split(k) {
				So(e, ShouldEqual, v[i])
			}
		}
	})

	Convey("Should fill in field of struct", t, func() {
		type Code struct {
			NameOfField *Sql
		}
		c := &Code{}
		m := CodeLookup{
			"NameOfField": {
				Name: "name of field",
				Sql: "code",
			},
		}
		m.LoadSqlFields(c)
		So(c.NameOfField, ShouldNotBeNil)
		So(c.NameOfField.Sql, ShouldEqual, "code")
		So(c.NameOfField.Name, ShouldEqual, "name of field")
	})
}

