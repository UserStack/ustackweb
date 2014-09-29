<div class="container">
  <div class="row">
    <div class="col-md-12">
      <form role="form" class="form-horizontal">
        <div class="form-group">
          <label class="col-md-3 control-label">Username</label>
          <div class="col-md-6 input-group">
            <input type="text" class="form-control" style="height: 33px;" disabled Bvalue="{{.user.Name}}">
            <span class="input-group-btn">
              <a class='btn btn-default' href='{{urlfor "UsersController.EditUsername" ":id" (printf "%d" .user.Uid) }}'>
                <span class="glyphicon glyphicon-pencil"></span>
              </a>
            </span>
          </div>
        </div>
        <div class="form-group">
          <label class="col-md-3 control-label" for="inputUser-password">Password</label>
          <div class="col-md-6 input-group">
            <input type="password" class="form-control" style="height: 33px;" disabled value="********">
            <span class="input-group-btn">
              <a class='btn btn-default' href='{{urlfor "UsersController.EditPassword" ":id" (printf "%d" .user.Uid) }}'>
                <span class="glyphicon glyphicon-pencil"></span>
              </a>
            </span>
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
