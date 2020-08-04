## Parsing Template
1. text/template package
2. ```tpl,err:=template.ParseFiles("filename.gohtml")```
3. we can pass any no. of filename int templates.ParseFiles method.
4. filename can be with any extension
5. To execute -> ```err:=tpl.Execute(os.Stdout,nil)``` or ```err:=tpl.ExecuteTemplate(os.Stdout,"filename.gohtml",nil)```
6.  ExecuteTemplate is used if the tpl variable have more than one template then we have to use ExecuteTemplate to specify which template is to be used
7.  ParseGlob method -> ```tpl,err:=template.ParseGlob("directoryname/*.gohtml")```
8.  To initialise template variable -> ```func  init(){
  tpl=template.Must(template.ParseGlob("DirectoryName/*"))}```

```var tpl *template.Template```

9. "." dot is current data in template
10. To pass data into a template ->
```tpl,err:=template.ExecuteTemplate(os.Stduot,"filename.gohtml",dataTopass)```
To access the data we have passed ->
```{{.}}```
to assign variable -> ```{{varaibleName:=.}}```

11. to range over a list ->

``` {{range .}}<li>{{.}}</li>{{end}}```

12. Structs can be accessed by using field names for eg. -> for a struct having fielname Name of type string and Age of type int, then in the template file we can accessthe fields -> ```{{.Name}}{{.Age}}```