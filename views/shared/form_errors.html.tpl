<ul>
  {{range . }}
  <li>
    <mark>{{ .Key  }}</mark>
    {{ .Message  }}
  </li>
  {{end}}
</ul>
