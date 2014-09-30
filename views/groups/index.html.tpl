<div class="container">
  <div class="row">
    <div class="col-md-12 table-responsive">
      <table class="table table-condensed">
        <col style="width: 8.3333%">
        <col style="width:91.6667%">
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th>
              <a href="{{urlfor "GroupsController.New"}}" class="btn btn-primary btn-xs">
                <span class="glyphicon glyphicon-plus"></span>
                Add Group
              </a>
            </th>
          </tr>
        </thead>
        <tbody>
        {{range .groups}}
          <tr>
            <td>{{.Gid}}</td>
            <td>{{.Name}}</td>
            <td>
            </td>
          </tr>
        {{end}}
        </tbody>
      </table>
      {{template "shared/paginator.html.tpl" .}}
    </div>
  </div>
</div>
