<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Edit" ":id" (printf "%d" .user.Uid) }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Edit
      </a>
      {{template "users/form_username.html.tpl" .UsernameForm}}
    </div>
  </div>
</div>
