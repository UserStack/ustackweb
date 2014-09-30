<div class="container">
  <div class="row">
    <div class="col-md-12">
      <form action="{{urlfor "GroupsController.Create" }}" method="post" class="form-horizontal" role="form">
        {{.xsrf_html | str2html}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Create Group</h1>
            {{template "shared/form_errors.html.tpl" .NewGroupError.Errors}}
          </div>
        </div>
        {{template "shared/form/fields.html.tpl" .NewGroupFormSets}}
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <button type="submit" class="btn btn-default">Create</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
