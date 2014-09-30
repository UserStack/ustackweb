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
              <a href="#"
                class="btn btn-link btn-xs"
                tabindex="0"
                data-toggle="popover"
                data-trigger="focus"
                data-placement="left"
                data-html="true"
                data-content="<a class='btn btn-xs btn-danger' href='{{urlfor "GroupsController.Delete" ":id" (printf "%d" .Gid) }}'>Yes</a> <a class='btn btn-xs btn-default' href='#'>No</a>">
                <span class="glyphicon glyphicon-remove"></span>
                Delete
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
