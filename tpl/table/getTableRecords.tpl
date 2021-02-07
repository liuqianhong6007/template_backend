<table border="1" id="{{.TableName}}">
    <tr>
    {{range .Columns}}
        <th>{{.}}</th>
    {{end}}
    <th>操作</th>
    </tr>

    {{range .Records}}
    <tr>
         {{range .RecordColumns}}
            {{if .Editable}}
                <td class="{{.ColumnName}}" contentEditable="true">{{.Val}}</td>
            {{else}}
                <td class="{{.ColumnName}}">{{.Val}}</td>
            {{end}}
         {{end}}
         <td>
            <button type="submit" id="updateTableRecordBtn">修改</button>
            <button type="submit" id="deleteTableRecordBtn">删除</button>
         </td>
    </tr>
    {{end}}
    <tr>
        <td>
            <button type="submit">新增</button>
        </td>
    </tr>
</table>