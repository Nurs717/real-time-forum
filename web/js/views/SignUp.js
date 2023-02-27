import AbstractView from "./AbstractView.js";
import {getError500} from "./Shared.js";

const regHTML = `
<div class="sign-up-container">
    <form id="sign-up-form" action="/">
        <h1>Registration</h1>
        <div class="input-control">
            <label for="username">Username</label>
            <input id="username" name="username" type="text">
            <div class="error"></div>
        </div>
        <div class="input-control">
            <label for="age">Age</label>
            <input id="age" name="age" type="number" min="18" max="100" autocomplete="off">
            <div class="error"></div>
        </div>
        <div class="input-control">
            <label for="gender">Gender</label>
            <select name="gender" id="gender">
                <option value="male">Male</option>
                <option value="female">Female</option>
            </select>
            <div class="error"></div>
        </div>
        <div class="input-control">
            <label for="firstname">First Name</label>
            <input id="firstname" name="firstname" type="text">
            <div class="error"></div>
        </div>
        <div class="input-control">
            <label for="lastname">Last Name</label>
            <input id="lastname" name="lastname" type="text">
            <div class="error"></div>
        </div>
        <div class="input-control">
            <label for="email">Email</label>
            <input id="email" name="email" type="text">
            <div class="error"></div>
        </div>
        <div class="input-control">
            <label for="password">Password</label>
            <input id="password" name="password" type="password">
            <div class="error"></div>
        </div>
        <div class="input-control">
            <label for="password2">Password again</label>
            <input id="password2" name="password2" type="password">
            <div class="error"></div>
        </div>
        <button type="submit">Sign Up</button>
    </form>
</div>
`

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Register");
    }

    async getHtml() {
        return `
        <header id="top">
            <div class="row">
                <div class="column lpad">
                    <div class="logo">
                        <a href="http://localhost:8081/" data-link>MyForum</a>
                    </div>
                </div>
                <div class="column ar lpad">
                    <nav class="menu">
                        <a href="/" id="h_posts" class="current nav_link" data-link>Posts</a>
                    </nav>
                </div>
            </div>
        </header>
        <div class="row">
        </div>
        `;
    }

    async Init() {
        await this.drawSignUpForm()
        await this.signUp();
    }

    async drawSignUpForm() {
        await this.drawNavMenuLoggedOut()
        let app = document.getElementById("app");
        let body_container = document.createElement("div")
        body_container.id = "sign-up"
        body_container.className = "sign-up"
        app.appendChild(body_container)
        body_container.innerHTML = regHTML
    }

    async drawNavMenuLoggedOut() {
        let menu = document.getElementsByClassName("menu");
        let login = document.createElement('a');
        login.setAttribute('href', '/login');
        login.setAttribute('data-link', '');
        login.innerHTML = 'Log In';
        menu[0].appendChild(login);
    }

    async signUp() {
        const url = "http://localhost:8080/signup"

        const form = document.getElementById('sign-up-form');
        const username = document.getElementById('username');
        const age = document.getElementById('age');
        const gender = document.getElementById('gender');
        const firstname = document.getElementById('firstname');
        const lastname = document.getElementById('lastname')
        const email = document.getElementById('email');
        const password = document.getElementById('password');
        const password2 = document.getElementById('password2');

        let unique = false;

        form.addEventListener('submit', e => {
            e.preventDefault();

            if (validateInputs()) {
                const formdata = new FormData(form)
                fetch(url, {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        username: formdata.get("username"),
                        age: formdata.get("age"),
                        gender: formdata.get("gender"),
                        firstname: formdata.get("firstname"),
                        lastname: formdata.get("lastname"),
                        email: formdata.get("email"),
                        password: formdata.get("password") }),
                })
                    .then(
                        (response) => {
                            if (response.status === 201) {
                                console.log("201", response.status);
                                const sign_up = document.getElementById('sign-up');
                                sign_up.innerHTML = successHTML;
                            } else if (response.status === 409) {
                                unique = true;
                            } else if (response.status === 500) {
                                let app = document.getElementById("app");
                                app.innerHTML = getError500;
                                return
                            }
                            return response.json();
                        }
                    )
                    .then(
                        (data) => {
                            if (unique === true) {
                                if (data.error_type === 'user') {
                                    setError(username, data.error_message);
                                } else if (data.error_type === 'mail') {
                                    setError(email, data.error_message);
                                }
                            }
                        }
                    )
                    .catch(
                        (error) => console.error(error)
                );
            }
        });

        const setError = (element, message) => {
            const inputControl = element.parentElement;
            const errorDisplay = inputControl.querySelector('.error');

            errorDisplay.innerText = message;
            inputControl.classList.add('error');
            inputControl.classList.remove('success')
        }

        const setSuccess = element => {
            const inputControl = element.parentElement;
            const errorDisplay = inputControl.querySelector('.error');

            errorDisplay.innerText = '';
            inputControl.classList.add('success');
            inputControl.classList.remove('error');
        };

        const isValidEmail = email => {
            const regex_mail = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            return regex_mail.test(String(email).toLowerCase());
        }

        const isValidPassword = password => {
            return /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[A-Za-z\d]/.test(password);
        }

        const validateInputs = () => {
            let flag = true;
            const usernameValue = username.value.trim();
            const ageValue = age.value.trim();
            const genderValue = gender.value.trim();
            const firstnameValue = firstname.value.trim();
            const lastnameValue = lastname.value.trim();
            const emailValue = email.value.trim();
            const passwordValue = password.value.trim();
            const password2Value = password2.value.trim();

            if(usernameValue === '') {
                setError(username, 'Username is required');
                flag = false;
            } else {
                setSuccess(username);
            }

            if (ageValue === '') {
                setError(age, 'Age is required')
                flag = false;
            } else {
                setSuccess(age);
            }

            if(genderValue === '') {
                setError(gender, 'Gender is required');
                flag = false;
            } else {
                setSuccess(gender);
            }

            if(firstnameValue === '') {
                setError(firstname, 'First Name is required');
                flag = false;
            } else {
                setSuccess(firstname)
            }

            if(lastnameValue === '') {
                setError(lastname, 'Last Name is required');
                flag = false;
            } else {
                setSuccess(lastname)
            }

            if(emailValue === '') {
                setError(email, 'Email is required');
                flag = false;
            } else if (!isValidEmail(emailValue)) {
                setError(email, 'Provide a valid email address');
                flag = false;
            } else {
                setSuccess(email);
            }

            if (passwordValue === '') {
                setError(password, 'Password is required');
                flag = false;
            } else if (passwordValue.length < 8 ) {
                setError(password, 'Password must be at least 8 character.')
                flag = false;
            } else if (!isValidPassword(passwordValue)){
                setError(password, 'Required at least one uppercase, one lowercase and one number')
                flag = false;
            } else {
                setSuccess(password);
            }

            if(password2Value === '') {
                setError(password2, 'Please confirm your password');
                flag = false;
            } else if (password2Value !== passwordValue) {
                setError(password2, "Passwords doesn't match");
                flag = false;
            } else {
                setSuccess(password2);
            }

            return flag;
        };
    }
}

const successHTML= `
<div class="success-body">
    <div class="success-container">
        <div style="border-radius:200px; height:200px; width:200px; background: #F8FAF5; margin:0 auto;">
            <i class="success-checkmark">✓</i>
        </div>
        <h1>Success</h1>
        <p>Congratulations, your account<br/>has been successfully created.</p>
        <a href="http://localhost:8081/login" class="contBtn">Login</a>
    </div>
</div>
`