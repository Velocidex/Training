function toggleLeaf(e) {
    let parent = $(e).parent();
    let button = parent.children("span");
    let open = button.children("i.fa-angle-right");

    if (open.length) {
        button.find("i").
            addClass("fa-angle-down").
            removeClass("fa-angle-right");
        parent.addClass("toc_open").removeClass("toc_close");
    } else {
        button.find("i").addClass("fa-angle-right").
            removeClass("fa-angle-down");
        parent.addClass("toc_close").removeClass("toc_open");
    };
};

function toggleAll() {
    $.each($(".toc").find("i.category-icon"), function(idx, x) {
        toggleLeaf($(x).parent());
    });
}

function setTheme(name) {
    $("html").attr("data-bs-theme", name);
}
