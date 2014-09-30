<div class="container">
  <div class="row">
    <div class="col-md-offset-3 col-md-6 page-header">
      <h1>Configuration</h1>
      <a class='btn btn-default' href='{{urlfor "InstallController.Index" }}'>
        <span class="glyphicon glyphicon-refresh"></span>
        Retry
      </a>
    </div>
  </div>
  <div class="row">
    <div class="col-md-offset-3 col-md-6">
      <div class="list-group">
        {{if compare .rootUserError nil}}
        <div class="list-group-item success">
          <h4 class="list-group-item-heading">
            <span class="glyphicon glyphicon-ok-sign"></span>
            Root User
          </h4>
          <p class="list-group-item-text">
            {{.rootUser.Name}}
          </p>
        </div>
        {{else}}
        <div class="list-group-item list-group-item-danger">
          <h4 class="list-group-item-heading">
            <span class="glyphicon glyphicon-exclamation-sign"></span>
            Root User
            <a href="{{urlfor "InstallController.CreateRootUser" }}" class="btn btn-xs btn-default pull-right">
              Create
            </a>
          </h4>
          <p class="list-group-item-text">
            {{.rootUserError.Message}}
          </p>
        </div>
        {{end}}
        {{if compare .groupsError nil}}
        <div class="list-group-item">
          <h4 class="list-group-item-heading">
            <span class="glyphicon glyphicon-ok-sign"></span>
            Groups
          </h4>
          <p class="list-group-item-text">
            {{range .groups}}
              {{.Name}}
            {{end}}
          </p>
        </div>
        {{else}}
        <div class="list-group-item list-group-item-danger">
          <h4 class="list-group-item-heading">
            <span class="glyphicon glyphicon-exclamation-sign"></span>
            Groups
          </h4>
          <p class="list-group-item-text">
            {{.groupsError.Message}}
          </p>
        </div>
        {{end}}
      </div>
    </div>
  </div>
</div>

