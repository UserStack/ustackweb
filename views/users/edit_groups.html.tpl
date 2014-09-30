<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Edit" ":id" (printf "%d" .user.Uid) }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Back
      </a>
      <form class="form-horizontal">
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Change Membership of {{.user.Name}}</h1>
          </div>
        </div>
        <div class="form-group">
          <label class="col-md-3 control-label">Member in</label>
          <div class="col-md-6">
            {{$user := .user}}
            {{range .userGroups}}
            <div class="btn-group">
              <a href="{{urlfor "GroupsController.Edit" ":id" (printf "%d" .Gid) }}" class="btn btn-default">
                {{.Name}}
              </a>
              <a href="{{urlfor "UsersController.RemoveUserFromGroup" ":id" (printf "%d" $user.Uid) ":groupId" (printf "%d" .Gid) }}" class="btn btn-danger">
                <span class="glyphicon glyphicon-remove"></span>
              </a>
            </div>
            {{end}}
          </div>
        </div>
      </form>
      <form class="form-horizontal">
        <div class="form-group">
          <label class="col-md-3 control-label">Not Member in</label>
          <div class="col-md-6">
            {{$user := .user}}
            {{range .allGroups}}
            <div class="btn-group">
              <a href="{{urlfor "GroupsController.Edit" ":id" (printf "%d" .Gid) }}" class="btn btn-default">
                {{.Name}}
              </a>
              <a href="{{urlfor "UsersController.AddUserToGroup" ":id" (printf "%d" $user.Uid) ":groupId" (printf "%d" .Gid) }}" class="btn btn-default">
                <span class="glyphicon glyphicon-ok"></span>
              </a>
            </div>
            {{end}}
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
