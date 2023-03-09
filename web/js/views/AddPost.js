import AbstractView from "./AbstractView.js";
import Login from "./LogIn.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("AddPost");
    }

    async getHtml() {
        return `
        <h1>Creat Post</h1>
        <form id="inputForm" onSubmit="return false;">
       
        <div class="multiselect">
        <div id="click" class="selectBox">
        <select>
        <option>Categories</option>
        </select>
        <div class="overSelect"></div>
        </div>
        <div id="checkboxes">
        <label for="one">
        <input type="checkbox" name="sport" id="sport" value="sport" />Sport</label>
        <label for="two">
        <input type="checkbox" name="religion" id="religion" value="religion" />Religion</label>
        <label for="three">
        <input type="checkbox" name="programming" id="programming" value="programming" />Programming</label>
        </div>
        </div>

        <div>
        <label for="post_body">Title</label>
        <input id="title" name="post_title" type="text">
        </div>

        <div>
        <label for="post_body">Post</label>
        <input id="post" name="post_body" value="hello!!" type="text">
        </div>
        <button type="submit" style="width:100px;">Create Post</button>
        </form>
        `;
    }

    async Init() {
        const url = "http://localhost:8080/create-post"

        const inputForm = document.getElementById("inputForm")

        fetch(url, {
            mode: 'cors',
            credentials: 'include'
        }).then(async(resp) => {
            console.log("creat post resp:", resp.status)
            if (resp.status == 401) {
                window.history.pushState("", "", '/login');
                let view = new Login;
                document.querySelector("#app").innerHTML = await view.getHtml();
                view.logIn();
                return
            }
        })

        inputForm.addEventListener("submit", (e) => {

            //prevent auto submission
            e.preventDefault()

            const formdata = new FormData(inputForm)
            const body = JSON.stringify({
                post_body: formdata.get("post_body"),
                post_title: formdata.get("post_title"),
                category: [
                             formdata.get("sport") ?? null,
                             formdata.get("religion") ?? null,
                             formdata.get("programming") ?? null,
                            ].filter(val => val !== null)
            });

            fetch(url, {
                credentials: 'include',
                method: "POST",
                body: body,
            }).catch(
                error => console.error(error)
            )
        })

        let expanded = false;

        document.getElementById("click").addEventListener("click", (e) => {

            e.preventDefault()

            function showCheckboxes() {
                const checkboxes = document.getElementById("checkboxes");
                if (!expanded) {
                    checkboxes.style.display = "block";
                    expanded = true;
                } else {
                    checkboxes.style.display = "none";
                    expanded = false;
                }
            }
            showCheckboxes();
        });
    }
}