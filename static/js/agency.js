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
})(); // End of use strict



let attention = Prompt();

// Range picker for Reservation Date selection
const elem = document.getElementById('reservation-dates');
const rangepicker = new DateRangePicker(elem, {
  // ...options
  format: "yyyy-mm-dd",
  autohide: true
});


//notie alerts
function Notify(msg, msgType) {
  notie.alert({
    type: msgType,
    text: msg,
    position: 'bottom'
  });
}

//sendMessageButton
const reservationalert = document.getElementById('resmodal');
reservationalert.addEventListener("click", function (){
  //attention.toast({title:"Hello World",icon:"success",position:"top"});
  //attention.sweetalert({title:"Hello World",text:"I am more text.",icon:"success",confirmButtonText:"Yes"});

  //Set the reservation form html into a var.
  let formhtml= `
    <form id="contactForm2" name="contactForm2" method="post" action="" novalidate class="needs-validation m-auto">
        <div class="container px-3">
          <div class="row">
              <div class="col-12 text-center">
                  <h4 class="mb-3">Occupant Information</h4>
              </div>
              <div class="col-12">
                  <div class="form-group mb-3"><input class="form-control" type="text" id="name" placeholder="Your Name *" required=""><small class="form-text text-danger flex-grow-1 lead"></small></div>
                  <div class="form-group mb-3"><input class="form-control" type="email" id="email" placeholder="Your Email *" required=""><small class="form-text text-danger lead"></small></div>
                  <div class="form-group mb-3"><input class="form-control" type="tel" placeholder="Your Phone *" required=""><small class="form-text text-danger lead"></small></div>
                  <div class="form-group mb-3"><textarea class="form-control" id="message" placeholder="Additional Accommodations" rows="6" style="height:178px;"></textarea><small class="form-text text-danger lead"></small></div>
              </div>
          </div>
          <div class="row form-row" id="reservation-dates2">
              <div class="col-lg-12 text-center mt-5">
                  <h4 class="mb-3">Reservation Dates</h4>          
                  <input disabled required class="form-control-lg mb-3" type="text" id="res-start" name="res-start" autocomplete="off" placeholder="Arrival Date*">          
                  <input disabled required class="form-control-lg" type="text" id="res-end" name="res-end" autocomplete="off" placeholder="Departure Date*">
              </div>             
          </div>
        </div>
    </form>
  `;

  attention.sweetform({
    title:"Make a Reservation",
    html: formhtml
  });
})

//Use Toast
function Prompt(){
  let toast = function(c) {

    const {
        title = "",
        icon = "success",
        position = "top-center",
    } = c;

    const Toast = Swal.mixin({
      toast: true,
      title: title,
      position: position,
      icon: icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
      }
    })

    Toast.fire({})
  }

  //use SweetAlert
  let sweetalert = function(c) {
    const {
      title = "",
      text = "",
      icon = "success",
      confirmButtonText = "Yes",
      position = "top-center",
    } = c;

    Swal.fire({
      title: title,
      text: text,
      icon: icon,
      confirmButtonText: confirmButtonText
    })
  }

  //Use SweetAlert with Multi Input Form or custom html
  async function sweetform(c) {
    const {
      title = "",
      html = "",
      focusConfirm = true,
      showCancelButton= true,
      icon = "success",
      confirmButtonText = "Yes",
      position = "top-center",
    } = c;

    const { value: formValues } = await Swal.fire({
      title: title,
      html: html,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      willOpen: () => {
        const elem = document.getElementById('reservation-dates2');
        const rangepicker = new DateRangePicker(elem, {
          format: "yyyy-mm-dd",

          autohide: true
        });
      },
      preConfirm: () => {
        return [
          document.getElementById('res-start').value,
          document.getElementById('res-end').value
        ]
      },
      didOpen: () => {
        document.getElementById('res-start').removeAttribute('disabled');
        document.getElementById('res-end').removeAttribute('disabled');
      }
    })

    if (formValues) {
      Swal.fire(JSON.stringify(formValues))
    }
  }


  return {
    toast: toast,
    sweetalert: sweetalert,
    sweetform: sweetform,
  }
}

