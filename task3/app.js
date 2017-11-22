$(document).ready(function(){
    $("#root").append("<ul></ul>");
    $("ul").append("<li></li>");
    $("li").append("<span>Сделать задание #3 по web-программированию</span>");
    $("li").append("<button> Удалить</button>").click(function() {
        $(this).remove()
    });
    $("#root").append($("<input>").attr("id", "add_task_input"));
    $("#root").append($("<button>Добавить</button>").attr("id", "add_task"));
    var i = 0;    
    $("#add_task").click(function() {
        i++;
        var j = $("input").val();
        $("ul").append($("<li></li>").attr("id", i));
        j = "<span>" + j + "</span>";
        $("#" + i).append(j);
        $("#" + i).append("<button> Удалить</button>").click(function() {
            $(this).remove()
        });
    });
});



