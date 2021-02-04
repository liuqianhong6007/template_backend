<table border="1">
    <tr>
    {{range .Columns}}
        <th>{{.}}</th>
    {{end}}
    <th>操作</th>
    </tr>

    {{range .Records}}
    <tr>
         {{range .Values}}
         <td>{{.}}</td>
         {{end}}
         <td>
            <button type="submit">修改</button>
            <button type="submit">删除</button>
         </td>
    </tr>
    {{end}}
    <tr>
        <td>
            <button type="submit">新增</button>
        </td>
    </tr>
</table>