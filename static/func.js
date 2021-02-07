var domain = "http://127.0.0.1:8081";

// jquery 入口
$(document).ready(function(){
    // table 列表
    $("#getTablesSpan").click(function(){
        url = domain + "/getTables";
        $.get(url,function(data,status){
            $("aside").html(data);
        })
    });

    // 弹出选项卡
    $("aside").on("mouseenter",".table-name",function(){
        $("#table-option-pop").html(`
            <button type="submit" id="getTableBtn">查看表结构</button>
            <button type="submit" id="getTableRecordBtn">查看数据</button>
        `);
    });
    $("aside").on("mouseleave",".table-name",function(){
        setTimeout(function () {
            $("#table-option-pop").html("");
        },4000);
    });

    // 查看表结构
    $("aside").on("click","#getTableBtn",function () {
        tableName = $(this).parent().parent().find(".table-name").text().trim();
        url = domain+"/getTable?tableName=" + tableName;
        $.get(url,function(data,status){
            $("main").html(data);
        });
    });

    // 查看表数据
    $("aside").on("click","#getTableRecordBtn",function () {
        tableName = $(this).parent().parent().find(".table-name").text().trim();
        url = domain+"/getTableRecords?tableName=" + tableName
        $.get(url,function(data,status){
            $("main").html(data);
        });
    });

    // 更新表数据
    $("main").on("click","#updateTableRecordBtn",function () {
        tableName = $(this).parent().parent().parent().parent().attr("id").trim();
        url = domain+"/updateTableRecord";

        columns = $(this).parent().siblings();
        var data = "";
        for (i = 0;i < columns.length;i++){
            data = data + $(columns[i]).attr("class").trim() + "=" + $(columns[i]).text() + ";"
        }
        data = data.substr(0,data.length-1);

        $.post(url, {
            "tableName": tableName,
            "data": data,
        },function(data,status){
            $("main").html(data);
        });
    });

    // 表格样式
    $("tr").hover(function(){
        $(this).css("background-color","yellow");
    },function(){
        $(this).css("background-color","white");
    });
});

function getTableColumn(table) {

}
