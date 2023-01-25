function removeError() {
  let instructorcode = document.getElementById("ic");
  let instructorname = document.getElementById("in");
  let department = document.getElementById("dp");
  let coursename = document.getElementById("cn_drop_down");
  let error_log = document.getElementById("error_log");
  instructorcode.classList.remove("error");
  instructorname.classList.remove("error");
  department.classList.remove("error");
  coursename.classList.remove("error");
  error_log.style.visibility = "hidden";
}
async function InsertInstructorValues() {
  let instructorcode = document.getElementById("ic");
  let instructorname = document.getElementById("in");
  let department = document.getElementById("dp");
  let coursename = document.getElementById("cn_drop_down");
  let error_log = document.getElementById("error_log");
  let redirect_to_login = document.getElementById("redirect_to_login");
  let inline_buttons = document.getElementById("inline_buttons");

  if (instructorcode.value === "") {
    instructorcode.classList.add("error");
  }

  if (instructorname.value === "") {
    instructorname.classList.add("error");
  }

  if (department.value === "") {
    department.classList.add("error");
  }

  if (coursename.value == "Choose Course") {
    coursename.classList.add("error");
  }
  let cookie_token = getCookie("token");
  console.log(cookie_token);
  let createInstructor = await fetch(
    `http://localhost:5050/insert-instructor-details`,
    {
      method: "POST",
      headers: { Token: cookie_token },

      body: JSON.stringify({
        instructor_code: instructorcode.value,
        instructor_name: instructorname.value,
        department: department.value,
        course_name: coursename.value,
      }),
    }
  );
  let response = await createInstructor.json();
  if (createInstructor.ok != true) {
    error_log.style.visibility = "visible";
    error_log.innerHTML = "Error submitting!<br>" + response.Err;
  } else if (createInstructor.ok == true) {
    document.getElementById("responseBody").innerHTML =
      "Added<br> Please Create Account";
    redirect_to_login.classList.add("diplay-property");
    inline_buttons.classList.add("inline_buttons_css");
    redirect_to_login.innerHTML = "Create Account";
    let URL = `http://localhost:5050` + response.URl;
    document.cookie = `url=${URL}`;
    localStorage.setItem("URL_Create_Login", URL);
  }
}

async function create_account_new() {
  let instructorcode = document.getElementById("ic").value;
  let instructorname = document.getElementById("in").value;
  let department = document.getElementById("dp").value;
  let coursename = document.getElementById("cn_drop_down").value;
  if (
    !instructorcode ||
    !instructorname ||
    !department ||
    coursename == "Choose Course"
  ) {
    alert("Please fill details");
    return;
  }
  window.location.replace("createInstructorAccount.html");
}
function setbackpage() {
  window.location.replace("instructorDetails.html");
}

function setdashboard() {
  window.location.replace("allinstructor.html");
}

function getCookie(name) {
  var nameEQ = name + "=";
  var ca = document.cookie.split(";");
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == " ") c = c.substring(1, c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
}
