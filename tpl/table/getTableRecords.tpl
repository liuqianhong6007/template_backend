<table border="1" id="{{.TableName}}">
    <tr>
    {{range .Columns}}
        <th>{{.}}</th>
    {{end}}
    <th>operation</th>
    </tr>

    {{range .Records}}
    <tr>
         {{range .RecordColumns}}
         <td class="{{.ColumnName}}" contentEditable="{{.Editable}}">{{.Val}}</td>
         {{end}}
         <td>
            <button type="submit" id="updateTableRecordBtn">修改</button>
            <button type="submit" id="deleteTableRecordBtn">删除</button>
         </td>
    </tr>
    {{end}}
</table>
<button type="submit" id="addTableRecordBtn">新增</button>
