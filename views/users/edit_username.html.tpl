<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Edit" ":id" (printf "%d" .user.Uid) }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Edit
      </a>
      <form action="{{urlfor "UsersController.UpdateUsername" ":id" (printf "%d" .user.Uid)}}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Change Username</h1>
            {{template "shared/form_errors.html.tpl" .EditUsernameError.Errors}}
          </div>
        </div>
        {{with .EditUsernameFormSets.Fields.Username}}
        <div class="form-group {{if .Error}} has-error{{end}}">
          <label class="col-md-3 control-label" for="{{.Id}}">{{.LabelText}}</label>
          <div class="col-md-6">
            {{call .Field}}
          </div>
        </div>
        {{end}}
        {{with .EditUsernameFormSets.Fields.ConfirmPassword}}
        <div class="form-group {{if .Error}} has-error{{end}}">
          <label class="col-md-3 control-label" for="{{.Id}}">{{.LabelText}}</label>
          <div class="col-md-6">
            {{call .Field}}
          </div>
        </div>
        {{end}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <button type="submit" class="btn btn-default">Change Username</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
