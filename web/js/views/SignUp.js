import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("SignUp");
    }

    async getHtml() {
        return `
        <h1>Please Register</h1>
        <form id="signUpInputForm" onSubmit="return false;">
        <div>
        <label for="username">UserName</label>
        <input id="username" name="username" type="text">
        </div>
        <label for="firstname">First Name</label>
        <input id="firstname" name="firstname" type="text">
        </div>
        <div>
        <label for="lastname">Last Name</label>
        <input id="lastname" name="lastname" type="text">
        </div>
        <div>
        <label for="email">Email</label>
        <input id="email" name="email" type="text">
        </div>
        <div>
        <label for="password">Password</label>
        <input id="password" name="" type="text">
        </div>
        <div>
        <button type="submit" style="width:100px;">Register</button>
        </div>
        </form>
        `;
    }

    async Init() {
        this.signUp();
    }

    async signUp() {
        const url = "http://localhost:8080/signup"

        var inputForm = document.getElementById("signUpInputForm")

        inputForm.addEventListener("submit", (e) => {

            //prevent auto submission
            e.preventDefault()

            const formdata = new FormData(inputForm)
            fetch(url, {

                method: "POST",
                body: JSON.stringify({ username: formdata.get("username"), firstname: formdata.get("firstname"), lastname: formdata.get("lastname"), email: formdata.get("email"), password: formdata.get("password") }),
            }).catch(
                error => console.error(error)
            )
        })
    }
}