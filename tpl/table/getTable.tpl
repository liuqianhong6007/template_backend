<table border="1">
    <tr>
        <th>字段名称</th>
        <th>字段类型</th>
        <th>数据类型</th>
        <th>描述</th>
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

<script>
$(document).ready(function(){
    // 表格样式
    $("tr").hover(function(){
        $(this).css("background-color","yellow");
    },function(){
         $(this).css("background-color","white");
    });
 });
</script>