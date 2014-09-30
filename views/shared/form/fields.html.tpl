{{range $field := .FieldList}}
    {{template "shared/form/group.html.tpl" $field}}
{{end}}
