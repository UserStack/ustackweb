<div class="container">
  <div class="row">
    <div class="col-md-12">
      <form action="{{urlfor "UsersController.Create" }}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <p>Create User</p>
          </div>
        </div>
        <div class="form-group">
          <label class="col-md-3 control-label" for="inputUser-username">Username</label>
          <div class="col-md-6">
            <input type="text" class="form-control" id="inputUser-username" name="username" placeholder="Username">
          </div>
        </div>
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <button type="submit" class="btn btn-default">Create</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
