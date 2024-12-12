let pass = document.getElementById("pass");
let confirmPass = document.getElementById("confirmPass");
let user = document.getElementById("user");
let email = document.getElementById("email");
let registerBtn = document.getElementById("registerBtn");
let usernameMessage = document.getElementById("usernameMessage");
let emailMessage = document.getElementById("emailMessage");
let confirmPassMessage = document.getElementById("confirmPassMessage");

function checkPassword() {
    const patterns = {
        lower: /(?=.*[a-z])/,
        upper: /(?=.*[A-Z])/,
        number: /(?=.*\d)/,
        symbol: /(?=.*[\W_])/,
        length: /^.{8,}$/
    };

    let Divs = {
        lower: document.querySelector('.lower'),
        upper: document.querySelector('.upper'),
        number: document.querySelector('.number'),
        symbol: document.querySelector('.symbol'),
        length: document.querySelector('.length'),
        passChecker: document.querySelector('.passChecker'),
    }

    Divs.passChecker.style.display = pass.value ? "block" : "none";

    for (const [k, v] of Object.entries(patterns)) {
        if (pass.value && pass.value.match(patterns[k])) {
            Divs[k].style.color = "green";
        } else {
            Divs[k].style.color = "red";
        }
    }
}

function validateForm() {
    let usernameValid = /^[\w]+$/.test(user.value);
    let emailValid = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(email.value);
    let confirmPasswordValid = pass.value === confirmPass.value;
    let passwordValid = pass.value.match(/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}/);

    if (!usernameValid && user.value !== "") {
        usernameMessage.textContent = "Username must contain only letters and numbers.";
        usernameMessage.classList.add("error-message");
        usernameMessage.style.display = "block";
    } else {
        usernameMessage.style.display = "none";
    }

    if (!emailValid && email.value !== "") {
        emailMessage.textContent = "Invalid email format.";
        emailMessage.classList.add("error-message");
        emailMessage.style.display = "block";
    } else {
        emailMessage.style.display = "none";
    }

    if (!confirmPasswordValid && confirmPass.value !== "") {
        confirmPassMessage.textContent = "Confirmation password does not match.";
        confirmPassMessage.classList.add("error-message");
        confirmPassMessage.style.display = "block";
    } else {
        confirmPassMessage.style.display = "none";
    }
    if (usernameValid && emailValid && confirmPasswordValid && passwordValid) {
        registerBtn.disabled = false;
    } else {
        registerBtn.disabled = true;
    }
}
pass.addEventListener("input", () => {
    checkPassword();
    validateForm();
});
user.addEventListener("input", validateForm);
email.addEventListener("input", validateForm);
confirmPass.addEventListener("input", validateForm);