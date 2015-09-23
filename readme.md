# Introduction

`ycode` is a small library intended to read a yaml file mapping a list of sql
name to code pairs to fields in a go structure.  So long as a mapping can be
derived of "FieldName" to "SqlCode" then the go structure can be filled in by
this code.  The code mapping can be carried out in any way and this library
will handle the final step of filling in the go structure.

## Usage

Two usage scenarios come to mind when using this library.  A project could
have sql code organized into .sql files for editing via a sql management
editor.  The mapping could then be from the name of the .sql file to the 
source code contained in those files -- producing something like this:

```
mapping := map[string]*Sql{
  "CreateProductTable":{
    Name: "create_product_table.sql",
    Sql: "create table Product (...)"
  }
}
```

Another approach might be to have a single yaml file with a number of embedded
sql scripts. Those scripts could then be mapped to the above structure.  An
example might look like this:

```
---
scripts:
  - name: "create product table"
    sql: |
      create table Product (...)
```

So, long as a map can be produced in the form of the first example this library
could then fill a Struct of this form:

```
type ProductScripts struct {
  CreateProductTable *ycode.Sql
  ...
}
```


## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt

