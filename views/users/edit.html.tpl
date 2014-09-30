<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Index" }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Back
      </a>
      {{template "users/form_username_and_password.html.tpl" .}}
      <div class="col-md-offset-3 col-md-6">
        <table class="table">
          <tbody>
          {{range .userGroups}}
            <tr>
              <td>{{.Gid}}</td>
              <td>{{.Name}}</td>
            </tr>
          {{end}}
          </tbody>
        </table>
      </div>
      <div class="clearfix"></div>
      <form action="{{urlfor "UsersController.AddUserToGroup" ":id" (printf "%d" .user.Uid)}}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        {{with .AddUserToGroupFormSets.Fields.GroupId}}
        <div class="form-group {{if .Error}} has-error{{end}}">
          <label class="col-md-3 control-label">{{.LabelText}}</label>
          <div class="col-md-6">
            {{call .Field}}
            <span class="input-group-btn">
              <button type="submit" class="btn btn-default">Add to Group</button>
            </span>
          </div>
        </div>
        {{end}}
      </form>
      <form class="form-horizontal">
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            {{if .user.Active}}
            <a class='btn btn-warning' href='{{urlfor "UsersController.Disable" ":id" (printf "%d" .user.Uid) }}'>
              <span class="glyphicon glyphicon-eye-close"></span>
              Lock
            </a>
            {{else}}
            <a class='btn btn-success' href='{{urlfor "UsersController.Enable" ":id" (printf "%d" .user.Uid) }}'>
              <span class="glyphicon glyphicon-eye-open"></span>
              Unlock
            </a>
            {{end}}
            <a href="#"
              class="btn btn-danger pull-right"
              tabindex="0"
              data-toggle="popover"
              data-trigger="focus"
              data-placement="left"
              data-html="true"
              data-content="<a class='btn btn-xs btn-danger' href='{{urlfor "UsersController.Destroy" ":id" (printf "%d" .user.Uid) }}'>Yes</a> <a class='btn btn-xs btn-default' href='#'>No</a>">
              <span class="glyphicon glyphicon-remove"></span>
              Delete
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
