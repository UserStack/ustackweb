<div class="container">
  <div class="row">
    <div class="col-md-12">
      <form action="{{urlfor "UsersController.Create" }}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Create User</h1>
            {{template "shared/form_errors.html.tpl" .Errors}}
          </div>
        </div>
        {{with .NewUserFormSets.Fields.Username}}
        <div class="form-group {{if .Error}} has-error{{end}}">
          <label class="col-md-3 control-label" for="{{.Id}}">{{.LabelText}}</label>
          <div class="col-md-6">
            {{call .Field}}
          </div>
        </div>
        {{end}}
        {{with .NewUserFormSets.Fields.Password}}
        <div class="form-group {{if .Error}} has-error{{end}}">
          <label class="col-md-3 control-label" for="{{.Id}}">{{.LabelText}}</label>
          <div class="col-md-6">
            {{call .Field}}
          </div>
        </div>
        {{end}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <button type="submit" class="btn btn-default">Create</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
