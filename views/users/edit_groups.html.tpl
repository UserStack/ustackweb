<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Edit" ":id" (printf "%d" .user.Uid) }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Back
      </a>
      <h1>Group Membership of {{.user.Name}}</h1>
      <div class="list-group">
        {{$user := .user}}
        {{range .groupMemberships}}
          {{if .IsMember}}
          <a href="{{urlfor "UsersController.RemoveUserFromGroup" ":id" (printf "%d" $user.Uid) ":groupId" .Group.Name }}" class="list-group-item">
            {{.Group.Name}}
            <span class="badge badge-primary">
              <span class="glyphicon glyphicon-ok"></span>
            </span>
          </a>
          {{else}}
          <a href="{{urlfor "UsersController.AddUserToGroup" ":id" (printf "%d" $user.Uid) ":groupId" .Group.Name }}" class="list-group-item">
            {{.Group.Name}}
          </a>
          {{end}}
        {{end}}
      </div>
    </div>
  </div>
</div>
