const loginCreationLabel = document.getElementById("loginCreationLabel"),
    emailCreationLabel = document.getElementById("emailCreationLabel"),
    passwordCreationLabel = document.getElementById("passwordCreationLabel"),
    passwordConfirmationLabel = document.getElementById("passwordConfirmationLabel");

// (Dis)appear signup/login form
login.addEventListener("click", () => {
    loginContainer.removeAttribute("hidden");
    signUpContainer.setAttribute("hidden", "");
});

// Submitting signup form
signUpForm.addEventListener("submit", (e) => {
    e.preventDefault();

    let login = document.getElementById("loginCreation").value;
    let email = document.getElementById("emailCreation").value;
    let password = document.getElementById("passwordCreation").value;
    let passwordConfirmation = document.getElementById("passwordConfirmation").value;
    console.log(JSON.stringify({email: login, password: password}))

    if (login === "") {
        loginCreationLabel.innerText = "Login can't be empty"
        loginCreationLabel.classList.add("error")
    } else if (email === "") {
        emailCreationLabel.innerText = "Email can't be empty"
        emailCreationLabel.classList.add("error")
    } else if (password === "") {
        passwordCreationLabel.innerText = "Password can't be empty"
        passwordCreationLabel.classList.add("error")
    } else if (password !== passwordConfirmation) {
        passwordConfirmationLabel.innerText = "Passwords must match"
        passwordConfirmationLabel.classList.add("error")
    } else {
        fetch("http://localhost:8080/signup", {
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
