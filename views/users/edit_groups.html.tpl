<div class="container">
  <div class="row">
    <div class="col-md-12">
      <a class='btn btn-default' href='{{urlfor "UsersController.Edit" ":id" (printf "%d" .user.Uid) }}'>
        <span class="glyphicon glyphicon-arrow-left"></span>
        Back
      </a>
      <h1>Edit Group Membership of {{.user.Name}}</h1>
      <table class="table table-condensed">
        <col style="width:70%">
        <col style="width:20%">
        <col style="width:10%">
        <thead>
          <tr>
            <th>Group</th>
            <th>Member?</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {{$user := .user}}
          {{range .groupMemberships}}
          <tr>
            <td>
              <a href="{{urlfor "GroupsController.Edit" ":id" (printf "%d" .Group.Gid) }}">
                {{.Group.Name}}
              </a>
            </td>
            <td>
            {{if .IsMember}}
              <span class="glyphicon glyphicon-ok"></span>
            {{end}}
            </td>
            <td>
            {{if .IsMember}}
            <a href="{{urlfor "UsersController.RemoveUserFromGroup" ":id" (printf "%d" $user.Uid) ":groupId" (printf "%d" .Group.Gid) }}" class="btn btn-danger btn-xs">
              <span class="glyphicon glyphicon-remove"></span>
            </a>
            {{else}}
            <a href="{{urlfor "UsersController.AddUserToGroup" ":id" (printf "%d" $user.Uid) ":groupId" (printf "%d" .Group.Gid) }}" class="btn btn-success btn-xs">
              <span class="glyphicon glyphicon-ok"></span>
            </a>
            {{end}}
            </td>
          </tr>
          {{end}}
        </tbody>
      </table>
    </div>
  </div>
</div>
