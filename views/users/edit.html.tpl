<div class="container">
  <div class="row">
    <div class="col-md-12">
      <form role="form" class="form-horizontal">
        <div class="form-group">
          <label class="col-md-3 control-label">Username</label>
          <div class="col-md-4">
            <fieldset disabled>
              <input type="text" class="form-control" value="{{.user.Name}}">
            </fieldset>
          </div>
          <div class="col-md-2">
            <a class='btn btn-default pull-right' href='{{urlfor "UsersController.EditUsername" ":id" (printf "%d" .user.Uid) }}'>
              Change Username
            </a>
          </div>
        </div>
        <div class="form-group">
          <label class="col-md-3 control-label" for="inputUser-password">Password</label>
          <div class="col-md-4">
            <fieldset disabled>
              <input type="password" class="form-control" value="********">
            </fieldset>
          </div>
          <div class="col-md-2">
            <a class='btn btn-default pull-right' href='{{urlfor "UsersController.EditPassword" ":id" (printf "%d" .user.Uid) }}'>
              Change Password
            </a>
          </div>
        </div>
      </div>

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
