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




// Use this to call alerts and modals

// use like:
// attention.toast({title:"Hello World",icon:"success",position:"top"});
// attention.sweetalert({title:"Hello World",text:"I am more text.",icon:"success",confirmButtonText:"Yes"});


let attention = Prompt();

// Notie alerts
function Notify(msg, msgType) {
  notie.alert({
    type: msgType,
    text: msg,
    position: 'bottom'
  });
}


function Prompt(){

  // Use Toast
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

  // Use SweetAlert
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

  // Use SweetAlert with Multi Input Form or custom html
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
      backdrop: true,
      focusConfirm: false,
      showCancelButton: true,
      width: '100%',
      allowOutsideClick: false,
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


  // Use SweetAlert with Multi Input Form or custom html
  async function sweetformroom(c) {
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
      backdrop: true,
      showConfirmButton: false,
      showCancelButton: true,
      width: '100%',
      allowOutsideClick: false,
      focusCancel: true,
      willOpen: () => {
        const elem = document.getElementById('reservation-dates');
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
      //console.log("form values:" + formValues);
      Swal.fire(JSON.stringify(formValues))
    }
  }

  return {
    toast: toast,
    sweetalert: sweetalert,
    sweetform: sweetform,
    sweetformroom: sweetformroom,
  }
}

