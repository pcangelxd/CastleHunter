var a = false;
$(document).ready(function() {
    $(".lang-select-arrow").click(function(){
        if (a == false){
          $(".lang-select-options").addClass("show");
          a = true;
        }else if (a == true) {
          $(".lang-select-options").removeClass("show");
          a = false;
        }
    });
    $("#eng").click(function(){
        $("#lang").attr("src","static/images/eng.png");
        $("#lang").attr("alt","eng");
        $(".lang-select-options").removeClass("show");
    });
    $("#rus").click(function(){
        $("#lang").attr("src","static/images/rus.png");
        $("#lang").attr("alt","ru");
        $(".lang-select-options").removeClass("show");
    });
});
