<div class="container">
  <div class="row">
    <div class="col-md-offset-3 col-md-6 page-header">
      <h1>Configuration</h1>
      <a class='btn btn-default' href='{{urlfor "InstallController.Index" }}'>
        <span class="glyphicon glyphicon-refresh"></span>
        Refresh
      </a>
      <a href="#"
        class="btn btn-warning pull-right"
        tabindex="0"
        data-toggle="popover"
        data-trigger="focus"
        data-placement="left"
        data-html="true"
        data-content="<a class='btn btn-xs btn-danger' href='{{urlfor "InstallController.DropDatabase" }}'>Yes</a> <a class='btn btn-xs btn-default' href='#'>No</a>">
        <span class="glyphicon glyphicon-remove"></span>
        Drop Database
      </a>
    </div>
  </div>
  <div class="row">
    <div class="col-md-offset-3 col-md-6">
      <h4>
        Root User
        <a href="{{urlfor "InstallController.CreateRootUser" }}" class="btn btn-xs btn-default pull-right">
          Recreate Root User
        </a>
      </h4>
      <div class="list-group">
        {{if .hasRootUser }}
        <div class="list-group-item list-group-item-success">
          <span class="glyphicon glyphicon-ok-sign"></span>
          {{.rootUser.Name}}
        </div>
        {{else if .hasRootUserError}}
        <div class="list-group-item list-group-item-danger">
          <span class="glyphicon glyphicon-exclamation-sign"></span>
          {{.rootUserError.Message}}
        </div>
        {{else}}
        <div class="list-group-item list-group-item-danger">
          <span class="glyphicon glyphicon-exclamation-sign"></span>
          No Root User
        </div>
        {{end}}
      </div>
      <h4>
        Permissions
        <a href="{{urlfor "InstallController.CreatePermissions" }}" class="btn btn-xs btn-default pull-right">
          Recreate Permissions
        </a>
      </h4>
      <div class="list-group">
        {{if compare .groupsError nil}}
        {{range .permissionRequirements}}
        <div class="list-group-item {{if .Exists}}list-group-item-success{{else}}list-group-item-danger{{end}}">
          {{if .Exists}}
          <span class="glyphicon glyphicon-ok-sign"></span>
          {{else}}
          <span class="glyphicon glyphicon-exclamation-sign"></span>
          {{end}}
          {{.Permission.Name}}
          <span class="pull-right">{{.Permission.GroupName}}</span>
        </div>
        {{end}}
        {{else}}
        <div class="list-group-item list-group-item-danger">
          <span class="glyphicon glyphicon-exclamation-sign"></span>
          {{.groupsError.Message}}
        </div>
        {{end}}
      </div>
      <h4>
        Root User Permissions
        <a href="{{urlfor "InstallController.AssignPermissions" }}" class="btn btn-xs btn-default pull-right">
          Reassign Permissions
        </a>
      </h4>
      <div class="list-group">
        {{if compare .groupsError nil}}
        {{range .permissionRequirements}}
        <div class="list-group-item {{if .Assigned}}list-group-item-success{{else}}list-group-item-danger{{end}}">
          {{if .Assigned}}
          <span class="glyphicon glyphicon-ok-sign"></span>
          {{else}}
          <span class="glyphicon glyphicon-exclamation-sign"></span>
          {{end}}
          {{.Permission.Name}}
          <span class="pull-right">{{.Permission.GroupName}}</span>
        </div>
        {{end}}
        {{else}}
        <div class="list-group-item list-group-item-danger">
          <span class="glyphicon glyphicon-exclamation-sign"></span>
          {{.groupsError.Message}}
        </div>
        {{end}}
      </div>
    </div>
  </div>
</div>

