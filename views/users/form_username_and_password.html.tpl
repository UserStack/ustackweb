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
