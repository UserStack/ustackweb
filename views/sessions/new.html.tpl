<div class="container">
  <div class="col-md-offset-3 col-md-6">
    <form action="{{urlfor "SessionsController.Create"}}" method="post" role="form" class="form-horizontal">
      {{.xsrf_html | str2html}}
      <input type="hidden" name="origin" value="{{.origin}}" />
      <div class="form-group">
        <div class="col-md-offset-2 col-md-10 ">
          <h1>Sign In</h1>
        </div>
      </div>
      <div class="form-group">
        <label for="inputUsername" class="col-md-2 control-label">Username</label>
        <div class="col-md-10">
          <input type="text" class="form-control" id="inputUsername" name="Username" placeholder="Username or Email">
        </div>
      </div>
      <div class="form-group">
        <label for="inputPassword" class="col-md-2 control-label">Password</label>
        <div class="col-md-10">
          <input type="password" class="form-control" id="inputPassword" name="Password" placeholder="Password">
        </div>
      </div>
      <div class="form-group">
        <div class="col-md-offset-2 col-md-10">
          <button type="submit" class="btn btn-default">Sign In</button>
        </div>
      </div>
      <div class="form-group">
        <div class="col-md-offset-2 col-md-10">
          <span class="help-block">
            <a href="{{urlfor "RegistrationsController.New"}}">Register</a>
            if you don't have an account.
          </span>
        </div>
      </div>
    </form>
  </div>
</div>
