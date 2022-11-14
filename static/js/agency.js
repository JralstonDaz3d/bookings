(function() {
  "use strict"; // Start of use strict

  var mainNav = document.querySelector('#mainNav');

  if (mainNav) {

    var navbarCollapse = mainNav.querySelector('.navbar-collapse');
    
    if (navbarCollapse) {
      
      var collapse = new bootstrap.Collapse(navbarCollapse, {
        toggle: false
      });
      
      var navbarItems = navbarCollapse.querySelectorAll('a');
      
      // Closes responsive menu when a scroll trigger link is clicked
      for (var item of navbarItems) {
        item.addEventListener('click', function (event) {
          collapse.hide();
        });
      }
    }

    // Collapse Navbar
    var collapseNavbar = function() {

      var scrollTop = (window.pageYOffset !== undefined) ? window.pageYOffset : (document.documentElement || document.body.parentNode || document.body).scrollTop;

      if (scrollTop > 100) {
        mainNav.classList.add("navbar-shrink");
      } else {
        mainNav.classList.remove("navbar-shrink");
      }
    };
    // Collapse now if page is not at top
    collapseNavbar();
    // Collapse the navbar when page is scrolled
    document.addEventListener("scroll", collapseNavbar);

    // Hide navbar when modals trigger
    var modals = document.querySelectorAll('.portfolio-modal');
      
    for (var modal of modals) {
      
      modal.addEventListener('shown.bs.modal', function (event) {
        mainNav.classList.add('d-none');
      });
        
      modal.addEventListener('hidden.bs.modal', function (event) {
        mainNav.classList.remove('d-none');
      });
    }
  }

  // Range picker for Reservation Date selection
  const elem = document.getElementById('reservation-dates');
  const rangepicker = new DateRangePicker(elem, {
    // ...options
    format: "yyyy-mm-dd"
  });


  //notie alerts
  function Notify(msg, msgType) {
    notie.alert({
      type: msgType,
      text: msg,
      position: 'bottom'
    });
  }

  const reservationalert = document.getElementById('sendMessageButton');
  reservationalert.addEventListener("click", function (){
    //Notify("Reservation has been made!", "success");
    Swal.fire({
      title: 'Reservation Ready!',
      text: 'Do you want to reserve it?',
      icon: 'question',
      confirmButtonText: 'Yes'
    })
  })

})(); // End of use strict


