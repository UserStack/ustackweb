{{if eq .Type "hidden"}}
    {{call .Field}}
{{else}}
  <label class="col-md-3 control-label" for="{{.Id}}">{{.LabelText}}</label>
  <div class="col-md-6">
    {{call .Field}}
    {{if .Error}}<p class="error-block">{{.Error}}</p>{{end}}
    {{if .Help}}<p class="help-block">{{.Help}}</p>{{end}}
  </div>
{{end}}
