<div class="container">
  <div class="row">
    <div class="col-md-12">
      <form action="{{urlfor "UsersController.UpdateUsername" ":id" (printf "%d" .user.Uid)}}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Change Username</h1>
            {{ template "shared/form_errors.html.tpl" .UpdateUsernameErrors }}
          </div>
        </div>
        <div class="form-group {{hasFormError .UpdateUsernameErrors "Username"}}">
          <label class="col-md-3 control-label" for="inputUser-username">New Username</label>
          <div class="col-md-6">
            <input type="text" class="form-control" id="inputUser-username" name="Username" value="{{.user.Name }}" placeholder="Username">
          </div>
        </div>
        <div class="form-group {{hasFormError .UpdateUsernameErrors "Password"}}">
          <label class="col-md-3 control-label" for="inputUser-password">Confirm</label>
          <div class="col-md-6">
            <input type="password" class="form-control" id="inputUser-password" name="Password" placeholder="Password">
          </div>
        </div>
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <button type="submit" class="btn btn-default">Change Username</button>
          </div>
        </div>
      </form>
      <form action="{{urlfor "UsersController.UpdatePassword" ":id" (printf "%d" .user.Uid)}}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Change Password</h1>
            {{ template "shared/form_errors.html.tpl" .UpdatePasswordErrors }}
          </div>
        </div>
        <div class="form-group {{hasFormError .UpdatePasswordErrors "Password"}}">
          <label class="col-md-3 control-label" for="inputUser-password">New Password</label>
          <div class="col-md-6">
            <input type="password" class="form-control" id="inputUser-password" name="Password" placeholder="Password">
          </div>
        </div>
        <div class="form-group {{hasFormError .UpdatePasswordErrors "OldPassword"}}">
          <label class="col-md-3 control-label" for="inputUser-oldPassword">Confirm</label>
          <div class="col-md-6">
            <input type="password" class="form-control" id="inputUser-oldPassword" name="OldPassword" placeholder="Old Password">
          </div>
        </div>
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <button type="submit" class="btn btn-default">Change Password</button>
          </div>
        </div>
      </form>
      <div class="form-horizontal">
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
          </div>
        </div>
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <a href="#"
              class="btn btn-danger"
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
