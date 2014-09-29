<div class="container">
  <div class="row">
    <div class="col-md-12 table-responsive">
      <table class="table table-condensed">
        <col style="width:10%">
        <col style="width:5%">
        <col style="width:65%">
        <col style="width:20%">
        <thead>
          <tr>
            <th>#</th>
            <th></th>
            <th>Name</th>
            <th>
              <a href="{{urlfor "UsersController.New"}}" class="btn btn-primary btn-xs">
                <span class="glyphicon glyphicon-plus"></span>
                Add User
              </a>
            </th>
          </tr>
        </thead>
        <tbody>
        {{range .users}}
          <tr>
            <td>
              <a href="{{urlfor "UsersController.Edit" ":id" (printf "%d" .Uid) }}">
                {{.Uid}}
              </a>
            </td>
            <td>
              {{if .Active }}{{else}}
                <span class="text-muted glyphicon glyphicon-eye-close" data-placement="top" title="User is locked and cannot sign in."></span>
              {{end}}
            </td>
            <td>
              <a href="{{urlfor "UsersController.Edit" ":id" (printf "%d" .Uid) }}">
                {{.Name}}
              </a>
            <td>
              <div class="btn-group">
                <a href="{{urlfor "UsersController.Edit" ":id" (printf "%d" .Uid) }}" class="btn btn-link btn-xs">
                  <span class="glyphicon glyphicon-pencil"></span>
                  Change
                </a>
                {{if .Active }}
                <a href="{{urlfor "UsersController.Disable" ":id" (printf "%d" .Uid) }}" class="btn btn-link btn-xs">
                  <span class="glyphicon glyphicon-eye-close"></span>
                  Lock
                </a>
                {{else}}
                <a href="{{urlfor "UsersController.Enable" ":id" (printf "%d" .Uid) }}" class="btn btn-link btn-xs">
                  <span class="glyphicon glyphicon-eye-open"></span>
                  Unlock
                </a>
                {{end}}
                <a href="{{urlfor "UsersController.Destroy" ":id" (printf "%d" .Uid) }}" class="btn btn-link btn-xs">
                  <span class="glyphicon glyphicon-remove"></span>
                  Delete
                </a>
              </div>
            </td>
          </tr>
        {{end}}
        </tbody>
      </table>
      {{template "shared/paginator.html.tpl" .}}
    </div>
  </div>
</div>
