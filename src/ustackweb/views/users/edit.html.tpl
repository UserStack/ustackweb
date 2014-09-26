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
        <div class="form-group">
          <label class="col-md-3 control-label" for="inputUser-username">New Username</label>
          <div class="col-md-6">
            <input type="text" class="form-control" id="inputUser-username" name="Username" value="{{.user.Name }}" placeholder="Username">
          </div>
        </div>
        <div class="form-group">
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
        <div class="form-group">
          <label class="col-md-3 control-label" for="inputUser-password">New Password</label>
          <div class="col-md-6">
            <input type="password" class="form-control" id="inputUser-password" name="Password" placeholder="Password">
          </div>
        </div>
        <div class="form-group">
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
    </div>
  </div>
</div>
