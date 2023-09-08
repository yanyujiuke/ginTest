$(function () {
    baseApp.init();
})

var baseApp = {
    init: function () {
        this.initAside()
        this.confirmDelete()
        this.resizeIframe()
        this.changeStatus()
        this.changeNum()
    },
    initAside: function () {
        $('.aside h4').click(function () {
            $(this).siblings('ul').slideToggle();
        })
    },
    resizeIframe: function () {
        $('#rightMain').height($('window').height() - 80)
    },
    // 删除提示
    confirmDelete: function () {
        $('.delete').click(function () {
            var flag = confirm('您确定要删除吗？')
            return flag
        })
    },
    // 改变转态值
    changeStatus: function () {
        $(".chStatus").click(function () {
            var dataId = $(this).attr("data-id")
            var dataTable = $(this).attr("data-table")
            var dataField = $(this).attr("data-field")
            var el = $(this)
            $.get("/admin/changeStatus", {
                id: dataId,
                table: dataTable,
                field: dataField,
            }, function (res) {
                if (res.success) {
                    if (el.attr("src").indexOf("yes") != -1) {
                        el.attr("src", "/static/admin/images/no.gif")
                    } else {
                        el.attr("src", "/static/admin/images/yes.gif")
                    }
                }
            })
        })
    },
    // 改变排序
    changeNum: function () {
        $(".chSpanNum").click(function () {
            // 1 获取el的属性值
            var id = $(this).attr("data-id")
            var table = $(this).attr("data-table")
            var field = $(this).attr("data-field")
            var num = $(this).html().trim()
            var spanEl = $(this)

            // 2 创建一个dom节点
            var input = $("<input style='width: 60px' value='' />")

            // 3 把input 放在el里
            $(this).html(input)

            // 4 让 input 获取焦点 给 input 赋值
            $(input).trigger("focus").val(num)

            // 5 点击input的时候阻止冒泡
            $(input).click(function (e) {
                e.stopPropagation()
            })

            // 6 鼠标离开的时候给span赋值，并触发ajax请求
            $(input).blur(function () {
                // 获取最新的input 的值
                var inputNum = $(this).val()
                // 给原来的span赋值
                spanEl.html(inputNum)
                // 触发ajax请求，更新数据
                $.get("/admin/changeNum", {
                    id: id,
                    table: table,
                    field: field,
                    num: inputNum,
                }, function(res) {
                    console.log(res)
                })
            })
        })
    }
}

