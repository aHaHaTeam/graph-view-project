const signUp = document.querySelector(".signup-link"),
    login = document.querySelector(".login-link"),
    signUpContainer = document.querySelector(".form.signup"),
    loginContainer = document.querySelector(".form.login"),
    signUpForm = document.getElementById("signUpForm"),
    loginForm = document.getElementById("loginForm"),
    userPasswordLabel = document.getElementById("userPasswordLabel"),
    userLoginLabel = document.getElementById("userLoginLabel");

// (Dis)appear signup/login form
signUp.addEventListener("click", () => {
    loginContainer.setAttribute("hidden", "");
    signUpContainer.removeAttribute("hidden");
});

// Submitting login form
loginForm.addEventListener("submit", (e) => {
    e.preventDefault();

    let login = document.getElementById("userLogin").value;
    let password = document.getElementById("userPassword").value;

    if (login === "") {
        userLoginLabel.innerText = "Login can't be empty"
        userLoginLabel.classList.add("error")
    } else if (password === "") {
        userPasswordLabel.innerText = "Password can't be empty"
        userPasswordLabel.classList.add("error")
    } else {
        fetch("http://localhost:8080/login", {
            method: "POST",
            body: JSON.stringify({login: login, password: password}),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        }).then(response => {
            if (response.headers.get("success") === "User logged in") {
                document.location.reload();
                document.location.replace("http://localhost:8080/")
                return null
            } else {
                userPasswordLabel.innerText = response.headers.get("success")
                userPasswordLabel.classList.add("error")
                return response
            }
        })
    }
});
