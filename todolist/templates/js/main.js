$(function () {
    $.fn.createNewItems = function (user_input, item_id, is_checked) {
        let $new_item = $("<div\>");
        let $new_item_container = $("<div\>");
        let $content_container = $("<div\>");
        let $button_container = $("<div\>");
        let $del_btn = $("<button\>");
        let $chk_box = $("<input>");
        $new_item.addClass("mx-2 my-2 to_do_list_title");
        $new_item.css("word-break", "break-all");
        $new_item_container.addClass("row w-100");
        $content_container.addClass("col-10");
        $content_container.text(user_input);
        $button_container.addClass("col-2 p-0 text-right");
        $del_btn.text("Del");
        $del_btn.addClass("btn btn-danger");
        $del_btn.css("whiteSpace", "nowrap");
        $del_btn.attr("name", item_id);
        $del_btn.click(function (e) {
            $.ajax({
                type: "DELETE",
                url: "/todolist/" + $(this).attr("name"),
                // data: JSON.stringify({ status: status }),
                // dataType: "JSON",
                // contentType: "application/json",
                beforeSend: function () {

                },
                complete: function () {

                },
                fail: function (e) {

                },
                success: function () {
                    $del_btn.hide().parent().parent().parent().fadeOut(500, function () {
                        $del_btn.remove();
                    });
                },
            });
            // $(this).hide().parent().parent().parent().fadeOut(500, function () {
            //     $(this).remove();
            // });
        });
        $chk_box.attr({
            "type": "checkbox",
            "name": item_id
        });
        $chk_box.prop("checked", is_checked);
        $chk_box.click(function (e) {
            let status = $(this).prop("checked");
            $(this).prop("checked", !status);
            $.ajax({
                type: "PUT",
                url: "/todolist/" + $(this).attr("name"),
                data: JSON.stringify({ status: status }),
                dataType: "JSON",
                contentType: "application/json",
                beforeSend: function () {

                },
                complete: function () {

                },
                fail: function (e) {

                },
                success: function () {
                    $chk_box.prop("checked", status);
                },
            });
        });
        $button_container.append($del_btn);
        $content_container.prepend($chk_box);
        $new_item_container.append($content_container)
        $new_item_container.append($button_container);
        $new_item.append($new_item_container);
        $(this).prepend($new_item.hide().fadeIn(600));
    };
    $("body").css("background-color", "#9400D3");
    $("#new_item").click(function (e) {
        let user_input = $("#input_area").val();
        if (user_input.trim() != "") {
            $.ajax({
                type: "POST",
                contentType: "application/json",
                data: JSON.stringify({ user_input: user_input }),
                dataType: "JSON",
                url: "/todolist",
                beforeSend: function () {
                    //
                },
                complete: function () {
                    //alert("request sent")
                },
                fail: function (e) {
                    alert("li way server is die die")
                },
                success: function (data) {
                    e.preventDefault();
                    $("#input_area").val('');
                    $("#to_do_list_container").createNewItems(user_input, data.item_id, false);
                },
            });

            // //test
            // e.preventDefault();
            // $("#input_area").val('');
            // $("#to_do_list_container").createNewItems(user_input, 0);
            // //test
        }
        else {
            alert("please input something!")
        }
    });
    $.get("/todolist/lists", function (data) {
        for (let i = data.length - 1; i >= 0; i--) {
            $("#to_do_list_container").createNewItems(data[i].user_input, data[i].item_id, data[i].status);
        }
    });
});