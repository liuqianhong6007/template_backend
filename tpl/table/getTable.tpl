<table border="1">
    <tr>
        <th>field_name</th>
        <th>data_type</th>
        <th>field_type</th>
        <th>desc</th>
    </tr>

    {{range .Columns}}
    <tr>
        <td>{{.ColumnName}}</td>
        <td>{{.DataType}}</td>
        <td>{{.ColumnType}}</td>
        <td>{{.ColumnComment}}</td>
    </tr>
    {{end}}
</table>