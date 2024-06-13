
function initializeSlides() {
    let handle = $("<div class=\"revealer\">Reveal Solution</div>").
        click(function() {
            $(this).next().toggleClass("solution-closed");
        });
   $(".solution").before(handle);
}
