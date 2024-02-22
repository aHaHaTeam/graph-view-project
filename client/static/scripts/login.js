const container = document.querySelector(".container"),
    pwShowHide = document.querySelectorAll(".showHidePw"),
    pwFields = document.querySelectorAll(".password"),
    signUp = document.querySelector(".signup-link"),
    login = document.querySelector(".login-link"),
    signUpContainer = document.querySelector(".form.signup"),
    loginContainer = document.querySelector(".form.login"),
    signUpForm = document.getElementById("signUpForm"),
    loginForm = document.getElementById("loginForm");

// Show/hide password and change icon
pwShowHide.forEach(eyeIcon => {
    eyeIcon.addEventListener("click", () => {
        pwFields.forEach(pwField => {
            if (pwField.type === "password") {
                pwField.type = "text";

                pwShowHide.forEach(icon => {
                    icon.classList.replace("uil-eye-slash", "uil-eye");
                })
            } else {
                pwField.type = "password";

                pwShowHide.forEach(icon => {
                    icon.classList.replace("uil-eye", "uil-eye-slash");
                })
            }
        })
    })
})

// (Dis)appear signup/login form
signUp.addEventListener("click", () => {
    loginContainer.setAttribute("hidden", "");
    signUpContainer.removeAttribute("hidden");
});

login.addEventListener("click", () => {
    loginContainer.removeAttribute("hidden");
    signUpContainer.setAttribute("hidden", "");
});


// Submitting login form
loginForm.addEventListener("submit", (e) => {
    e.preventDefault();

    let login = document.getElementById("userLogin").value;
    let password = document.getElementById("userPassword").value;

    if (login.value === "" || password.value === "") {
        throw new Error()
    } else {
        console.log("submitting")
        SubmitForm("http://localhost:8080/login", {login: login, password: password})
    }
});

// Submitting signup form
signUpForm.addEventListener("submit", (e) => {
    e.preventDefault();

    let login = document.getElementById("loginCreation").value;
    let email = document.getElementById("emailCreation").value;
    let password = document.getElementById("passwordCreation").value;
    let passwordConfirmation = document.getElementById("passwordConfirmation").value;
    console.log(JSON.stringify({email: login, password: password}))
    if (login.value === "" || password.value !== passwordConfirmation.value) {
        throw new Error()
    } else {
        SubmitForm("http://localhost:8080/signup", {login: login, email: email, password: password})
    }
});


function SubmitForm(url, data) {
    console.log(JSON.stringify(data))
    let response = fetch(url, {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    }).then(response => {
        if (response.headers.get("success") === "user logged in" ||
            response.headers.get("success") === "user created") {
            document.location.replace("http://localhost:8080/")
        }
    })
}