function gaBurgerChange(x) {
  var r = document.querySelector('body');
  var rs = getComputedStyle(r);

  if (rs.getPropertyValue('--sidebar-width') == '200px') {
    r.style.setProperty('--sidebar-width', '50px');
    r.style.setProperty('--sideMenuItem-Display', 'none');
    r.style.setProperty('--sideMenuTooltip', 'visible');
    document.getElementById("toggleLeftMenu").innerHTML = "<i class='fas fa-arrow-circle-right'></i>";
  } else {
    r.style.setProperty('--sidebar-width', '200px');
    r.style.setProperty('--sideMenuItem-Display', 'inline');
    r.style.setProperty('--sideMenuTooltip', 'hidden');
    document.getElementById("toggleLeftMenu").innerHTML = "<i class='fas fa-arrow-circle-left'></i>";
  }
}

var acc = document.getElementsByClassName("accordion");
var i;
for (i = 0; i < acc.length; i++) {
  acc[i].addEventListener("click", function() {
    this.classList.toggle("active");
    var panel = this.nextElementSibling;
    if (panel.style.maxHeight) {
      panel.style.maxHeight = null;
    } else {
      panel.style.maxHeight = panel.scrollHeight + "px";
    } 
  });
}

var acc = document.getElementsByClassName("mainMenu");
var i;
for (i = 0; i < acc.length; i++) {
  acc[i].addEventListener("click", function() {
    this.classList.toggle("active");
    var panel = this.nextElementSibling;
    if (panel.style.maxHeight) {
      panel.style.maxHeight = null;
    } else {
      panel.style.maxHeight = panel.scrollHeight + "px";
    } 
  });
}

function openCity(evt, cityName) {
  var i, tabcontent, tablinks;
  tabcontent = document.getElementsByClassName("tabcontent");
  for (i = 0; i < tabcontent.length; i++) {
    tabcontent[i].style.display = "none";
  }
  tablinks = document.getElementsByClassName("tablinks");
  for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace(" active", "");
  }
  document.getElementById(cityName).style.display = "block";
  evt.currentTarget.className += " active";
}

// document.getElementById("defaultOpen").click();