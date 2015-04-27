var main = function() {
    //initialize smoothscroll.js
    smoothScroll.init();

    $('#carousel-testimonials').flexslider({
        animation: "slide",
        smoothHeight: true,
        slideshowSpeed: 3000,
        pauseOnHover: true,
        controlNav: false,
        directionNav: false
    });

    //navigation hamburger icon
    $(".button-collapse").sideNav();
};

$(document).ready(main);