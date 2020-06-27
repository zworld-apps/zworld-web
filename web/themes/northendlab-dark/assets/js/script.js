(function ($) {
  'use strict';

  let search_open = false;

  // Preloader js
  $(window).on('load', function () {
    $('.preloader').fadeOut(100);
  });

  // navigation
  $(window).scroll(function () {
    if ($('.navigation').offset().top > 1) {
      $('.navigation').addClass('nav-bg');
    } else {
      $('.navigation').removeClass('nav-bg');
    }
  });

  // Search Form Open
  $('#searchOpen').on('click', function () {
    $('.search-wrapper').addClass('open');
	setTimeout(function() {
	  $('.search-box').focus();
	}, 400)

	search_open = true;
  });
  $('#searchClose').on('click', function () {
    $('.search-wrapper').removeClass('open');

	search_open = false;
  });

})(jQuery);
