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
            {{range .groupMemberships}}
            <div class="btn-group">
              <a href="{{urlfor "GroupsController.Edit" ":id" (printf "%d" .Group.Gid) }}" class="btn btn-default">
                {{.Group.Name}}
              </a>
              {{if .IsMember}}
              <a href="{{urlfor "UsersController.RemoveUserFromGroup" ":id" (printf "%d" $user.Uid) ":groupId" (printf "%d" .Group.Gid) }}" class="btn btn-danger">
                <span class="glyphicon glyphicon-remove"></span>
              </a>
              {{else}}
              <a href="{{urlfor "UsersController.AddUserToGroup" ":id" (printf "%d" $user.Uid) ":groupId" (printf "%d" .Group.Gid) }}" class="btn btn-default">
                <span class="glyphicon glyphicon-ok"></span>
              </a>
              {{end}}
            </div>
            {{end}}
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
