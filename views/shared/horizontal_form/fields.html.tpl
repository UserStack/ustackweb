{{range $field := .FieldList}}
    {{template "shared/horizontal_form/group.html.tpl" $field}}
{{end}}
