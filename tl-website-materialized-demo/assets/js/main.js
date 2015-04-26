var main = function() {
    //initialize smoothscroll.js
    smoothScroll.init();

    //start the carousels
    $('.carousel').carousel();

    //navigation hamburger icon
    $(".button-collapse").sideNav();
};

$(document).ready(main);