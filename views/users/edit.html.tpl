<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Index" }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Back
      </a>
      {{template "users/form_username_and_password.html.tpl" .}}
      <form class="form-horizontal">
        <div class="form-group">
          <label class="col-md-3 control-label">Memberships</label>
          <div class="col-md-6">
            {{$user := .user}}
            {{range .userGroups}}
            <a href="{{urlfor "GroupsController.Edit" ":id" (printf "%d" .Gid) }}" class="btn btn-link">
              {{.Name}}
            </a>
            {{end}}
            <a href="{{urlfor "UsersController.EditGroups" ":id" (printf "%d" .user.Uid) }}" class="btn btn-default">
              <span class="glyphicon glyphicon-wrench"></span>
              Change
            </a>
          </div>
        </div>
      </form>
      <form class="form-horizontal">
        <div class="form-group">
          <label class="col-md-3 control-label">Permissions</label>
          <div class="col-md-6">
            {{$user := .user}}
            {{range .userPermissions}}
            {{if .Granted}}
            <a class="btn">
              {{.Permission.Name}}
            </a>
            {{end}}
            {{end}}
            <a href="{{urlfor "UsersController.EditPermissions" ":id" (printf "%d" .user.Uid) }}" class="btn btn-default">
              <span class="glyphicon glyphicon-wrench"></span>
              Change
            </a>
          </div>
        </div>
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
