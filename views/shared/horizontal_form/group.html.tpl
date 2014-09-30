{{if eq .Type "hidden"}}
    {{call .Field}}
{{else}}
    <div class="form-group{{if .Error}} has-error{{end}}">
      {{template "shared/horizontal_form/field.html.tpl" .}}
    </div>
{{end}}
