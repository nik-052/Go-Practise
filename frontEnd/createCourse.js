function setbackpage() {
  window.location.replace("showCourse.html");
}

function setdashboard() {
  window.location.replace("dashboard-v2.html");
}

async function InsertCourseValues() {
  let course_name = document.getElementById("cn").value;
  let response_for_creation = document.getElementById("response_for_creation");
  if (!course_name) {
    alert("Please enter Course Name.");
    return;
  }
  let cookie_token = getCookie("token");
  let createCourse = await fetch(`http://localhost:5050/insert-course`, {
    method: "POST",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      course_name: course_name,
    }),
  });
  let response = await createCourse.json();
  if (!createCourse.ok) {
    response_for_creation.innerHTML = response + "!!";
  } else {
    let response_for_creation = document.getElementById(
      "response_for_creation"
    );
    response_for_creation.innerHTML = "Successfully Created";
  }
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

setInterval(checkTokenValidity, 300000);

async function checkTokenValidity() {
  let api_error;
  let cookie_token = getCookie("token");
  let api_response = await fetch(`http://localhost:5050/check-token-status`, {
    method: "GET",
    headers: { Token: cookie_token },
  }).catch((err) => {
    api_error = err;
  });
  if (api_error == "TypeError: Failed to fetch") {
    alert("Internal Server Error Please Login Again");
    window.location.replace("index.html");
    return "";
  }
  let response = await api_response.json();
  if (response == "token expired! Generate new token") {
    alert("Timed-out re login");
    window.location.replace("index.html");
    return;
  }
}