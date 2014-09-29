<form action="{{urlfor "UsersController.UpdatePassword" ":id" (printf "%d" .User.Uid)}}" method="post" class="form-horizontal" role="form">
  {{.XsrfHtml | str2html}}
  <div class="form-group">
    <div class="col-md-offset-3 col-md-6">
      <h1>Change Password</h1>
      {{ template "shared/form_errors.html.tpl" .ValidationErrors }}
    </div>
  </div>
  <div class="form-group {{hasFormError .ValidationErrors "Password"}}">
    <label class="col-md-3 control-label" for="inputUser-password">New Password</label>
    <div class="col-md-6">
      <input type="password" class="form-control" id="inputUser-password" name="Password" placeholder="Password">
    </div>
  </div>
  <div class="form-group {{hasFormError .ValidationErrors "OldPassword"}}">
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
