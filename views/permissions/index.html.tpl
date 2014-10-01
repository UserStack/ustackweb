<div class="container">
  <div class="row">
    <div class="col-md-12 table-responsive">
      <table class="table table-condensed">
        <col style="width:90%">
        <col style="width:20%">
        <thead>
          <tr>
            <th>Name</th>
            <th>
              <a href="{{urlfor "PermissionsController.New"}}" class="btn btn-primary btn-xs">
                <span class="glyphicon glyphicon-plus"></span>
                Add Permission
              </a>
            </th>
          </tr>
        </thead>
        <tbody>
        {{range .permissions}}
          <tr>
            <td>{{.Name}}</td>
            <td>
              <a href="{{urlfor "UsersController.Index" ":groupId" .GroupName }}" class="btn btn-link btn-xs">
                <span class="glyphicon glyphicon-list"></span>
                Users
              </a>
            </td>
          </tr>
        {{end}}
        </tbody>
      </table>
      {{template "shared/paginator.html.tpl" .}}
    </div>
  </div>
</div>
