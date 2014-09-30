<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Edit" ":id" (printf "%d" .user.Uid) }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Edit
      </a>
      <form action="{{urlfor "UsersController.UpdateUsername" ":id" (printf "%d" .user.Uid)}}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Change Username</h1>
          </div>
        </div>
        {{template "shared/horizontal_form/fields.html.tpl" .EditUsernameFormSets}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <button type="submit" class="btn btn-default">Change Username</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
